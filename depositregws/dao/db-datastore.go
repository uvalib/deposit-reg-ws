package dao

import (
	"database/sql"
	"fmt"
	"github.com/uvalib/deposit-reg-ws/depositregws/api"
	"github.com/uvalib/deposit-reg-ws/depositregws/config"
	"github.com/uvalib/deposit-reg-ws/depositregws/logger"
	"time"

	// needed
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type storage struct {
	*sql.DB
}

//
// newDBStore -- create the database singletomn
//
func newDBStore() (Storage, error) {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowOldPasswords=1&tls=%s&timeout=%s&readTimeout=%s&writeTimeout=%s",
		config.Configuration.DbUser,
		config.Configuration.DbPassphrase,
		config.Configuration.DbHost,
		config.Configuration.DbName,
		config.Configuration.DbSecure,
		config.Configuration.DbTimeout,
		config.Configuration.DbTimeout,
		config.Configuration.DbTimeout)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	//taken from https://github.com/go-sql-driver/mysql/issues/461
	db.SetConnMaxLifetime( time.Minute * 5 )
	db.SetMaxIdleConns(2 )
	db.SetMaxOpenConns(2 )

	return &storage{db}, nil
}

//
// CheckDB -- check our database health
//
func (s *storage) Check() error {
	return s.Ping()
}

func (s *storage) GetDepositRequest(id string) ([]*api.Registration, error) {

	rows, err := s.Query("SELECT * FROM depositrequest WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositRequestResults(rows)
}

func (s *storage) SearchDepositRequest(id string) ([]*api.Registration, error) {

	rows, err := s.Query("SELECT * FROM depositrequest WHERE id > ? ORDER BY id ASC", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositRequestResults(rows)
}

func (s *storage) CreateDepositRequest(reg api.Registration) (*api.Registration, error) {

	stmt, err := s.Prepare("INSERT INTO depositrequest( requester, user, department, degree ) VALUES(?,?,?,?)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(reg.Requester, reg.For, reg.Department, reg.Degree)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	reg.ID = strconv.FormatInt(lastID, 10)
	return &reg, nil
}

func (s *storage) DeleteDepositRequest(id string) (int64, error) {

	stmt, err := s.Prepare("DELETE FROM depositrequest WHERE id = ? LIMIT 1")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}

func (s *storage) GetMappedOptions() ([]StringPair, error) {

	// first get all the mapped options
	rows, err := s.Query("SELECT d1.field_value, d2.field_value FROM fieldvalues d1, fieldvalues d2, fieldmaps m WHERE m.source_id = d1.id AND m.map_id = d2.id ORDER BY 1, 2 ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mapped, err := optionsMapResults(rows)
	if err != nil {
		return nil, err
	}

	// then get the unmapped
	rows, err = s.Query("SELECT field_value, '' FROM fieldvalues f1 WHERE f1.field_name = 'department' AND f1.id NOT IN (SELECT source_id FROM fieldmaps) ORDER BY 1 ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	unmapped, err := optionsMapResults(rows)
	if err != nil {
		return nil, err
	}

	// merge them and return
	return append(mapped, unmapped...), nil
}

func (s *storage) GetAllOptions() ([]StringPair, error) {
	rows, err := s.Query("SELECT field_name, field_value from fieldvalues ORDER BY 1, 2 ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return optionsResults(rows)
}

func (s *storage) CreateOption(option api.Option) error {

	stmt, err := s.Prepare("INSERT INTO fieldvalues( field_name, field_value ) VALUES(?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(option.Option, option.Value)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) UpdateOptionMap(optionMap api.DepartmentMap) error {

	optionsSet, err := s.GetAllOptions()
	if err != nil {
		return err
	}

	// check the department is already known
	found := false
	for _, sp := range optionsSet {
		if sp.A == "department" && sp.B == optionMap.Department {
			found = true
			break
		}
	}

	if found == false {
		return fmt.Errorf("department %s does not exist", optionMap.Department)
	}

	// check each degree is already known
	for _, degree := range optionMap.Degrees {

		found = false
		for _, sp := range optionsSet {
			if sp.A == "degree" && sp.B == degree {
				found = true
				break
			}
		}

		if found == false {
			return fmt.Errorf("degree %s does not exist", degree)
		}
	}

	// delete any existing mappings
	err = s.deleteAllOptionMaps(optionMap.Department)
	if err != nil {
		return err
	}

	// add a mapping for each degree
	for _, degree := range optionMap.Degrees {
		err = s.addOptionMap(optionMap.Department, degree)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *storage) deleteAllOptionMaps(department string) error {

	stmt, err := s.Prepare("DELETE FROM fieldmaps WHERE source_id = ( SELECT id from fieldvalues where field_name = ? AND field_value = ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec("department", department)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) addOptionMap(department string, degree string) error {

	stmt, err := s.Prepare("INSERT INTO fieldmaps( source_id, map_id ) VALUES( ( SELECT id from fieldvalues where field_name = ? AND field_value = ?), ( SELECT id from fieldvalues where field_name = ? AND field_value = ?) )")
	if err != nil {
		return err
	}

	_, err = stmt.Exec("department", department, "degree", degree)
	if err != nil {
		return err
	}

	return nil
}

func depositRequestResults(rows *sql.Rows) ([]*api.Registration, error) {

	var optional sql.NullString

	results := make([]*api.Registration, 0)
	for rows.Next() {
		reg := new(api.Registration)
		err := rows.Scan(&reg.ID,
			&reg.Requester,
			&reg.For,
			&reg.Department,
			&reg.Degree,
			&reg.Status,
			&reg.RequestDate,
			&optional)
		if err != nil {
			return nil, err
		}
		if optional.Valid {
			reg.DepositDate = optional.String
		}
		results = append(results, reg)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	logger.Log(fmt.Sprintf("Deposit registration request returns %d row(s)", len(results)))
	return results, nil
}

func optionsMapResults(rows *sql.Rows) ([]StringPair, error) {

	results := make([]StringPair, 0)
	for rows.Next() {
		var school string
		var degree string
		err := rows.Scan(&school, &degree)
		if err != nil {
			return nil, err
		}
		results = append(results, StringPair{A: school, B: degree})
	}

	logger.Log(fmt.Sprintf("Options map request returns %d row(s)", len(results)))
	return results, nil
}

func optionsResults(rows *sql.Rows) ([]StringPair, error) {

	results := make([]StringPair, 0)
	for rows.Next() {
		var name string
		var value string
		err := rows.Scan(&name, &value)
		if err != nil {
			return nil, err
		}
		results = append(results, StringPair{A: name, B: value})
	}

	logger.Log(fmt.Sprintf("Options request returns %d row(s)", len(results)))
	return results, nil
}

//
// end of file
//

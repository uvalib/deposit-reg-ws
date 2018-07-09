package dao

import (
	"database/sql"
	"depositregws/api"
	"depositregws/logger"
	"fmt"
	// needed
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type dbStruct struct {
	*sql.DB
}

//
// StringPair -- used for some results; not idiomatic
//
type StringPair struct {
	A string
	B string
}

//
// DB -- the database instance
//
var DB *dbStruct

//
// NewDB -- create the database singletomn
//
func NewDB(dbHost string, dbName string, dbUser string, dbPassword string, dbTimeout string) error {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowOldPasswords=1&timeout=%s&readTimeout=%s&writeTimeout=%s",
		dbUser,
		dbPassword,
		dbHost,
		dbName,
		dbTimeout, dbTimeout, dbTimeout)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	DB = &dbStruct{db}
	return nil
}

//
// CheckDB -- check our database health
//
func (db *dbStruct) CheckDB() error {
	return db.Ping()
}

func (db *dbStruct) GetDepositRequest(id string) ([]*api.Registration, error) {

	rows, err := db.Query("SELECT * FROM depositrequest WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositRequestResults(rows)
}

func (db *dbStruct) SearchDepositRequest(id string) ([]*api.Registration, error) {

	rows, err := db.Query("SELECT * FROM depositrequest WHERE id > ? ORDER BY id ASC", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositRequestResults(rows)
}

func (db *dbStruct) CreateDepositRequest(reg api.Registration) (*api.Registration, error) {

	stmt, err := db.Prepare("INSERT INTO depositrequest( requester, user, department, degree ) VALUES(?,?,?,?)")
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

func (db *dbStruct) DeleteDepositRequest(id string) (int64, error) {

	stmt, err := db.Prepare("DELETE FROM depositrequest WHERE id = ? LIMIT 1")
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

func (db *dbStruct) GetMappedOptions() ([]StringPair, error) {
	rows, err := db.Query("SELECT d1.field_value, d2.field_value FROM fieldvalues d1, fieldvalues d2, fieldmaps m WHERE m.source_id = d1.id AND m.map_id = d2.id ORDER BY 1, 2 ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return optionsMapResults(rows)
}

func (db *dbStruct) GetAllOptions() ([]StringPair, error) {
	rows, err := db.Query("SELECT field_name, field_value from fieldvalues ORDER BY 1, 2 ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return optionsResults(rows)
}

func (db *dbStruct) CreateOption( option api.Option ) error {

	return nil
}


func (db *dbStruct) UpdateOptionMap( optionMap api.DepartmentMap ) error {

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

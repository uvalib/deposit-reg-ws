package dao

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "depositregws/api"
    "log"
    "strconv"
)

type Datastore interface {
    Check( ) error
    Get( id string ) ( [] * api.Registration, error )
//    Search( ) ( [ ] * api.Registration, error )
    Create( reg api.Registration ) ( * api.Registration, error )
    Delete( id string ) error
}

type DB struct {
    *sql.DB
}

var Database * DB

func NewDB( dataSourceName string ) error {
    db, err := sql.Open( "mysql", dataSourceName )
    if err != nil {
        return err
    }
    if err = db.Ping( ); err != nil {
        return err
    }
    Database = &DB{ db }
    return nil
}

func ( db *DB ) Check( ) error {
    if err := db.Ping( ); err != nil {
        return err
    }
    return nil
}

func ( db *DB ) Get( id string ) ( [] * api.Registration, error ) {

    rows, err := db.Query( "SELECT * FROM depositrequest WHERE id = ? LIMIT 1", id )
    if err != nil {
        return nil, err
    }
    defer rows.Close( )

    return makeResults( rows )
}

func ( db *DB ) Create( reg api.Registration ) ( * api.Registration, error ) {

    stmt, err := db.Prepare( "INSERT INTO depositrequest( user, school, degree ) VALUES(?,?,?)" )
    if err != nil {
        return nil, err
    }

    res, err := stmt.Exec( reg.For, reg.School, reg.Degree )
    if err != nil {
        return nil, err
    }

    lastId, err := res.LastInsertId( )
    if err != nil {
        return nil, err
    }

    reg.Id = strconv.FormatInt( lastId, 10 )
    return &reg, nil
}

func ( db *DB ) Delete( id string ) ( int64, error ) {

    stmt, err := db.Prepare( "DELETE FROM depositrequest WHERE id = ? LIMIT 1" )
    if err != nil {
        return 0, err
    }

    res, err := stmt.Exec( id )
    if err != nil {
        return 0, err
    }

    rowCount, err := res.RowsAffected( )
    if err != nil {
        return 0, err
    }

    return rowCount, nil
}

func makeResults( rows * sql.Rows ) ( [] * api.Registration, error ) {

    var optional sql.NullString

    results := make([ ] * api.Registration, 0 )
    for rows.Next() {
        reg := new( api.Registration )
        err := rows.Scan( &reg.Id, &reg.For, &reg.School, &reg.Degree, &reg.Status, &reg.RequestDate, &optional )
        if err != nil {
            return nil, err
        }
        if optional.Valid {
            reg.DepositDate = optional.String
        }
        results = append( results, reg )
    }
    if err := rows.Err( ); err != nil {
        return nil, err
    }

    log.Printf( "Returning %d row(s)", len( results ) )
    return results, nil
}


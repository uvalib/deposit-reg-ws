package dao

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "depositregws/api"
    "log"
)

type Datastore interface {
    Check( ) error
    Get( id string ) ( [] * api.Registration, error )
//    Search( ) ( [ ] * api.Registration, error )
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


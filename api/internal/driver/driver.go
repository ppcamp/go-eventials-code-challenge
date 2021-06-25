package driver

//#region: Imports
import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

//#endregion

// Database is used to create a connection pool with sql database
type Database struct {
	SQL *sql.DB // With this approach, i can change it later
}

//#region: database connection
var dbconn = &Database{}

//#endregion

//#region: constants
const maxOpenDbConn = 10              // openned connections
const maxIleDbConn = 5                // iled connections
const maxDbLifetime = 5 * time.Minute // connections lifetime
//#endregion

// creates a db connection pool
func ConnectSQL(dsn string) (*Database, error) {
	d, err := newDatabase(dsn)
	if err != nil {
		log.Panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbconn.SQL = d

	// testing after modifications
	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbconn, nil
}

// Creates a new connection with database and check if its working
func newDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Check if database is online
func testDB(d *sql.DB) error {
	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}

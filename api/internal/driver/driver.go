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

// Database é usado na fila de conexão com o banco
type Database struct {
	SQL *sql.DB // Usando desta forma, posso alterar o banco futuramente
}

//#region: Conexão com o banco
var dbconn = &Database{}

//#endregion

//#region: Constantes
const maxOpenDbConn = 10              // conexões abertas
const maxIleDbConn = 5                // conexões inativas
const maxDbLifetime = 5 * time.Minute // lifetime da conexão
//#endregion

// Cria uma pool de conexões com o banco
func ConnectSQL(dsn string) (*Database, error) {
	d, err := newDatabase(dsn)
	if err != nil {
		log.Panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbconn.SQL = d

	// testando novamente a conexão após as modificações
	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbconn, nil
}

// Cria uma nova conexão com o banco e testa se funcionou
//
// @panic Se não conseguiu conectar com o banco
func newDatabase(dsn string) (*sql.DB, error) {
	//#region: Cria a conexão com o banco
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	//#endregion

	//#region: testa a conexão
	if err = db.Ping(); err != nil {
		return nil, err
	}

	//#endregion

	return db, nil
}

// Dá um ping no banco para ver se ele está online
func testDB(d *sql.DB) error {
	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/subosito/gotenv"

	"yawoen.com/app/internal/config"
	"yawoen.com/app/internal/driver"
)

// Configurando o número da porta que irá rodar o servidor
const port_number = ":3000"

// Configurações globais da aplicação
var app config.AppConfig

func main() {
	// configurando a aplicação
	db, err := setUp()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close() // fecha o banco quando a main for finalizada

	// Criando o servidor
	srv := &http.Server{
		Addr:    port_number,
		Handler: setUpRoutes(&app),
	}

	log.Println(fmt.Sprintf("Staring application on http://localhost%s", port_number))

	// Iniciando servidor
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// Startup routine for this module
func init() {
	gotenv.Load()
}

// Realiza algumas configurações antes de iniciar o servidor
//
// Como por exemplo:
// - conexão com banco;
// - criação e configuração de loggers
func setUp() (*driver.Database, error) {
	//#region: configurando o modo de execução
	app.InProduction = false
	//#endregion

	//#region: lendo arquivo env
	dbconfig := driver.LoadDatabaseConfig()
	//#endregion

	//#region: configurando os loggers
	app.InfoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime) // console
	app.InfoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime) // console
	//#endregion

	//#region: conexão com o banco
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL(dbconfig.GetDNS())
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")
	//#endregion

	return db, nil
}

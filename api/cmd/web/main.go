package main

import (
	"fmt"
	"log"
	"net/http"

	"yawoen.com/app/pkg/configs"
)

// Configurando o número da porta que irá rodar o servidor
const port_number = ":3000"

// Configurações globais da aplicação
var app configs.AppConfig

func main() {
	app.InProduction = false

	// Criando o servidor
	srv := &http.Server{
		Addr:    port_number,
		Handler: setUpRoutes(&app),
	}

	log.Println(fmt.Sprintf("Staring application on http://localhost%s", port_number))

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

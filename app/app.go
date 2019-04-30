package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"./commons"
	"./migration"
	"./routes"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la DB")
	flag.IntVar(&commons.Port, "port", 8080, "Puerto para el server web")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la migración...")
		migration.Migrate()
		log.Println("Finalizo la migración.")
	}

	// Inicia las rutas
	router := routes.InitRoutes()

	// Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}
	log.Printf("Inicia el servidor el http://localhost:%d", commons.Port)
	log.Println(server.ListenAndServe())
	log.Println("Final del programa")
}

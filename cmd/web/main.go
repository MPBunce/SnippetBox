package main

import (
	"log"
	"net/http"
	"flag"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {

	//Defining input address
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//Logging Info
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime);
	errorLog :=log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Initialize app
	app := &application {
		errorLog: errorLog,
		infoLog: infoLog,
	}

	//make a server struct so i can use my logger
	srv := &http.Server {
		Addr:
		*addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

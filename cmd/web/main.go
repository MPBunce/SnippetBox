package main

import (
	"MPBunce/SnippetBox/pkg/models/sqlite"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *sqlite.SnippetModel
}

func main() {

	//Defining input address
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//Logging Info
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB("database/snippetbox.db")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	//Initialize app
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &sqlite.SnippetModel{DB: db},
	}

	//make a server struct so i can use my logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

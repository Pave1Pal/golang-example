package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &Application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	port := flag.String("port", ":4000", "port")
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	log.Println("start web server on \"localhost" + *port + "\"")
	err := http.ListenAndServe(*port, mux)
	log.Fatal(err)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
}

func main() {
	portPtr := flag.Int("p", 9000, "Port to run server on")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		infoLog: infoLog,
		errorLog: errorLog,
	}

	srv := app.routes()

	err := srv.Run(fmt.Sprintf("0.0.0.0:%v", *portPtr))
	errorLog.Fatal(err.Error())
}
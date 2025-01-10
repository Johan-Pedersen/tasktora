package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"time"

	"tasktora/internals/config"
	customlogger "tasktora/internals/customLogger"
	"tasktora/internals/handlers"
	"tasktora/internals/models"

	_ "github.com/go-sql-driver/mysql" // New import
)

func main() {
	app := &config.App{
		InfoLogger:  customlogger.NewInfoLogger(),
		ErrorLogger: customlogger.NewErrorLogger(),
	}
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "taskDev:dev@/tasktora?parseTime=true", "MySQL data source name")
	flag.Parse()
	db, err := openDB(*dsn)
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}
	defer db.Close()

	tm := models.TaskModel{DB: db}
	app.TaskModel = tm
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.ErrorLogger,
		Handler:  handlers.Routes(app),
		// Add Idle, Read and Write timeouts to the server.
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	customlogger.InfoLog(app, fmt.Sprintf("Starting server on %s", *addr))
	err = srv.ListenAndServe()
	app.ErrorLogger.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

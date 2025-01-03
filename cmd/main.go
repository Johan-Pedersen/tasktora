package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"tasktora/internals/config"
	customlogger "tasktora/internals/customLogger"
	"tasktora/internals/handlers"
)

func main() {
	app := &config.App{
		InfoLogger:  customlogger.NewInfoLogger(),
		ErrorLogger: customlogger.NewErrorLogger(),
	}
	addr := flag.String("addr", ":4000", "HTTP network address")
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
	err := srv.ListenAndServe()
	app.ErrorLogger.Fatal(err)
}

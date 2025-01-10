package customlogger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"tasktora/internal/config"
)

func ServerError(app *config.App, w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLogger.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClinteError(app *config.App, w http.ResponseWriter, status int) {
	app.ErrorLogger.Output(2, http.StatusText(status))
	http.Error(w, http.StatusText(status), status)
}

func InfoLog(app *config.App, msg string) {
	app.InfoLogger.Println(msg)
}

func NewInfoLogger() *log.Logger {
	return log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
}

func NewErrorLogger() *log.Logger {
	return log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

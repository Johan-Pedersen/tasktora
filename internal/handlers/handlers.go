package handlers

import (
	"fmt"
	"net/http"
	"tasktora/internal/config"
	"tasktora/internal/middleware"

	customlogger "tasktora/internal/customLogger"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func Routes(app *config.App) http.Handler {
	router := httprouter.New()

	home := home(app)
	router.HandlerFunc(http.MethodGet, "/", home)

	standard := alice.New(middleware.RecoverPanic(app), middleware.LogRequest(app), middleware.SecureHeaders)

	return standard.Then(router)
}

func home(app *config.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := app.TaskModel
		task, err := tm.GetAll()
		if err != nil {
			customlogger.ServerError(app, w, err)
		}

		for _, t := range task {
			fmt.Fprintf(w, "hello, %s!\n", t.Title)
		}
	}
}

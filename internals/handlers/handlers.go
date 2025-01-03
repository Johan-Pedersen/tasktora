package handlers

import (
	"fmt"
	"net/http"

	"tasktora/internals/config"

	"tasktora/internals/middleware"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func Routes(app *config.App) http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", home)

	standard := alice.New(middleware.RecoverPanic(app), middleware.LogRequest(app), middleware.SecureHeaders)

	return standard.Then(router)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s!\n", "home")
}

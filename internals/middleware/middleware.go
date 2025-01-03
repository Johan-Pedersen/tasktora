package middleware

import (
	"fmt"
	"net/http"

	"tasktora/internals/config"
	customlogger "tasktora/internals/customLogger"
)

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

func LogRequest(app *config.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		// customlogger := customlogger.NewCliLogger()
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			customlogger.InfoLog(app, fmt.Sprintf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

			next.ServeHTTP(w, r)
		})
	}
}

func RecoverPanic(app *config.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a deferred function (which will always be run in the event
			// of a panic as Go unwinds the stack).
			defer func() {
				// Use the builtin recover function to check if there has been a
				// panic or not. If there has...
				if err := recover(); err != nil {
					// Set a "Connection: close" header on the response.
					w.Header().Set("Connection", "close")
					// Call the app.serverError helper method to return a 500
					// Internal Server response.
					customlogger.ServerError(app, w, fmt.Errorf("%s", err))
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

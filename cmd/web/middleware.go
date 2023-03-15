package main

import (
	"net/http"
	"github.com/justinas/nosurf"
)

// gets CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InPoduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and save the session in every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)

}

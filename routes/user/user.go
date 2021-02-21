package user

import (
	"auth0-tutorial/app"
	"auth0-tutorial/pages"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pages.Render(w, "user", session.Values["profile"])
}

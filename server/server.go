package server

import (
	"auth0-tutorial/app"
	"auth0-tutorial/callback"
	"auth0-tutorial/pages"
	"auth0-tutorial/routes/login"
	"auth0-tutorial/routes/logout"
	"auth0-tutorial/routes/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func NewServer(r *mux.Router) {
	r.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)

	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))

	log.Println("Listening on port [::]:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func IsAuthenticated(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, ok := session.Values["profile"]; !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		next(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	pages.Render(w, "home", nil)
}

package logout

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	domain := os.Getenv("AUTH0_DOMAIN")
	logoutURL, err := url.Parse("https://" + domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logoutURL.Path += "/v2/logout"
	params := url.Values{}
	scheme := "http"
	if r.TLS != nil {
		scheme = scheme + "s"
	}

	returnURL := fmt.Sprintf("%s://%s", scheme, r.Host)
	returnTo, err := url.Parse(returnURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params.Add("returnTo", returnTo.String())
	params.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutURL.RawQuery = params.Encode()

	http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
}

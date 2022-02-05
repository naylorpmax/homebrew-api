package api

import (
	"net/http"

	"golang.org/x/oauth2"
)

type(
	Login struct{
		OAuth2Config *oauth2.Config
	}
)

func (l *Login) Handler(w http.ResponseWriter, r *http.Request) error {
	url := l.OAuth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
}

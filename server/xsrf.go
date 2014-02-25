package server

import (
	"appengine"
	"appengine/user"
	"code.google.com/p/xsrftoken"
	"net/http"
)

const COOKIE_NAME = "XSRF-TOKEN"
const HEADER_NAME = "X-Xsrf-Token"

type XSRFEntity struct {
	Token string
}

func CreateXSRFToken(w http.ResponseWriter, r *http.Request) string {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		return ""
	}

	// first parameter should be changed to a secret key
	token := xsrftoken.Generate(appengine.AppID(c), u.ID, "API")

	cookie := http.Cookie{
		Name:  COOKIE_NAME,
		Value: token,
	}

	http.SetCookie(w, &cookie)

	return token
}

func ValidateXSRFToken(w http.ResponseWriter, r *http.Request) bool {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		return false
	}

	cookie, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		// err means no cookie found
		return false
	}

	cookieToken := cookie.Value
	if !xsrftoken.Valid(cookieToken, appengine.AppID(c), u.ID, "API") {
		return false
	}

	headerTokens := r.Header[HEADER_NAME]
	for _, headerToken := range headerTokens {
		if xsrftoken.Valid(headerToken, appengine.AppID(c), u.ID, "API") {
			if cookieToken == headerToken {
				return true
			}
		}
	}

	return false
}

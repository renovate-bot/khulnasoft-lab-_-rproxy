package auth

import (
	"crypto/subtle"
	"net/http"

	"github.com/khulnasoft-lab/rproxy/config"
)

func Middleware(handler http.HandlerFunc, config config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			respondUnauthorised(w)
			return
		}

		allow := false
		for _, user := range config.Users {
			if subtle.ConstantTimeCompare([]byte(username), []byte(user.Username)) == 1 && subtle.ConstantTimeCompare([]byte(password), []byte(user.Password)) == 1 {
				allow = true
				break
			}
		}

		if !allow {
			respondUnauthorised(w)
			return
		}

		handler(w, r)
	}
}

func respondUnauthorised(w http.ResponseWriter) {
	w.WriteHeader(401)
	w.Write([]byte("Unauthorized\n"))
}

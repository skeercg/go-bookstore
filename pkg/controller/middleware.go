package controller

import (
	"net/http"
	"strings"
)

func (c *Controller) userIdentity(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 {
			w.WriteHeader(401)
			return
		}

		_, err := c.services.AuthService.ParseToken(headerParts[1])

		if err != nil {
			w.WriteHeader(401)
			return
		}

		h.ServeHTTP(w, r)
	})
}

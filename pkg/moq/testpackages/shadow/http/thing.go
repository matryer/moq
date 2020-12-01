package http

import "net/http"

type Thing interface {
	Blah(w http.ResponseWriter, r *http.Request)
}

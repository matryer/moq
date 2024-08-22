package http

import "net/http"

// Thing is a test interface.
type Thing interface {
	Blah(w http.ResponseWriter, r *http.Request)
}

package shadow

import (
	"io"
	"net/http"
)

// Shadower is an interface, with a method, with a parameter whose name
// shadows an import name.
type Shadower interface {
	Shadow(io io.Reader)
	ShadowTwo(r io.Reader, io interface{})
	ShadowThree(http interface{}, srv *http.Server)
}

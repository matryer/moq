package genparamname

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type (
	// Go is a test function.
	Go func()

	myType struct{}
)

// Interface is a test interface.
type Interface interface {
	Method(
		*myType,
		[3]json.Number,
		[]byte,
		map[sql.NullString]io.Reader,
		func(conn net.Conn),
		Go,
		chan *httputil.BufferPool,
		struct{ URL *url.URL },
		interface {
			fmt.Stringer
			CookieJar() http.CookieJar
		},
	)
}

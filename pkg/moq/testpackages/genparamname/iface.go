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
	Go     func()
	myType struct{}
)

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

package genparamname

import (
	"database/sql"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Interface interface {
	Method(
		[3]json.Number,
		[]byte,
		map[sql.NullString]io.Reader,
		func(conn net.Conn),
		chan *httputil.BufferPool,
		struct{ URL *url.URL },
		interface{ CookieJar() http.CookieJar },
	)
}

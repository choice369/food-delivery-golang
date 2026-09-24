package skio

import (
	"net"
	"net/http"
	"net/url"

	"food_delivery/common"
)

type Conn interface {
	// ID returns session id
	ID() string
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header

	Context() interface{}
	SetContext(v interface{})
	Namespace() string
	Emit(msg string, v ...interface{})

	Join(room string)
	Leave(room string)
	LeaveAll()
	Rooms() []string
}

type AppSocket interface {
	Conn
	common.Requester
}

type appSocket struct {
	Conn
	common.Requester
}

func NewAppSocket(conn Conn, requester common.Requester) *appSocket {
	return &appSocket{conn, requester}
}

package websocket

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Context holds the information related to a WebSocket connection and
// the associated request context
type Context struct {
	Conn *websocket.Conn
	Url  *url.URL
	Vars map[string]string
}

// Socket Connection handler
type Handler interface {
	ServeWsConn(ctx *Context)
}

type Server struct {
	http.Handler
	upgrader websocket.Upgrader
	handler  Handler
}

// ServeHTTP http.Handler interface; handles incoming connections
// for accessing given websocket connections
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// upgrade connection to websocket
	c, err := s.upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer c.Close()
	if s.handler != nil {
		ctx := &Context{
			Conn: c,
			Url:  req.URL,
			// Makes webSocket handler contextually aware of the request params
			Vars: mux.Vars(req),
		}
		s.handler.ServeWsConn(ctx)
	}
}

func AllocateServer(handler Handler) *Server {
	return &Server{
		upgrader: websocket.Upgrader{
			// Accept all incoming requests irrespective of context
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		handler: handler,
	}
}

package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"net/http"
)

var stream = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebSocketStream(ws *websocket.Conn) *WebSocketStream {
	return &WebSocketStream{ws: ws}
}

type WebSocketStream struct {
	grpc.ServerStream
	ws *websocket.Conn
}

func (w *WebSocketStream) Send(rsp *rancher.Project) error {
	return w.ws.WriteJSON(rsp)
}

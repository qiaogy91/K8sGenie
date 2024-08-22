package api

import (
	"gitee.com/qiaogy91/K8sGenie/common"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
)

func (h *Handler) SyncProject(req *restful.Request, rsp *restful.Response) {
	conn, err := stream.Upgrade(rsp, req.Request, nil)
	if err != nil {
		common.L().Error().Msgf("upgrade websocket failed")
		return
	}

	if err := h.svc.SyncProject(nil, NewWebSocketStream(conn)); err != nil {
		common.L().Error().Msgf("sync project failed, %v", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, []byte("资源同步结束")); err != nil {
		common.L().Error().Msgf("ws write EOF Message failed: %v", err)
	}
	if err := conn.Close(); err != nil {
		common.L().Error().Msgf("ws close conn failed: %v", err)
	}
}

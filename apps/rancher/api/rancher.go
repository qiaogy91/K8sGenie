package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/common"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"net/http"

	"strconv"
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

func (h *Handler) QueryProject(req *restful.Request, rsp *restful.Response) {
	tp, err := strconv.ParseInt(req.QueryParameter("type"), 10, 32)
	if err != nil {
		common.SendFailed(rsp, http.StatusBadRequest, err)
		return
	}
	ins := &rancher.QueryProjectReq{
		QueryType: rancher.QUERY_TYPE(tp),
		KeyWord:   req.QueryParameter("keyword"),
	}

	out, err := h.svc.QueryProject(req.Request.Context(), ins)
	if err != nil {
		common.SendFailed(rsp, http.StatusBadRequest, err)
		return
	}

	common.SendSuccess(rsp, out)
}

package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/common"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

func (h *Handler) SyncK8SWorkload(req *restful.Request, rsp *restful.Response) {
	conn, err := stream.Upgrade(rsp, req.Request, nil)
	if err != nil {
		common.L().Error().Msgf("upgrade websocket failed")
		return
	}

	if err := h.svc.SyncK8SWorkload(nil, NewWebSocketStream(conn)); err != nil {
		common.L().Error().Msgf("sync workload failed, %v", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, []byte("资源同步结束")); err != nil {
		common.L().Error().Msgf("ws write EOF Message failed: %v", err)
	}
	if err := conn.Close(); err != nil {
		common.L().Error().Msgf("ws close conn failed: %v", err)
	}
}

func (h *Handler) QueryK8SWorkload(req *restful.Request, rsp *restful.Response) {
	tp, err := strconv.ParseInt(req.QueryParameter("type"), 10, 64)
	if err != nil {
		common.SendFailed(rsp, http.StatusBadRequest, err)
		return
	}

	inst := &k8s.QueryK8SWorkloadReq{Type: k8s.SEARCH_TYPE(tp), Keyword: req.QueryParameter("keyword")}
	outs, err := h.svc.QueryK8SWorkload(req.Request.Context(), inst)
	if err != nil {
		common.SendFailed(rsp, http.StatusBadRequest, err)
		return
	}

	common.SendSuccess(rsp, outs)
}

package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/alert"
	"gitee.com/qiaogy91/K8sGenie/common"
	"github.com/emicklei/go-restful/v3"
	"net/http"
)

func (h *Handler) HandlerAlert(req *restful.Request, rsp *restful.Response) {

	ins := &alert.HandlerAlertReq{}
	if err := req.ReadEntity(ins); err != nil {
		common.SendFailed(rsp, http.StatusBadRequest, err)
		return
	}
	// todo 这里要处理，不论什么情况都要返回 200 ，避免让 alertManager 收到错误
	_, _ = h.svc.HandlerAlert(req.Request.Context(), ins)
	common.SendSuccess(rsp, "ok")
}

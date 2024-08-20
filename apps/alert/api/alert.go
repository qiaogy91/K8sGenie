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

	res, err := h.svc.HandlerAlert(req.Request.Context(), ins)
	if err != nil {
		common.SendFailed(rsp, http.StatusBadRequest, err)
		return
	}

	common.SendSuccess(rsp, res)
}

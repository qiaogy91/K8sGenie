package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/router"
	"gitee.com/qiaogy91/K8sGenie/common"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"strconv"
)

func (h *Handler) CreateRoute(req *restful.Request, rsp *restful.Response) {
	ins := &router.Spec{}
	if err := req.ReadEntity(ins); err != nil {
		common.SendFailed(rsp, http.StatusBadRequest, err)
		return
	}

	out, err := h.svc.CreateRoute(req.Request.Context(), ins)
	if err != nil {
		common.SendFailed(rsp, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(rsp, out)
}

func (h *Handler) QueryRoute(req *restful.Request, rsp *restful.Response) {
	t, err := strconv.Atoi(req.QueryParameter("type"))
	if err != nil {
		common.SendFailed(rsp, http.StatusInternalServerError, err)
		return
	}

	ins := &router.QueryRouteReq{Type: router.QUERY_TYPE(t), Keyword: req.QueryParameter("keyword")}

	out, err := h.svc.QueryRoute(req.Request.Context(), ins)
	if err != nil {
		common.SendFailed(rsp, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(rsp, out)
}

func (h *Handler) DeleteRoute(req *restful.Request, rsp *restful.Response) {
	ins := &router.DeleteRouteReq{Id: req.PathParameter("id")}
	out, err := h.svc.DeleteRoute(req.Request.Context(), ins)
	if err != nil {
		common.SendFailed(rsp, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(rsp, out)
}

func (h *Handler) UpdateRoute(r *restful.Request, w *restful.Response) {
	spec := &router.Spec{}
	if err := r.ReadEntity(spec); err != nil {
		common.SendFailed(w, http.StatusBadRequest, err)
		return
	}
	req := &router.UpdateRouteReq{
		Id:   r.PathParameter("id"),
		Spec: spec,
	}

	ins, err := h.svc.UpdateRoute(r.Request.Context(), req)
	if err != nil {
		common.SendFailed(w, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(w, ins)
}

func (h *Handler) UrgentChange(r *restful.Request, w *restful.Response) {
	req := &router.UrgentChangeReq{}
	if err := r.ReadEntity(req); err != nil {
		common.SendFailed(w, http.StatusBadRequest, err)
		return
	}

	ins, err := h.svc.UrgentChange(r.Request.Context(), req)
	if err != nil {
		common.SendFailed(w, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(w, ins)
}

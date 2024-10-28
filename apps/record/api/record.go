package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/record"
	"gitee.com/qiaogy91/K8sGenie/common"
	"github.com/emicklei/go-restful/v3"
	"net/http"
)

/*
------------------------- CreateRecordService ------------------------------------------------------
*/

func (h *Handler) CreateTable(r *restful.Request, w *restful.Response) {
	if err := h.svc.CreateTable(r.Request.Context()); err != nil {
		common.SendFailed(w, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(w, "table create success")
}
func (h *Handler) CreateNamespaceRecord(r *restful.Request, w *restful.Response) {
	req := &record.CreateNamespaceRecordRequest{}
	if err := h.svc.CreateNamespaceRecord(r.Request.Context(), req); err != nil {
		common.SendFailed(w, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(w, "namespace record create success")
}
func (h *Handler) CreateProjectRecord(r *restful.Request, w *restful.Response) {
	req := &record.CreateProjectRecordRequest{}

	if err := r.ReadEntity(req); err != nil {
		common.SendFailed(w, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.CreateProjectRecord(r.Request.Context(), req); err != nil {
		common.SendFailed(w, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(w, "project record create success")
}
func (h *Handler) CreateLineRecord(r *restful.Request, w *restful.Response) {
	req := &record.CreateLineRecordRequest{}

	if err := r.ReadEntity(req); err != nil {
		common.SendFailed(w, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.CreateLineRecord(r.Request.Context(), req); err != nil {
		common.SendFailed(w, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(w, "cluster record create success")
}

/*
------------------------- QueryRecordService ------------------------------------------------------
*/

func (h *Handler) QueryNamespaceRecord(r *restful.Request, w *restful.Response) {
	req := &record.QueryNamespaceRecordRequest{
		CreatedAt:   r.QueryParameter("createdAt"),
		ProjectCode: r.QueryParameter("projectCode"),
		ProjectLine: r.QueryParameter("projectLine"),
		ClusterName: r.QueryParameter("clusterName"),
	}

	ins, err := h.svc.QueryNamespaceRecord(r.Request.Context(), req)
	if err != nil {
		common.SendFailed(w, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(w, ins)
}
func (h *Handler) QueryProjectRecord(r *restful.Request, w *restful.Response) {
	req := &record.QueryProjectRecordRequest{
		Month:       r.QueryParameter("month"),
		ProjectLine: r.QueryParameter("projectLine"),
		ClusterName: r.QueryParameter("clusterName"),
	}

	ins, err := h.svc.QueryProjectRecord(r.Request.Context(), req)
	if err != nil {
		common.SendFailed(w, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(w, ins)
}
func (h *Handler) QueryLineRecord(r *restful.Request, w *restful.Response) {
	req := &record.QueryLineRecordRequest{
		Month:       r.QueryParameter("month"),
		ClusterName: r.QueryParameter("clusterName"),
	}

	ins, err := h.svc.QueryLineRecord(r.Request.Context(), req)
	if err != nil {
		common.SendFailed(w, http.StatusInternalServerError, err)
		return
	}
	common.SendSuccess(w, ins)
}

package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"github.com/emicklei/go-restful/v3"
	"path"
)

type Handler struct {
	svc rancher.Service
}

func (h *Handler) Name() string {
	return rancher.AppName
}

func (h *Handler) Init() error {
	h.svc = ioc.GetController(rancher.AppName).(rancher.Service)
	return nil
}

func (h *Handler) RegistryRoute(pro string) {
	ws := new(restful.WebService)
	ws.Path(path.Join("/api/v1", pro, rancher.AppName)).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.GET("sync").To(h.SyncProject).
		Metadata("auth", false).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", rancher.AppName).
		Metadata("action", "create").
		Doc("同步项目信息"))

	ws.Route(ws.GET("").To(h.QueryProject).
		Metadata("auth", true).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", rancher.AppName).
		Metadata("action", "list").
		Doc("获取项目列表"))

	restful.Add(ws)
}

func init() {
	ioc.RegistryHandlerObject(&Handler{})
}

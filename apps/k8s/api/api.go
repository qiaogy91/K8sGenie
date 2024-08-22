package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"github.com/emicklei/go-restful/v3"
	"path"
)

type Handler struct {
	svc k8s.Service
}

func (h *Handler) Name() string {
	return k8s.AppName
}

func (h *Handler) Init() error {
	h.svc = ioc.GetController(k8s.AppName).(k8s.Service)
	return nil
}

func (h *Handler) RegistryRoute(p string) {
	ws := new(restful.WebService)
	ws.Path(path.Join("/api/v1", p, k8s.AppName)).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.GET("sync").To(h.SyncK8SWorkload).
		Metadata("auth", true).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", k8s.AppName).
		Metadata("action", "sync").
		Doc("同步工作负载"))

	restful.Add(ws)

}

func init() {
	ioc.RegistryHandlerObject(&Handler{})
}

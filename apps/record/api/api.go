package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/record"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"github.com/emicklei/go-restful/v3"
	"path"
)

type Handler struct {
	svc record.Service
}

func (h *Handler) Name() string {
	return record.AppName
}

func (h *Handler) Init() error {
	h.svc = ioc.GetController(record.AppName).(record.Service)
	return nil
}

func (h *Handler) RegistryRoute(p string) {
	ws := new(restful.WebService)
	ws.Path(path.Join("/api/v1", p, record.AppName)).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.POST("namespace").To(h.CreateNamespaceRecord).
		Metadata("auth", true).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", k8s.AppName).
		Metadata("action", "sync").
		Doc("同步名称空间级用量"))

	ws.Route(ws.POST("project").To(h.CreateProjectRecord).
		Metadata("auth", true).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", k8s.AppName).
		Metadata("action", "sync").
		Doc("同步集群项目级用量"))

	ws.Route(ws.POST("cluster").To(h.CreateLineRecord).
		Metadata("auth", true).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", k8s.AppName).
		Metadata("action", "sync").
		Doc("同步集群产线级用量"))

	ws.Route(ws.GET("namespace").To(h.QueryNamespaceRecord).
		Metadata("auth", false).
		Metadata("perm", false).
		Metadata("audit", false).
		Metadata("resource", k8s.AppName).
		Metadata("action", "sync").
		Doc("获取名称空间级用量"))

	ws.Route(ws.GET("project").To(h.QueryProjectRecord).
		Metadata("auth", false).
		Metadata("perm", false).
		Metadata("audit", false).
		Metadata("resource", k8s.AppName).
		Metadata("action", "sync").
		Doc("获取集群项目级用量"))

	ws.Route(ws.GET("cluster").To(h.QueryLineRecord).
		Metadata("auth", false).
		Metadata("perm", false).
		Metadata("audit", false).
		Metadata("resource", k8s.AppName).
		Metadata("action", "sync").
		Doc("获取集群产线级用量"))

	restful.Add(ws)

}

func init() {
	ioc.RegistryHandlerObject(&Handler{})
}

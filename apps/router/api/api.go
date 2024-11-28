package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/router"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"github.com/emicklei/go-restful/v3"
	"path"
)

type Handler struct {
	svc router.Service
}

func (h *Handler) Name() string {
	return router.AppName
}

func (h *Handler) Init() error {
	h.svc = ioc.GetController(router.AppName).(router.Service)
	return nil
}

func (h *Handler) RegistryRoute(p string) {
	ws := new(restful.WebService)
	ws.Path(path.Join("/api/v1", p, router.AppName)).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.POST("").To(h.CreateRoute).
		Metadata("auth", true).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", router.AppName).
		Metadata("action", "create").
		Doc("创建告警路由"))

	ws.Route(ws.GET("").To(h.QueryRoute).
		Param(ws.QueryParameter("type", "query type")).
		Param(ws.QueryParameter("keyword", "keyword")).
		Metadata("auth", true).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", router.AppName).
		Metadata("action", "list").
		Doc("查询告警路由"))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteRoute).
		Metadata("auth", true).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", router.AppName).
		Metadata("action", "delete").
		Doc("删除告警路由"))

	ws.Route(ws.PUT("/{id}").To(h.UpdateRoute).
		Metadata("auth", true).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", router.AppName).
		Metadata("action", "update").
		Doc("更新告警路由"))

	ws.Route(ws.POST("/UrgentChange").To(h.UrgentChange).
		Metadata("auth", true).
		Metadata("perm", true).
		Metadata("audit", true).
		Metadata("resource", router.AppName).
		Metadata("action", "update").
		Doc("转换加急状态"))

	restful.Add(ws)
}

func init() {
	ioc.RegistryHandlerObject(&Handler{})
}

package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/alert"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"github.com/emicklei/go-restful/v3"
	"path"
)

type Handler struct {
	svc alert.Service
}

func (h *Handler) Name() string {
	return alert.AppName
}

func (h *Handler) Init() error {
	h.svc = ioc.GetController(alert.AppName).(alert.Service)
	return nil
}

func (h *Handler) RegistryRoute(proName string) {
	ws := new(restful.WebService)

	ws.Path(path.Join("/api/v1", proName, alert.AppName))
	ws.Route(ws.POST("").To(h.HandlerAlert).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON).
		Metadata("auth", false).
		Metadata("perm", false).
		Metadata("audit", false).
		Metadata("resource", alert.AppName).
		Metadata("action", "create").
		Doc("告警处理"))

	restful.Add(ws)

}

func init() {
	ioc.RegistryHandlerObject(&Handler{})
}

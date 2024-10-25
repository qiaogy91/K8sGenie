package impl

import (
	"context"
	"errors"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/common"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *Impl) GetDeploymentSummary(ctx context.Context, req *k8s.GetDeploymentSummaryReq) error {
	c, ok := i.cs[req.ClusterName]
	if !ok {
		return errors.New("no such cluster, " + req.ClusterName)
	}

	// 3. 如果存在，则遍历该名称空间下所有dep，进行统计计算
	deps, err := c.AppsV1().Deployments(req.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		common.L().Info().Msgf("handlerDeployment list ns/%s deployment failed %v", req.Namespace, err)
		return err
	}

	// 3. 插入数据 deployment
	for _, dep := range deps.Items {
		var cpuTotal int64
		var ramTotal int64

		for _, c := range dep.Spec.Template.Spec.Containers {
			cpuTotal += c.Resources.Limits.Cpu().Value()
			ramTotal += c.Resources.Limits.Memory().Value()
		}
		ins := &k8s.WorkLoad{
			Spec: &k8s.Spec{
				Type:       k8s.Type_TYPE_DEPLOYMENT,
				CreateTime: dep.GetCreationTimestamp().Format("2006/01/02"),
				Namespace:  dep.Namespace,
				Name:       dep.Name,
				Replicas:   *dep.Spec.Replicas,
				Ram:        ramTotal / 1024 / 1024,
				Cpu:        cpuTotal,
			},
		}

		common.L().Info().Msgf("recv %+v", ins)
	}
	return nil
}

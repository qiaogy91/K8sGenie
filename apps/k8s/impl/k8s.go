package impl

import (
	"context"
	"errors"
	"gitee.com/qiaogy91/InfraGenie/apps/resourcer"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/common"
	cv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

func (i *Impl) CreateTable(ctx context.Context, empty *k8s.Empty) (*k8s.Empty, error) {
	if err := i.db.WithContext(ctx).AutoMigrate(&k8s.WorkLoad{}); err != nil {
		return nil, err
	}
	return nil, nil
}

// SyncK8SWorkload 将Kubernetes 数据同步到DB
// 1. 内部调用 handler 				数据同步入口，来负责上下文切换
// 2. 内部调用 handlerDeployment   	处理dep， 而后记录进数据库
// 3. 内部调用 handlerDaemonSet 		处理ds，  而后记录进数据库
// 4. 内部调用 handlerStatefulSet  	处理sts， 而后记录进数据库
func (i *Impl) SyncK8SWorkload(empty *k8s.Empty, stream k8s.Rpc_SyncK8SWorkloadServer) error {
	// 数据清洗
	ctx := context.Background()
	sTime := time.Now().Truncate(24 * time.Hour)
	eTime := sTime.Add(24 * time.Hour)
	if err := i.db.WithContext(ctx).Where("created_at >= ? AND created_at < ?", sTime.Unix(), eTime.Unix()).Delete(&k8s.WorkLoad{}).Error; err != nil {
		return err
	}

	for cluster, client := range i.cs {
		common.L().Info().Msgf("=============== %s =======================", cluster)
		if err := i.handler(ctx, client, stream); err != nil {
			return err
		}
	}
	return nil
}

func (i *Impl) handler(ctx context.Context, c *kubernetes.Clientset, stream k8s.Rpc_SyncK8SWorkloadServer) error {
	// 获取所有namespace
	nsSet, err := c.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		return err
	}

	for _, ns := range nsSet.Items {
		// 1. 获取名称空间的注解
		pid, ok := ns.Annotations["field.cattle.io/projectId"]
		if !ok {
			common.L().Info().Msgf("K8SResourceSync namespace has no annotation, %s", ns.Name)
			continue
		}

		// 2. 根据注解信息查找项目
		pro, err := i.DescRancherProject(ctx, &k8s.DescRancherProjectReq{ProjectID: pid, DescType: k8s.FromProjectID})
		if err != nil {
			common.L().Info().Msgf("K8SResourceSync namespace has no projectId, %s , %v", ns.Name, err)
			continue
		}

		// - 处理 deployment
		if err := i.handlerDeployment(ctx, c, pro, ns, stream); err != nil {
			return err
		}

		// - 处理 daemonSet
		if err := i.handlerDaemonSet(ctx, c, pro, ns, stream); err != nil {
			return err
		}

		// - 处理 statefulSet
		if err := i.handlerStatefulSet(ctx, c, pro, ns, stream); err != nil {
			return err
		}
	}
	return nil
}

func (i *Impl) handlerDeployment(ctx context.Context, c *kubernetes.Clientset, pro *resourcer.Project, ns cv1.Namespace, stream k8s.Rpc_SyncK8SWorkloadServer) error {
	// 3. 如果存在，则遍历该名称空间下所有dep，进行统计计算
	deps, err := c.AppsV1().Deployments(ns.Name).List(ctx, v1.ListOptions{})
	if err != nil {
		common.L().Info().Msgf("handlerDeployment list deployment failed %v", err)
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
				Type:      k8s.Type_TYPE_DEPLOYMENT,
				Namespace: dep.Namespace,
				Name:      dep.Name,
				Replicas:  *dep.Spec.Replicas,
				Ram:       ramTotal / 1024 / 1024 / 1024,
				Cpu:       cpuTotal,
				ProjectId: pro.Spec.ProjectId,
			},
		}

		if err := i.db.WithContext(ctx).Model(&k8s.WorkLoad{}).Create(ins).Error; err != nil {
			return err
		}

		if err := stream.Send(ins); err != nil {
			return err
		}
	}
	return nil
}

func (i *Impl) handlerDaemonSet(ctx context.Context, c *kubernetes.Clientset, pro *resourcer.Project, ns cv1.Namespace, stream k8s.Rpc_SyncK8SWorkloadServer) error {
	// 获取节点数量
	nodes, err := c.CoreV1().Nodes().List(ctx, v1.ListOptions{})
	if err != nil {
		common.L().Info().Msgf("handlerDaemonSet list nodes failed %v", err)
		return err
	}

	// 3. 如果存在，则遍历该名称空间下所有dep，进行统计计算
	ds, err := c.AppsV1().DaemonSets(ns.Name).List(ctx, v1.ListOptions{})
	if err != nil {
		common.L().Info().Msgf("handlerDaemonSet list ds failed %v", err)
		return err
	}

	// 3. 插入数据 deployment
	for _, d := range ds.Items {
		var cpuTotal int64
		var ramTotal int64
		for _, c := range d.Spec.Template.Spec.Containers {
			cpuTotal += c.Resources.Limits.Cpu().Value()
			ramTotal += c.Resources.Limits.Memory().Value()
		}
		ins := &k8s.WorkLoad{
			Spec: &k8s.Spec{
				Type:      k8s.Type_TYPE_DAEMON_SET,
				Namespace: d.Namespace,
				Name:      d.Name,
				Replicas:  int32(len(nodes.Items)),
				Ram:       ramTotal / 1024 / 1024 / 1024,
				Cpu:       cpuTotal,
				ProjectId: pro.Spec.ProjectId,
			},
		}

		if err := i.db.WithContext(ctx).Model(&k8s.WorkLoad{}).Create(ins).Error; err != nil {
			return err
		}

		if err := stream.Send(ins); err != nil {
			return err
		}
	}
	return nil
}

func (i *Impl) handlerStatefulSet(ctx context.Context, c *kubernetes.Clientset, pro *resourcer.Project, ns cv1.Namespace, stream k8s.Rpc_SyncK8SWorkloadServer) error {
	// 1. 如果存在，则遍历该名称空间下所有sts，进行统计计算
	sts, err := c.AppsV1().StatefulSets(ns.Name).List(ctx, v1.ListOptions{})
	if err != nil {
		common.L().Info().Msgf("handlerDaemonSet list ds failed %v", err)
		return err
	}

	// 2. 插入数据 deployment
	for _, st := range sts.Items {
		var cpuTotal int64
		var ramTotal int64
		for _, c := range st.Spec.Template.Spec.Containers {
			cpuTotal += c.Resources.Limits.Cpu().Value()
			ramTotal += c.Resources.Limits.Memory().Value()
		}
		ins := &k8s.WorkLoad{
			Spec: &k8s.Spec{
				Type:      k8s.Type_TYPE_STATEFUL_SET,
				Namespace: st.Namespace,
				Name:      st.Name,
				Replicas:  *st.Spec.Replicas,
				Ram:       ramTotal / 1024 / 1024 / 1024,
				Cpu:       cpuTotal,
				ProjectId: pro.Spec.ProjectId,
			},
		}

		if err := i.db.WithContext(ctx).Model(&k8s.WorkLoad{}).Create(ins).Error; err != nil {
			return err
		}

		if err := stream.Send(ins); err != nil {
			return err
		}
	}
	return nil
}

// DescRancherProject 根据名称空间，反查项目
func (i *Impl) DescRancherProject(ctx context.Context, req *k8s.DescRancherProjectReq) (*resourcer.Project, error) {
	switch req.DescType {
	case k8s.FromProjectID:
		ins := &resourcer.Project{}
		if err := i.db.WithContext(ctx).Model(&resourcer.Project{}).Where("project_id = ?", req.ProjectID).First(ins).Error; err != nil {
			return nil, err
		}
		return ins, nil

	case k8s.FromClusterAndNamespace:
		// 获取集群客户端
		c, ok := i.cs[req.ClusterName]
		if !ok {
			return nil, errors.New("no such cluster")
		}

		// 找到名称空间对象
		ns, err := c.CoreV1().Namespaces().Get(ctx, req.Namespace, v1.GetOptions{})
		if err != nil {
			return nil, err
		}
		// 获取注解
		pid, ok := ns.Annotations["field.cattle.io/projectId"]
		if !ok {
			return nil, errors.New("not belong to any project")
		}

		// 查找Project
		ins := &resourcer.Project{}
		if err := i.db.WithContext(ctx).Model(&resourcer.Project{}).Where("project_id = ?", pid).First(ins).Error; err != nil {
			return nil, err
		}
		return ins, nil
	}

	return nil, errors.New("search type err")

}

package impl

import (
	"context"
	"gitee.com/qiaogy91/InfraGenie/apps/resourcer"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/common"
	cv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *Impl) TableCreate(ctx context.Context, empty *k8s.Empty) (*k8s.Empty, error) {
	if err := i.db.WithContext(ctx).AutoMigrate(&resourcer.Rancher{}, &k8s.WorkLoad{}); err != nil {
		return nil, err
	}
	return nil, nil
}

// RancherResourceSync 同步所有项目信息到DB
func (i *Impl) RancherResourceSync(ctx context.Context) error {
	// 每次都是重新获取数据，确保数据与当前完全一致
	i.db.Exec("TRUNCATE TABLE ranchers")
	i.db.Exec("ALTER TABLE ranchers AUTO_INCREMENT = 1")

	stream, err := i.rc.RancherResourceSync(ctx, nil)
	if err != nil {
		return err
	}

	for {
		rsp, err := stream.Recv()
		if err != nil {
			common.L().Error().Msgf("stream recv err: %v", err)
			break
		}
		if err := i.db.WithContext(ctx).Create(rsp).Error; err != nil {
			return err
		}
	}
	common.L().Info().Msgf("recv finished")
	return nil
}

// DescRancherProject 项目查询
func (i *Impl) DescRancherProject(ctx context.Context, pid string) (*resourcer.Rancher, error) {
	ins := &resourcer.Rancher{}
	if err := i.db.WithContext(ctx).Model(&resourcer.Rancher{}).Where("project_id = ?", pid).First(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// K8SResourceSync 同步所有容器信息到DB
func (i *Impl) K8SResourceSync(empty *k8s.Empty, stream k8s.Rpc_K8SResourceSyncServer) error {
	// 获取所有namespace
	ctx := context.Background()
	nsSet, err := i.kc.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
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
		pro, err := i.DescRancherProject(ctx, pid)
		if err != nil {
			common.L().Info().Msgf("K8SResourceSync namespace has no projectId, %s , %v", ns.Name, err)
			continue
		}

		// - 处理deployment
		if err := i.handlerDeployment(ctx, pro, ns, stream); err != nil {
			return err
		}

		// 4. 插入数据 daemonSet
		if err := i.handlerDaemonSet(ctx, pro, ns, stream); err != nil {
			return err
		}
	}
	return nil
}

func (i *Impl) handlerDeployment(ctx context.Context, pro *resourcer.Rancher, ns cv1.Namespace, stream k8s.Rpc_K8SResourceSyncServer) error {
	// 3. 如果存在，则遍历该名称空间下所有dep，进行统计计算
	deps, err := i.kc.AppsV1().Deployments(ns.Name).List(ctx, v1.ListOptions{})
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

		if err := stream.Send(ins); err != nil {
			return err
		}
	}
	return nil
}

func (i *Impl) handlerDaemonSet(ctx context.Context, pro *resourcer.Rancher, ns cv1.Namespace, stream k8s.Rpc_K8SResourceSyncServer) error {
	// 获取节点数量
	nodes, err := i.kc.CoreV1().Nodes().List(ctx, v1.ListOptions{})
	if err != nil {
		common.L().Info().Msgf("handlerDaemonSet list nodes failed %v", err)
		return err
	}

	// 3. 如果存在，则遍历该名称空间下所有dep，进行统计计算
	ds, err := i.kc.AppsV1().DaemonSets(ns.Name).List(ctx, v1.ListOptions{})
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

		if err := stream.Send(ins); err != nil {
			return err
		}
	}
	return nil
}

// todo handler for sts, handler for job cronjob

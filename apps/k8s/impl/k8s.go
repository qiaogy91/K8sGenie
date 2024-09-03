package impl

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/common"
	"github.com/go-playground/validator"
	cv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sort"
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
	// 每次都是重新获取数据，确保数据与当前完全一致
	i.db.Exec("TRUNCATE TABLE work_loads")
	i.db.Exec("ALTER TABLE work_loads AUTO_INCREMENT = 1")

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
		annotation, ok := ns.Annotations["field.cattle.io/projectId"]
		if !ok {
			common.L().Info().Msgf("K8SResourceSync namespace has no annotation, %s", ns.Name)
			continue
		}

		// 2. 根据注解信息查找项目
		pro, err := i.rc.DescProject(ctx, &rancher.DescProjectReq{DescType: rancher.DESC_TYPE_DESC_TYPE_PROJECT_ID, KeyWord: annotation})
		if err != nil {
			common.L().Info().Msgf("no such project, %s , %v", ns.Name, err)
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

func (i *Impl) handlerDeployment(ctx context.Context, c *kubernetes.Clientset, pro *rancher.Project, ns cv1.Namespace, stream k8s.Rpc_SyncK8SWorkloadServer) error {
	// 3. 如果存在，则遍历该名称空间下所有dep，进行统计计算
	deps, err := c.AppsV1().Deployments(ns.Name).List(ctx, v1.ListOptions{})
	if err != nil {
		common.L().Info().Msgf("handlerDeployment list ns/%s deployment failed %v", ns.Name, err)
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

func (i *Impl) handlerDaemonSet(ctx context.Context, c *kubernetes.Clientset, pro *rancher.Project, ns cv1.Namespace, stream k8s.Rpc_SyncK8SWorkloadServer) error {
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

func (i *Impl) handlerStatefulSet(ctx context.Context, c *kubernetes.Clientset, pro *rancher.Project, ns cv1.Namespace, stream k8s.Rpc_SyncK8SWorkloadServer) error {
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

func (i *Impl) QueryK8SWorkload(ctx context.Context, req *k8s.QueryK8SWorkloadReq) (*k8s.WorkLoadSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	sql := i.db.WithContext(ctx).Model(&k8s.WorkLoad{})

	switch req.Type {
	case k8s.SEARCH_TYPE_SEARCH_TYPE_ALL:
		sql = sql.Where("1 = 1")
	case k8s.SEARCH_TYPE_SEARCH_TYPE_PROJECT_CODE:
		pro, err := i.rc.DescProject(ctx, &rancher.DescProjectReq{DescType: rancher.DESC_TYPE_DESC_TYPE_PROJECT_CODE, KeyWord: req.Keyword})
		if err != nil {
			return nil, err
		}
		sql = sql.Where("project_id = ?", pro.Spec.ProjectId)
	case k8s.SEARCH_TYPE_SEARCH_TYPE_NAMESPACE:
		sql = sql.Where("namespace = ?", req.Keyword)
	case k8s.SEARCH_TYPE_SEARCH_TYPE_WORKLOAD_NAME:
		sql = sql.Where("name like ?", "%"+req.Keyword+"%")
	}

	ins := &k8s.WorkLoadSet{}
	if err := sql.Find(&ins.Items).Error; err != nil {
		return nil, err
	}

	ins.Total = int64(len(ins.Items))

	return ins, nil
}

// DescNamespace 根据名称空间，反查项目
func (i *Impl) DescNamespace(ctx context.Context, req *k8s.DescNamespaceReq) (*rancher.Project, error) {
	// 获取集群客户端
	c, ok := i.cs[req.ClusterName]
	if !ok {
		return nil, errors.New("no such cluster")
	}

	// 找到名称空间对象
	ns, err := c.CoreV1().Namespaces().Get(ctx, req.NamespaceName, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 获取注解
	pid, ok := ns.Annotations["field.cattle.io/projectId"]
	if !ok {
		return nil, errors.New("not belong to any project")
	}

	// 查找Project
	ins := &rancher.Project{}
	if err := i.db.WithContext(ctx).Model(&rancher.Project{}).Where("project_id = ?", pid).First(&ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// GetPodRamUsage 获取资源利用率前10的pod
func (i *Impl) GetPodRamUsage(ctx context.Context, req *k8s.GetPodRamUsageReq) (*k8s.GetPodRamUsageRsp, error) {
	mc, ok := i.ms[req.ClusterName]
	cc, ok := i.cs[req.ClusterName]

	if !ok {
		return nil, fmt.Errorf("cluster name err: %s", req.ClusterName)
	}

	// 获取节点 10.0.0.100 上的所有 Pod
	pods, err := cc.CoreV1().Pods("").List(ctx, v1.ListOptions{FieldSelector: fmt.Sprintf("spec.nodeName=%s", req.NodeName)})
	if err != nil {
		return nil, err
	}

	// 获取这些 Pod 的 metrics
	ins := &k8s.GetPodRamUsageRsp{}

	for _, pod := range pods.Items {

		podMetrics, err := mc.MetricsV1beta1().PodMetricses(pod.Namespace).Get(ctx, pod.Name, v1.GetOptions{})
		if err != nil {
			common.L().Error().Msgf("Error getting metrics for pod %s: %v", pod.Name, err)
			continue
		}

		for _, c := range podMetrics.Containers {
			ins.Items = append(ins.Items, &k8s.PodMetricDetail{
				PodName:       pod.Name,
				NamespaceName: pod.Namespace,
				CpuCores:      c.Usage.Cpu().MilliValue(),
				RamMbi:        c.Usage.Memory().Value() / 1024 / 1024,
			})
			ins.Count += 1
		}
	}

	sort.Sort(ins)
	return ins, nil
}

func (i *Impl) KillTop1Pod(ctx context.Context, req *k8s.KillTop1PodReq) (*k8s.KillTop1PodRsp, error) {
	cc, ok := i.cs[req.ClusterName]
	if !ok {
		return nil, fmt.Errorf("cluster name err: %s", req.ClusterName)
	}

	if err := cc.CoreV1().Pods(req.NamespaceName).Delete(ctx, req.PodName, v1.DeleteOptions{}); err != nil {
		return nil, err
	}
	return nil, nil
}

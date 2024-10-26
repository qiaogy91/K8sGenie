package impl

import (
	"context"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/apps/record"
	"time"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&record.NamespaceRecord{}, &record.ProjectRecord{}, &record.LineRecord{})
}

func (i *Impl) GetCluster(ctx context.Context) (*record.ClustersSet, error) {
	sql := i.db.WithContext(ctx)

	// 查询所有集群
	inst := &record.ClustersSet{}
	if err := sql.Model(&rancher.Project{}).Select("cluster_name").Distinct().Find(&inst).Error; err != nil {
		return nil, err
	}
	return inst, nil
}

func (i *Impl) GetProject(ctx context.Context, req *record.GetProjectReq) (*rancher.ProjectSet, error) {
	sql := i.db.WithContext(ctx)

	inst := &rancher.ProjectSet{}
	if err := sql.Model(&rancher.Project{}).
		Where("cluster_name = ? AND project_line != '' ", req.ClusterName).Find(&inst.Items).Error; err != nil {
		return nil, err
	}
	return inst, nil
}

func (i *Impl) HavingNamespace(ctx context.Context, req *rancher.Project) error {
	sql := i.db.WithContext(ctx)

	if req.Spec.ProjectLine == "" {
		return nil
	}

	// 聚合当前项目下各名称空间的资源使用情况
	ns := &record.NamespaceWightSet{}
	if err := sql.Model(&k8s.WorkLoad{}).Select("namespace, SUM(ram * replicas) as weight").Where("project_id = ?", req.Spec.ProjectId).Group("namespace").Find(&ns).Error; err != nil {
		return err
	}

	for _, nw := range *ns {
		r := &record.NamespaceRecord{
			ClusterName:    req.Spec.ClusterName,
			ProjectLine:    req.Spec.ProjectLine,
			ProjectDesc:    req.Spec.ProjectDesc,
			ProjectCode:    req.Spec.ProjectCode,
			NamespaceWight: nw,
		}

		fmt.Printf("%s/%s/%s/%s/  %s %d\n",
			req.Spec.ClusterName,
			req.Spec.ProjectLine,
			req.Spec.ProjectCode,
			req.Spec.ProjectDesc,
			nw.Namespace,
			nw.Weight,
		)
		if err := sql.Create(r).Error; err != nil {
			return err
		}
	}
	return nil
}

// CreateNamespaceRecord 根据当前（当天）Project、Workload 情况，绘制Record 数据并入库
func (i *Impl) CreateNamespaceRecord(ctx context.Context, req *record.CreateNamespaceRecordRequest) error {
	// 清理掉当天的数据
	start := time.Now().Truncate(24 * time.Hour).Unix()
	end := start + 24*60*60
	if err := i.db.WithContext(ctx).Where("created_at >= ? AND created_at < ?", start, end).Delete(&record.NamespaceRecord{}).Error; err != nil {
		return err
	}

	// 获取所有集群
	cls, err := i.GetCluster(ctx)
	if err != nil {
		return err
	}

	// 遍历这个集群下所有的项目
	for _, cluster := range *cls {
		pros, err := i.GetProject(ctx, &record.GetProjectReq{ClusterName: cluster.ClusterName})
		if err != nil {
			return err
		}

		// 聚合每个项目下的名称空间资源使用情况
		for _, pro := range pros.Items {
			if err := i.HavingNamespace(ctx, pro); err != nil {
				return err
			}
		}
	}
	return nil
}

// CreateProjectRecord 汇总项目用量分布
func (i *Impl) CreateProjectRecord(ctx context.Context, req *record.CreateProjectRecordRequest) error {
	sql := i.db.WithContext(ctx).Model(&record.NamespaceRecord{})
	// 时间解析
	start, end, month, err := req.TimeRage()
	if err != nil {
		return err
	}

	// 删除当月数据
	if err := i.db.WithContext(ctx).Where("month = ?", month).Delete(&record.ProjectRecord{}).Error; err != nil {
		return err
	}

	// 过滤出数据
	nsSet := &record.ProjectRecordSet{}
	ql := fmt.Sprintf("cluster_name, project_line, project_desc, project_code, SUM(weight) as weight, '%s' as month", month)

	if err := sql.Debug().
		Select(ql).
		Where("created_at >= ? AND created_at <=?", start, end).
		Group("cluster_name, project_line, project_desc, project_code").
		Find(nsSet).
		Error; err != nil {
		return err
	}

	// 入库
	if err := i.db.WithContext(ctx).Model(&record.ProjectRecord{}).Create(nsSet).Error; err != nil {
		return err
	}
	return nil
}

// CreateLineRecord 汇总产线用量情况
func (i *Impl) CreateLineRecord(ctx context.Context, req *record.CreateLineRecordRequest) error {
	sql := i.db.WithContext(ctx).Model(&record.NamespaceRecord{})
	// 时间解析
	start, end, month, err := req.TimeRage()
	if err != nil {
		return err
	}

	// 删除当月数据
	if err := i.db.WithContext(ctx).Where("month = ?", month).Delete(&record.LineRecord{}).Error; err != nil {
		return err
	}

	// 过滤出数据
	lineSet := &record.LineRecordSet{}
	ql := fmt.Sprintf("cluster_name, project_line, SUM(weight) as weight, '%s' as month", month)

	if err := sql.Debug().
		Select(ql).
		Where("created_at >= ? AND created_at <=?", start, end).
		Group("cluster_name, project_line").
		Find(lineSet).
		Error; err != nil {
		return err
	}

	// 入库
	if err := i.db.WithContext(ctx).Model(&record.LineRecord{}).Create(lineSet).Error; err != nil {
		return err
	}
	return nil
}

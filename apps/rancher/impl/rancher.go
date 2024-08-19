package impl

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/common"
	"github.com/rancher/norman/types"
	"k8s.io/utils/strings/slices"
	"strings"
)

func (i *Impl) CreateTable(ctx context.Context, empty *rancher.Empty) (*rancher.Empty, error) {
	return nil, i.db.WithContext(ctx).AutoMigrate(&rancher.Project{})
}
func (i *Impl) SyncProject(empty *rancher.Empty, server rancher.Rpc_SyncProjectServer) error {
	// 列出所有集群
	clusterSet, err := i.c.Cluster.List(&types.ListOpts{})
	if err != nil {
		return err
	}

	// 每次都是重新获取数据，确保数据与当前完全一致
	i.db.Exec("TRUNCATE TABLE projects")
	i.db.Exec("ALTER TABLE projects AUTO_INCREMENT = 1")

	// 开始同步
	for _, cluster := range clusterSet.Data {
		// 项目信息
		projectSet, err := i.c.Project.List(&types.ListOpts{Filters: map[string]any{"clusterId": cluster.ID}})
		if err != nil {
			return err
		}

		for _, project := range projectSet.Data {
			// 用户信息
			roleSet, err := i.c.ProjectRoleTemplateBinding.List(&types.ListOpts{Filters: map[string]any{"projectId": project.ID}})
			if err != nil {
				return err
			}

			var users []string
			for _, binding := range roleSet.Data {
				user, err := i.c.User.ByID(binding.UserID)
				if err != nil {
					return err
				}
				if ok := slices.Contains(i.adm, user.Username); ok {
					continue
				}
				users = append(users, user.Username)
			}
			if users == nil {
				common.L().Warn().Msgf("[%s/%s] project has no active user, skip", cluster.Name, project.Name)
				continue // 如果项目（Default、System）中没有用户，则不要入库
			}

			ins := &rancher.Project{
				Spec: &rancher.Spec{
					ClusterName: cluster.Name,
					ProjectId:   project.ID,
					ProjectCode: project.Name,
					ProjectDesc: project.Annotations["projectDesc"],
					ProjectLine: project.Annotations["businessLine"],
					User:        strings.Join(users, ","),
				},
			}

			if err := i.db.WithContext(context.Background()).Create(ins).Error; err != nil {
				return err
			}

			if err := server.Send(ins); err != nil {
				return err
			}
		}
	}
	return nil
}

func (i *Impl) QueryProject(ctx context.Context, req *rancher.QueryProjectReq) (*rancher.ProjectSet, error) {
	ins := &rancher.ProjectSet{}
	sql := i.db.WithContext(ctx).Model(&rancher.Project{})

	switch req.SearchType {
	case rancher.SEARCH_TYPE_SEARCH_TYPE_ID:
		sql = sql.Where("id = ?", req.KeyWord)
	case rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_CODE:
		sql = sql.Where("project_code like ?", "%"+req.KeyWord+"%")
	case rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_LINE:
		sql = sql.Where("project_line like ?", "%"+req.KeyWord+"%")
	case rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_DESC:
		sql = sql.Where("project_desc like ?", "%"+req.KeyWord+"%")
	case rancher.SEARCH_TYPE_SEARCH_TYPE_CLUSTER_NAME:
		sql = sql.Where("cluster_name = ?", req.KeyWord)
	case rancher.SEARCH_TYPE_SEARCH_TYPE_ANNOTATION:
		sql = sql.Where("project_id = ?", req.KeyWord)
	}

	if err := sql.Find(&ins.Items).Error; err != nil {
		return nil, err
	}
	ins.Total = int64(len(ins.Items))
	return ins, nil
}

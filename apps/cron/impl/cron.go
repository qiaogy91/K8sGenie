package impl

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/record"
	"gitee.com/qiaogy91/K8sGenie/common"
)

func (i *Impl) Run() {
	//i.RancherTask()
	//i.K8sTask()
	i.RecordTask()
}

func (i *Impl) RancherTask() {
	common.L().Info().Msgf("RancherTask start")
	if err := i.rancherSvc.SyncProject(nil, &RancherStream{}); err != nil {
		common.L().Error().Msgf("RancherTask run err, %v", err)
	}
	common.L().Info().Msgf("RancherTask finished")
}

func (i *Impl) K8sTask() {
	common.L().Info().Msgf("K8sTask start")
	if err := i.k8sSvc.SyncK8SWorkload(nil, &K8sStream{}); err != nil {
		common.L().Error().Msgf("RancherTask run err, %v", err)
	}
	common.L().Info().Msgf("K8sTask finished")
}

// RecordTask 执行数据统计
func (i *Impl) RecordTask() {
	common.L().Info().Msgf("RecordTask start")
	if err := i.recordSvc.CreateNamespaceRecord(context.Background(), &record.CreateNamespaceRecordRequest{}); err != nil {
		common.L().Error().Msgf("job run err, %v", err)
	}
	common.L().Info().Msgf("RecordTask run namespace record finished")
	if err := i.recordSvc.CreateProjectRecord(context.Background(), &record.CreateProjectRecordRequest{}); err != nil {
		common.L().Error().Msgf("job run err, %v", err)
	}
	common.L().Info().Msgf("RecordTask run project record finished")
	if err := i.recordSvc.CreateLineRecord(context.Background(), &record.CreateLineRecordRequest{}); err != nil {
		common.L().Error().Msgf("job run err, %v", err)
	}
	common.L().Info().Msgf("RecordTask finished")
}

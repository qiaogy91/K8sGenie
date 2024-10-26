package impl

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/record"
	"gitee.com/qiaogy91/K8sGenie/common"
	"time"
)

func (i *Impl) Run() {
	common.L().Info().Msgf("Cron job start at %s", time.Now().Format("2006-01-02 15:04:05"))

	i.RancherTask()
	i.K8sTask()
	i.RecordTask()

	common.L().Info().Msgf("Cron job finished at %s", time.Now().Format("2006-01-02 15:04:05"))
}

func (i *Impl) RancherTask() {
}

func (i *Impl) K8sTask() {}

// RecordTask 执行数据统计
func (i *Impl) RecordTask() {
	if err := i.recordSvc.CreateNamespaceRecord(context.Background(), &record.CreateNamespaceRecordRequest{}); err != nil {
		common.L().Error().Msgf("job run err, %v", err)
	}
	common.L().Info().Msgf("Cron job namespace record finished")
	if err := i.recordSvc.CreateProjectRecord(context.Background(), &record.CreateProjectRecordRequest{}); err != nil {
		common.L().Error().Msgf("job run err, %v", err)
	}
	common.L().Info().Msgf("Cron job project record finished")
	if err := i.recordSvc.CreateLineRecord(context.Background(), &record.CreateLineRecordRequest{}); err != nil {
		common.L().Error().Msgf("job run err, %v", err)
	}
	common.L().Info().Msgf("Cron job line record finished")

}

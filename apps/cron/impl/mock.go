package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/common"
	"google.golang.org/grpc"
)

type RancherStream struct {
	grpc.ServerStream
}

func (s *RancherStream) Send(r *rancher.Project) error {
	common.L().Info().Msgf("stream: %+v", r)
	return nil
}

type K8sStream struct {
	grpc.ServerStream
}

func (s *K8sStream) Send(r *k8s.WorkLoad) error {
	common.L().Info().Msgf("stream: %+v", r)
	return nil
}

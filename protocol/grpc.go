package protocol

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/common"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"slices"
)

func TokenAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Aborted, "未获取到meta 数据")
	}

	// token 校验
	token := md.Get("token")
	if !slices.Contains(token, conf.C().Grpc.Token) {
		return nil, status.Errorf(codes.PermissionDenied, "非法token")
	}

	return handler(ctx, req) // 通过则放行
}

type GrpcServer struct {
	s *grpc.Server
}

func (g *GrpcServer) Start() error {
	// 服务注册
	for name, o := range ioc.LoadController() {
		o.RegistrySvc(g.s)
		common.L().Info().Msgf("%s controller object registry success", name)
	}

	// 启动监听
	ln, err := net.Listen("tcp", conf.C().GrpcAddr())
	if err != nil {
		return err
	}

	common.L().Info().Msgf("grpc server started, at %s", conf.C().GrpcAddr())
	if err := g.s.Serve(ln); err != nil {
		return err
	}
	return nil
}

func (g *GrpcServer) Stop() {
	g.s.GracefulStop()
	common.L().Info().Msgf("grpc server stopped")
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{
		s: grpc.NewServer(grpc.UnaryInterceptor(TokenAuth)),
	}
}

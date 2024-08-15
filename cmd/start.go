package cmd

import (
	_ "gitee.com/qiaogy91/K8sGenie/apps"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"gitee.com/qiaogy91/K8sGenie/protocol"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var (
	configType string
	configFile string
)
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// 配置文件加载
		switch configType {
		case "file":
			conf.LoadFromToml(configFile)
		default:
			conf.LoadFromEnv()
		}

		// ioc 对象初始化，附加数据库连接池、附加业务实现对象
		ioc.InitController()
		ioc.InitHandler()

		// protocol 协议服务器启动
		httpServer := protocol.NewHttpServer()
		grpcServer := protocol.NewGrpcServer()
		go func() {
			cobra.CheckErr(httpServer.Start())
		}()
		go func() {
			cobra.CheckErr(grpcServer.Start())
		}()

		// 优雅关闭
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
		select {
		case <-ch:
			httpServer.Stop()
			grpcServer.Stop()
			conf.C().CloseConnection()
		}
	},
}

func init() {
	startCmd.Flags().StringVarP(&configType, "config-type", "t", "", "config type")
	startCmd.Flags().StringVarP(&configFile, "config-file", "f", "", "config file")
}

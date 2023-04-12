package subset

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-aiart/service"
)

var serverAct string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "服务端管理",
	Run: func(cmd *cobra.Command, args []string) {
		service.Control("server", serverAct)
	},
}

func WithServer() *cobra.Command {

	serverCmd.Flags().BoolP("help", "h", false, "查看帮助")
	serverCmd.Flags().MarkHidden("help")

	serverCmd.Flags().StringVarP(&serverAct, "service", "s", "", "管理系统服务")
	serverCmd.Flags().StringP("listen", "l", ":7700", "服务端监听的IP地址和端口")

	viper.BindPFlag("server.listen", serverCmd.Flags().Lookup("listen"))

	return serverCmd

}

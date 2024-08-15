package main

import (
	"gitee.com/qiaogy91/K8sGenie/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}

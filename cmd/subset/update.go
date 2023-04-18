package subset

import (
	"github.com/open-tdp/go-helper/upgrade"
	"github.com/spf13/cobra"

	"tdp-aiart/cmd/args"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update assistant",
	Long:  "TDP Aiart Update Assistant",
	Run: func(cmd *cobra.Command, rq []string) {
		ExecUpdate()
	},
}

func WithUpdate() *cobra.Command {

	return updateCmd

}

func ExecUpdate() error {

	err := upgrade.Apply(&upgrade.RequesParam{
		UpdateUrl: args.UpdateUrl,
		Version:   args.Version,
	})

	return err

}

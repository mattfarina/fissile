package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	flagShowLayerFrom string
)

// showLayerCmd represents the layer command
var showLayerCmd = &cobra.Command{
	Use:   "layer",
	Short: "Displays information about all the docker layers used by fissile.",
	RunE: func(cmd *cobra.Command, args []string) error {
		flagShowLayerFrom = viper.GetString("from")

		return fissile.ShowBaseImage(
			flagShowLayerFrom,
			flagRepository,
		)
	},
}

func init() {
	showCmd.AddCommand(showLayerCmd)

	showLayerCmd.PersistentFlags().StringP(
		"from",
		"F",
		"ubuntu:14.04",
		"Docker image used as a base for the layers",
	)

	viper.BindPFlags(showLayerCmd.PersistentFlags())

}

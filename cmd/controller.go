package cmd

import (
	"github.com/logand22/service-mesh-from-scratch/pkg/controller"
	"github.com/logand22/service-mesh-from-scratch/pkg/k8s"
	"github.com/spf13/cobra"
)

var controllerCmd = &cobra.Command{
	Use: "controller",
	RunE: func(cmd *cobra.Command, args []string) error {
		kubeConfig, _ := cmd.PersistentFlags().GetString("kubeconfig")

		client, err := k8s.NewClient(kubeConfig)
		if err != nil {
			return err
		}

		control := controller.NewController(client)
		control.Start(cmd.Context().Done())

		return nil
	},
}

func init() {
	controllerCmd.PersistentFlags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests")

	rootCmd.AddCommand(controllerCmd)
}

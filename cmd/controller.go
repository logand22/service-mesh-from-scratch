package cmd

import (
	"fmt"
	"time"

	"github.com/logand22/service-mesh-from-scratch/pkg/controller"
	"github.com/logand22/service-mesh-from-scratch/pkg/k8s"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/labels"
)

var controllerCmd = &cobra.Command{
	Use: "controller",
	RunE: func(cmd *cobra.Command, args []string) error {
		kubeConfig, _ := cmd.PersistentFlags().GetString("kubeconfig")

		client, err := k8s.NewClient(kubeConfig)
		if err != nil {
			return err
		}

		controller := controller.NewController(client)
		controller.Start(cmd.Context().Done())
		time.Sleep(time.Second * 15)
		fmt.Println(controller.GetServices("", "", labels.Everything()))

		return nil
	},
}

func init() {
	controllerCmd.PersistentFlags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests")

	rootCmd.AddCommand(controllerCmd)
}

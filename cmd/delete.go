package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/TwinProduction/go-color"
	"github.com/digitalocean/godo"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Override PaaS Delete Cluster",
		Run: func(cmd *cobra.Command, args []string) {
			if len(cfgFile) < 1 && len(provider) < 1 {
				cmd.Help()
				return
			}
			if len(provider) < 1 {
				provider = viper.GetString("provider")
			}
			if provider == "do" {
				token := os.Getenv("DO_TOKEN")
				if token == "" {
					fmt.Println(color.Colorize(color.Red, "\u274c DO_TOKEN env is required when using provider: do"))
					return
				}
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Do you want to delete %v?", clusterName),
					IsConfirm: true,
					Default:   "N",
				}
				_, err := prompt.Run()
				if err != nil {
					fmt.Println(color.Colorize(color.Blue, "\u2139 No resource is deleted"))
				} else {
					if len(cfgFile) > 0 {
						clusterName = viper.GetString("name")
					}
					msg := fmt.Sprintf("\u2139 Deleting %v cluster from %v", clusterName, "DigitalOcean")
					fmt.Println(color.Colorize(color.Blue, msg))
					client := godo.NewFromToken(token)
					ctx := context.TODO()
					clusters, _, errCluster := client.Kubernetes.List(ctx, &godo.ListOptions{})
					if errCluster != nil {
						fmt.Println(errCluster)
					}
					var clusterID string
					for _, value := range clusters {
						if value.Name == clusterName {
							clusterID = value.ID
						}
					}
					_, err := client.Kubernetes.Delete(ctx, clusterID)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println(color.Colorize(color.Green, "\u2713 Delete cluster: "+clusterName+" successful"))
					}
				}
			}
		},
	}
)

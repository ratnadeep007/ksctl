package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/TwinProduction/go-color"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	services "ratnadeep007/override-pass/m/v2/services"
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Override PaaS Create Cluster",
		Long: `Creates kubernetes cluster.
--config will be searched first if not given then rest of the flags are all needed.
		`,
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
				fmt.Print(color.Colorize(color.Blue, "\u2139 Using Provider: DigitalOcean"))
				var nodes []struct {
					Type  string
					Count int
				}
				if len(cfgFile) > 0 {
					clusterName = viper.GetString("name")
					region = viper.GetString("region")
					errFromNode := viper.UnmarshalKey("nodes", &nodes)
					if errFromNode != nil {
						fmt.Println(color.Colorize(color.Red, "\u274c Error occured"))
						fmt.Println(errFromNode.Error())
						return
					}
					nodePools := make([]*godo.KubernetesNodePoolCreateRequest, len(nodes))
					for i := 0; i < len(nodes); i++ {
						nodePools[i] = &godo.KubernetesNodePoolCreateRequest{
							Size:  nodes[i].Type,
							Count: nodes[i].Count,
							Name:  clusterName + "-" + nodes[i].Type,
						}
					}
					services.CreateCluster(token, clusterName, region, nodePools)
				} else if len(provider) > 0 {
					if len(provider) == 0 || len(nodeCount) == 0 || len(nodeType) == 0 || len(clusterName) == 0 {
						fmt.Println(color.Colorize(color.Red, "If --config file not provided then all are required"))
						cmd.Help()
						return
					}
					fmt.Println(color.Colorize(color.Blue, "\u2139 Using provider: DigitalOcean\n"))
					nodePools := make([]*godo.KubernetesNodePoolCreateRequest, 1)
					count, _ := strconv.Atoi(nodeCount)
					nodePools[0] = &godo.KubernetesNodePoolCreateRequest{
						Size:  nodeType,
						Count: count,
						Name:  clusterName + "-" + nodeType,
					}
					services.CreateCluster(token, clusterName, region, nodePools)
				}
			}
		},
	}
)

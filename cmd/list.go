package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/TwinProduction/go-color"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

var (
	provider string
	listCmd  = &cobra.Command{
		Use:   "list",
		Short: "List k8s or k3s cluster",
		Run: func(cmd *cobra.Command, args []string) {
			providerCode := provider
			var token string
			if providerCode == "do" {
				token = os.Getenv("DO_TOKEN")
				if token == "" {
					fmt.Println(color.Colorize(color.Red, "\u274c DO_TOKEN env is required when using provider: do"))
					return
				}
				client := godo.NewFromToken(token)
				ctx := context.TODO()
				clusters, _, err := client.Kubernetes.List(ctx, &godo.ListOptions{})
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(color.Colorize(color.Blue, "Found "+strconv.Itoa(len(clusters))+" cluster(s)"))
				for _, value := range clusters {
					fmt.Println(value.Name + " (" + value.ID + ") " + " -> " + value.VersionSlug + " (" + value.RegionSlug + ")")
				}
			}
		},
	}
)

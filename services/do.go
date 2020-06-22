package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/TwinProduction/go-color"
	"github.com/digitalocean/godo"
	"github.com/manifoldco/promptui"
)

// CreateCluster create clauser
func CreateCluster(token string, clusterName string, region string, nodePools []*godo.KubernetesNodePoolCreateRequest) {
	client := godo.NewFromToken(token)
	createRequest := &godo.KubernetesClusterCreateRequest{
		Name:        clusterName,
		NodePools:   nodePools,
		RegionSlug:  region,
		VersionSlug: "1.17",
	}
	ctx := context.TODO()
	newk8s, _, err := client.Kubernetes.Create(ctx, createRequest)
	fmt.Println(color.Colorize(color.Blue, "\u2139 Creating cluster"+clusterName))
	if err != nil {
		println(err.Error())
	} else {
		ticker := time.NewTicker(5 * time.Second)
		quit := make(chan struct{})
		for {
			select {
			case <-ticker.C:
				stat, _, _ := client.Kubernetes.Get(ctx, newk8s.ID)
				if stat.Status.State == "running" {
					config, resp, errConfig := client.Kubernetes.GetKubeConfig(ctx, newk8s.ID)
					if errConfig != nil {
						fmt.Println(errConfig)
						fmt.Println(color.Colorize(color.Red, "\u274c Error while getting kube config\n"))
						os.Exit(0)
					}
					fmt.Print("\n")
					fmt.Print(color.Colorize(color.Blue, "\u2139 Kubeconfig for "+clusterName+"\n"))
					fmt.Println(string(config.KubeconfigYAML))
					ioutil.WriteFile(clusterName+"-kubecofig", config.KubeconfigYAML, 0644)
					fmt.Println(color.Colorize(color.Blue, "\u2139 Save kubeconfig as "+clusterName+"-kubeconfig\n"))
					fmt.Println(resp)
					fmt.Println(color.Colorize(color.Green, "\u2713 Creating cluster: "+clusterName))
					close(quit)
					os.Exit(0)
				}
			case <-quit:
				ticker.Stop()
				// return
			}
		}
	}
}

// DeleteCluster deletes cluster
func DeleteCluster(token string, clusterName string) {
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Do you want to delete %v?", clusterName),
		IsConfirm: true,
		Default:   "N",
	}
	result, _ := prompt.Run()
	if len(result) < 0 {
		fmt.Println(color.Colorize(color.Blue, "\u2139 No resource is deleted"))
		return
	}
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

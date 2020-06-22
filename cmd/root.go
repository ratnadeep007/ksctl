package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	clusterName string
	nodeType    string
	nodeCount   string
	region      string
	rootCmd     = &cobra.Command{
		Use:   "ovr",
		Short: "Override PaaS",
		Long: `Override PaaS is a single binary Platform running on kubernetes.
Supports k8s and k3s both for different cloud providers.
Availble cloud provider: AWS, Azure, DigitalOcean
		`,
		Version: "0.0.1-alpha",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	createCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is do.yaml)")
	createCmd.PersistentFlags().StringVar(&clusterName, "cluster-name", "", "name of cluster")
	createCmd.PersistentFlags().StringVar(&nodeType, "node-type", "", "type of node to use")
	createCmd.PersistentFlags().StringVar(&provider, "provider", "", "provider to use to create cluster")
	createCmd.PersistentFlags().StringVar(&nodeCount, "node-count", "", "number of node for given type")
	createCmd.PersistentFlags().StringVar(&region, "region", "", "region to create cluster")

	deleteCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is do.yaml)")
	deleteCmd.PersistentFlags().StringVar(&provider, "provider", "", "provider to delete cluster from")
	deleteCmd.PersistentFlags().StringVar(&clusterName, "cluster-name", "", "name to cluster to delete")

	listCmd.PersistentFlags().StringVar(&provider, "provider", "", "provider to list cluster")
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createCmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".yaml")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using Config file: ", viper.ConfigFileUsed())
	}
}

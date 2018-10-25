// Copyright © 2018 Brandon Height <bmheight@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/brandon-height/kray/kube"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
)

// KubeClient ...
var KubeClient kubernetes.Interface

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kray",
	Short: "kray is used to interact with multiple cloud UDE environments",
	Long: `kray is designed to manage UDE environments within a multi cloud space.
This includes the creation, deletion, and listing of resources in a UDE. For specific cloud environments this information is not just Kubernetes specific, but also any VM instances, DNS Records, etc that may be associated to a given UDE environment.`,
	PersistentPreRun: kubeConnect,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("cloud", "c", "", "cloud provider to use for this transaction.")
	rootCmd.PersistentFlags().StringP("namespace", "n", "", "namespace name to use for this transaction.")

	rootCmd.MarkPersistentFlagRequired("namespace")
	rootCmd.MarkPersistentFlagRequired("cloud")
	rootCmd.MarkFlagRequired("namespace")
	rootCmd.MarkFlagRequired("cloud")
}

// initConfig reads in config file and ENV variables if set.
/*func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cmd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}*/

func kubeConnect(cmd *cobra.Command, args []string) {
	var err error
	KubeClient, err = kube.NewClient()
	if err != nil {
		log.Fatalln(err)
	}
}

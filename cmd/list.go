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
	"log"

	"github.com/brandon-height/kray/kube"
	"github.com/brandon-height/kray/ude"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "perform a list action for a ude resource",
	Long:  `This command performs the steps to list the resources of a given UDE environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the flags
		cloud, err := cmd.Parent().PersistentFlags().GetString("cloud")
		if err != nil {
			log.Fatalln(err)
		}

		namespace, err := cmd.Parent().PersistentFlags().GetString("namespace")
		if err != nil {
			log.Fatalln(err)
		}

		// Setup a config
		config := kube.NewConfig(namespace, KubeClient)

		// Create the resource
		switch cloud {
		case "gcp":
			if err := list(namespace, config, ude.NewGCP(namespace)); err != nil {
				log.Fatalln(err)
			}
		case "alicloud":
			if err := list(namespace, config, ude.NewAlicloud(namespace)); err != nil {
				log.Fatalln(err)
			}
		case "aws":
			if err := list(namespace, config, ude.NewAWS(namespace)); err != nil {
				log.Fatalln(err)
			}
		default:
			log.Fatalln("Unknown input")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func list(name string, c *kube.Config, u ude.Interface) error {
	return u.List(name, c)
}

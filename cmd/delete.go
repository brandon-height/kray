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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "perform a delete action for a ude resource",
	Long:  `This command peforms the steps to delete a specific or set of resources within a UDE.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the flags
		cloud, _ := cmd.Parent().PersistentFlags().GetString("cloud")
		namespace, _ := cmd.Parent().PersistentFlags().GetString("namespace")

		// Setup a config
		config := kube.NewConfig(namespace, KubeClient)

		// delete the resource
		switch cloud {
		case "gcp":
			if err := delete(config, ude.NewGCP(namespace)); err != nil {
				log.Fatalln(err)
			}
		case "alicloud":
			if err := delete(config, ude.NewAlicloud(namespace)); err != nil {
				log.Fatalln(err)
			}
		case "aws":
			if err := delete(config, ude.NewAWS(namespace)); err != nil {
				log.Fatalln(err)
			}
		default:
			log.Fatalln("Unknown input")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete(c *kube.Config, u ude.Interface) error {
	return u.Delete(c)
}

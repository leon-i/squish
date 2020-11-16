/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"squish/config"
	"squish/optimizer"
	"squish/util"
	"time"
)

// onlyCmd represents the only command
var onlyCmd = &cobra.Command{
	Use:   "only",
	Short: "optimize specific files by name in your current directory",
	Long: `optimize specific files by name in the current directory - separate files by space and if the name of the file contains spaces, wrap in quotes. Ex. squish only "example file.jpg" `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--------------------------------------------------")
		fmt.Println("squishing desired image(s) in current directory...")
		fmt.Println("--------------------------------------------------")

		q, _ := cmd.Flags().GetUint("quality")
		d, _ := cmd.Flags().GetString("destination")

		config.SetValues(q, d)

		util.Startup()

		start := time.Now()
		optimizeOnly(args)

		util.LogDuration(start)

		util.Cleanup()
	},
}

func optimizeOnly(fileNames []string) {
	for _, n := range fileNames {
		f, err := os.Open(n)
		util.Check(err)

		t, err := util.GetFileContentType(f)
		if err != nil {
			util.LogError(err)
			continue
		}

		if t != "image/jpeg" && t != "image/png" {
			fmt.Println(fmt.Sprintf("ERROR - %s is not a valid image", n))
			continue
		}

		optimizer.ToMozillaJpeg(f, t)
	}
}

func init() {
	rootCmd.AddCommand(onlyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onlyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	onlyCmd.Flags().UintP("quality", "q", viper.GetUint("quality"), "Set output image quality")
	onlyCmd.Flags().StringP("destination", "d", viper.GetString("destination"), "Set destination/new directory name for outputted images")
}

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
	"io/ioutil"
	"os"
	"squish/util"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("all called")

		if _, err := os.Stat("./squished"); os.IsNotExist(err) {
			if err := os.Mkdir("squished", 0755); err != nil {
				panic(err)
			}
		}

		files, err := ioutil.ReadDir(".")

		if (err != nil ) {
			fmt.Println(err)
		}

		walkFiles(files)

		util.Cleanup()
	},
}

func walkFiles(files []os.FileInfo) {
	for _, file := range files {
		n := file.Name()

		if (n == "squished") {
			continue
		}

		f, err := os.Open(n)
		defer f.Close()

		if err != nil {
			fmt.Println(err)
			continue
		}

		t, err := util.GetFileContentType(f)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if t != "image/jpeg" && t != "image/png" {
			fmt.Println(fmt.Sprintf("%s is not a valid image", n))
			continue
		}
		
		util.OptimizeImage(f, t)
	}
}

func init() {
	rootCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

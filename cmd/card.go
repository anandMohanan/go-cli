/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
)

// cardCmd represents the card command
var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("card called")
		name, _ := cmd.Flags().GetString("name")
		occasion, _ := cmd.Flags().GetString("occassion")
		language, _ := cmd.Flags().GetString("language")
		fmt.Println("value of the flag name :" + name)
		fmt.Println("value of the flag occasion :" + occasion)
		fmt.Println("value of the flag language :" + language)

	},
}

func init() {
	createCmd.AddCommand(cardCmd)
	cardCmd.PersistentFlags().StringP("occasion", "o", "", "Possible values: newyear, thanksgiving, birthday")
	cardCmd.PersistentFlags().StringP("language", "l", "en", "Possible values: en, fr")
	cardCmd.PersistentFlags().StringP("name", "n", "", "Name of the user to whom you want to greet")
	cardCmd.MarkPersistentFlagRequired("name")
	cardCmd.MarkPersistentFlagRequired("occassion")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

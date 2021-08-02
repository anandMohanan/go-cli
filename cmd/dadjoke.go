/*
Copyright © 2021 ANAND MOHANAN ananthkvmohanan@gmail.com

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
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/spf13/cobra"
)

// dadjokeCmd represents the dadjoke command
var dadjokeCmd = &cobra.Command{
	Use:   "dadjoke",
	Short: "Returns a dadjoke",
	Long:  `Returns a dadjoke from an api`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(dadjokeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dadjokeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dadjokeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJoke(url)
	joke := joke{}
	err := json.Unmarshal(responseBytes, &joke)
	ErrorHandling(err)

	fmt.Println(joke.Joke)

}

func getJoke(baseAPI string) []byte {
	request, err := http.NewRequest(http.MethodGet, baseAPI, nil)
	ErrorHandling(err)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke cli")

	response, ero := http.DefaultClient.Do(request)
	ErrorHandling(ero)
	responseBytes, erro := ioutil.ReadAll(response.Body)
	ErrorHandling(erro)

	return responseBytes
}

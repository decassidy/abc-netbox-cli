/*
Copyright © 2024 Derrick Cassidy - Metropolis Technologies, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package core

import (
	"github.com/decassidy/metropolis-netbox-cli/cmd/dcim"
	"github.com/spf13/cobra"
)

type dataFiles struct {
	Count    uint   `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Id      uint   `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Source  struct {
			dcim.CommonFieldsNoSlug
		} `json:"source"`
		Path        string `json:"path"`
		LastUpdated string `json:"last_updated"`
		Size        uint   `json:"size"`
		Hash        string `json:"hash"`
	} `json:"results"`
}

// GetCoreDataFilesCmd represents the getCoreDataFiles command
var GetCoreDataFilesCmd = &cobra.Command{
	Use:   "getCoreDataFiles",
	Short: "Get Metropolis Netbox core data file objects.",
	Long:  `Get Metropolis Netbox core data file objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(dataFiles)
		dcim.ApiConnectionNonID(responseObject, "GET", "cmd.core.core_api_url.data_files")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	//GetCoreDataFilesCmd.Flags().StringVarP()

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCoreDataFilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCoreDataFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

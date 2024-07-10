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

package circuits

import (
	"github.com/decassidy/metropolis-netbox-cli/cmd/dcim"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
)

type providers struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		dcim.CommonFieldsSlug
		Accounts []struct {
			dcim.CommonFieldsNoSlug
			Account string `json:"account"`
		} `json:"accounts"`
		Description string `json:"description"`
		Comments    string `json:"comments"`
		Asns        []int  `json:"asns"`
		Tags        []struct {
			dcim.CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created      string `json:"created"`
		LastUpdated  string `json:"last_updated"`
		CircuitCount int    `json:"circuit_count"`
	} `json:"results"`
}

// GetCircuitsProvidersCmd represents the getCircuitsProviders command
var GetCircuitsProvidersCmd = &cobra.Command{
	Use:   "getCircuitsProviders",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(providers)
		ApiConnectionNonID(responseObject, "GET", "cmd.circuits.circuits_api_url.providers")

		if responseObject.Count != 0 {
			color.Cyan("\nMetropolis Total Providers in Netbox: "+color.YellowString("%d"), responseObject.Count)
			for _, provider := range responseObject.Results {
				color.Cyan("\n============================================================================")
				color.Cyan("\n\tMetropolis Provider Name: "+color.YellowString("%s\n"), provider.Display)
				color.Cyan("============================================================================\n")
				color.Cyan("\tID: "+color.YellowString("%d"), provider.Id)
				color.Cyan("\tURL: "+color.YellowString("%s"), provider.Url)
				color.Cyan("\tDisplay: "+color.YellowString("%s"), provider.Display)
				color.Cyan("\tName: "+color.YellowString("%s"), provider.Name)
				color.Cyan("\tSlug: "+color.YellowString("%s"), provider.Slug)
				for _, account := range provider.Accounts {
					if account.Id != 0 {
						color.Cyan("\tAccount: ")
						color.Cyan("\t  ID: "+color.YellowString("%d"), account.Id)
						color.Cyan("\t  URL: "+color.YellowString("%s"), account.Url)
						color.Cyan("\t  Display: "+color.YellowString("%s"), account.Display)
						color.Cyan("\t  Name: "+color.YellowString("%s"), account.Name)
						color.Cyan("\t  Account: "+color.YellowString("%s"), account.Account)
					} else {
						color.Cyan("\tAccount: " + color.RedString("No provider account found for provider: %s", color.YellowString("%s", provider.Display)))
					}
				}
				if provider.Description != "" {
					color.Cyan("\tDescription: "+color.YellowString("%s"), provider.Description)
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description found for provider: %s", color.YellowString("%s", provider.Display)))
				}
				if provider.Comments != "" {
					color.Cyan("\tComments: "+color.YellowString("%s"), provider.Comments)
				} else {
					color.Cyan("\tComments: " + color.RedString("No comments found for provider: %s", color.YellowString("%s", provider.Display)))
				}
				for _, asn := range provider.Asns {
					if asn != 0 {
						color.Cyan("\tAccount: "+color.YellowString("%d"), asn)
					} else {
						color.Cyan("\tAccount: " + color.RedString("No provider account found for provider: %s", color.YellowString("%s", provider.Display)))
					}
				}
				for _, tag := range provider.Tags {
					if tag.Id == 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: "+color.YellowString("%d"), tag.Id)
						color.Cyan("\t  URL: "+color.YellowString("%s"), tag.Url)
						color.Cyan("\t  Display: "+color.YellowString("%s"), tag.Display)
						color.Cyan("\t  Name: "+color.YellowString("%s"), tag.Name)
						color.Cyan("\t  Slug: "+color.YellowString("%s"), tag.Slug)
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags found for provider: %s", color.YellowString("%s", provider.Display)))
					}
				}
				color.Cyan("\tCreated: "+color.YellowString("%s"), provider.Created)
				color.Cyan("\tLast Updated: "+color.YellowString("%s"), provider.LastUpdated)
				color.Cyan("\tCircuit Count: "+color.YellowString("%d"), provider.CircuitCount)
			}
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetCircuitsProvidersCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetCircuitsProvidersCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetCircuitsProvidersCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCircuitsProvidersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCircuitsProvidersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

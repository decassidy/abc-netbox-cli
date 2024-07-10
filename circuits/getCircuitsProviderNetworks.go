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

type providerNetworks struct {
	Count    uint   `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Id       uint   `json:"id"`
		Url      string `json:"url"`
		Display  string `json:"display"`
		Provider struct {
			dcim.CommonFieldsSlug
		} `json:"provider"`
		Name        string `json:"name"`
		ServiceId   string `json:"service_id"`
		Description string `json:"description"`
		Comments    string `json:"comments"`
		Tags        []struct {
			dcim.CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created     string `json:"created"`
		LastUpdated string `json:"last_updated"`
	} `json:"results"`
}

// GetCircuitsProviderNetworksCmd represents the circuitProviderNetworks command
var GetCircuitsProviderNetworksCmd = &cobra.Command{
	Use:   "getCircuitsProviderNetworks",
	Short: "GET a list of Provider Network objects.",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of Provider Network objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(providerNetworks)
		ApiConnectionNonID(responseObject, "GET", "cmd.circuits.circuits_api_url.provider_networks")

		if responseObject.Count != 0 {
			color.Cyan("\nMetropolis Total Provider Networks in Netbox: "+color.YellowString("%d"), responseObject.Count)
			for _, network := range responseObject.Results {
				color.Cyan("\n============================================================================")
				color.Cyan("\n\tMetropolis Provider Network Name: "+color.YellowString("%s\n"), network.Display)
				color.Cyan("============================================================================\n")
				color.Cyan("\tID: "+color.YellowString("%d"), network.Id)
				color.Cyan("\tURL: "+color.YellowString("%s"), network.Url)
				color.Cyan("\tDisplay: "+color.YellowString("%s"), network.Display)
				if network.Provider.Id != 0 {
					color.Cyan("\tProvider: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), network.Provider.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), network.Provider.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), network.Provider.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), network.Provider.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%s"), network.Provider.Slug)
				} else {
					color.Cyan("\tProvider: " + color.RedString("No provider found for provider network: %s", color.YellowString("%s", network.Display)))
				}
				if network.Name != "" {
					color.Cyan("\tName: "+color.YellowString("%s"), network.Name)
				} else {
					color.Cyan("\tName: " + color.RedString("No name found for provider account: %s", color.YellowString("%s", network.Display)))
				}
				if network.ServiceId != "" {
					color.Cyan("\tService ID: "+color.YellowString("%d"), network.ServiceId)
				} else {
					color.Cyan("\tService ID: " + color.RedString("No service ID found for provider network: %s", color.YellowString("%s", network.Display)))
				}
				if network.Description != "" {
					color.Cyan("\tDescription: "+color.YellowString("%s"), network.Description)
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description found for provider account: %s", color.YellowString("%s", network.Display)))
				}
				if network.Comments != "" {
					color.Cyan("\tComments: "+color.YellowString("%s"), network.Comments)
				} else {
					color.Cyan("\tComments: " + color.RedString("No comments found for provider account: %s", color.YellowString("%s", network.Display)))
				}
				for _, tag := range network.Tags {
					if tag.Id == 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: "+color.YellowString("%d"), tag.Id)
						color.Cyan("\t  URL: "+color.YellowString("%s"), tag.Url)
						color.Cyan("\t  Display: "+color.YellowString("%s"), tag.Display)
						color.Cyan("\t  Name: "+color.YellowString("%s"), tag.Name)
						color.Cyan("\t  Slug: "+color.YellowString("%s"), tag.Slug)
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags found for provider account: %s", color.YellowString("%s", network.Display)))
					}
				}
				color.Cyan("\tCreated: "+color.YellowString("%s"), network.Created)
				color.Cyan("\tLast Updated: "+color.YellowString("%s"), network.LastUpdated)
			}
		} else {
			color.Cyan("  Metropolis Total Provider Networks in Netbox: " + color.RedString("No provider networks found on the server\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetCircuitsProviderNetworksCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetCircuitsProviderNetworksCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetCircuitsProviderNetworksCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// circuitProviderNetworksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// circuitProviderNetworksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

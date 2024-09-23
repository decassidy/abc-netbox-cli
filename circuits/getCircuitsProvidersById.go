/*
Copyright © 2024 Derrick Cassidy.

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
	"github.com/decassidy/abc-netbox-cli/cmd/dcim"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
)

type providersByID struct {
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
}

// GetCircuitsProvidersByIdCmd represents the getCircuitsProvidersById command
var GetCircuitsProvidersByIdCmd = &cobra.Command{
	Use:   "getCircuitsProvidersById",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(providersByID)
		apiConnectionID(responseObject, "GET", "cmd.circuits.circuits_api_url.providers_id")

		if responseObject.Id > 0 {
			color.Cyan("\n============================================================================")
			color.Cyan("\n\tABC Provider Name: "+color.YellowString("%s\n"), responseObject.Display)
			color.Cyan("============================================================================\n")
			color.Cyan("\tID: "+color.YellowString("%d"), responseObject.Id)
			color.Cyan("\tURL: "+color.YellowString("%s"), responseObject.Url)
			color.Cyan("\tDisplay: "+color.YellowString("%s"), responseObject.Display)
			color.Cyan("\tName: "+color.YellowString("%s"), responseObject.Name)
			color.Cyan("\tSlug: "+color.YellowString("%s"), responseObject.Slug)
			for _, account := range responseObject.Accounts {
				if account.Id != 0 {
					color.Cyan("\tAccount: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), account.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), account.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), account.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), account.Name)
					color.Cyan("\t  Account: "+color.YellowString("%s"), account.Account)
				} else {
					color.Cyan("\tAccount: " + color.RedString("No provider account found for provider: %s", color.YellowString("%s", responseObject.Display)))
				}
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: "+color.YellowString("%s"), responseObject.Description)
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description found for provider: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Comments != "" {
				color.Cyan("\tComments: "+color.YellowString("%s"), responseObject.Comments)
			} else {
				color.Cyan("\tComments: " + color.RedString("No comments found for provider: %s", color.YellowString("%s", responseObject.Display)))
			}
			for _, asn := range responseObject.Asns {
				if asn != 0 {
					color.Cyan("\tAccount: "+color.YellowString("%d"), asn)
				} else {
					color.Cyan("\tAccount: " + color.RedString("No provider account found for provider: %s", color.YellowString("%s", responseObject.Display)))
				}
			}
			for _, tag := range responseObject.Tags {
				if tag.Id == 0 {
					color.Cyan("\tTags: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), tag.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), tag.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), tag.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), tag.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%s"), tag.Slug)
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags found for provider: %s", color.YellowString("%s", responseObject.Display)))
				}
			}
			color.Cyan("\tCreated: "+color.YellowString("%s"), responseObject.Created)
			color.Cyan("\tLast Updated: "+color.YellowString("%s"), responseObject.LastUpdated)
			color.Cyan("\tCircuit Count: "+color.YellowString("%d\n"), responseObject.CircuitCount)
		} else {
			color.Red("  Doh! No provider object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetCircuitsProvidersByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetCircuitsProvidersByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetCircuitsProvidersByIdCmd", err)
	}

	GetCircuitsProvidersByIdCmd.Flags().IntVarP(&id, "id", "", 0, "Provider ID")
	err = GetCircuitsProvidersByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetCircuitsProvidersByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCircuitsProvidersByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCircuitsProvidersByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

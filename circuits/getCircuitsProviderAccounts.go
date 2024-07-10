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

type providerAccounts struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Id       int    `json:"id"`
		Url      string `json:"url"`
		Display  string `json:"display"`
		Provider struct {
			dcim.CommonFieldsSlug
		} `json:"provider"`
		Name        string `json:"name"`
		Account     string `json:"account"`
		Description string `json:"description"`
		Comments    string `json:"comments"`
		Tags        []struct {
			Id      int    `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Name    string `json:"name"`
			Slug    string `json:"slug"`
			Color   string `json:"color"`
		} `json:"tags"`
		Created     string `json:"created"`
		LastUpdated string `json:"last_updated"`
	} `json:"results"`
}

// GetCircuitsProviderAccountsCmd represents the circuitProviderAccounts command
var GetCircuitsProviderAccountsCmd = &cobra.Command{
	Use:   "getCircuitsProviderAccounts",
	Short: "GET a list of Provider Account objects.",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of Provider Account objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(providerAccounts)
		ApiConnectionNonID(responseObject, "GET", "cmd.circuits.circuits_api_url.provider_accounts")

		if responseObject.Count != 0 {
			color.Cyan("\nMetropolis Total Provider Accounts in Netbox: "+color.YellowString("%d"), responseObject.Count)
			for _, account := range responseObject.Results {
				color.Cyan("\n============================================================================")
				color.Cyan("\n\tMetropolis Provider Account Name: "+color.YellowString("%s\n"), account.Display)
				color.Cyan("============================================================================\n")
				color.Cyan("\tID: "+color.YellowString("%d"), account.Id)
				color.Cyan("\tURL: "+color.YellowString("%s"), account.Url)
				color.Cyan("\tDisplay: "+color.YellowString("%s"), account.Display)
				if account.Provider.Id != 0 {
					color.Cyan("\tProvider: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), account.Provider.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), account.Provider.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), account.Provider.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), account.Provider.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%s"), account.Provider.Slug)
				} else {
					color.Cyan("\tProvider: " + color.RedString("No provider found for provider account: %s", color.YellowString("%s", account.Display)))
				}
				if account.Name != "" {
					color.Cyan("\tName: "+color.YellowString("%s"), account.Name)
				} else {
					color.Cyan("\tName: " + color.RedString("No name found for provider account: %s", color.YellowString("%s", account.Display)))
				}
				if account.Account != "" {
					color.Cyan("\tAccount: "+color.YellowString("%s"), account.Account)
				} else {
					color.Cyan("\tAccount: " + color.RedString("No account found for provider account: %s", color.YellowString("%s", account.Display)))
				}
				if account.Description != "" {
					color.Cyan("\tDescription: "+color.YellowString("%s"), account.Description)
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description found for provider account: %s", color.YellowString("%s", account.Display)))
				}
				if account.Comments != "" {
					color.Cyan("\tComments: "+color.YellowString("%s"), account.Comments)
				} else {
					color.Cyan("\tComments: " + color.RedString("No comments found for provider account: %s", color.YellowString("%s", account.Display)))
				}
				for _, tag := range account.Tags {
					if tag.Id == 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: "+color.YellowString("%d"), tag.Id)
						color.Cyan("\t  URL: "+color.YellowString("%s"), tag.Url)
						color.Cyan("\t  Display: "+color.YellowString("%s"), tag.Display)
						color.Cyan("\t  Name: "+color.YellowString("%s"), tag.Name)
						color.Cyan("\t  Slug: "+color.YellowString("%s"), tag.Slug)
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags found for provider account: %s", color.YellowString("%s", account.Display)))
					}
				}
				color.Cyan("\tCreated: "+color.YellowString("%s"), account.Created)
				color.Cyan("\tLast Updated: "+color.YellowString("%s"), account.LastUpdated)
			}
		} else {
			color.Cyan("\nMetropolis Total Provider Accounts in Netbox: " + color.RedString("No provider accounts found on the server"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetCircuitsProviderAccountsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetCircuitsProviderAccountsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetCircuitsProviderAccountsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// circuitProviderAccountsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// circuitProviderAccountsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

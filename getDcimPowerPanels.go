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

package dcim

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type powerPanels struct {
	Count    uint   `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Id      uint   `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Site    struct {
			CommonFieldsSlug
		} `json:"site"`
		Location struct {
			CommonFieldsSlug
			Depth uint `json:"_depth"`
		} `json:"location"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Comments    string `json:"comments"`
		Tags        []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		PowerfeedCount uint   `json:"powerfeed_count"`
		Created        string `json:"created"`
		LastUpdated    string `json:"last_updated"`
	} `json:"results"`
}

// GetDcimPowerPanelsCmd represents the getDcimPowerPanels command
var GetDcimPowerPanelsCmd = &cobra.Command{
	Use:   "getDcimPowerPanels",
	Short: "GET a list of power panel objects",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of power panel objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(powerPanels)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.power_panels")

		if responseObject.Count != 0 {
			color.Cyan("\n  Metropolis Power Panels: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    Metropolis Power Panel: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.Site.Id != 0 {
					color.Cyan("\tSite: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Site.Id))
					color.Cyan("\t  URL: " + color.YellowString("%d", result.Site.Url))
					color.Cyan("\t  Display: " + color.YellowString("%d", result.Site.Display))
					color.Cyan("\t  Name: " + color.YellowString("%d", result.Site.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%d", result.Site.Slug))
				} else {
					color.Cyan("\tSite" + color.RedString("No site entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Location.Id != 0 {
					color.Cyan("\tLocation: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Location.Id))
					color.Cyan("\t  URL: " + color.YellowString("%d", result.Location.Url))
					color.Cyan("\t  Display: " + color.YellowString("%d", result.Location.Display))
					color.Cyan("\t  Name: " + color.YellowString("%d", result.Location.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%d", result.Location.Slug))
					color.Cyan("\t  Depth: " + color.YellowString("%d", result.Location.Depth))
				} else {
					color.Cyan("\tLocation" + color.RedString("No location entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Name != "" {
					color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				} else {
					color.Cyan("\tName" + color.RedString("No name entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription" + color.RedString("No description entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Comments != "" {
					color.Cyan("\tComments: " + color.YellowString("%s", result.Comments))
				} else {
					color.Cyan("\tComments" + color.RedString("No comments entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.PowerfeedCount != 0 {
					color.Cyan("\tPowerfeed Count: " + color.YellowString("%d", result.PowerfeedCount))
				} else {
					color.Cyan("\tPowerfeed Count" + color.RedString("No powerfeed count entry found for ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
			}
		} else {
			color.Cyan("  Metropolis Power Panels: " + color.RedString("No power panels found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimPowerPanelsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimPowerPanelsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPowerPanelsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimPowerPanelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimPowerPanelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

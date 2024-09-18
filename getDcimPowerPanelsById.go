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

type powerPanelsByID struct {
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
}

// GetDcimPowerPanelsByIdCmd represents the getDcimPowerPanelsById command
var GetDcimPowerPanelsByIdCmd = &cobra.Command{
	Use:   "getDcimPowerPanelsById",
	Short: "GET an power panel object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an power outlet panel by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(powerPanelsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.power_panels_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    Metropolis Power Panel: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.Site.Id != 0 {
				color.Cyan("\tSite: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Site.Id))
				color.Cyan("\t  URL: " + color.YellowString("%d", responseObject.Site.Url))
				color.Cyan("\t  Display: " + color.YellowString("%d", responseObject.Site.Display))
				color.Cyan("\t  Name: " + color.YellowString("%d", responseObject.Site.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%d", responseObject.Site.Slug))
			} else {
				color.Cyan("\tSite" + color.RedString("No site entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Location.Id != 0 {
				color.Cyan("\tLocation: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Location.Id))
				color.Cyan("\t  URL: " + color.YellowString("%d", responseObject.Location.Url))
				color.Cyan("\t  Display: " + color.YellowString("%d", responseObject.Location.Display))
				color.Cyan("\t  Name: " + color.YellowString("%d", responseObject.Location.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%d", responseObject.Location.Slug))
				color.Cyan("\t  Depth: " + color.YellowString("%d", responseObject.Location.Depth))
			} else {
				color.Cyan("\tLocation" + color.RedString("No location entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tName" + color.RedString("No name entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription" + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Comments != "" {
				color.Cyan("\tComments: " + color.YellowString("%s", responseObject.Comments))
			} else {
				color.Cyan("\tComments" + color.RedString("No comments entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.PowerfeedCount != 0 {
				color.Cyan("\tPowerfeed Count: " + color.YellowString("%d", responseObject.PowerfeedCount))
			} else {
				color.Cyan("\tPowerfeed Count" + color.RedString("No powerfeed count entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s\n", responseObject.LastUpdated))
		} else {
			color.Red("  Doh! No power panel object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimPowerPanelsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimPowerPanelsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking dcim power panel flag as required, %s - for GetDcimPowerPanelsByIdCmd", err)
	}

	GetDcimPowerPanelsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the power panel object")
	err = GetDcimPowerPanelsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking id flag as required, %s - for GetDcimPowerPanelsByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimPowerPanelsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimPowerPanelsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

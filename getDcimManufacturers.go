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

type manufacturers struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		CommonFieldsSlug
		Description string `json:"description"`
		Tags        []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created            string `json:"created"`
		LastUpdated        string `json:"last_updated"`
		DevicetypeCount    int    `json:"devicetype_count"`
		InventoryitemCount int    `json:"inventoryitem_count"`
		PlatformCount      int    `json:"platform_count"`
	} `json:"results"`
}

// GetDcimManufacturersCmd represents the getDcimManufacturers command
var GetDcimManufacturersCmd = &cobra.Command{
	Use:   "getDcimManufacturers",
	Short: "GET a list of manufacturer objects",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of manufacturer objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(manufacturers)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.manufacturers")

		if responseObject.Count != 0 {
			color.Cyan("\n  Metropolis Manufacturers: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    Metropolis Manufacturers: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				color.Cyan("\tSlug: " + color.YellowString("%s", result.Slug))
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", result.Display))
				}
				for _, tag := range result.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
						color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
						color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
						color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
						color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
						color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s", result.Display))
					}
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
				if result.DevicetypeCount != 0 {
					color.Cyan("\tDevice Type Count: " + color.YellowString("%d", result.DevicetypeCount))
				} else {
					color.Cyan("\tDevice Type Count: " + color.RedString("No device type count entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.InventoryitemCount != 0 {
					color.Cyan("\tInventory Item Count: " + color.YellowString("%d", result.InventoryitemCount))
				} else {
					color.Cyan("\tInventory Item Count: " + color.RedString("No inventory item count entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.PlatformCount != 0 {
					color.Cyan("\tPlatform Count: " + color.YellowString("%d\n", result.PlatformCount))
				} else {
					color.Cyan("\tPlatform Count: " + color.RedString("No platform count entry found for ") + color.YellowString("%s\n", result.Display))
				}
			}
		} else {
			color.Cyan("  Metropolis Manufacturers: " + color.RedString("No manufacturers found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimManufacturersCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimManufacturersCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimManufacturersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimManufacturersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

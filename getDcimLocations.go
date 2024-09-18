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

type locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous string  `json:"previous"`
	Results  []struct {
		CommonFieldsSlug
		Site struct {
			CommonFieldsSlug
		} `json:"site"`
		Parent struct {
			CommonFieldsSlug
			Depth int `json:"_depth"`
		} `json:"parent"`
		Status struct {
			ValueLabel
		} `json:"status"`
		Tenant struct {
			CommonFieldsSlug
		} `json:"tenant"`
		Description string `json:"description"`
		Tags        []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created     string `json:"created"`
		LastUpdated string `json:"last_updated"`
		RackCount   int    `json:"rack_count"`
		DeviceCount int    `json:"device_count"`
		Depth       int    `json:"_depth"`
	} `json:"results"`
}

// GetDcimLocationsCmd represents the getDcimLocations command
var GetDcimLocationsCmd = &cobra.Command{
	Use:   "getDcimLocations",
	Short: "GET a list of location objects",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of location objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(locations)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.locations")

		if responseObject.Count != 0 {
			color.Cyan("\n  Metropolis Locations: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    Metropolis Locations: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				color.Cyan("\tSlug: " + color.YellowString("%s", result.Slug))
				if result.Parent.Id != 0 {
					color.Cyan("\tParent: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Parent.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Parent.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Parent.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Parent.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Parent.Slug))
				} else {
					color.Cyan("\tParent: " + color.RedString("No parent entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Status.Value != "" {
					color.Cyan("\tStatus: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Status.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Status.Label))
				} else {
					color.Cyan("\tStatus: " + color.RedString("No status entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Tenant.Id != 0 {
					color.Cyan("\tTenant: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Tenant.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Tenant.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Tenant.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Tenant.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Tenant.Slug))
				} else {
					color.Cyan("\tTenant: " + color.RedString("No tenant entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%d", result.Tenant.Id))
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
				if result.RackCount != 0 {
					color.Cyan("\tRack Count: " + color.YellowString("%d", result.RackCount))
				} else {
					color.Cyan("\tRack Count: " + color.RedString("No rack count entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.DeviceCount != 0 {
					color.Cyan("\tDevice Count: " + color.YellowString("%d", result.DeviceCount))
				} else {
					color.Cyan("\tDevice Count: " + color.RedString("No device count entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Depth != 0 {
					color.Cyan("\tDepth: " + color.YellowString("%d\n", result.Depth))
				} else {
					color.Cyan("\tDepth: " + color.RedString("No depth entry found for ") + color.YellowString("%s\n", result.Display))
				}
			}
		} else {
			color.Cyan("  Metropolis Locations: " + color.RedString("No locations found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimLocationsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimLocationsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimLocationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimLocationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

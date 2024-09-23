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

package dcim

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type platforms struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		CommonFieldsSlug
		Manufacturer struct {
			CommonFieldsSlug
		} `json:"manufacturer"`
		ConfigTemplate struct {
			CommonFieldsNoSlug
		} `json:"config_template"`
		Description string `json:"description"`
		Tags        []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created             string `json:"created"`
		LastUpdated         string `json:"last_updated"`
		DeviceCount         int    `json:"device_count"`
		VirtualmachineCount int    `json:"virtualmachine_count"`
	} `json:"results"`
}

// GetDcimPlatformsCmd represents the getDcimPlatforms command
var GetDcimPlatformsCmd = &cobra.Command{
	Use:   "getDcimPlatforms",
	Short: "GET a list of platform objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of platform objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(platforms)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.platforms")

		if responseObject.Count != 0 {
			color.Cyan("\n  ABC Platforms: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    ABC Platform: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				color.Cyan("\tSlug: " + color.YellowString("%s", result.Slug))
				if result.Manufacturer.Id != 0 {
					color.Cyan("\tManufacturer: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Manufacturer.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Manufacturer.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Manufacturer.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Manufacturer.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Manufacturer.Slug))
				} else {
					color.Cyan("\tManufacturer: " + color.RedString("No manufacturer entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.ConfigTemplate.Id != 0 {
					color.Cyan("\tConfig Template: ")
					color.Cyan("\t  ID: " + color.YellowString("%s", result.ConfigTemplate.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.ConfigTemplate.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.ConfigTemplate.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.ConfigTemplate.Name))
				} else {
					color.Cyan("\tConfig Template: " + color.RedString("No config template entry found for ") + color.YellowString("%s", result.Display))
				}
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
				if result.DeviceCount != 0 {
					color.Cyan("\tDevice Count: " + color.YellowString("%d", result.DeviceCount))
				} else {
					color.Cyan("\tDevice Count: " + color.RedString("No device count entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.VirtualmachineCount != 0 {
					color.Cyan("\tVirtualmachine Count: " + color.YellowString("%s\n", result.VirtualmachineCount))
				} else {
					color.Cyan("\tVirtualmachine Count: " + color.RedString("No virtualmachine count entry found for ") + color.YellowString("%s\n", result.Display))
				}
			}
		} else {
			color.Cyan("  ABC Platforms: " + color.RedString("No platforms found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimPlatformsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimPlatformsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPlatformsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimPlatformsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimPlatformsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

type platformsByID struct {
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
}

// GetDcimPlatformsByIdCmd represents the getDcimPlatformsById command
var GetDcimPlatformsByIdCmd = &cobra.Command{
	Use:   "getDcimPlatformsById",
	Short: "GET an platform object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an platform object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(platformsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.platforms_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    ABC Platform: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			color.Cyan("\tSlug: " + color.YellowString("%s", responseObject.Slug))
			if responseObject.Manufacturer.Id != 0 {
				color.Cyan("\tManufacturer: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Manufacturer.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Manufacturer.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Manufacturer.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Manufacturer.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Manufacturer.Slug))
			} else {
				color.Cyan("\tManufacturer: " + color.RedString("No manufacturer entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.ConfigTemplate.Id != 0 {
				color.Cyan("\tConfig Template: ")
				color.Cyan("\t  ID: " + color.YellowString("%s", responseObject.ConfigTemplate.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.ConfigTemplate.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.ConfigTemplate.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.ConfigTemplate.Name))
			} else {
				color.Cyan("\tConfig Template: " + color.RedString("No config template entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			for _, tag := range responseObject.Tags {
				if tag.Id != 0 {
					color.Cyan("\tTags: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
					color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
			if responseObject.DeviceCount != 0 {
				color.Cyan("\tDevice Count: " + color.YellowString("%d", responseObject.DeviceCount))
			} else {
				color.Cyan("\tDevice Count: " + color.RedString("No device count entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.VirtualmachineCount != 0 {
				color.Cyan("\tVirtualmachine Count: " + color.YellowString("%s\n", responseObject.VirtualmachineCount))
			} else {
				color.Cyan("\tVirtualmachine Count: " + color.RedString("No virtualmachine count entry found for ") + color.YellowString("%s\n", responseObject.Display))
			}
		} else {
			color.Red("  Doh! No platform object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimPlatformsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimPlatformsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPlatformsByIdCmd", err)
	}

	GetDcimPlatformsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the platform object")
	err = GetDcimPlatformsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPlatformsByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimPlatformsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimPlatformsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

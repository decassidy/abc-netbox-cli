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

type deviceBaysByID struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Device  struct {
		CommonFieldsNoSlug
	} `json:"device"`
	Name            string `json:"name"`
	Label           string `json:"label"`
	Description     string `json:"description"`
	InstalledDevice struct {
		CommonFieldsNoSlug
	} `json:"installed_device"`
	Tags        []CommonFieldsSlug `json:"tags"`
	Created     string             `json:"created"`
	LastUpdated string             `json:"last_updated"`
}

// GetDcimDeviceBaysByIdCmd represents the getDcimDeviceBaysById command
var GetDcimDeviceBaysByIdCmd = &cobra.Command{
	Use:   "getDcimDeviceBaysById",
	Short: "GET an device bay object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an device bay object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(deviceBaysByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.device_bays_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    ABC Device Bay: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.Device.Id != 0 {
				color.Cyan("\tDevice: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Device.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Device.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Device.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Device.Name))
			} else {
				color.Cyan("\tDevice: " + color.RedString("No device entry for device bay: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tName: " + color.RedString("No name entry for device bay: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Label != "" {
				color.Cyan("\tLabel: " + color.YellowString("%s", responseObject.Label))
			} else {
				color.Cyan("\tLabel: " + color.RedString("No label entry for device bay: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry for device bay: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.InstalledDevice.Id != 0 {
				color.Cyan("\tInstalled Device: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.InstalledDevice.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.InstalledDevice.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.InstalledDevice.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.InstalledDevice.Name))
			} else {
				color.Cyan("\tInstalled Device: " + color.RedString("No installed device entry for device bay: ") + color.YellowString("%s", responseObject.Display))
			}
			for _, tag := range responseObject.Tags {
				if tag.Id != 0 {
					color.Cyan("\tTags: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags entry for device bay: ") + color.YellowString("%s", responseObject.Display))
				}
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s\n", responseObject.LastUpdated))
		} else {
			color.Red("Doh! No Device Bay Object was found with ID: " + color.YellowString("%d", id))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimDeviceBaysByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimDeviceBaysByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimDeviceBaysByIdCmd", err)
	}

	GetDcimDeviceBaysByIdCmd.Flags().IntVarP(&id, "id", "", 0, "Device Bay ID")
	err = GetDcimDeviceBaysByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimDeviceBaysByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimDeviceBaysByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

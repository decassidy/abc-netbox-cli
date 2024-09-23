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

type deviceBayTemplatesByID struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	Display    string `json:"display"`
	DeviceType struct {
		Id           int              `json:"id"`
		Url          string           `json:"url"`
		Display      string           `json:"display"`
		Manufacturer CommonFieldsSlug `json:"manufacturer"`
		Model        string           `json:"model"`
		Slug         string           `json:"slug"`
	} `json:"device_type"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

// GetDcimDeviceBayTemplatesByIdCmd represents the getDcimDeviceBayTemplatesById command
var GetDcimDeviceBayTemplatesByIdCmd = &cobra.Command{
	Use:   "getDcimDeviceBayTemplatesById",
	Short: "GET an device bay template object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an device bay template object by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(deviceBayTemplatesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.device_bay_templates_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    ABC Device Bay Template: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.DeviceType.Id != 0 {
				color.Cyan("\tDevice Type: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.DeviceType.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.DeviceType.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.DeviceType.Display))
				color.Cyan("\t  Manufacturer: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.DeviceType.Manufacturer.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Slug))
			} else {
				color.Cyan("\t  Manufacturer: " + color.RedString("No manufacturer entry for device bay template: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tName: " + color.RedString("No name entry for device bay template: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Label != "" {
				color.Cyan("\tLabel: " + color.YellowString("%s", responseObject.Label))
			} else {
				color.Cyan("\tLabel: " + color.RedString("No label entry for device bay template: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry for device bay template: ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s\n", responseObject.LastUpdated))
		} else {
			color.Cyan("  ABC Device Bay Templates: " + color.RedString("No device bay template entries found on server for ID: %d. Exiting...\n", id))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimDeviceBayTemplatesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimDeviceBayTemplatesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	GetDcimDeviceBayTemplatesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the device bay template")
	err = GetDcimDeviceBayTemplatesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimDeviceBayTemplatesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimDeviceBayTemplatesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

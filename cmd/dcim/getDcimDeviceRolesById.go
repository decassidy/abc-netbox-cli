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

type deviceRolesByID struct {
	CommonFieldsSlug
	Color          string `json:"color"`
	VmRole         bool   `json:"vm_role"`
	ConfigTemplate struct {
		CommonFieldsNoSlug
	} `json:"config_template"`
	Description         string             `json:"description"`
	Tags                []CommonFieldsSlug `json:"tags"`
	Created             string             `json:"created"`
	LastUpdated         string             `json:"last_updated"`
	DeviceCount         int                `json:"device_count"`
	VirtualmachineCount int                `json:"virtualmachine_count"`
}

// GetDcimDeviceRolesByIdCmd represents the getDcimDeviceRolesById command
var GetDcimDeviceRolesByIdCmd = &cobra.Command{
	Use:   "getDcimDeviceRolesById",
	Short: "GET an device role object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an device role object by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(deviceRolesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.device_roles_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    ABC Device Role: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			color.Cyan("\tSlug: " + color.YellowString("%s", responseObject.Slug))
			if responseObject.Color != "" {
				color.Cyan("\tColor: " + color.YellowString("%s", responseObject.Color))
			} else {
				color.Cyan("\tColor: " + color.RedString("No color entry for device role: ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tVM Role: " + color.YellowString("%v", responseObject.VmRole))
			if responseObject.ConfigTemplate.Id != 0 {
				color.Cyan("\tConfig Template: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.ConfigTemplate.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.ConfigTemplate.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.ConfigTemplate.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.ConfigTemplate.Name))
			} else {
				color.Cyan("\tConfig Template: " + color.RedString("No config template entry for device role: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry for device role: ") + color.YellowString("%s", responseObject.Display))
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
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
			if responseObject.DeviceCount != 0 {
				color.Cyan("\tDevice Count: " + color.YellowString("%d", responseObject.DeviceCount))
			} else {
				color.Cyan("\tDevice Count: " + color.RedString("No device count entry for device bay: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.VirtualmachineCount != 0 {
				color.Cyan("\tVirtual Machine Count: " + color.YellowString("%d", responseObject.VirtualmachineCount))
			} else {
				color.Cyan("\tVirtual Machine Count: " + color.RedString("No virtual machine count entry for device bay: ") + color.YellowString("%s\n", responseObject.Display))
			}
		} else {
			color.Cyan("  ABC Device Role: " + color.RedString("No device role entries found on server for ID: %d. Exiting...\n", id))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimDeviceRolesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimDeviceRolesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	GetDcimDeviceRolesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "Device Role ID")
	err = GetDcimDeviceRolesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimDeviceRolesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimDeviceRolesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

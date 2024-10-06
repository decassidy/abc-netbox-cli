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

type deviceTypesByID struct {
	Id           int    `json:"id"`
	Url          string `json:"url"`
	Display      string `json:"display"`
	Manufacturer struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
	} `json:"manufacturer"`
	DefaultPlatform struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
	} `json:"default_platform"`
	Model                  string  `json:"model"`
	Slug                   string  `json:"slug"`
	PartNumber             string  `json:"part_number"`
	UHeight                float32 `json:"u_height"`
	ExcludeFromUtilization bool    `json:"exclude_from_utilization"`
	IsFullDepth            bool    `json:"is_full_depth"`
	SubdeviceRole          struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"subdevice_role"`
	Airflow struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"airflow"`
	Weight     float32 `json:"weight"`
	WeightUnit struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"weight_unit"`
	FrontImage  string `json:"front_image"`
	RearImage   string `json:"rear_image"`
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
	Created                        string `json:"created"`
	LastUpdated                    string `json:"last_updated"`
	DeviceCount                    int    `json:"device_count"`
	ConsolePortTemplateCount       int    `json:"console_port_template_count"`
	ConsoleServerPortTemplateCount int    `json:"console_server_port_template_count"`
	PowerPortTemplateCount         int    `json:"power_port_template_count"`
	PowerOutletTemplateCount       int    `json:"power_outlet_template_count"`
	InterfaceTemplateCount         int    `json:"interface_template_count"`
	FrontPortTemplateCount         int    `json:"front_port_template_count"`
	RearPortTemplateCount          int    `json:"rear_port_template_count"`
	DeviceBayTemplateCount         int    `json:"device_bay_template_count"`
	ModuleBayTemplateCount         int    `json:"module_bay_template_count"`
	InventoryItemTemplateCount     int    `json:"inventory_item_template_count"`
}

// GetDcimDeviceTypesByIdCmd represents the getDcimDeviceTypesById command
var GetDcimDeviceTypesByIdCmd = &cobra.Command{
	Use:   "getDcimDeviceTypesById",
	Short: "GET an device type object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an device type object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(deviceTypesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.device_types_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    ABC Device Type: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.Manufacturer.Id != 0 {
				color.Cyan("\tManufacturer: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Manufacturer.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Manufacturer.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Manufacturer.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Manufacturer.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Manufacturer.Slug))
			} else {
				color.Cyan("\tManufacturer: " + color.RedString("No manufacturer found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.DefaultPlatform.Id != 0 {
				color.Cyan("\tDefault Platform: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.DefaultPlatform.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.DefaultPlatform.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.DefaultPlatform.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.DefaultPlatform.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.DefaultPlatform.Slug))
			} else {
				color.Cyan("\tDefault Platform: " + color.RedString("No default platform found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.Model != "" {
				color.Cyan("\tModel: " + color.YellowString("%s", responseObject.Model))
			} else {
				color.Cyan("\tModel: " + color.RedString("No model found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.Slug != "" {
				color.Cyan("\tSlug: " + color.YellowString("%s", responseObject.Slug))
			} else {
				color.Cyan("\tSlug: " + color.RedString("No slug found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.PartNumber != "" {
				color.Cyan("\tPart Number: " + color.YellowString("%s", responseObject.PartNumber))
			} else {
				color.Cyan("\tPart Number: " + color.RedString("No part number found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.UHeight != 0.0 {
				color.Cyan("\tUnit Height: " + color.YellowString("%v", responseObject.UHeight))
			} else {
				color.Cyan("\tUnit Height: " + color.RedString("No unit height found for device: ") + color.YellowString(responseObject.Display))
			}
			color.Cyan("\tExcluded From Utilization: " + color.YellowString("%v", responseObject.ExcludeFromUtilization))
			color.Cyan("\tIs Full Depth: " + color.YellowString("%v", responseObject.IsFullDepth))
			if responseObject.SubdeviceRole.Value != "" {
				color.Cyan("\tSubdevice Role: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.SubdeviceRole.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.SubdeviceRole.Label))
			} else {
				color.Cyan("\tSubdevice Role: " + color.RedString("No subdevice role found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.Airflow.Value != "" {
				color.Cyan("\tAirflow: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Airflow.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Airflow.Label))
			} else {
				color.Cyan("\tAirflow: " + color.RedString("No airflow found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.Weight != 0.0 {
				color.Cyan("\tWeight: " + color.YellowString("%d", responseObject.Weight))
			} else {
				color.Cyan("\tWeight: " + color.RedString("No weight found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.WeightUnit.Value != "" {
				color.Cyan("\tWeight Unit: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.WeightUnit.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.WeightUnit.Label))
			} else {
				color.Cyan("\tWeight Unit: " + color.RedString("No Weight unit found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.FrontImage != "" {
				color.Cyan("\tFront Image: " + color.YellowString("%s", responseObject.FrontImage))
			} else {
				color.Cyan("\tFront Image: " + color.RedString("No front image found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.RearImage != "" {
				color.Cyan("\tRear Image: " + color.YellowString("%s", responseObject.RearImage))
			} else {
				color.Cyan("\tRear Image: " + color.RedString("No rear image found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description found for device: ") + color.YellowString(responseObject.Description))
			}
			if responseObject.Comments != "" {
				color.Cyan("\tComments: " + color.YellowString("%s", responseObject.Comments))
			} else {
				color.Cyan("\tComments: " + color.RedString("No comments found for device: ") + color.YellowString(responseObject.Comments))
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
					color.Cyan("\tTags: " + color.RedString("No tags found for device: ") + color.YellowString(responseObject.Display))
				}
			}

			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
			if responseObject.DeviceCount != 0 {
				color.Cyan("\tDevice Count: " + color.YellowString("%d", responseObject.DeviceCount))
			} else {
				color.Cyan("\tDevice Count: " + color.RedString("No device count found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.ConsolePortTemplateCount != 0 {
				color.Cyan("\tConsole Port Template Count: " + color.YellowString("%d", responseObject.ConsolePortTemplateCount))
			} else {
				color.Cyan("\tConsole Port Template Count: " + color.RedString("No console port template count found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.ConsoleServerPortTemplateCount != 0 {
				color.Cyan("\tConsole Server Port Template Count: " + color.YellowString("%d", responseObject.ConsoleServerPortTemplateCount))
			} else {
				color.Cyan("\tConsole Server Port Template Count: " + color.RedString("No console server port template count found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.PowerPortTemplateCount != 0 {
				color.Cyan("\tPower Port Template Count: " + color.YellowString("%d", responseObject.PowerPortTemplateCount))
			} else {
				color.Cyan("\tPower Port Template Count: " + color.RedString("No power port template count found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.PowerOutletTemplateCount != 0 {
				color.Cyan("\tPower Outlet Template Count: " + color.YellowString("%d", responseObject.PowerOutletTemplateCount))
			} else {
				color.Cyan("\tPower Outlet Template Count: " + color.RedString("No power outlet template count found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.InterfaceTemplateCount != 0 {
				color.Cyan("\tInterface Template Count: " + color.YellowString("%d", responseObject.InterfaceTemplateCount))
			} else {
				color.Cyan("\tInterface Template Count: " + color.RedString("No interface template count found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.FrontPortTemplateCount != 0 {
				color.Cyan("\tFront Port Template Count: " + color.YellowString("%d", responseObject.FrontPortTemplateCount))
			} else {
				color.Cyan("\tFront Port Template Count: " + color.RedString("No front port template count found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.RearPortTemplateCount != 0 {
				color.Cyan("\tRear Port Template Count: " + color.YellowString("%d", responseObject.RearPortTemplateCount))
			} else {
				color.Cyan("\tRear Port Template Count: " + color.RedString("No rear port template count found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.DeviceBayTemplateCount != 0 {
				color.Cyan("\tDevice Bay Template Count: " + color.YellowString("%d", responseObject.DeviceBayTemplateCount))
			} else {
				color.Cyan("\tDevice Bay Template Count: " + color.RedString("No device bay template count found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.ModuleBayTemplateCount != 0 {
				color.Cyan("\tModule Bay Template Count: " + color.YellowString("%d", responseObject.ModuleBayTemplateCount))
			} else {
				color.Cyan("\tModule Bay Template Count: " + color.RedString("No module bay template count found for device: ") + color.YellowString(responseObject.Display))
			}
			if responseObject.InventoryItemTemplateCount != 0 {
				color.Cyan("\tInventory Item Template Count: " + color.YellowString("%d\n", responseObject.InventoryItemTemplateCount))
			} else {
				color.Cyan("\tInventory Item Template Count: " + color.RedString("No inventory template count found for device: ") + color.YellowString(responseObject.Display+"\n"))
			}
		} else {
			color.Red("Doh! No device type object was found with ID: " + color.YellowString("%s\n", id))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimDeviceTypesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimDeviceTypesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	GetDcimDeviceTypesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the device type")
	err = GetDcimDeviceTypesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimDeviceTypesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimDeviceTypesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

// GetDcimDeviceByQueryCmd represents the getDcimDeviceByQuery command
var GetDcimDeviceByQueryCmd = &cobra.Command{
	Use:   "getDcimDeviceByQuery",
	Short: "GET a list of device objects by string query",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of device objects by string query`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(devices)
		ApiConnectionQuery(responseObject, "GET", "cmd.dcim.dcim_api_url.devices_id")

		if responseObject.Count > 0 {
			color.Cyan("\n  Total Metropolis Devices Return from Query: "+color.YellowString("%d"), responseObject.Count)
			for _, device := range responseObject.Results {
				display := fmt.Sprintf("    Metropolis Devices: %s\n", color.YellowString(device.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", device.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", device.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", device.Display))
				color.Cyan("\tName: " + color.YellowString("%s", device.Name))
				if device.DeviceType.Id != 0 {
					color.Cyan("\tDevice Type: \n")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.DeviceType.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.DeviceType.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.DeviceType.Display))
				} else {
					color.Cyan("\tDevice Type: No device type entry found for device: " + color.YellowString("%s", device.Name))
				}
				if device.DeviceType.Manufacturer.Id > 0 {
					color.Cyan("\t  Manufacturer: \n")
					color.Cyan("\t    ID: " + color.YellowString("%d", device.DeviceType.Manufacturer.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", device.DeviceType.Manufacturer.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", device.DeviceType.Manufacturer.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", device.DeviceType.Manufacturer.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%s", device.DeviceType.Manufacturer.Slug))
					color.Cyan("\t    Description: " + color.YellowString("%s", device.DeviceType.Manufacturer.Description))
					color.Cyan("\t    Device Type Count: " + color.YellowString("%d", device.DeviceType.Manufacturer.DevicetypeCount))
					color.Cyan("\t  Model: " + color.YellowString("%s", device.DeviceType.Model))
					color.Cyan("\t  Slug: " + color.YellowString("%s", device.DeviceType.Slug))
					color.Cyan("\t  Description: " + color.YellowString("%s", device.DeviceType.Description))
					color.Cyan("\t  Device Count: " + color.YellowString("%d", device.DeviceType.DeviceCount))
				} else {
					color.Cyan("\t  Manufacturer: No manufacturer entry found for device: " + color.YellowString("%s", device.Name))
				}
				if device.Role.Id != 0 {
					color.Cyan("\tRole: \n")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.Role.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.Role.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.Role.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", device.Role.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", device.Role.Slug))
					color.Cyan("\t  Description: " + color.YellowString("%s", device.Role.Description))
					color.Cyan("\t  Device Count: " + color.YellowString("%d", device.Role.DeviceCount))
					color.Cyan("\t  Virtual Machine Count: " + color.YellowString("%d", device.Role.VirtualmachineCount))
				} else {
					color.Cyan("\tRole: No role entry found for device: " + color.YellowString("%s", device.Name))
				}
				if device.Tenant.Id != 0 {
					color.Cyan("\tTenant: \n")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.Tenant.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.Tenant.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.Tenant.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", device.Tenant.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", device.Tenant.Slug))
				} else {
					color.Cyan("\tTenant: " + color.RedString("No tenant entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Platform.Id != 0 {
					color.Cyan("\tPlatform: \n")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.Platform.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.Platform.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.Platform.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", device.Platform.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", device.Platform.Slug))
				} else {
					color.Cyan("\tPlatform: " + color.RedString("No platform entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Location.Id != 0 {
					color.Cyan("\tLocation: \n")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.Location.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.Location.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.Location.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", device.Location.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", device.Location.Slug))
					color.Cyan("\t  Depth: " + color.YellowString("%d", device.Location.Depth))
				} else {
					color.Cyan("\tLocation: " + color.RedString("No location entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Rack.Id != 0 {
					color.Cyan("\tRack: \n")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.Rack.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.Rack.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.Rack.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", device.Rack.Name))
				} else {
					color.Cyan("\tRack: " + color.RedString("No rack entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Position != 0 {
					color.Cyan("\tPosition: " + color.YellowString("%.2f", device.Position))
				} else {
					color.Cyan("\tPosition: " + color.RedString("No position entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Face.Value != "" {
					color.Cyan("\tFace: ")
					color.Cyan("\t  Value: " + color.YellowString("%v", device.Face.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", device.Face.Label))
				} else {
					color.Cyan("\tFace: " + color.RedString("No face entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Latitude != 0 {
					color.Cyan("\tLatitude: " + color.YellowString("%v", device.Latitude))
				} else {
					color.Cyan("\tLatitude: " + color.RedString("No latitude entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Longitude != 0 {
					color.Cyan("\tLongitude: " + color.YellowString("%v", device.Longitude))
				} else {
					color.Cyan("\tLongitude: " + color.RedString("No longitude entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.ParentDevice.Id != 0 {
					color.Cyan("\tParent Device: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.ParentDevice.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.ParentDevice.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.ParentDevice.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", device.ParentDevice.Name))
				} else {
					color.Cyan("\tParent Device: " + color.RedString("No parent device entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Status.Value != "" {
					color.Cyan("\tStatus: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", device.Status.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", device.Status.Label))
				} else {
					color.Cyan("\tStatus: " + color.RedString("No status entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Airflow.Value != "" {
					color.Cyan("\tAirflow: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", device.Airflow.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", device.Airflow.Label))
				} else {
					color.Cyan("\tAirflow: " + color.RedString("No airflow entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.PrimaryIp.Id != 0 {
					color.Cyan("\tPrimary IP: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.PrimaryIp.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.PrimaryIp.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.PrimaryIp.Display))
					color.Cyan("\t  Family: ")
					color.Cyan("\t    Value: " + color.YellowString("%d", device.PrimaryIp.Family.Value))
					color.Cyan("\t    Label: " + color.YellowString("%s", device.PrimaryIp.Family.Label))
					color.Cyan("\t  Address: " + color.YellowString("%s", device.PrimaryIp.Address))
				} else {
					color.Cyan("\tPrimary IP: " + color.RedString("No primary ip entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.PrimaryIp4.Id != 0 {
					color.Cyan("\tPrimary IPv4: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.PrimaryIp4.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.PrimaryIp4.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.PrimaryIp4.Display))
					color.Cyan("\t  Family: ")
					color.Cyan("\t    Value: " + color.YellowString("%d", device.PrimaryIp4.Family.Value))
					color.Cyan("\t    Label: " + color.YellowString("%s", device.PrimaryIp4.Family.Label))
					color.Cyan("\t  Address: " + color.YellowString("%s", device.PrimaryIp4.Address))
				} else {
					color.Cyan("\tPrimary IP: " + color.RedString("No primary ip entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.PrimaryIp6.Id != 0 {
					color.Cyan("\tPrimary IPv6: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.PrimaryIp6.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.PrimaryIp6.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.PrimaryIp6.Display))
					color.Cyan("\t  Family: " + color.YellowString("%d", device.PrimaryIp6.Family))
					color.Cyan("\t  Address: " + color.YellowString("%s", device.PrimaryIp6.Address))
				} else {
					color.Cyan("\tPrimary IPv6: " + color.RedString("No primary ipv6 entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.OobIp.Id != 0 {
					color.Cyan("\tOOB IP: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.OobIp.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.OobIp.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.OobIp.Display))
					color.Cyan("\t  Family: " + color.YellowString("%d", device.OobIp.Family))
					color.Cyan("\t  Address: " + color.YellowString("%s", device.OobIp.Address))
				} else {
					color.Cyan("\tOOB IP: " + color.RedString("No oop ip entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.VirtualChassis.Id != 0 {
					color.Cyan("\tVirtual Chassis: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.VirtualChassis.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.VirtualChassis.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.VirtualChassis.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", device.VirtualChassis.Name))
					color.Cyan("\t  Master: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", device.VirtualChassis.Master.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", device.VirtualChassis.Master.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", device.VirtualChassis.Master.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", device.VirtualChassis.Master.Name))
				} else {
					color.Cyan("\tVirtual Chassis: " + color.RedString("No virtual chassis entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.VcPosition != 0 {
					color.Cyan("\tVC Position: " + color.YellowString("%d", device.VcPosition))
				} else {
					color.Cyan("\tVC Position: " + color.RedString("No vc position entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", device.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.Comments != "" {
					color.Cyan("\tComments: " + color.YellowString("%s", device.Description))
				} else {
					color.Cyan("\tComments: " + color.RedString("No comments entry found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.ConfigTemplate.Id != 0 {
					color.Cyan("\tConfig Template: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", device.ConfigTemplate.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", device.ConfigTemplate.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", device.ConfigTemplate.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", device.ConfigTemplate.Name))
				} else {
					color.Cyan("\tConfig Template: " + color.RedString("No config template entry found for device: ") + color.YellowString("%s", device.Name))
				}
				color.Cyan("\tConfig Context: " + color.YellowString("%v", device.ConfigContext))
				if device.LocalContextData != "" {
					color.Cyan("\tLocal Context Data: " + color.YellowString("%s", device.LocalContextData))
				} else {
					color.Cyan("\tLocal Context Data: " + color.RedString("No local context data entry found for device: ") + color.YellowString("%s", device.Name))
				}
				for _, tag := range device.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
						color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
						color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
						color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
						color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
						color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags entry found for device: ") + color.YellowString("%s", device.Name))
					}
				}

				color.Cyan("\tCreated: " + color.YellowString("%s", device.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", device.LastUpdated))
				if device.ConsolePortCount != 0 {
					color.Cyan("\tConsole Port Count: " + color.YellowString("%d", device.ConsolePortCount))
				} else {
					color.Cyan("\tConsole Port Count: " + color.RedString("No console port count found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.ConsoleServerPortCount != 0 {
					color.Cyan("\tConsole Server Port Count: " + color.YellowString("%d", device.ConsolePortCount))
				} else {
					color.Cyan("\tConsole Server Port Count: " + color.RedString("No console server port count found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.PowerPortCount != 0 {
					color.Cyan("\tPower Port Count: " + color.YellowString("%d", device.PowerPortCount))
				} else {
					color.Cyan("\tPower Port Count: " + color.RedString("No power port count found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.PowerOutletCount != 0 {
					color.Cyan("\tPower Outlet Count: " + color.YellowString("%d", device.PowerOutletCount))
				} else {
					color.Cyan("\tPower Outlet Count: " + color.RedString("No power outlet count found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.InterfaceCount != 0 {
					color.Cyan("\tInterface Count: " + color.YellowString("%d", device.InterfaceCount))
				} else {
					color.Cyan("\tPower Outlet Count: " + color.RedString("No interface count found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.FrontPortCount != 0 {
					color.Cyan("\tFront Port Count: " + color.YellowString("%d", device.FrontPortCount))
				} else {
					color.Cyan("\tFront Port Count: " + color.RedString("No front port count found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.RearPortCount != 0 {
					color.Cyan("\tRear Port Count: " + color.YellowString("%d", device.RearPortCount))
				} else {
					color.Cyan("\tRear Port Count: " + color.RedString("No rear port count found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.DeviceBayCount != 0 {
					color.Cyan("\tDevice Bay Count: " + color.YellowString("%d", device.DeviceBayCount))
				} else {
					color.Cyan("\tDevice Bay Count: " + color.RedString("No device bay count found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.ModuleBayCount != 0 {
					color.Cyan("\tModule Bay Count: " + color.YellowString("%d", device.ModuleBayCount))
				} else {
					color.Cyan("\tModule Bay Count: " + color.RedString("No module bay count found for device: ") + color.YellowString("%s", device.Name))
				}
				if device.InventoryItemCount != 0 {
					color.Cyan("\tInventory Item Count: " + color.YellowString("%d\n", device.InventoryItemCount))
				} else {
					color.Cyan("\tInventory Item Count: " + color.RedString("No inventory item count found for device: ") + color.YellowString("%s\n", device.Name))
				}
			}
		} else {
			color.Cyan("  Metropolis Device: " + color.RedString("No device entries found on server with serial number: "+color.YellowString("%s", serial)+color.RedString(" Exiting...\n")))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimDeviceByQueryCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimDeviceByQueryCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for GetDcimDeviceByQueryCmd", err)
	}

	GetDcimDeviceByQueryCmd.Flags().StringVarP(&query, "query", "q", "", "string query of object you want to get")
	err = GetDcimDeviceByQueryCmd.MarkFlagRequired("query")
	if err != nil {
		log.Fatalf("Error marking query flag as required: %s - for GetDcimDeviceByQueryCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimDeviceByQueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimDeviceByQueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

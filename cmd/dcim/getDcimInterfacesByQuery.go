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

// GetDcimInterfacesByQueryCmd represents the getDcimInterfacesByQuery command
var GetDcimInterfacesByQueryCmd = &cobra.Command{
	Use:   "getDcimInterfacesByQuery",
	Short: "GET a interface object(s) by string query",
	Long: `
Netbox Automation Tools:
  GET a interface object(s) by string query`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionQuery(responseObjectInterfaces, "GET", "cmd.dcim.dcim_api_url.interfaces_id")

		if responseObjectInterfaces.Count != 0 {
			color.Cyan("\n  Total ABC Interfaces: "+color.YellowString("%d, for Query: %s"), responseObjectInterfaces.Count, query)
			for _, result := range responseObjectInterfaces.Results {
				display := fmt.Sprintf("    ABC Interface Name: %s\n", color.YellowString(result.Display)+color.CyanString(" Device: ")+color.YellowString(result.Device.Name))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.Device.Id > 0 {
					color.Cyan("\tDevice: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Device.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Device.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Device.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Device.Name))
				} else {
					color.Cyan("\tDevice: " + color.RedString("No device found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Module.Id > 0 {
					color.Cyan("\tModule: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Module.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Module.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Module.Display))
					color.Cyan("\t  Device: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Module.Device.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Module.Device.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Module.Device.Display))
					color.Cyan("\t    Device Type: " + color.YellowString("%s", result.Module.Device.DeviceType))
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.DeviceType.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.DeviceType.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.DeviceType.Display))
					color.Cyan("\t      Manufacturer: ")
					color.Cyan("\t        ID: " + color.YellowString("%d", result.Module.Device.DeviceType.Manufacturer.Id))
					color.Cyan("\t        URL: " + color.YellowString("%s", result.Module.Device.DeviceType.Manufacturer.Url))
					color.Cyan("\t        Display: " + color.YellowString("%s", result.Module.Device.DeviceType.Manufacturer.Display))
					color.Cyan("\t        Name: " + color.YellowString("%s", result.Module.Device.DeviceType.Manufacturer.Name))
					color.Cyan("\t        Slug: " + color.YellowString("%s", result.Module.Device.DeviceType.Manufacturer.Slug))
					color.Cyan("\t        Description: " + color.YellowString("%s", result.Module.Device.DeviceType.Manufacturer.Description))
					color.Cyan("\t    Model: " + color.YellowString("%s", result.Module.Device.DeviceType.Model))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.Module.Device.DeviceType.Slug))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Module.Device.DeviceType.Description))
					color.Cyan("\t    Device Count: " + color.YellowString("%d", result.Module.Device.DeviceType.DeviceCount))
					color.Cyan("\t    Role: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.Role.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.Role.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.Role.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Module.Device.Role.Name))
					color.Cyan("\t      Slug: " + color.YellowString("%s", result.Module.Device.Role.Slug))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.Role.Description))
					color.Cyan("\t      Device Count: " + color.YellowString("%d", result.Module.Device.Role.DeviceCount))
					color.Cyan("\t      Virtual Machine Count: " + color.YellowString("%d", result.Module.Device.Role.VirtualmachineCount))
					color.Cyan("\t    Tenant: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.Tenant.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.Tenant.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.Tenant.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Module.Device.Tenant.Name))
					color.Cyan("\t      Slug: " + color.YellowString("%s", result.Module.Device.Tenant.Slug))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.Tenant.Description))
					color.Cyan("\t    Platform: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.Platform.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.Platform.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.Platform.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Module.Device.Platform.Name))
					color.Cyan("\t      Slug: " + color.YellowString("%s", result.Module.Device.Platform.Slug))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.Platform.Description))
					color.Cyan("\t      Device Count: " + color.YellowString("%d", result.Module.Device.Platform.DeviceCount))
					color.Cyan("\t      Virtual Machine Count: " + color.YellowString("%d", result.Module.Device.Platform.VirtualmachineCount))
					color.Cyan("\t    Serial: " + color.YellowString("%s", result.Module.Device.Serial))
					color.Cyan("\t    Asset Tag: " + color.YellowString("%s", result.Module.Device.AssetTag))
					color.Cyan("\t    Site: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.Site.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.Site.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.Site.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Module.Device.Site.Name))
					color.Cyan("\t      Slug: " + color.YellowString("%s", result.Module.Device.Site.Slug))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.Site.Description))
					color.Cyan("\t    Location: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.Location.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.Location.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.Location.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Module.Device.Location.Name))
					color.Cyan("\t      Slug: " + color.YellowString("%s", result.Module.Device.Location.Slug))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.Location.Description))
					color.Cyan("\t      Rack Count: " + color.YellowString("%d", result.Module.Device.Location.RackCount))
					color.Cyan("\t      Depth: " + color.YellowString("%d", result.Module.Device.Location.Depth))
					color.Cyan("\t    Rack: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.Rack.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.Rack.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.Rack.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Module.Device.Rack.Name))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.Rack.Description))
					color.Cyan("\t    Position: " + color.YellowString("%d", result.Module.Device.Position))
					color.Cyan("\t    Face: ")
					color.Cyan("\t      Value: " + color.YellowString("%s", result.Module.Device.Face.Value))
					color.Cyan("\t      Label: " + color.YellowString("%s", result.Module.Device.Face.Label))
					color.Cyan("\t    Latitude: " + color.YellowString("%f", result.Module.Device.Latitude))
					color.Cyan("\t    Longitude: " + color.YellowString("%f", result.Module.Device.Longitude))
					color.Cyan("\t    Parent Device: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.ParentDevice.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.ParentDevice.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.ParentDevice.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Module.Device.ParentDevice.Name))
					color.Cyan("\t    Status: ")
					color.Cyan("\t      Value: " + color.YellowString("%s", result.Module.Device.Status.Value))
					color.Cyan("\t      Label: " + color.YellowString("%s", result.Module.Device.Status.Label))
					color.Cyan("\t    Airflow: ")
					color.Cyan("\t      Value: " + color.YellowString("%s", result.Module.Device.Airflow.Value))
					color.Cyan("\t      Label: " + color.YellowString("%s", result.Module.Device.Airflow.Label))
					color.Cyan("\t    Primary IP: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.PrimaryIp.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.PrimaryIp.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.PrimaryIp.Display))
					color.Cyan("\t      Family: ")
					color.Cyan("\t        Value: " + color.YellowString("%d", result.Module.Device.PrimaryIp.Family.Value))
					color.Cyan("\t        Label: " + color.YellowString("%s", result.Module.Device.PrimaryIp.Family.Label))
					color.Cyan("\t      Address: " + color.YellowString("%s", result.Module.Device.PrimaryIp.Address))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.PrimaryIp.Description))
					color.Cyan("\t    Primary IPv4: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.PrimaryIp4.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.PrimaryIp4.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.PrimaryIp4.Display))
					color.Cyan("\t      Family: ")
					color.Cyan("\t        Value: " + color.YellowString("%d", result.Module.Device.PrimaryIp4.Family.Value))
					color.Cyan("\t        Label: " + color.YellowString("%s", result.Module.Device.PrimaryIp4.Family.Label))
					color.Cyan("\t      Address: " + color.YellowString("%s", result.Module.Device.PrimaryIp4.Address))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.PrimaryIp4.Description))
					color.Cyan("\t    Primary IPv6: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.PrimaryIp6.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.PrimaryIp6.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.PrimaryIp6.Display))
					color.Cyan("\t      Family: ")
					color.Cyan("\t        Value: " + color.YellowString("%d", result.Module.Device.PrimaryIp6.Family.Value))
					color.Cyan("\t        Label: " + color.YellowString("%s", result.Module.Device.PrimaryIp6.Family.Label))
					color.Cyan("\t      Address: " + color.YellowString("%s", result.Module.Device.PrimaryIp6.Address))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.PrimaryIp6.Description))
					color.Cyan("\t    OOB IP: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.OobIp.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.OobIp.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.OobIp.Display))
					color.Cyan("\t      Family: ")
					color.Cyan("\t        Value: " + color.YellowString("%d", result.Module.Device.OobIp.Family.Value))
					color.Cyan("\t        Label: " + color.YellowString("%s", result.Module.Device.OobIp.Family.Label))
					color.Cyan("\t      Address: " + color.YellowString("%s", result.Module.Device.OobIp.Address))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.OobIp.Description))
					color.Cyan("\t    Cluster: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.Cluster.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.Cluster.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.Cluster.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Module.Device.Cluster.Name))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.Cluster.Description))
					color.Cyan("\t    Virtual Chassis: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.VirtualChassis.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.VirtualChassis.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.VirtualChassis.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Module.Device.VirtualChassis.Name))
					color.Cyan("\t      Master: " + color.YellowString("%s", result.Module.Device.VirtualChassis.Name))
					color.Cyan("\t        ID: " + color.YellowString("%d", result.Module.Device.VirtualChassis.Master.Id))
					color.Cyan("\t        URL: " + color.YellowString("%s", result.Module.Device.VirtualChassis.Master.Url))
					color.Cyan("\t        Display: " + color.YellowString("%s", result.Module.Device.VirtualChassis.Master.Display))
					color.Cyan("\t        Name: " + color.YellowString("%s", result.Module.Device.VirtualChassis.Master.Name))
					color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.VirtualChassis.Description))
					color.Cyan("\t      Member Count: " + color.YellowString("%d", result.Module.Device.VirtualChassis.MemberCount))
					color.Cyan("\t    VC Position: " + color.YellowString("%d", result.Module.Device.VcPosition))
					color.Cyan("\t    VC Priority: " + color.YellowString("%d", result.Module.Device.VcPriority))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Module.Device.Description))
					color.Cyan("\t    Comments: " + color.YellowString("%s", result.Module.Device.Comments))
					if result.Module.Device.ConfigTemplate.Id > 0 {
						color.Cyan("\t    Config Template: ")
						color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.Device.ConfigTemplate.Id))
						color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.Device.ConfigTemplate.Url))
						color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.Device.ConfigTemplate.Display))
						color.Cyan("\t      Name: " + color.YellowString("%s", result.Module.Device.ConfigTemplate.Name))
						if result.Module.Device.ConfigTemplate.Description != "" {
							color.Cyan("\t      Description: " + color.YellowString("%s", result.Module.Device.ConfigTemplate.Description))
						} else {
							color.Cyan("\t      Description: " + color.RedString("No description found for interface: %s", color.YellowString("%s", result.Display)))
						}
					} else {
						color.Cyan("\t    Config Template: " + color.RedString("No config template found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.Device.LocalContextData != "" {
						color.Cyan("\t    Local Context Data: " + color.YellowString("%s", result.Module.Device.LocalContextData))
					} else {
						color.Cyan("\t    Local Context Data: " + color.RedString("No local context data found for interface: %s", color.YellowString("%s", result.Display)))
					}
					for _, tag := range result.Module.Device.Tags {
						if tag.Id != 0 {
							color.Cyan("\t    Tags: ")
							color.Cyan("\t      ID: " + color.YellowString("%d", tag.Id))
							color.Cyan("\t      URL: " + color.YellowString("%s", tag.Url))
							color.Cyan("\t      Display: " + color.YellowString("%s", tag.Display))
							color.Cyan("\t      Name: " + color.YellowString("%s", tag.Name))
							color.Cyan("\t      Slug: " + color.YellowString("%s", tag.Slug))
							color.Cyan("\t      Color: " + color.YellowString("%s", tag.Color))
						} else {
							color.Cyan("\tTags: " + color.RedString("No tags found for interface: %s", color.YellowString("%s", result.Display)))
						}
					}
					color.Cyan("\t    Created: " + color.YellowString("%s", result.Module.Device.Created))
					color.Cyan("\t    Last Updated: " + color.YellowString("%s", result.Module.Device.LastUpdated))
					if result.Module.Device.ConsolePortCount > 0 {
						color.Cyan("\t    Console Port Count: " + color.YellowString("%d", result.Module.Device.ConsolePortCount))
					} else {
						color.Cyan("\t    Console Port Count: " + color.RedString("No console port count found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.Device.ConsoleServerPortCount > 0 {
						color.Cyan("\t    Console Server Port Count: " + color.YellowString("%d", result.Module.Device.ConsoleServerPortCount))
					} else {
						color.Cyan("\t    Console Server Port Count: " + color.RedString("No console server port count found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.Device.PowerPortCount > 0 {
						color.Cyan("\t    Power Port Count: " + color.YellowString("%d", result.Module.Device.PowerPortCount))
					} else {
						color.Cyan("\t    Power Port Count: " + color.RedString("No power port count found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.Device.PowerOutletCount > 0 {
						color.Cyan("\t    Power Outlet Count: " + color.YellowString("%d", result.Module.Device.PowerOutletCount))
					} else {
						color.Cyan("\t    Power Outlet Count: " + color.RedString("No power outlet count found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.Device.InterfaceCount > 0 {
						color.Cyan("\t    Interface Count: " + color.YellowString("%d", result.Module.Device.InterfaceCount))
					} else {
						color.Cyan("\t    Interface Count: " + color.RedString("No interface count found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.Device.FrontPortCount > 0 {
						color.Cyan("\t    Front Port Count: " + color.YellowString("%d", result.Module.Device.FrontPortCount))
					} else {
						color.Cyan("\t    Front Port Count: " + color.RedString("No front port count found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.Device.RearPortCount > 0 {
						color.Cyan("\t    Rear Port Count: " + color.YellowString("%d", result.Module.Device.RearPortCount))
					} else {
						color.Cyan("\t    Rear Port Count: " + color.RedString("No rear port count found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.Device.DeviceBayCount > 0 {
						color.Cyan("\t    Device Bay Count: " + color.YellowString("%d", result.Module.Device.DeviceBayCount))
					} else {
						color.Cyan("\t    Device Bay Count: " + color.RedString("No device bay count found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.Device.ModuleBayCount > 0 {
						color.Cyan("\t    Module Bay Count: " + color.YellowString("%d", result.Module.Device.ModuleBayCount))
					} else {
						color.Cyan("\t    Module Bay Count: " + color.RedString("No module bay count found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.Device.InventoryItemCount > 0 {
						color.Cyan("\t    Inventory Item Count: " + color.YellowString("%d", result.Module.Device.InventoryItemCount))
					} else {
						color.Cyan("\t    Inventory Item Count: " + color.RedString("No inventory item count found for interface: %s", color.YellowString("%s", result.Display)))
					}
				} else {
					color.Cyan("\tModule: " + color.RedString("No module found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Module.ModuleBay.Id > 0 {
					color.Cyan("\t  Module Bay: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Module.ModuleBay.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Module.ModuleBay.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Module.ModuleBay.Display))
					if result.Module.ModuleBay.InstalledModule.Id > 0 {
						color.Cyan("\t    Installed Module: ")
						color.Cyan("\t      ID: " + color.YellowString("%d", result.Module.ModuleBay.InstalledModule.Id))
						color.Cyan("\t      URL: " + color.YellowString("%s", result.Module.ModuleBay.InstalledModule.Url))
						if result.Module.ModuleBay.InstalledModule.Display != "" {
							color.Cyan("\t      Display: " + color.YellowString("%s", result.Module.ModuleBay.InstalledModule.Display))
						} else {
							color.Cyan("\t      Display: " + color.RedString("No Display found for interface: %s", color.YellowString("%s", result.Display)))
						}
						if result.Module.ModuleBay.InstalledModule.Serial != "" {
							color.Cyan("\t      Serial: " + color.YellowString("%s", result.Module.ModuleBay.InstalledModule.Serial))
						} else {
							color.Cyan("\t      Serial: " + color.RedString("No serial found for interface: %s", color.YellowString("%s", result.Display)))
						}
					} else {
						color.Cyan("\t    Installed Module: " + color.RedString("No installed module found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Module.ModuleBay.Name != "" {
						color.Cyan("\t  Name: " + color.YellowString("%s", result.Module.ModuleBay.Name))
					} else {
						color.Cyan("\t  Name: " + color.RedString("No name found for interface: %s", color.YellowString("%s", result.Display)))
					}
				} else {
					color.Cyan("\t  Module Bay: " + color.RedString("No module bay found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Label != "" {
					color.Cyan("\tLabel: " + color.YellowString("%s", result.Label))
				} else {
					color.Cyan("\tLabel: " + color.RedString("No label found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Type.Value != "" {
					color.Cyan("\tType: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Type.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Type.Label))
				} else {
					color.Cyan("\tType: " + color.RedString("No type found for interface: %s", color.YellowString("%s", result.Display)))
				}
				color.Cyan("\tEnabled: " + color.YellowString("%t", result.Enabled))
				if result.Parent.Id != 0 {
					color.Cyan("\tParent: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Parent.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Parent.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Parent.Display))
					color.Cyan("\t  Device: " + color.YellowString("%s", result.Parent.Device))
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Parent.Device.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Parent.Device.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Parent.Device.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Parent.Device.Name))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Parent.Name))
					color.Cyan("\t  Cable: " + color.YellowString("%s", result.Parent.Cable))
					color.Cyan("\t  Occupied: " + color.YellowString("%t", result.Parent.Occupied))
				} else {
					color.Cyan("\tParent: " + color.RedString("No parent found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Bridge.Id != 0 {
					color.Cyan("\tBridge: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Bridge.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Bridge.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Bridge.Display))
					color.Cyan("\t  Device: " + color.YellowString("%s", result.Bridge.Device))
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Bridge.Device.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Bridge.Device.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Bridge.Device.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Bridge.Device.Name))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Bridge.Name))
					color.Cyan("\t  Cable: " + color.YellowString("%s", result.Bridge.Cable))
					color.Cyan("\t  Occupied: " + color.YellowString("%t", result.Bridge.Occupied))
				} else {
					color.Cyan("\tBridge: " + color.RedString("No bridge found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Lag.Id != 0 {
					color.Cyan("\tLink Aggregation Groups: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Lag.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Lag.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Lag.Display))
					color.Cyan("\t  Device: " + color.YellowString("%s", result.Lag.Device))
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Lag.Device.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Lag.Device.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Lag.Device.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Lag.Device.Name))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Lag.Name))
					color.Cyan("\t  Cable: " + color.YellowString("%s", result.Lag.Cable))
					color.Cyan("\t  Occupied: " + color.YellowString("%t", result.Lag.Occupied))
				} else {
					color.Cyan("\tLink Aggregation Groups: " + color.RedString("No LAGs found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Mtu != 0 {
					color.Cyan("\tMTU: " + color.YellowString("%d", result.Mtu))
				} else {
					color.Cyan("\tMTU: " + color.RedString("No MTU found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.MacAddress != "" {
					color.Cyan("\tMac Address: " + color.YellowString("%s", result.MacAddress))
				} else {
					color.Cyan("\tMac Address: " + color.RedString("No mac address found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Speed != 0 {
					color.Cyan("\tSpeed: " + color.YellowString("%d", result.Speed))
				} else {
					color.Cyan("\tSpeed: " + color.RedString("No Speed found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Duplex.Value != "" {
					color.Cyan("\tDuplex: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Duplex.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Duplex.Label))
				} else {
					color.Cyan("\tDuplex: " + color.RedString("No duplex found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Wwn != "" {
					color.Cyan("\tWorld Wide Name: " + color.YellowString("%s", result.Wwn))
				} else {
					color.Cyan("\tWorld Wide Name: " + color.RedString("No world wide name found for interface: %s", color.YellowString("%s", result.Display)))
				}
				color.Cyan("\tMgmt Only: " + color.YellowString("%t", result.MgmtOnly))

				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Mode.Value != "" {
					color.Cyan("\tMode: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Mode.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Mode.Label))
				} else {
					color.Cyan("\tMode: " + color.RedString("No mode found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.RfRole.Value != "" {
					color.Cyan("\tRadio Frequency (RF) Role: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.RfRole.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.RfRole.Label))
				} else {
					color.Cyan("\tRadio Frequency (RF) Role: " + color.RedString("No RF role found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.RfChannel.Value != "" {
					color.Cyan("\tRadio Frequency (RF) Channel: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.RfChannel.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.RfChannel.Label))
				} else {
					color.Cyan("\tRadio Frequency (RF) Channel: " + color.RedString("No RF channel found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.PoeMode.Value != "" {
					color.Cyan("\tPower Over Ethernet (PoE) Mode: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.PoeMode.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.PoeMode.Label))
				} else {
					color.Cyan("\tPower Over Ethernet (PoE) Mode: " + color.RedString("No PoE mode found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.PoeType.Value != "" {
					color.Cyan("\tPower Over Ethernet (PoE) Type: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.PoeType.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.PoeType.Label))
				} else {
					color.Cyan("\tPower Over Ethernet (PoE) Type: " + color.RedString("No PoE type found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.RfChannelFrequency != 0 {
					color.Cyan("\tRF Channel Frequency: " + color.YellowString("%d", result.RfChannelFrequency))
				} else {
					color.Cyan("\tRF Channel Frequency: " + color.RedString("No RF channel frequency found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.RfChannelWidth != 0 {
					color.Cyan("\tRF Channel Width: " + color.YellowString("%d", result.RfChannelWidth))
				} else {
					color.Cyan("\tRF Channel Width: " + color.RedString("No RF channel width found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.TxPower != 0 {
					color.Cyan("\tTx Power: " + color.YellowString("%d", result.TxPower))
				} else {
					color.Cyan("\tTx Power: " + color.RedString("No Tx power found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.UntaggedVlan.Id != 0 {
					color.Cyan("\tUntagged Vlan: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.UntaggedVlan.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.UntaggedVlan.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.UntaggedVlan.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.UntaggedVlan.Name))
				} else {
					color.Cyan("\tUntagged Vlan: " + color.RedString("No untagged vlan found for interface: %s", color.YellowString("%s", result.Display)))
				}
				color.Cyan("\tMark Connected: " + color.YellowString("%t", result.MarkConnected))
				if result.Cable.Id != 0 {
					color.Cyan("\tCable: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Cable.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Cable.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Cable.Display))
					color.Cyan("\t    Label: " + color.YellowString("%s", result.Cable.Label))
				} else {
					color.Cyan("\tCable: " + color.RedString("No cable found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.CableEnd != "" {
					color.Cyan("\tCable End: " + color.YellowString("%s", result.CableEnd))
				} else {
					color.Cyan("\tCable End: " + color.RedString("No cable end found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.WirelessLink.Id != 0 {
					color.Cyan("\tWireless Link: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.WirelessLink.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.WirelessLink.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.WirelessLink.Display))
					color.Cyan("\t    SSID: " + color.YellowString("%s", result.WirelessLink.Ssid))
				} else {
					color.Cyan("\tWireless Link: " + color.RedString("No wireless link found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.LinkPeers == nil || len(result.LinkPeers) == 0 {
					color.Cyan("\tLink Peers: " + color.RedString("No link peers found for interface: %s", color.YellowString("%s", result.Display)))
				} else {
					color.Cyan("\tLink Peers: " + color.YellowString("%d", result.LinkPeers))
				}
				if result.LinkPeersType != "" {
					color.Cyan("\tLink Peers Type: " + color.YellowString("%d", result.LinkPeers))
				} else {
					color.Cyan("\tLink Peers Type: " + color.RedString("No link peers type found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.WirelessLans == nil || len(result.WirelessLans) == 0 {
					color.Cyan("\tWireless LANs: " + color.RedString("No wireless lans found for interface: %s", color.YellowString("%s", result.Display)))
				} else {
					color.Cyan("\tWireless LANs: " + color.YellowString("%d", result.WirelessLans))
				}
				if result.Vrf.Id != 0 {
					color.Cyan("\tVirtual Router Forwarding (vrf): ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Vrf.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Vrf.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Vrf.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Vrf.Name))
					color.Cyan("\t    RD: " + color.YellowString("%s", result.Vrf.Rd))
				} else {
					color.Cyan("\tVirtual Router Forwarding (vrf): " + color.RedString("No vrf found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.L2VpnTermination.Id != 0 {
					color.Cyan("\tL2VPN Termination: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.L2VpnTermination.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.L2VpnTermination.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.L2VpnTermination.Display))
					color.Cyan("\t  L2VPN: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.L2VpnTermination.L2Vpn.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.L2VpnTermination.L2Vpn.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.L2VpnTermination.L2Vpn.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.L2VpnTermination.L2Vpn.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.L2VpnTermination.L2Vpn.Slug))
					color.Cyan("\t    Identifier: " + color.YellowString("%f", result.L2VpnTermination.L2Vpn.Identifier))
					color.Cyan("\t    Type: " + color.YellowString("%s", result.L2VpnTermination.L2Vpn.Type))
				} else {
					color.Cyan("\tL2VPN Termination: " + color.RedString("No l2vpn termination found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.ConnectedEndpoints == nil || len(result.ConnectedEndpoints) == 0 {
					color.Cyan("\tConnected Enpoints: " + color.RedString("No connected endpoints found for interface: %s", color.YellowString("%s", result.Display)))
				} else {
					color.Cyan("\tConnected Enpoints: " + color.YellowString("%d", result.LinkPeers))
				}
				if result.ConnectedEndpointsType != "" {
					color.Cyan("\tConnected Endpoints Type: " + color.YellowString("%d", result.ConnectedEndpointsType))
				} else {
					color.Cyan("\tConnected Endpoints Type: " + color.RedString("No connected endpoints type found for interface: %s", color.YellowString("%s", result.Display)))
				}
				color.Cyan("\tConnected Endpoints Reachable: " + color.YellowString("%t", result.ConnectedEndpointsReachable))
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
						color.Cyan("\tTags: " + color.RedString("No tags found for interface: %s", color.YellowString("%s", result.Display)))
					}
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
				if result.CountIpaddresses != 0 {
					color.Cyan("\tCount IP Addresses: " + color.YellowString("%d", result.CountIpaddresses))
				} else {
					color.Cyan("\tCount IP Addresses: " + color.RedString("No count IP addresses found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.CountFhrpGroups != 0 {
					color.Cyan("\tCount FHRP Groups: " + color.YellowString("%d", result.CountFhrpGroups))
				} else {
					color.Cyan("\tCount FHRP Groups: " + color.RedString("No count fhrp groups found for interface: %s", color.YellowString("%s", result.Display)))
				}
				color.Cyan("\tOccupied: " + color.YellowString("%t\n", result.Occupied))
			}
			for responseObjectInterfaces.Next != nil {
				nextPageInterfaces()
			}
			if responseObjectInterfaces.Next == nil {
				display := color.HiGreenString("\tAll Netbox interface objects have been successfully displayed for Query: " + color.YellowString("%s", query))
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n\n")
			}
		} else {
			color.Cyan("  Total ABC Interfaces: " + color.RedString("No interfaces found on the server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimInterfacesByQueryCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimInterfacesByQueryCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for GetDcimInterfacesByQueryCmd", err)
	}

	GetDcimInterfacesByQueryCmd.Flags().StringVarP(&query, "query", "q", "", "string query of object you want to get")
	err = GetDcimInterfacesByQueryCmd.MarkFlagRequired("query")
	if err != nil {
		log.Fatalf("Error marking query flag as required: %s - for GetDcimInterfacesByQueryCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimInterfacesByQueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimInterfacesByQueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

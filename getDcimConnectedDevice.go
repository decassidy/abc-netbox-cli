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

type connectedDevice struct {
	Id         uint   `json:"id"`
	Url        string `json:"url"`
	Display    string `json:"display"`
	Name       string `json:"name"`
	DeviceType struct {
		Id           uint   `json:"id"`
		Url          string `json:"url"`
		Display      string `json:"display"`
		Manufacturer struct {
			CommonFieldsSlug
		} `json:"manufacturer"`
		Model string `json:"model"`
		Slug  string `json:"slug"`
	} `json:"device_type"`
	Role struct {
		CommonFieldsSlug
	} `json:"role"`
	DeviceRole struct {
		CommonFieldsSlug
	} `json:"device_role"`
	Tenant struct {
		CommonFieldsSlug
	} `json:"tenant"`
	Platform struct {
		CommonFieldsSlug
	} `json:"platform"`
	Serial   string `json:"serial"`
	AssetTag string `json:"asset_tag"`
	Site     struct {
		CommonFieldsSlug
	} `json:"site"`
	Location struct {
		CommonFieldsSlug
		Depth int `json:"_depth"`
	} `json:"location"`
	Rack struct {
		CommonFieldsNoSlug
	} `json:"rack"`
	Position uint `json:"position"`
	Face     struct {
		ValueLabel
	} `json:"face"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
	ParentDevice struct {
		CommonFieldsNoSlug
	} `json:"parent_device"`
	Status struct {
		ValueLabel
	} `json:"status"`
	Airflow struct {
		ValueLabel
	} `json:"airflow"`
	PrimaryIp struct {
		Id      uint   `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Family  uint   `json:"family"`
		Address string `json:"address"`
	} `json:"primary_ip"`
	PrimaryIp4 struct {
		Id      uint   `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Family  uint   `json:"family"`
		Address string `json:"address"`
	} `json:"primary_ip4"`
	PrimaryIp6 struct {
		Id      uint   `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Family  uint   `json:"family"`
		Address string `json:"address"`
	} `json:"primary_ip6"`
	OobIp struct {
		Id      uint   `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Family  uint   `json:"family"`
		Address string `json:"address"`
	} `json:"oob_ip"`
	Cluster struct {
		Id      uint   `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
	} `json:"cluster"`
	VirtualChassis struct {
		CommonFieldsNoSlug
		Master struct {
			CommonFieldsNoSlug
		} `json:"master"`
	} `json:"virtual_chassis"`
	VcPosition     uint   `json:"vc_position"`
	VcPriority     uint   `json:"vc_priority"`
	Description    string `json:"description"`
	Comments       string `json:"comments"`
	ConfigTemplate struct {
		CommonFieldsNoSlug
	} `json:"config_template"`
	LocalContextData string `json:"local_context_data"`
	Tags             []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created                string `json:"created"`
	LastUpdated            string `json:"last_updated"`
	ConsolePortCount       uint   `json:"console_port_count"`
	ConsoleServerPortCount uint   `json:"console_server_port_count"`
	PowerPortCount         uint   `json:"power_port_count"`
	PowerOutletCount       uint   `json:"power_outlet_count"`
	InterfaceCount         uint   `json:"interface_count"`
	FrontPortCount         uint   `json:"front_port_count"`
	RearPortCount          uint   `json:"rear_port_count"`
	DeviceBayCount         uint   `json:"device_bay_count"`
	ModuleBayCount         uint   `json:"module_bay_count"`
	InventoryItemCount     uint   `json:"inventory_item_count"`
}

// GetDcimConnectedDeviceCmd represents the getDcimConnectedDevice command
var GetDcimConnectedDeviceCmd = &cobra.Command{
	Use:   "getDcimConnectedDevice",
	Short: "GET a list of connected device objects",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of connected device objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(connectedDevice)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.connected_device")

		if responseObject.Id > 0 {
			color.Cyan("\n  Metropolis Connected Devices: "+color.YellowString("for %s"), responseObject.Display)

			display := fmt.Sprintf("    Metropolis Connected Device: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			if responseObject.DeviceType.Id > 0 {
				color.Cyan("\tDevice Type: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Display))
			} else {
				color.Cyan("\tDevice Type: " + color.RedString("No device type found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.DeviceType.Manufacturer.Id > 0 {
				color.Cyan("\t  Manufacturer: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.DeviceType.Manufacturer.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Slug))
			} else {
				color.Cyan("\t  Manufacturer: " + color.RedString("No manufacturer found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.DeviceType.Model != "" {
				color.Cyan("\t  Model: " + color.YellowString("%s", responseObject.DeviceType.Model))
			} else {
				color.Cyan("\t  Model: " + color.RedString("No model found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.DeviceType.Slug != "" {
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.DeviceType.Slug))
			} else {
				color.Cyan("\t  Slug: " + color.RedString("No slug found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Role.Id > 0 {
				color.Cyan("\tRole: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Role.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Role.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Role.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Role.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Role.Slug))
			} else {
				color.Cyan("\tRole: " + color.RedString("No role found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.DeviceRole.Id > 0 {
				color.Cyan("\tDevice Role: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.DeviceRole.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.DeviceRole.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.DeviceRole.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.DeviceRole.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.DeviceRole.Slug))
			} else {
				color.Cyan("\tDevice Role: " + color.RedString("No device role found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Tenant.Id > 0 {
				color.Cyan("\tTenant: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Tenant.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Tenant.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Tenant.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Tenant.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Tenant.Slug))
			} else {
				color.Cyan("\tTenant: " + color.RedString("No tenant found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Platform.Id > 0 {
				color.Cyan("\tPlatform: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Platform.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Platform.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Platform.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Platform.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Platform.Slug))
			} else {
				color.Cyan("\tPlatform: " + color.RedString("No platform found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Serial != "" {
				color.Cyan("\tSerial: " + color.YellowString("%s", responseObject.Serial))
			} else {
				color.Cyan("\tSerial: " + color.RedString("No serial found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.AssetTag != "" {
				color.Cyan("\tAsset Tag: " + color.YellowString("%s", responseObject.AssetTag))
			} else {
				color.Cyan("\tAsset Tag: " + color.RedString("No asset tag found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Site.Id > 0 {
				color.Cyan("\tSite: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Site.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Site.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Site.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Site.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Site.Slug))
			} else {
				color.Cyan("\tSite: " + color.RedString("No site found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Location.Id > 0 {
				color.Cyan("\tSite: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Location.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Location.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Location.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Location.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Location.Slug))
				if responseObject.Location.Depth > 0 {
					color.Cyan("\t  Depth: " + color.YellowString("%d", responseObject.Location.Depth))
				} else {
					color.Cyan("\t  Depth: " + color.RedString("No depth configured for connected device: %s", color.YellowString(responseObject.Display)))
				}
			} else {
				color.Cyan("\tSite: " + color.RedString("No site found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Rack.Id > 0 {
				color.Cyan("\tRack: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Rack.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Rack.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Rack.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Rack.Name))
			} else {
				color.Cyan("\tRack: " + color.RedString("No rack found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Position > 0 {
				color.Cyan("\tPosition: " + color.YellowString("%d", responseObject.Position))
			} else {
				color.Cyan("\tPosition: " + color.RedString("No position configured for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Face.Value != "" {
				color.Cyan("\tFace: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Face.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Face.Label))
			} else {
				color.Cyan("\tFace: " + color.RedString("No face found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Latitude > 0 {
				color.Cyan("\tLatitude: " + color.YellowString("%f", responseObject.Latitude))
			} else {
				color.Cyan("\tLatitude: " + color.RedString("No latitude found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Longitude > 0 {
				color.Cyan("\tLongitude: " + color.YellowString("%f", responseObject.Longitude))
			} else {
				color.Cyan("\tLongitude: " + color.RedString("No longitude found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.ParentDevice.Id > 0 {
				color.Cyan("\tParent Device: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.ParentDevice.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.ParentDevice.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.ParentDevice.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.ParentDevice.Name))
			} else {
				color.Cyan("\tParent Device: " + color.RedString("No parent device found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Status.Value != "" {
				color.Cyan("\tStatus: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Status.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Status.Label))
			} else {
				color.Cyan("\tStatus: " + color.RedString("No status found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Airflow.Value != "" {
				color.Cyan("\tAirflow: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Airflow.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Airflow.Label))
			} else {
				color.Cyan("\tAirflow: " + color.RedString("No airflow settings found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.PrimaryIp.Id > 0 {
				color.Cyan("\tPrimary IP: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.PrimaryIp.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.PrimaryIp.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.PrimaryIp.Display))
				color.Cyan("\t  Family: " + color.YellowString("%d", responseObject.PrimaryIp.Family))
				color.Cyan("\t  Address: " + color.YellowString("%s", responseObject.PrimaryIp.Address))
			} else {
				color.Cyan("\tPrimary IP: " + color.RedString("No primary IP found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.PrimaryIp4.Id > 0 {
				color.Cyan("\tPrimary IPv4: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.PrimaryIp4.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.PrimaryIp4.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.PrimaryIp4.Display))
				color.Cyan("\t  Family: " + color.YellowString("%d", responseObject.PrimaryIp4.Family))
				color.Cyan("\t  Address: " + color.YellowString("%s", responseObject.PrimaryIp4.Address))
			} else {
				color.Cyan("\tPrimary IPv4: " + color.RedString("No primary IPv4 found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.PrimaryIp6.Id > 0 {
				color.Cyan("\tPrimary IPv6: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.PrimaryIp6.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.PrimaryIp6.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.PrimaryIp6.Display))
				color.Cyan("\t  Family: " + color.YellowString("%d", responseObject.PrimaryIp6.Family))
				color.Cyan("\t  Address: " + color.YellowString("%s", responseObject.PrimaryIp6.Address))
			} else {
				color.Cyan("\tPrimary IPv6: " + color.RedString("No primary IPv6 found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.OobIp.Id > 0 {
				color.Cyan("\tOut Of Band - OOB IP: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.OobIp.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.OobIp.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.OobIp.Display))
				color.Cyan("\t  Family: " + color.YellowString("%d", responseObject.OobIp.Family))
				color.Cyan("\t  Address: " + color.YellowString("%s", responseObject.OobIp.Address))
			} else {
				color.Cyan("\tOut Of Band - OOB IP: " + color.RedString("No OOB IP found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Cluster.Id > 0 {
				color.Cyan("\tCluster: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Cluster.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Cluster.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Cluster.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Cluster.Name))
			} else {
				color.Cyan("\tCluster: " + color.RedString("No cluster found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.VirtualChassis.Id > 0 {
				color.Cyan("\tVirtual Chassis: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.VirtualChassis.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.VirtualChassis.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.VirtualChassis.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.VirtualChassis.Name))
				if responseObject.VirtualChassis.Master.Id > 0 {
					color.Cyan("\t  Master: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", responseObject.VirtualChassis.Master.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", responseObject.VirtualChassis.Master.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", responseObject.VirtualChassis.Master.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", responseObject.VirtualChassis.Master.Name))
				} else {
					color.Cyan("\t  Master: " + color.RedString("No virtual chassis master found for connected device: %s", color.YellowString(responseObject.Display)))
				}
			} else {
				color.Cyan("\tVirtual Chassis: " + color.RedString("No virtual chassis found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.VcPosition > 0 {
				color.Cyan("\tVC Position: " + color.YellowString("%d", responseObject.VcPosition))
			} else {
				color.Cyan("\tVC Position: " + color.RedString("No vc position found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.VcPriority > 0 {
				color.Cyan("\tVC Priority: " + color.YellowString("%d", responseObject.VcPriority))
			} else {
				color.Cyan("\tVC Priority: " + color.RedString("No vc priority found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.Comments != "" {
				color.Cyan("\tComments: " + color.YellowString("%s", responseObject.Comments))
			} else {
				color.Cyan("\tComments: " + color.RedString("No comments found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.ConfigTemplate.Id > 0 {
				color.Cyan("\tConfig Template: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.ConfigTemplate.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.ConfigTemplate.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.ConfigTemplate.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.ConfigTemplate.Name))
			} else {
				color.Cyan("\tConfig Template: " + color.RedString("No config template found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.LocalContextData != "" {
				color.Cyan("\tLocal Context Data: " + color.YellowString("%s", responseObject.LocalContextData))
			} else {
				color.Cyan("\tLocal Context Data: " + color.RedString("No local context data found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			for _, tag := range responseObject.Tags {
				if tag.Id > 0 {
					color.Cyan("\tTags: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
					color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
				} else {
					color.Cyan("\tTags: " + color.RedString("No Tags entry found for console port name: ") + color.YellowString("%s", responseObject.Display))
				}
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
			if responseObject.ConsolePortCount > 0 {
				color.Cyan("\tConsole Port Count: " + color.YellowString("%d", responseObject.ConsolePortCount))
			} else {
				color.Cyan("\tConsole Port Count: " + color.RedString("No console port count found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.ConsoleServerPortCount > 0 {
				color.Cyan("\tConsole Server Port Count: " + color.YellowString("%d", responseObject.ConsoleServerPortCount))
			} else {
				color.Cyan("\tConsole Server Port Count: " + color.RedString("No console server port count found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.PowerPortCount > 0 {
				color.Cyan("\tPower Port Count: " + color.YellowString("%d", responseObject.PowerPortCount))
			} else {
				color.Cyan("\tPower Port Count: " + color.RedString("No power port count found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.PowerOutletCount > 0 {
				color.Cyan("\tPower Outlet Count: " + color.YellowString("%d", responseObject.PowerOutletCount))
			} else {
				color.Cyan("\tPower Outlet Count: " + color.RedString("No power outlet count found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.InterfaceCount > 0 {
				color.Cyan("\tInterface Count: " + color.YellowString("%d", responseObject.InterfaceCount))
			} else {
				color.Cyan("\tInterface Count: " + color.RedString("No interface count found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.FrontPortCount > 0 {
				color.Cyan("\tFront Port Count: " + color.YellowString("%d", responseObject.FrontPortCount))
			} else {
				color.Cyan("\tFront Port Count: " + color.RedString("No front port count found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.RearPortCount > 0 {
				color.Cyan("\tRear Port Count: " + color.YellowString("%d", responseObject.RearPortCount))
			} else {
				color.Cyan("\tRear Port Count: " + color.RedString("No rear port count found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.DeviceBayCount > 0 {
				color.Cyan("\tDevice Bay Count: " + color.YellowString("%d", responseObject.DeviceBayCount))
			} else {
				color.Cyan("\tDevice Bay Count: " + color.RedString("No device bay port count found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.ModuleBayCount > 0 {
				color.Cyan("\tModule Bay Count: " + color.YellowString("%d", responseObject.ModuleBayCount))
			} else {
				color.Cyan("\tModule Bay Count: " + color.RedString("No module bay count found for connected device: %s", color.YellowString(responseObject.Display)))
			}
			if responseObject.InventoryItemCount > 0 {
				color.Cyan("\tInventory Item Count: " + color.YellowString("%d", responseObject.InventoryItemCount))
			} else {
				color.Cyan("\tInventory Item Count: " + color.RedString("No inventory item count found for connected device: %s", color.YellowString(responseObject.Display)))
			}
		} else {
			color.Cyan("  Metropolis Connected Device: " + color.RedString("No connected device found on server. Exiting..."))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimConnectedDeviceCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimConnectedDeviceCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimConnectedDeviceCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimConnectedDeviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimConnectedDeviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

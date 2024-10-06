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
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

type interfaces struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		idUrlDisplay
		Device struct {
			CommonFieldsNoSlug
			DeviceType struct {
				idUrlDisplay
				Manufacturer struct {
					CommonFieldsSlug
					Description     string `json:"description"`
					DevicetypeCount int    `json:"devicetype_count"`
				} `json:"manufacturer"`
				Model       string `json:"model"`
				Slug        string `json:"slug"`
				Description string `json:"description"`
				DeviceCount int    `json:"device_count"`
			} `json:"device_type,omitempty"`
			Role struct {
				CommonFieldsSlug
				Description         string `json:"description"`
				DeviceCount         int    `json:"device_count"`
				VirtualmachineCount int    `json:"virtualmachine_count"`
			} `json:"role,omitempty"`
			Tenant struct {
				CommonFieldsSlug
				Description string `json:"description"`
			} `json:"tenant,omitempty"`
			Platform struct {
				CommonFieldsSlug
				Description         string `json:"description"`
				DeviceCount         int    `json:"device_count"`
				VirtualmachineCount int    `json:"virtualmachine_count"`
			} `json:"platform,omitempty"`
			Serial   string `json:"serial,omitempty"`
			AssetTag string `json:"asset_tag,omitempty"`
			Site     struct {
				CommonFieldsSlug
				Description string `json:"description"`
			} `json:"site,omitempty"`
			Location struct {
				CommonFieldsSlug
				Description string `json:"description"`
				RackCount   int    `json:"rack_count"`
				Depth       int    `json:"_depth"`
			} `json:"location,omitempty"`
			Rack struct {
				CommonFieldsNoSlug
				Description string `json:"description"`
				DeviceCount int    `json:"device_count"`
			} `json:"rack,omitempty"`
			Position int `json:"position,omitempty"`
			Face     struct {
				ValueLabel
			} `json:"face,omitempty"`
			Latitude     float32 `json:"latitude,omitempty"`
			Longitude    float32 `json:"longitude,omitempty"`
			ParentDevice struct {
				CommonFieldsNoSlug
			} `json:"parent_device,omitempty"`
			Status struct {
				ValueLabel
			} `json:"status,omitempty"`
			Airflow struct {
				ValueLabel
			} `json:"airflow"`
			PrimaryIp struct {
				idUrlDisplay
				Family struct {
					family
				} `json:"family"`
				addressDescription
			} `json:"primary_ip"`
			PrimaryIp4 struct {
				idUrlDisplay
				Family struct {
					family
				} `json:"family"`
				addressDescription
			} `json:"primary_ip4"`
			PrimaryIp6 struct {
				idUrlDisplay
				Family struct {
					family
				} `json:"family"`
				addressDescription
			} `json:"primary_ip6"`
			OobIp struct {
				idUrlDisplay
				Family struct {
					family
				} `json:"family"`
				addressDescription
			} `json:"oob_ip"`
			Cluster struct {
				CommonFieldsNoSlug
				Description         string `json:"description"`
				VirtualmachineCount int    `json:"virtualmachine_count"`
			} `json:"cluster"`
			VirtualChassis struct {
				CommonFieldsNoSlug
				Master struct {
					CommonFieldsNoSlug
				} `json:"master"`
				Description string `json:"description"`
				MemberCount int    `json:"member_count"`
			} `json:"virtual_chassis"`
			VcPosition     int    `json:"vc_position"`
			VcPriority     int    `json:"vc_priority"`
			Description    string `json:"description"`
			Comments       string `json:"comments"`
			ConfigTemplate struct {
				CommonFieldsNoSlug
				Description string `json:"description"`
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
		} `json:"device"`
		Vdcs []struct {
			CommonFieldsNoSlug
			Device struct {
				CommonFieldsNoSlug
				DeviceType struct {
					idUrlDisplay
					Manufacturer struct {
						CommonFieldsSlug
						Description     string `json:"description"`
						DevicetypeCount int    `json:"devicetype_count"`
					} `json:"manufacturer"`
					Model       string `json:"model"`
					Slug        string `json:"slug"`
					Description string `json:"description"`
					DeviceCount int    `json:"device_count"`
				} `json:"device_type"`
				Role struct {
					CommonFieldsSlug
					Description         string `json:"description"`
					DeviceCount         int    `json:"device_count"`
					VirtualmachineCount int    `json:"virtualmachine_count"`
				} `json:"role"`
				Tenant struct {
					CommonFieldsSlug
					Description string `json:"description"`
				} `json:"tenant"`
				Platform struct {
					CommonFieldsSlug
					Description         string `json:"description"`
					DeviceCount         int    `json:"device_count"`
					VirtualmachineCount int    `json:"virtualmachine_count"`
				} `json:"platform"`
				Serial   string `json:"serial"`
				AssetTag string `json:"asset_tag"`
				Site     struct {
					CommonFieldsSlug
					Description string `json:"description"`
				} `json:"site"`
				Location struct {
					CommonFieldsSlug
					Description string `json:"description"`
					RackCount   int    `json:"rack_count"`
					Depth       int    `json:"_depth"`
				} `json:"location"`
				Rack struct {
					CommonFieldsNoSlug
					Description string `json:"description"`
					DeviceCount int    `json:"device_count"`
				} `json:"rack"`
				Position int `json:"position"`
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
					idUrlDisplay
					Family struct {
						family
					} `json:"family"`
					addressDescription
				} `json:"primary_ip"`
				PrimaryIp4 struct {
					idUrlDisplay
					Family struct {
						family
					} `json:"family"`
					addressDescription
				} `json:"primary_ip4"`
				PrimaryIp6 struct {
					idUrlDisplay
					Family struct {
						family
					} `json:"family"`
					addressDescription
				} `json:"primary_ip6"`
				OobIp struct {
					idUrlDisplay
					Family struct {
						family
					} `json:"family"`
					addressDescription
				} `json:"oob_ip"`
				Cluster struct {
					CommonFieldsNoSlug
					Description         string `json:"description"`
					VirtualmachineCount int    `json:"virtualmachine_count"`
				} `json:"cluster"`
				VirtualChassis struct {
					CommonFieldsNoSlug
					Master struct {
						CommonFieldsNoSlug
					} `json:"master"`
					Description string `json:"description"`
					MemberCount int    `json:"member_count"`
				} `json:"virtual_chassis"`
				VcPosition     int    `json:"vc_position"`
				VcPriority     int    `json:"vc_priority"`
				Description    string `json:"description"`
				Comments       string `json:"comments"`
				ConfigTemplate struct {
					CommonFieldsNoSlug
					Name        string `json:"name"`
					Description string `json:"description"`
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
			} `json:"device"`
			Identifier int `json:"identifier"`
			Tenant     struct {
				CommonFieldsSlug
				Description string `json:"description"`
			} `json:"tenant"`
			PrimaryIp struct {
				idUrlDisplay
				Family struct {
					family
				} `json:"family"`
				addressDescription
			} `json:"primary_ip"`
			PrimaryIp4 struct {
				idUrlDisplay
				Family struct {
					family
				} `json:"family"`
				addressDescription
			} `json:"primary_ip4"`
			PrimaryIp6 struct {
				idUrlDisplay
				Family struct {
					family
				} `json:"family"`
				addressDescription
			} `json:"primary_ip6"`
			Status struct {
				ValueLabel
			} `json:"status"`
			Description string `json:"description"`
			Comments    string `json:"comments"`
			Tags        []struct {
				CommonFieldsSlug
				Color string `json:"color"`
			} `json:"tags"`
			Created        string `json:"created"`
			LastUpdated    string `json:"last_updated"`
			InterfaceCount int    `json:"interface_count"`
		} `json:"vdcs"`
		Module struct {
			idUrlDisplay
			Device struct {
				CommonFieldsNoSlug
				Name       string `json:"name"`
				DeviceType struct {
					idUrlDisplay
					Manufacturer struct {
						CommonFieldsSlug
						Description     string `json:"description"`
						DevicetypeCount int    `json:"devicetype_count"`
					} `json:"manufacturer"`
					Model       string `json:"model"`
					Slug        string `json:"slug"`
					Description string `json:"description"`
					DeviceCount int    `json:"device_count"`
				} `json:"device_type"`
				Role struct {
					CommonFieldsSlug
					Description         string `json:"description"`
					DeviceCount         int    `json:"device_count"`
					VirtualmachineCount int    `json:"virtualmachine_count"`
				} `json:"role"`
				Tenant struct {
					CommonFieldsSlug
					Description string `json:"description"`
				} `json:"tenant"`
				Platform struct {
					CommonFieldsSlug
					Description         string `json:"description"`
					DeviceCount         int    `json:"device_count"`
					VirtualmachineCount int    `json:"virtualmachine_count"`
				} `json:"platform"`
				Serial   string `json:"serial"`
				AssetTag string `json:"asset_tag"`
				Site     struct {
					CommonFieldsSlug
					Description string `json:"description"`
				} `json:"site"`
				Location struct {
					CommonFieldsSlug
					Description string `json:"description"`
					RackCount   int    `json:"rack_count"`
					Depth       int    `json:"_depth"`
				} `json:"location"`
				Rack struct {
					CommonFieldsNoSlug
					Description string `json:"description"`
					DeviceCount int    `json:"device_count"`
				} `json:"rack"`
				Position int `json:"position"`
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
					idUrlDisplay
					Family struct {
						family
					} `json:"family"`
					addressDescription
				} `json:"primary_ip"`
				PrimaryIp4 struct {
					idUrlDisplay
					Family struct {
						family
					} `json:"family"`
					addressDescription
				} `json:"primary_ip4"`
				PrimaryIp6 struct {
					idUrlDisplay
					Family struct {
						family
					} `json:"family"`
					addressDescription
				} `json:"primary_ip6"`
				OobIp struct {
					idUrlDisplay
					Family struct {
						family
					} `json:"family"`
					addressDescription
				} `json:"oob_ip"`
				Cluster struct {
					CommonFieldsNoSlug
					Description         string `json:"description"`
					VirtualmachineCount int    `json:"virtualmachine_count"`
				} `json:"cluster"`
				VirtualChassis struct {
					CommonFieldsNoSlug
					Master struct {
						CommonFieldsNoSlug
					} `json:"master"`
					Description string `json:"description"`
					MemberCount int    `json:"member_count"`
				} `json:"virtual_chassis"`
				VcPosition     int    `json:"vc_position"`
				VcPriority     int    `json:"vc_priority"`
				Description    string `json:"description"`
				Comments       string `json:"comments"`
				ConfigTemplate struct {
					CommonFieldsNoSlug
					Description string `json:"description"`
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
			} `json:"device"`
			ModuleBay struct {
				idUrlDisplay
				InstalledModule struct {
					idUrlDisplay
					Serial string `json:"serial"`
				} `json:"installed_module"`
				Name string `json:"name"`
			} `json:"module_bay"`
		} `json:"module"`
		Name  string `json:"name"`
		Label string `json:"label"`
		Type  struct {
			ValueLabel
		} `json:"type"`
		Enabled bool `json:"enabled"`
		Parent  struct {
			idUrlDisplay
			Device struct {
				CommonFieldsNoSlug
			} `json:"device"`
			Name     string `json:"name"`
			Cable    int    `json:"cable"`
			Occupied bool   `json:"_occupied"`
		} `json:"parent"`
		Bridge struct {
			idUrlDisplay
			Device struct {
				CommonFieldsNoSlug
			} `json:"device"`
			Name     string `json:"name"`
			Cable    int    `json:"cable"`
			Occupied bool   `json:"_occupied"`
		} `json:"bridge"`
		Lag struct {
			idUrlDisplay
			Device struct {
				CommonFieldsNoSlug
			} `json:"device"`
			Name     string `json:"name"`
			Cable    int    `json:"cable"`
			Occupied bool   `json:"_occupied"`
		} `json:"lag"`
		Mtu        int    `json:"mtu"`
		MacAddress string `json:"mac_address"`
		Speed      int    `json:"speed"`
		Duplex     struct {
			ValueLabel
		} `json:"duplex"`
		Wwn         string `json:"wwn"`
		MgmtOnly    bool   `json:"mgmt_only"`
		Description string `json:"description"`
		Mode        struct {
			ValueLabel
		} `json:"mode"`
		RfRole struct {
			ValueLabel
		} `json:"rf_role"`
		RfChannel struct {
			ValueLabel
		} `json:"rf_channel"`
		PoeMode struct {
			ValueLabel
		} `json:"poe_mode"`
		PoeType struct {
			ValueLabel
		} `json:"poe_type"`
		RfChannelFrequency int `json:"rf_channel_frequency"`
		RfChannelWidth     int `json:"rf_channel_width"`
		TxPower            int `json:"tx_power"`
		UntaggedVlan       struct {
			Id          uint   `json:"id"`
			Url         string `json:"url"`
			Display     string `json:"display"`
			Vid         string `json:"vid"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"untagged_vlan"`
		TaggedVlans []struct {
			Id          uint   `json:"id"`
			Url         string `json:"url"`
			Display     string `json:"display"`
			Vid         string `json:"vid"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"tagged_vlans"`
		MarkConnected bool `json:"mark_connected"`
		Cable         struct {
			idUrlDisplay
			Label       string `json:"label"`
			Description string `json:"description"`
		} `json:"cable"`
		CableEnd     string `json:"cable_end"`
		WirelessLink struct {
			idUrlDisplay
			Ssid string `json:"ssid"`
		} `json:"wireless_link"`
		LinkPeers     []string `json:"link_peers"`
		LinkPeersType string   `json:"link_peers_type"`
		WirelessLans  []struct {
			idUrlDisplay
			Ssid        string `json:"ssid"`
			Description string `json:"description"`
			Group       struct {
				CommonFieldsSlug
				Description      string `json:"description"`
				WirelesslanCount int    `json:"wirelesslan_count"`
				Depth            int    `json:"_depth"`
			} `json:"group"`
			Status struct {
				ValueLabel
			} `json:"status"`
			Vlan struct {
				Id          uint   `json:"id"`
				Url         string `json:"url"`
				Display     string `json:"display"`
				Vid         string `json:"vid"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"vlan"`
			Tenant struct {
				CommonFieldsSlug
				Description string `json:"description"`
			} `json:"tenant"`
			AuthType struct {
				ValueLabel
			} `json:"auth_type"`
			AuthCipher struct {
				ValueLabel
			} `json:"auth_cipher"`
			AuthPsk  string `json:"auth_psk"`
			Comments string `json:"comments"`
			Tags     []struct {
				CommonFieldsSlug
				Color string `json:"color"`
			} `json:"tags"`
			Created     string `json:"created"`
			LastUpdated string `json:"last_updated,omitempty"`
		} `json:"wireless_lans,omitempty"`
		Vrf struct {
			Id          uint   `json:"id,omitempty"`
			Url         string `json:"url,omitempty"`
			Display     string `json:"display,omitempty"`
			Name        string `json:"name,omitempty"`
			Rd          string `json:"rd,omitempty"`
			Description string `json:"description,omitempty"`
			PrefixCount int    `json:"prefix_count,omitempty"`
		} `json:"vrf,omitempty"`
		L2VpnTermination struct {
			idUrlDisplay
			L2Vpn struct {
				idUrlDisplay
				Identifier float64 `json:"identifier,omitempty"`
				Name       string  `json:"name,omitempty"`
				Slug       string  `json:"slug,omitempty"`
				Type       struct {
					ValueLabel
				} `json:"type,omitempty"`
				Description string `json:"description,omitempty"`
			} `json:"l2vpn,omitempty"`
		} `json:"l2vpn_termination,omitempty"`
		ConnectedEndpoints          []string `json:"connected_endpoints"`
		ConnectedEndpointsType      string   `json:"connected_endpoints_type"`
		ConnectedEndpointsReachable bool     `json:"connected_endpoints_reachable"`
		Tags                        []struct {
			CommonFieldsSlug
			Color string `json:"color,omitempty"`
		} `json:"tags"`
		Created          string `json:"created"`
		LastUpdated      string `json:"last_updated"`
		CountIpaddresses int    `json:"count_ipaddresses,omitempty"`
		CountFhrpGroups  int    `json:"count_fhrp_groups,omitempty"`
		Occupied         bool   `json:"_occupied,omitempty"`
	} `json:"results"`
}

type idUrlDisplay struct {
	Id      int    `json:"id,omitempty"`
	Url     string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
}

type family struct {
	Value int    `json:"value,omitempty"`
	Label string `json:"label,omitempty"`
}

type addressDescription struct {
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
}

var responseObjectInterfaces = new(interfaces)

// GetDcimInterfacesCmd represents the getDcimInterfaces command
var GetDcimInterfacesCmd = &cobra.Command{
	Use:   "getDcimInterfaces",
	Short: "GET a list of interface objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of interface objects`,
	Run: func(cmd *cobra.Command, args []string) {

		ApiConnectionNonID(responseObjectInterfaces, "GET", "cmd.dcim.dcim_api_url.interfaces")

		if responseObjectInterfaces.Count > 0 {
			color.Cyan("\n  Total ABC Interfaces: "+color.YellowString("%d"), responseObjectInterfaces.Count)
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
				if result.Device.DeviceType.Id > 0 {
					color.Cyan("\t  Device Type: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.DeviceType.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.DeviceType.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.DeviceType.Display))
					color.Cyan("\t    Manufacturer: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Device.DeviceType.Manufacturer.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Device.DeviceType.Manufacturer.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Device.DeviceType.Manufacturer.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Device.DeviceType.Manufacturer.Name))
					color.Cyan("\t      Slug: " + color.YellowString("%s", result.Device.DeviceType.Manufacturer.Slug))
					if result.Device.DeviceType.Manufacturer.Description != "" {
						color.Cyan("\t      Description: " + color.YellowString("%s", result.Device.DeviceType.Manufacturer.Description))
					} else {
						color.Cyan("\t      Description: " + color.RedString("No description found for interface: %s", color.YellowString("%s", result.Display)))
					}
					if result.Device.DeviceType.Manufacturer.DevicetypeCount > 0 {
						color.Cyan("\t      Device Type Count: " + color.YellowString("%d", result.Device.DeviceType.Manufacturer.DevicetypeCount))
					} else {
						color.Cyan("\t      Device Type Count: " + color.RedString("No device type count found for interface: %s", color.YellowString("%s", result.Display)))
					}
					color.Cyan("\t    Model: " + color.YellowString("%s", result.Device.DeviceType.Model))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.DeviceType.Slug))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.DeviceType.Description))
					color.Cyan("\t    Device Count: " + color.YellowString("%d", result.Device.DeviceType.DeviceCount))
				} else {
					color.Cyan("\t  Device Type: " + color.RedString("No device type found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Role.Id > 0 {
					color.Cyan("\t  Role: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Role.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Role.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Role.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Role.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.Role.Slug))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Role.Description))
					color.Cyan("\t    Device Count: " + color.YellowString("%d", result.Device.Role.DeviceCount))
					color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", result.Device.Role.VirtualmachineCount))
				} else {
					color.Cyan("\t  Role: " + color.RedString("No role found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Tenant.Id > 0 {
					color.Cyan("\t  Tenant: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Tenant.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Tenant.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Tenant.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Tenant.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.Tenant.Slug))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Tenant.Description))
				} else {
					color.Cyan("\t  Tenant: " + color.RedString("No tenant found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Platform.Id > 0 {
					color.Cyan("\t  Platform: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Platform.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Platform.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Platform.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Platform.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.Platform.Slug))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Platform.Description))
					color.Cyan("\t    Device Count: " + color.YellowString("%d", result.Device.Platform.DeviceCount))
					color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", result.Device.Platform.VirtualmachineCount))
				} else {
					color.Cyan("\t  Platform: " + color.RedString("No platform found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Serial != "" {
					color.Cyan("\t  Serial: " + color.YellowString("%s", result.Device.Serial))
				} else {
					color.Cyan("\t  Serial: " + color.RedString("No serial found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.AssetTag != "" {
					color.Cyan("\t  Asset Tag: " + color.YellowString("%s", result.Device.AssetTag))
				} else {
					color.Cyan("\t  Asset Tag: " + color.RedString("No asset tag found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Site.Id > 0 {
					color.Cyan("\t  Site: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Site.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Site.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Site.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Site.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.Site.Slug))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Site.Description))
				} else {
					color.Cyan("\t  Site: " + color.RedString("No site found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Location.Id > 0 {
					color.Cyan("\t  Location: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Location.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Location.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Location.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Location.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.Location.Slug))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Location.Description))
					color.Cyan("\t    Rack Count: " + color.YellowString("%d", result.Device.Location.RackCount))
					color.Cyan("\t    Depth: " + color.YellowString("%d", result.Device.Location.Depth))
				} else {
					color.Cyan("\t  Location: " + color.RedString("No location found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Rack.Id > 0 {
					color.Cyan("\t  Rack: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Rack.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Rack.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Rack.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Rack.Name))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Rack.Description))
					color.Cyan("\t    Device Count: " + color.YellowString("%d", result.Device.Rack.DeviceCount))
				} else {
					color.Cyan("\t  Rack: " + color.RedString("No rack found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Position > 0 {
					color.Cyan("\t  Position: " + color.YellowString("%d", result.Device.Position))
				} else {
					color.Cyan("\t  Position: " + color.RedString("No position found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Face.Value != "" {
					color.Cyan("\t  Face: ")
					color.Cyan("\t    Value: " + color.YellowString("%s", result.Device.Face.Value))
					color.Cyan("\t    Label: " + color.YellowString("%s", result.Device.Face.Label))
				} else {
					color.Cyan("\t  Face: " + color.RedString("No face found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Latitude > 0 {
					color.Cyan("\t  Latitude: " + color.YellowString(".%f4", result.Device.Latitude))
				} else {
					color.Cyan("\t  Latitude: " + color.RedString("No latitude found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Longitude > 0 {
					color.Cyan("\t  Longitude: " + color.YellowString("%d", result.Device.Longitude))
				} else {
					color.Cyan("\t  Longitude: " + color.RedString("No longitude found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.ParentDevice.Id > 0 {
					color.Cyan("\t  Parent Device: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.ParentDevice.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.ParentDevice.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.ParentDevice.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.ParentDevice.Name))
				} else {
					color.Cyan("\t  Parent Device: " + color.RedString("No parent device found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Status.Value != "" {
					color.Cyan("\t  Status: ")
					color.Cyan("\t    Value: " + color.YellowString("%s", result.Device.Status.Value))
					color.Cyan("\t    Label: " + color.YellowString("%s", result.Device.Status.Label))
				} else {
					color.Cyan("\t  Status: " + color.RedString("No status found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Airflow.Value != "" {
					color.Cyan("\t  AirFlow: ")
					color.Cyan("\t    Value: " + color.YellowString("%s", result.Device.Airflow.Value))
					color.Cyan("\t    Label: " + color.YellowString("%s", result.Device.Airflow.Label))
				} else {
					color.Cyan("\t  AirFlow: " + color.RedString("No airflow found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.PrimaryIp.Id > 0 {
					color.Cyan("\t  Primary IP: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.PrimaryIp.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.PrimaryIp.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.PrimaryIp.Display))
					color.Cyan("\t    Family: ")
					color.Cyan("\t      Value: " + color.YellowString("%d", result.Device.PrimaryIp.Family.Value))
					color.Cyan("\t      Label: " + color.YellowString("%s", result.Device.PrimaryIp.Family.Label))
					color.Cyan("\t    Address: " + color.YellowString("%s", result.Device.PrimaryIp.Address))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.PrimaryIp.Description))
				} else {
					color.Cyan("\t  Primary IP: " + color.RedString("No primary IP found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.PrimaryIp4.Id > 0 {
					color.Cyan("\t  Primary IPv4: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.PrimaryIp4.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.PrimaryIp4.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.PrimaryIp4.Display))
					color.Cyan("\t    Family: ")
					color.Cyan("\t      Value: " + color.YellowString("%d", result.Device.PrimaryIp4.Family.Value))
					color.Cyan("\t      Label: " + color.YellowString("%s", result.Device.PrimaryIp4.Family.Label))
					color.Cyan("\t    Address: " + color.YellowString("%s", result.Device.PrimaryIp4.Address))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.PrimaryIp4.Description))
				} else {
					color.Cyan("\t  Primary IPv4: " + color.RedString("No primary IPv4 found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.PrimaryIp6.Id > 0 {
					color.Cyan("\t  Primary IPv6: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.PrimaryIp6.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.PrimaryIp6.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.PrimaryIp6.Display))
					color.Cyan("\t    Family: ")
					color.Cyan("\t      Value: " + color.YellowString("%d", result.Device.PrimaryIp6.Family.Value))
					color.Cyan("\t      Label: " + color.YellowString("%s", result.Device.PrimaryIp6.Family.Label))
					color.Cyan("\t    Address: " + color.YellowString("%s", result.Device.PrimaryIp6.Address))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.PrimaryIp6.Description))
				} else {
					color.Cyan("\t  Primary IPv6: " + color.RedString("No primary IPv6 found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.OobIp.Id > 0 {
					color.Cyan("\t  OOB IP: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.OobIp.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.OobIp.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.OobIp.Display))
					color.Cyan("\t    Family: ")
					color.Cyan("\t      Value: " + color.YellowString("%d", result.Device.OobIp.Family.Value))
					color.Cyan("\t      Label: " + color.YellowString("%s", result.Device.OobIp.Family.Label))
					color.Cyan("\t    Address: " + color.YellowString("%s", result.Device.OobIp.Address))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.OobIp.Description))
				} else {
					color.Cyan("\t  OOB IP: " + color.RedString("No oob IP found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Cluster.Id > 0 {
					color.Cyan("\t  Cluster: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Cluster.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Cluster.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Cluster.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Cluster.Name))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Cluster.Description))
					color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", result.Device.Cluster.VirtualmachineCount))
				} else {
					color.Cyan("\t  Cluster: " + color.RedString("No cluster found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.VirtualChassis.Id > 0 {
					color.Cyan("\t  Virtual Chassis: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.VirtualChassis.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.VirtualChassis.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.VirtualChassis.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.VirtualChassis.Name))
					color.Cyan("\t    Master: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", result.Device.VirtualChassis.Master.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", result.Device.VirtualChassis.Master.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", result.Device.VirtualChassis.Master.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", result.Device.VirtualChassis.Master.Name))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.VirtualChassis.Description))
					color.Cyan("\t    Member Count: " + color.YellowString("%d", result.Device.VirtualChassis.MemberCount))
				} else {
					color.Cyan("\t  Virtual Chassis: " + color.RedString("No virtual chassis found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.VcPosition > 0 {
					color.Cyan("\t  VC Position: " + color.YellowString("%d", result.Device.VcPosition))
				} else {
					color.Cyan("\t  VC Position: " + color.RedString("No vc position found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.VcPriority > 0 {
					color.Cyan("\t  VC Priority: " + color.YellowString("%d", result.Device.VcPriority))
				} else {
					color.Cyan("\t  VC Priority: " + color.RedString("No vc priority found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Description != "" {
					color.Cyan("\t  Description: " + color.YellowString("%s", result.Device.Description))
				} else {
					color.Cyan("\t  Description: " + color.RedString("No description found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.Comments != "" {
					color.Cyan("\t  Comments: " + color.YellowString("%s", result.Device.Comments))
				} else {
					color.Cyan("\t  Comments: " + color.RedString("No comments found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.ConfigTemplate.Id > 0 {
					color.Cyan("\t  Config Template: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.ConfigTemplate.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.ConfigTemplate.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.ConfigTemplate.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.ConfigTemplate.Name))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.ConfigTemplate.Description))
				} else {
					color.Cyan("\t  Config Template: " + color.RedString("No config template found for interface: %s", color.YellowString("%s", result.Display)))
				}
				color.Cyan("\t  Local Context Data: " + color.YellowString("%s", result.Device.LocalContextData))
				for _, tag := range result.Device.Tags {
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
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Device.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.Device.LastUpdated))
				if result.Device.ConsolePortCount > 0 {
					color.Cyan("\tConsole Port Count: " + color.YellowString("%d", result.Device.ConsolePortCount))
				} else {
					color.Cyan("\tConsole Port Count: " + color.RedString("No console port count found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.ConsoleServerPortCount > 0 {
					color.Cyan("\tConsole Server Port Count: " + color.YellowString("%d", result.Device.ConsoleServerPortCount))
				} else {
					color.Cyan("\tConsole Server Port Count: " + color.RedString("No console server port count found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.PowerPortCount > 0 {
					color.Cyan("\tPower Port Count: " + color.YellowString("%d", result.Device.PowerPortCount))
				} else {
					color.Cyan("\tPower Port Count: " + color.RedString("No power port count found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.PowerOutletCount > 0 {
					color.Cyan("\tPower Outlet Count: " + color.YellowString("%d", result.Device.PowerOutletCount))
				} else {
					color.Cyan("\tPower Outlet Count: " + color.RedString("No power outlet count found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.InterfaceCount > 0 {
					color.Cyan("\tInterface Count: " + color.YellowString("%d", result.Device.InterfaceCount))
				} else {
					color.Cyan("\tInterface Count: " + color.RedString("No interface count found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.FrontPortCount > 0 {
					color.Cyan("\tFront Port Count: " + color.YellowString("%d", result.Device.FrontPortCount))
				} else {
					color.Cyan("\tFront Port Count: " + color.RedString("No front port count found for interface: %s", color.YellowString("%s", result.Display)))
				}
				if result.Device.RearPortCount > 0 {
					color.Cyan("\tRear Port Count: " + color.YellowString("%d", result.Device.RearPortCount))
				} else {
					color.Cyan("\tRear Port Count: " + color.RedString("No rear port count found for interface: %s", color.YellowString("%s", result.Display)))
				}
				for _, vdc := range result.Vdcs {
					if vdc.Id > 0 {
						color.Cyan("\tVDC: ")
						color.Cyan("\t  ID: " + color.YellowString("%d", vdc.Id))
						color.Cyan("\t  URL: " + color.YellowString("%s", vdc.Url))
						color.Cyan("\t  Display: " + color.YellowString("%s", vdc.Display))
						color.Cyan("\t  Name: " + color.YellowString("%s", vdc.Name))
						color.Cyan("\t  Device: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Display))
						color.Cyan("\t  Device Type: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.DeviceType.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.DeviceType.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.DeviceType.Display))
						color.Cyan("\t    Manufacturer: ")
						color.Cyan("\t      ID: " + color.YellowString("%d", vdc.Device.DeviceType.Manufacturer.Id))
						color.Cyan("\t      URL: " + color.YellowString("%s", vdc.Device.DeviceType.Manufacturer.Url))
						color.Cyan("\t      Display: " + color.YellowString("%s", vdc.Device.DeviceType.Manufacturer.Display))
						color.Cyan("\t      Name: " + color.YellowString("%s", vdc.Device.DeviceType.Manufacturer.Name))
						color.Cyan("\t      Slug: " + color.YellowString("%s", vdc.Device.DeviceType.Manufacturer.Slug))
						color.Cyan("\t      Description: " + color.YellowString("%s", vdc.Device.DeviceType.Manufacturer.Description))
						color.Cyan("\t      Device Type Count: " + color.YellowString("%d", vdc.Device.DeviceType.Manufacturer.DevicetypeCount))
						color.Cyan("\t    Model: " + color.YellowString("%s", vdc.Device.DeviceType.Model))
						color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.DeviceType.Slug))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.DeviceType.Description))
						color.Cyan("\t    Device Count: " + color.YellowString("%d", vdc.Device.DeviceType.DeviceCount))
						color.Cyan("\t  Role: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Role.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Role.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Role.Display))
						color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Role.Name))
						color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.Role.Slug))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Role.Description))
						color.Cyan("\t    Device Count: " + color.YellowString("%d", vdc.Device.Role.DeviceCount))
						color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", vdc.Device.Role.VirtualmachineCount))
						color.Cyan("\t  Tenant: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Tenant.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Tenant.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Tenant.Display))
						color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Tenant.Name))
						color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.Tenant.Slug))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Tenant.Description))
						color.Cyan("\t  Platform: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Platform.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Platform.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Platform.Display))
						color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Platform.Name))
						color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.Platform.Slug))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Platform.Description))
						color.Cyan("\t    Device Count: " + color.YellowString("%d", vdc.Device.Platform.DeviceCount))
						color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", result.Device.Platform.VirtualmachineCount))
						color.Cyan("\t  Serial: " + color.YellowString("%s", vdc.Device.Serial))
						color.Cyan("\t  Asset Tag: " + color.YellowString("%s", vdc.Device.AssetTag))
						color.Cyan("\t  Site: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Site.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Site.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Site.Display))
						color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Site.Name))
						color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.Site.Slug))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Site.Description))
						color.Cyan("\t  Location: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Location.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Location.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Location.Display))
						color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Location.Name))
						color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.Location.Slug))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Location.Description))
						color.Cyan("\t    Rack Count: " + color.YellowString("%d", vdc.Device.Location.RackCount))
						color.Cyan("\t    Depth: " + color.YellowString("%d", vdc.Device.Location.Depth))
						color.Cyan("\t  Rack: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Rack.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Rack.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Rack.Display))
						color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Rack.Name))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Rack.Description))
						color.Cyan("\t    Device Count: " + color.YellowString("%d", vdc.Device.Rack.DeviceCount))
						color.Cyan("\t  Position: " + color.YellowString("%d", vdc.Device.Position))
						color.Cyan("\t  Face: ")
						color.Cyan("\t    Value: " + color.YellowString("%d", vdc.Device.Face.Value))
						color.Cyan("\t    Label: " + color.YellowString("%s", vdc.Device.Face.Label))
						color.Cyan("\t  Latitude: " + color.YellowString("%d", vdc.Device.Latitude))
						color.Cyan("\t  Longitude: " + color.YellowString("%d", vdc.Device.Longitude))
						color.Cyan("\t  Parent Device: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.ParentDevice.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.ParentDevice.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.ParentDevice.Display))
						color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.ParentDevice.Name))
						color.Cyan("\t  Status: ")
						color.Cyan("\t    Value: " + color.YellowString("%d", vdc.Device.Status.Value))
						color.Cyan("\t    Label: " + color.YellowString("%s", vdc.Device.Status.Label))
						color.Cyan("\t  AirFlow: ")
						color.Cyan("\t    Value: " + color.YellowString("%s", vdc.Device.Airflow.Value))
						color.Cyan("\t    Label: " + color.YellowString("%s", vdc.Device.Airflow.Label))
						color.Cyan("\t  Primary IP: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.PrimaryIp.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.PrimaryIp.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.PrimaryIp.Display))
						color.Cyan("\t    Family: ")
						color.Cyan("\t      Value: " + color.YellowString("%d", vdc.Device.PrimaryIp.Family.Value))
						color.Cyan("\t      Label: " + color.YellowString("%s", vdc.Device.PrimaryIp.Family.Label))
						color.Cyan("\t    Address: " + color.YellowString("%s", vdc.Device.PrimaryIp.Address))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.PrimaryIp.Description))
						color.Cyan("\t  Primary IPv4: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.PrimaryIp4.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.PrimaryIp4.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.PrimaryIp4.Display))
						color.Cyan("\t    Family: ")
						color.Cyan("\t      Value: " + color.YellowString("%d", vdc.Device.PrimaryIp4.Family.Value))
						color.Cyan("\t      Label: " + color.YellowString("%s", vdc.Device.PrimaryIp4.Family.Label))
						color.Cyan("\t    Address: " + color.YellowString("%s", vdc.Device.PrimaryIp4.Address))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.PrimaryIp4.Description))
						color.Cyan("\t  Primary IPv6: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.PrimaryIp6.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.PrimaryIp6.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.PrimaryIp6.Display))
						color.Cyan("\t    Family: ")
						color.Cyan("\t      Value: " + color.YellowString("%d", vdc.Device.PrimaryIp6.Family.Value))
						color.Cyan("\t      Label: " + color.YellowString("%s", vdc.Device.PrimaryIp6.Family.Label))
						color.Cyan("\t    Address: " + color.YellowString("%s", vdc.Device.PrimaryIp6.Address))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.PrimaryIp6.Description))
						color.Cyan("\t  OOP IP: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.OobIp.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.OobIp.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.OobIp.Display))
						color.Cyan("\t    Family: ")
						color.Cyan("\t      Value: " + color.YellowString("%d", vdc.Device.OobIp.Family.Value))
						color.Cyan("\t      Label: " + color.YellowString("%s", vdc.Device.OobIp.Family.Label))
						color.Cyan("\t    Address: " + color.YellowString("%s", vdc.Device.OobIp.Address))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.OobIp.Description))
						color.Cyan("\t  Cluster: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Cluster.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Cluster.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Cluster.Display))
						color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Cluster.Name))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Cluster.Description))
						color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", vdc.Device.Cluster.VirtualmachineCount))
						color.Cyan("\t  Virtual Chassis: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.VirtualChassis.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.VirtualChassis.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.VirtualChassis.Display))
						color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.VirtualChassis.Name))
						color.Cyan("\t    Master: ")
						color.Cyan("\t      ID: " + color.YellowString("%d", vdc.Device.VirtualChassis.Master.Id))
						color.Cyan("\t      URL: " + color.YellowString("%s", vdc.Device.VirtualChassis.Master.Url))
						color.Cyan("\t      Display: " + color.YellowString("%s", vdc.Device.VirtualChassis.Master.Display))
						color.Cyan("\t      Name: " + color.YellowString("%s", vdc.Device.VirtualChassis.Master.Name))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.VirtualChassis.Description))
						color.Cyan("\t    Member Count: " + color.YellowString("%d", vdc.Device.VirtualChassis.MemberCount))
						color.Cyan("\t  VC Position: " + color.YellowString("%d", vdc.Device.VcPosition))
						color.Cyan("\t  VC Priority: " + color.YellowString("%d", vdc.Device.VcPriority))
						color.Cyan("\t  Description: " + color.YellowString("%s", vdc.Device.Description))
						color.Cyan("\t  Comments: " + color.YellowString("%s", vdc.Device.Comments))
						color.Cyan("\t  Config Template: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.ConfigTemplate.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.ConfigTemplate.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.ConfigTemplate.Display))
						color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.ConfigTemplate.Name))
						color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.ConfigTemplate.Description))
						color.Cyan("\t  Local Context Data: " + color.YellowString("%s", vdc.Device.LocalContextData))
						for _, tag := range vdc.Device.Tags {
							if tag.Id != 0 {
								color.Cyan("\tTags: ")
								color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
								color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
								color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
								color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
								color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
								color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
							} else {
								color.Cyan("\tTags: " + color.RedString("No tags found for interface: %s", color.YellowString("%s", vdc.Display)))
							}
						}
						color.Cyan("\tCreated: " + color.YellowString("%s", vdc.Device.Created))
						color.Cyan("\tLast Updated: " + color.YellowString("%s", vdc.Device.LastUpdated))
						color.Cyan("\tInterface Count: " + color.YellowString("%d", vdc.Device.InterfaceCount))
					} else {
						color.Red("\tVirtual Device Context Not found for interface: " + color.YellowString("%s", vdc.Display))
					}
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
					color.Cyan("\t    Local Context Data: " + color.YellowString("%s", result.Module.Device.LocalContextData))
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
				display := color.HiGreenString("\tAll Netbox interface objects have been successfully displayed...")
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

func ApiConnectionNextPageInterfaces[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := *responseObjectInterfaces.Next

	color.Yellow("\n  Getting Netbox API objects from %s\n", suffix)
	token := config.GetString("cmd.token_key")

	err := CheckSSL(rootURL)
	if err != nil {
		fmt.Println("  SSL certificate is not valid: ", err)
	} else {
		color.Cyan("  SSL certificate is valid for: " + color.YellowString("%s", rootURL))
	}

	executeAPIRequest(httpMethod, fullAPIPath, token, r)
	if err != nil {
		log.Fatalf("Error getting Netbox API objects: %s\n", err)
	}
}

func nextPageInterfaces() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\tDo you want to continue to the next page of interface objects? ['Y' or 'yes'/'n' or 'no']: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageInterfaces(responseObjectInterfaces, "GET", *responseObjectInterfaces.Next)
		displayInterfacesOutput()
	case "n", "no":
		color.HiMagenta("\n\tExiting the ABC-netbox-cli application...\n\n")
		os.Exit(0)
	default:
		color.Cyan("\n\tInvalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func displayInterfacesOutput() {
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
		if result.Device.DeviceType.Id > 0 {
			color.Cyan("\t  Device Type: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.DeviceType.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.DeviceType.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.DeviceType.Display))
			color.Cyan("\t    Manufacturer: ")
			color.Cyan("\t      ID: " + color.YellowString("%d", result.Device.DeviceType.Manufacturer.Id))
			color.Cyan("\t      URL: " + color.YellowString("%s", result.Device.DeviceType.Manufacturer.Url))
			color.Cyan("\t      Display: " + color.YellowString("%s", result.Device.DeviceType.Manufacturer.Display))
			color.Cyan("\t      Name: " + color.YellowString("%s", result.Device.DeviceType.Manufacturer.Name))
			color.Cyan("\t      Slug: " + color.YellowString("%s", result.Device.DeviceType.Manufacturer.Slug))
			if result.Device.DeviceType.Manufacturer.Description != "" {
				color.Cyan("\t      Description: " + color.YellowString("%s", result.Device.DeviceType.Manufacturer.Description))
			} else {
				color.Cyan("\t      Description: " + color.RedString("No description found for interface: %s", color.YellowString("%s", result.Display)))
			}
			if result.Device.DeviceType.Manufacturer.DevicetypeCount > 0 {
				color.Cyan("\t      Device Type Count: " + color.YellowString("%d", result.Device.DeviceType.Manufacturer.DevicetypeCount))
			} else {
				color.Cyan("\t      Device Type Count: " + color.RedString("No device type count found for interface: %s", color.YellowString("%s", result.Display)))
			}
			color.Cyan("\t    Model: " + color.YellowString("%s", result.Device.DeviceType.Model))
			color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.DeviceType.Slug))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.DeviceType.Description))
			color.Cyan("\t    Device Count: " + color.YellowString("%d", result.Device.DeviceType.DeviceCount))
		} else {
			color.Cyan("\t  Device Type: " + color.RedString("No device type found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Role.Id > 0 {
			color.Cyan("\t  Role: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Role.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Role.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Role.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Role.Name))
			color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.Role.Slug))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Role.Description))
			color.Cyan("\t    Device Count: " + color.YellowString("%d", result.Device.Role.DeviceCount))
			color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", result.Device.Role.VirtualmachineCount))
		} else {
			color.Cyan("\t  Role: " + color.RedString("No role found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Tenant.Id > 0 {
			color.Cyan("\t  Tenant: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Tenant.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Tenant.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Tenant.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Tenant.Name))
			color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.Tenant.Slug))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Tenant.Description))
		} else {
			color.Cyan("\t  Tenant: " + color.RedString("No tenant found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Platform.Id > 0 {
			color.Cyan("\t  Platform: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Platform.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Platform.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Platform.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Platform.Name))
			color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.Platform.Slug))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Platform.Description))
			color.Cyan("\t    Device Count: " + color.YellowString("%d", result.Device.Platform.DeviceCount))
			color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", result.Device.Platform.VirtualmachineCount))
		} else {
			color.Cyan("\t  Platform: " + color.RedString("No platform found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Serial != "" {
			color.Cyan("\t  Serial: " + color.YellowString("%s", result.Device.Serial))
		} else {
			color.Cyan("\t  Serial: " + color.RedString("No serial found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.AssetTag != "" {
			color.Cyan("\t  Asset Tag: " + color.YellowString("%s", result.Device.AssetTag))
		} else {
			color.Cyan("\t  Asset Tag: " + color.RedString("No asset tag found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Site.Id > 0 {
			color.Cyan("\t  Site: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Site.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Site.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Site.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Site.Name))
			color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.Site.Slug))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Site.Description))
		} else {
			color.Cyan("\t  Site: " + color.RedString("No site found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Location.Id > 0 {
			color.Cyan("\t  Location: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Location.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Location.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Location.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Location.Name))
			color.Cyan("\t    Slug: " + color.YellowString("%s", result.Device.Location.Slug))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Location.Description))
			color.Cyan("\t    Rack Count: " + color.YellowString("%d", result.Device.Location.RackCount))
			color.Cyan("\t    Depth: " + color.YellowString("%d", result.Device.Location.Depth))
		} else {
			color.Cyan("\t  Location: " + color.RedString("No location found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Rack.Id > 0 {
			color.Cyan("\t  Rack: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Rack.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Rack.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Rack.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Rack.Name))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Rack.Description))
			color.Cyan("\t    Device Count: " + color.YellowString("%d", result.Device.Rack.DeviceCount))
		} else {
			color.Cyan("\t  Rack: " + color.RedString("No rack found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Position > 0 {
			color.Cyan("\t  Position: " + color.YellowString("%d", result.Device.Position))
		} else {
			color.Cyan("\t  Position: " + color.RedString("No position found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Face.Value != "" {
			color.Cyan("\t  Face: ")
			color.Cyan("\t    Value: " + color.YellowString("%s", result.Device.Face.Value))
			color.Cyan("\t    Label: " + color.YellowString("%s", result.Device.Face.Label))
		} else {
			color.Cyan("\t  Face: " + color.RedString("No face found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Latitude > 0 {
			color.Cyan("\t  Latitude: " + color.YellowString(".%f4", result.Device.Latitude))
		} else {
			color.Cyan("\t  Latitude: " + color.RedString("No latitude found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Longitude > 0 {
			color.Cyan("\t  Longitude: " + color.YellowString("%d", result.Device.Longitude))
		} else {
			color.Cyan("\t  Longitude: " + color.RedString("No longitude found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.ParentDevice.Id > 0 {
			color.Cyan("\t  Parent Device: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.ParentDevice.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.ParentDevice.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.ParentDevice.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.ParentDevice.Name))
		} else {
			color.Cyan("\t  Parent Device: " + color.RedString("No parent device found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Status.Value != "" {
			color.Cyan("\t  Status: ")
			color.Cyan("\t    Value: " + color.YellowString("%s", result.Device.Status.Value))
			color.Cyan("\t    Label: " + color.YellowString("%s", result.Device.Status.Label))
		} else {
			color.Cyan("\t  Status: " + color.RedString("No status found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Airflow.Value != "" {
			color.Cyan("\t  AirFlow: ")
			color.Cyan("\t    Value: " + color.YellowString("%s", result.Device.Airflow.Value))
			color.Cyan("\t    Label: " + color.YellowString("%s", result.Device.Airflow.Label))
		} else {
			color.Cyan("\t  AirFlow: " + color.RedString("No airflow found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.PrimaryIp.Id > 0 {
			color.Cyan("\t  Primary IP: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.PrimaryIp.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.PrimaryIp.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.PrimaryIp.Display))
			color.Cyan("\t    Family: ")
			color.Cyan("\t      Value: " + color.YellowString("%d", result.Device.PrimaryIp.Family.Value))
			color.Cyan("\t      Label: " + color.YellowString("%s", result.Device.PrimaryIp.Family.Label))
			color.Cyan("\t    Address: " + color.YellowString("%s", result.Device.PrimaryIp.Address))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.PrimaryIp.Description))
		} else {
			color.Cyan("\t  Primary IP: " + color.RedString("No primary IP found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.PrimaryIp4.Id > 0 {
			color.Cyan("\t  Primary IPv4: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.PrimaryIp4.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.PrimaryIp4.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.PrimaryIp4.Display))
			color.Cyan("\t    Family: ")
			color.Cyan("\t      Value: " + color.YellowString("%d", result.Device.PrimaryIp4.Family.Value))
			color.Cyan("\t      Label: " + color.YellowString("%s", result.Device.PrimaryIp4.Family.Label))
			color.Cyan("\t    Address: " + color.YellowString("%s", result.Device.PrimaryIp4.Address))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.PrimaryIp4.Description))
		} else {
			color.Cyan("\t  Primary IPv4: " + color.RedString("No primary IPv4 found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.PrimaryIp6.Id > 0 {
			color.Cyan("\t  Primary IPv6: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.PrimaryIp6.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.PrimaryIp6.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.PrimaryIp6.Display))
			color.Cyan("\t    Family: ")
			color.Cyan("\t      Value: " + color.YellowString("%d", result.Device.PrimaryIp6.Family.Value))
			color.Cyan("\t      Label: " + color.YellowString("%s", result.Device.PrimaryIp6.Family.Label))
			color.Cyan("\t    Address: " + color.YellowString("%s", result.Device.PrimaryIp6.Address))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.PrimaryIp6.Description))
		} else {
			color.Cyan("\t  Primary IPv6: " + color.RedString("No primary IPv6 found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.OobIp.Id > 0 {
			color.Cyan("\t  OOB IP: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.OobIp.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.OobIp.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.OobIp.Display))
			color.Cyan("\t    Family: ")
			color.Cyan("\t      Value: " + color.YellowString("%d", result.Device.OobIp.Family.Value))
			color.Cyan("\t      Label: " + color.YellowString("%s", result.Device.OobIp.Family.Label))
			color.Cyan("\t    Address: " + color.YellowString("%s", result.Device.OobIp.Address))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.OobIp.Description))
		} else {
			color.Cyan("\t  OOB IP: " + color.RedString("No oob IP found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Cluster.Id > 0 {
			color.Cyan("\t  Cluster: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.Cluster.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.Cluster.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.Cluster.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.Cluster.Name))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.Cluster.Description))
			color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", result.Device.Cluster.VirtualmachineCount))
		} else {
			color.Cyan("\t  Cluster: " + color.RedString("No cluster found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.VirtualChassis.Id > 0 {
			color.Cyan("\t  Virtual Chassis: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.VirtualChassis.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.VirtualChassis.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.VirtualChassis.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.VirtualChassis.Name))
			color.Cyan("\t    Master: ")
			color.Cyan("\t      ID: " + color.YellowString("%d", result.Device.VirtualChassis.Master.Id))
			color.Cyan("\t      URL: " + color.YellowString("%s", result.Device.VirtualChassis.Master.Url))
			color.Cyan("\t      Display: " + color.YellowString("%s", result.Device.VirtualChassis.Master.Display))
			color.Cyan("\t      Name: " + color.YellowString("%s", result.Device.VirtualChassis.Master.Name))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.VirtualChassis.Description))
			color.Cyan("\t    Member Count: " + color.YellowString("%d", result.Device.VirtualChassis.MemberCount))
		} else {
			color.Cyan("\t  Virtual Chassis: " + color.RedString("No virtual chassis found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.VcPosition > 0 {
			color.Cyan("\t  VC Position: " + color.YellowString("%d", result.Device.VcPosition))
		} else {
			color.Cyan("\t  VC Position: " + color.RedString("No vc position found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.VcPriority > 0 {
			color.Cyan("\t  VC Priority: " + color.YellowString("%d", result.Device.VcPriority))
		} else {
			color.Cyan("\t  VC Priority: " + color.RedString("No vc priority found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Description != "" {
			color.Cyan("\t  Description: " + color.YellowString("%s", result.Device.Description))
		} else {
			color.Cyan("\t  Description: " + color.RedString("No description found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.Comments != "" {
			color.Cyan("\t  Comments: " + color.YellowString("%s", result.Device.Comments))
		} else {
			color.Cyan("\t  Comments: " + color.RedString("No comments found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.ConfigTemplate.Id > 0 {
			color.Cyan("\t  Config Template: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Device.ConfigTemplate.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Device.ConfigTemplate.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Device.ConfigTemplate.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Device.ConfigTemplate.Name))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Device.ConfigTemplate.Description))
		} else {
			color.Cyan("\t  Config Template: " + color.RedString("No config template found for interface: %s", color.YellowString("%s", result.Display)))
		}
		color.Cyan("\t  Local Context Data: " + color.YellowString("%s", result.Device.LocalContextData))
		for _, tag := range result.Device.Tags {
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
		color.Cyan("\tCreated: " + color.YellowString("%s", result.Device.Created))
		color.Cyan("\tLast Updated: " + color.YellowString("%s", result.Device.LastUpdated))
		if result.Device.ConsolePortCount > 0 {
			color.Cyan("\tConsole Port Count: " + color.YellowString("%d", result.Device.ConsolePortCount))
		} else {
			color.Cyan("\tConsole Port Count: " + color.RedString("No console port count found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.ConsoleServerPortCount > 0 {
			color.Cyan("\tConsole Server Port Count: " + color.YellowString("%d", result.Device.ConsoleServerPortCount))
		} else {
			color.Cyan("\tConsole Server Port Count: " + color.RedString("No console server port count found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.PowerPortCount > 0 {
			color.Cyan("\tPower Port Count: " + color.YellowString("%d", result.Device.PowerPortCount))
		} else {
			color.Cyan("\tPower Port Count: " + color.RedString("No power port count found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.PowerOutletCount > 0 {
			color.Cyan("\tPower Outlet Count: " + color.YellowString("%d", result.Device.PowerOutletCount))
		} else {
			color.Cyan("\tPower Outlet Count: " + color.RedString("No power outlet count found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.InterfaceCount > 0 {
			color.Cyan("\tInterface Count: " + color.YellowString("%d", result.Device.InterfaceCount))
		} else {
			color.Cyan("\tInterface Count: " + color.RedString("No interface count found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.FrontPortCount > 0 {
			color.Cyan("\tFront Port Count: " + color.YellowString("%d", result.Device.FrontPortCount))
		} else {
			color.Cyan("\tFront Port Count: " + color.RedString("No front port count found for interface: %s", color.YellowString("%s", result.Display)))
		}
		if result.Device.RearPortCount > 0 {
			color.Cyan("\tRear Port Count: " + color.YellowString("%d", result.Device.RearPortCount))
		} else {
			color.Cyan("\tRear Port Count: " + color.RedString("No rear port count found for interface: %s", color.YellowString("%s", result.Display)))
		}
		for _, vdc := range result.Vdcs {
			if vdc.Id > 0 {
				color.Cyan("\tVDC: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", vdc.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", vdc.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", vdc.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", vdc.Name))
				color.Cyan("\t  Device: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Display))
				color.Cyan("\t  Device Type: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.DeviceType.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.DeviceType.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.DeviceType.Display))
				color.Cyan("\t    Manufacturer: ")
				color.Cyan("\t      ID: " + color.YellowString("%d", vdc.Device.DeviceType.Manufacturer.Id))
				color.Cyan("\t      URL: " + color.YellowString("%s", vdc.Device.DeviceType.Manufacturer.Url))
				color.Cyan("\t      Display: " + color.YellowString("%s", vdc.Device.DeviceType.Manufacturer.Display))
				color.Cyan("\t      Name: " + color.YellowString("%s", vdc.Device.DeviceType.Manufacturer.Name))
				color.Cyan("\t      Slug: " + color.YellowString("%s", vdc.Device.DeviceType.Manufacturer.Slug))
				color.Cyan("\t      Description: " + color.YellowString("%s", vdc.Device.DeviceType.Manufacturer.Description))
				color.Cyan("\t      Device Type Count: " + color.YellowString("%d", vdc.Device.DeviceType.Manufacturer.DevicetypeCount))
				color.Cyan("\t    Model: " + color.YellowString("%s", vdc.Device.DeviceType.Model))
				color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.DeviceType.Slug))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.DeviceType.Description))
				color.Cyan("\t    Device Count: " + color.YellowString("%d", vdc.Device.DeviceType.DeviceCount))
				color.Cyan("\t  Role: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Role.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Role.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Role.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Role.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.Role.Slug))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Role.Description))
				color.Cyan("\t    Device Count: " + color.YellowString("%d", vdc.Device.Role.DeviceCount))
				color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", vdc.Device.Role.VirtualmachineCount))
				color.Cyan("\t  Tenant: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Tenant.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Tenant.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Tenant.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Tenant.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.Tenant.Slug))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Tenant.Description))
				color.Cyan("\t  Platform: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Platform.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Platform.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Platform.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Platform.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.Platform.Slug))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Platform.Description))
				color.Cyan("\t    Device Count: " + color.YellowString("%d", vdc.Device.Platform.DeviceCount))
				color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", result.Device.Platform.VirtualmachineCount))
				color.Cyan("\t  Serial: " + color.YellowString("%s", vdc.Device.Serial))
				color.Cyan("\t  Asset Tag: " + color.YellowString("%s", vdc.Device.AssetTag))
				color.Cyan("\t  Site: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Site.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Site.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Site.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Site.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.Site.Slug))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Site.Description))
				color.Cyan("\t  Location: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Location.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Location.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Location.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Location.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", vdc.Device.Location.Slug))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Location.Description))
				color.Cyan("\t    Rack Count: " + color.YellowString("%d", vdc.Device.Location.RackCount))
				color.Cyan("\t    Depth: " + color.YellowString("%d", vdc.Device.Location.Depth))
				color.Cyan("\t  Rack: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Rack.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Rack.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Rack.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Rack.Name))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Rack.Description))
				color.Cyan("\t    Device Count: " + color.YellowString("%d", vdc.Device.Rack.DeviceCount))
				color.Cyan("\t  Position: " + color.YellowString("%d", vdc.Device.Position))
				color.Cyan("\t  Face: ")
				color.Cyan("\t    Value: " + color.YellowString("%d", vdc.Device.Face.Value))
				color.Cyan("\t    Label: " + color.YellowString("%s", vdc.Device.Face.Label))
				color.Cyan("\t  Latitude: " + color.YellowString("%d", vdc.Device.Latitude))
				color.Cyan("\t  Longitude: " + color.YellowString("%d", vdc.Device.Longitude))
				color.Cyan("\t  Parent Device: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.ParentDevice.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.ParentDevice.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.ParentDevice.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.ParentDevice.Name))
				color.Cyan("\t  Status: ")
				color.Cyan("\t    Value: " + color.YellowString("%d", vdc.Device.Status.Value))
				color.Cyan("\t    Label: " + color.YellowString("%s", vdc.Device.Status.Label))
				color.Cyan("\t  AirFlow: ")
				color.Cyan("\t    Value: " + color.YellowString("%s", vdc.Device.Airflow.Value))
				color.Cyan("\t    Label: " + color.YellowString("%s", vdc.Device.Airflow.Label))
				color.Cyan("\t  Primary IP: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.PrimaryIp.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.PrimaryIp.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.PrimaryIp.Display))
				color.Cyan("\t    Family: ")
				color.Cyan("\t      Value: " + color.YellowString("%d", vdc.Device.PrimaryIp.Family.Value))
				color.Cyan("\t      Label: " + color.YellowString("%s", vdc.Device.PrimaryIp.Family.Label))
				color.Cyan("\t    Address: " + color.YellowString("%s", vdc.Device.PrimaryIp.Address))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.PrimaryIp.Description))
				color.Cyan("\t  Primary IPv4: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.PrimaryIp4.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.PrimaryIp4.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.PrimaryIp4.Display))
				color.Cyan("\t    Family: ")
				color.Cyan("\t      Value: " + color.YellowString("%d", vdc.Device.PrimaryIp4.Family.Value))
				color.Cyan("\t      Label: " + color.YellowString("%s", vdc.Device.PrimaryIp4.Family.Label))
				color.Cyan("\t    Address: " + color.YellowString("%s", vdc.Device.PrimaryIp4.Address))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.PrimaryIp4.Description))
				color.Cyan("\t  Primary IPv6: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.PrimaryIp6.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.PrimaryIp6.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.PrimaryIp6.Display))
				color.Cyan("\t    Family: ")
				color.Cyan("\t      Value: " + color.YellowString("%d", vdc.Device.PrimaryIp6.Family.Value))
				color.Cyan("\t      Label: " + color.YellowString("%s", vdc.Device.PrimaryIp6.Family.Label))
				color.Cyan("\t    Address: " + color.YellowString("%s", vdc.Device.PrimaryIp6.Address))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.PrimaryIp6.Description))
				color.Cyan("\t  OOP IP: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.OobIp.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.OobIp.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.OobIp.Display))
				color.Cyan("\t    Family: ")
				color.Cyan("\t      Value: " + color.YellowString("%d", vdc.Device.OobIp.Family.Value))
				color.Cyan("\t      Label: " + color.YellowString("%s", vdc.Device.OobIp.Family.Label))
				color.Cyan("\t    Address: " + color.YellowString("%s", vdc.Device.OobIp.Address))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.OobIp.Description))
				color.Cyan("\t  Cluster: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.Cluster.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.Cluster.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.Cluster.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.Cluster.Name))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.Cluster.Description))
				color.Cyan("\t    Virtual Machine Count: " + color.YellowString("%d", vdc.Device.Cluster.VirtualmachineCount))
				color.Cyan("\t  Virtual Chassis: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.VirtualChassis.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.VirtualChassis.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.VirtualChassis.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.VirtualChassis.Name))
				color.Cyan("\t    Master: ")
				color.Cyan("\t      ID: " + color.YellowString("%d", vdc.Device.VirtualChassis.Master.Id))
				color.Cyan("\t      URL: " + color.YellowString("%s", vdc.Device.VirtualChassis.Master.Url))
				color.Cyan("\t      Display: " + color.YellowString("%s", vdc.Device.VirtualChassis.Master.Display))
				color.Cyan("\t      Name: " + color.YellowString("%s", vdc.Device.VirtualChassis.Master.Name))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.VirtualChassis.Description))
				color.Cyan("\t    Member Count: " + color.YellowString("%d", vdc.Device.VirtualChassis.MemberCount))
				color.Cyan("\t  VC Position: " + color.YellowString("%d", vdc.Device.VcPosition))
				color.Cyan("\t  VC Priority: " + color.YellowString("%d", vdc.Device.VcPriority))
				color.Cyan("\t  Description: " + color.YellowString("%s", vdc.Device.Description))
				color.Cyan("\t  Comments: " + color.YellowString("%s", vdc.Device.Comments))
				color.Cyan("\t  Config Template: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", vdc.Device.ConfigTemplate.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", vdc.Device.ConfigTemplate.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", vdc.Device.ConfigTemplate.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", vdc.Device.ConfigTemplate.Name))
				color.Cyan("\t    Description: " + color.YellowString("%s", vdc.Device.ConfigTemplate.Description))
				color.Cyan("\t  Local Context Data: " + color.YellowString("%s", vdc.Device.LocalContextData))
				for _, tag := range vdc.Device.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
						color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
						color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
						color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
						color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
						color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags found for interface: %s", color.YellowString("%s", vdc.Display)))
					}
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", vdc.Device.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", vdc.Device.LastUpdated))
				color.Cyan("\tInterface Count: " + color.YellowString("%d", vdc.Device.InterfaceCount))
			} else {
				color.Red("\tVirtual Device Context Not found for interface: " + color.YellowString("%s", vdc.Display))
			}
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
			color.Cyan("\t    Local Context Data: " + color.YellowString("%s", result.Module.Device.LocalContextData))
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
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimInterfacesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimInterfacesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimInterfacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimInterfacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

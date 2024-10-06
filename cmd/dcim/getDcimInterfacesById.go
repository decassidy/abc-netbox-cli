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

type interfacesByID struct {
	Id      int    `json:"id,omitempty"`
	Url     string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Device  struct {
		CommonFieldsNoSlug
	} `json:"device,omitempty"`
	Vdcs   []int `json:"vdcs,omitempty"`
	Module struct {
		Id        int    `json:"id,omitempty"`
		Url       string `json:"url,omitempty"`
		Display   string `json:"display,omitempty"`
		Device    int    `json:"device,omitempty"`
		ModuleBay struct {
			CommonFieldsNoSlug
		} `json:"module_bay,omitempty"`
	} `json:"module"`
	Name  string `json:"name,omitempty"`
	Label string `json:"label,omitempty"`
	Type  struct {
		ValueLabel
	} `json:"type,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Parent  struct {
		Id      int    `json:"id,omitempty"`
		Url     string `json:"url,omitempty"`
		Display string `json:"display,omitempty"`
		Device  struct {
			CommonFieldsNoSlug
		} `json:"device,omitempty"`
		Name     string `json:"name,omitempty"`
		Cable    int    `json:"cable,omitempty"`
		Occupied bool   `json:"_occupied,omitempty"`
	} `json:"parent,omitempty"`
	Bridge struct {
		Id      int    `json:"id,omitempty"`
		Url     string `json:"url,omitempty"`
		Display string `json:"display,omitempty"`
		Device  struct {
			CommonFieldsNoSlug
		} `json:"device"`
		Name     string `json:"name,omitempty"`
		Cable    int    `json:"cable,omitempty"`
		Occupied bool   `json:"_occupied,omitempty"`
	} `json:"bridge,omitempty"`
	Lag struct {
		Id      int    `json:"id,omitempty"`
		Url     string `json:"url,omitempty"`
		Display string `json:"display,omitempty"`
		Device  struct {
			CommonFieldsNoSlug
		} `json:"device,omitempty"`
		Name     string `json:"name,omitempty"`
		Cable    int    `json:"cable,omitempty"`
		Occupied bool   `json:"_occupied,omitempty"`
	} `json:"lag"`
	Mtu        int    `json:"mtu,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	Speed      int    `json:"speed,omitempty"`
	Duplex     struct {
		ValueLabel
	} `json:"duplex,omitempty"`
	Wwn         string `json:"wwn,omitempty"`
	MgmtOnly    bool   `json:"mgmt_only,omitempty"`
	Description string `json:"description,omitempty"`
	Mode        struct {
		ValueLabel
	} `json:"mode,omitempty"`
	RfRole struct {
		ValueLabel
	} `json:"rf_role,omitempty"`
	RfChannel struct {
		ValueLabel
	} `json:"rf_channel,omitempty"`
	PoeMode struct {
		ValueLabel
	} `json:"poe_mode,omitempty"`
	PoeType struct {
		ValueLabel
	} `json:"poe_type,omitempty"`
	RfChannelFrequency int `json:"rf_channel_frequency,omitempty"`
	RfChannelWidth     int `json:"rf_channel_width,omitempty"`
	TxPower            int `json:"tx_power,omitempty"`
	UntaggedVlan       struct {
		CommonFieldsNoSlug
		Vid int `json:"vid,omitempty"`
	} `json:"untagged_vlan,omitempty"`
	TaggedVlans   []int `json:"tagged_vlans,omitempty"`
	MarkConnected bool  `json:"mark_connected,omitempty"`
	Cable         struct {
		Id      int    `json:"id,omitempty"`
		Url     string `json:"url,omitempty"`
		Display string `json:"display,omitempty"`
		Label   string `json:"label,omitempty"`
	} `json:"cable,omitempty"`
	CableEnd     string `json:"cable_end,omitempty"`
	WirelessLink struct {
		Id      int    `json:"id,omitempty"`
		Url     string `json:"url,omitempty"`
		Display string `json:"display,omitempty"`
		Ssid    string `json:"ssid,omitempty"`
	} `json:"wireless_link,omitempty"`
	LinkPeers     []string `json:"link_peers,omitempty"`
	LinkPeersType string   `json:"link_peers_type,omitempty"`
	WirelessLans  []int    `json:"wireless_lans,omitempty"`
	Vrf           struct {
		CommonFieldsNoSlug
		Rd string `json:"rd,omitempty"`
	} `json:"vrf,omitempty"`
	L2VpnTermination struct {
		Id      int    `json:"id,omitempty"`
		Url     string `json:"url,omitempty"`
		Display string `json:"display,omitempty"`
		L2Vpn   struct {
			CommonFieldsSlug
			Identifier float64 `json:"identifier,omitempty"`
			Type       string  `json:"type,omitempty"`
		} `json:"l2vpn,omitempty"`
	} `json:"l2vpn_termination,omitempty"`
	ConnectedEndpoints          []string `json:"connected_endpoints,omitempty"`
	ConnectedEndpointsType      string   `json:"connected_endpoints_type,omitempty"`
	ConnectedEndpointsReachable bool     `json:"connected_endpoints_reachable,omitempty"`
	Tags                        []struct {
		CommonFieldsSlug
		Color string `json:"color,omitempty"`
	} `json:"tags,omitempty"`
	Created          string `json:"created,omitempty"`
	LastUpdated      string `json:"last_updated,omitempty"`
	CountIpaddresses int    `json:"count_ipaddresses,omitempty"`
	CountFhrpGroups  int    `json:"count_fhrp_groups,omitempty"`
	Occupied         bool   `json:"_occupied,omitempty"`
}

// GetDcimInterfacesByIdCmd represents the getDcimInterfacesById command
var GetDcimInterfacesByIdCmd = &cobra.Command{
	Use:   "getDcimInterfacesById",
	Short: "GET an interface object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an interface object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(interfacesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.interfaces_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    ABC Interface Name: %s\n", color.YellowString(responseObject.Display)+color.CyanString(" Device: ")+color.YellowString(responseObject.Device.Name))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			color.Cyan("\tDevice: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Device.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Device.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Device.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Device.Name))
			for _, vdc := range responseObject.Vdcs {
				if vdc != 0 {
					color.Cyan("\tVirtual Device Context: " + color.YellowString("%d", vdc))
				} else {
					color.Red("\tVirtual Device Context Not found for interface: " + color.YellowString("%s", responseObject.Display))
				}
			}
			if responseObject.Module.Id != 0 {
				color.Cyan("\tModule: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Module.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Module.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Module.Display))
				color.Cyan("\t  Device: " + color.YellowString("%d", responseObject.Module.Device))
			} else {
				color.Cyan("\tModule: " + color.RedString("No module found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Module.ModuleBay.Id != 0 {
				color.Cyan("\t  Module Bay: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Module.ModuleBay.Id))
				color.Cyan("\t    URL: " + color.YellowString("%d", responseObject.Module.ModuleBay.Url))
				color.Cyan("\t    Display: " + color.YellowString("%d", responseObject.Module.ModuleBay.Display))
				color.Cyan("\t    Name: " + color.YellowString("%d", responseObject.Module.ModuleBay.Name))
			} else {
				color.Cyan("\t  Module Bay: " + color.RedString("No module bay found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tName: " + color.RedString("No name found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Label != "" {
				color.Cyan("\tLabel: " + color.YellowString("%s", responseObject.Label))
			} else {
				color.Cyan("\tLabel: " + color.RedString("No label found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Type.Value != "" {
				color.Cyan("\tType: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Type.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Type.Label))
			} else {
				color.Cyan("\tType: " + color.RedString("No type found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			color.Cyan("\tEnabled: " + color.YellowString("%t", responseObject.Enabled))
			if responseObject.Parent.Id != 0 {
				color.Cyan("\tParent: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Parent.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Parent.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Parent.Display))
				color.Cyan("\t  Device: " + color.YellowString("%s", responseObject.Parent.Device))
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Parent.Device.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Parent.Device.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Parent.Device.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.Parent.Device.Name))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Parent.Name))
				color.Cyan("\t  Cable: " + color.YellowString("%s", responseObject.Parent.Cable))
				color.Cyan("\t  Occupied: " + color.YellowString("%t", responseObject.Parent.Occupied))
			} else {
				color.Cyan("\tParent: " + color.RedString("No parent found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Bridge.Id != 0 {
				color.Cyan("\tBridge: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Bridge.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Bridge.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Bridge.Display))
				color.Cyan("\t  Device: " + color.YellowString("%s", responseObject.Bridge.Device))
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Bridge.Device.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Bridge.Device.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Bridge.Device.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.Bridge.Device.Name))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Bridge.Name))
				color.Cyan("\t  Cable: " + color.YellowString("%s", responseObject.Bridge.Cable))
				color.Cyan("\t  Occupied: " + color.YellowString("%t", responseObject.Bridge.Occupied))
			} else {
				color.Cyan("\tBridge: " + color.RedString("No bridge found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Lag.Id != 0 {
				color.Cyan("\tLink Aggregation Groups: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Lag.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Lag.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Lag.Display))
				color.Cyan("\t  Device: " + color.YellowString("%s", responseObject.Lag.Device))
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Lag.Device.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Lag.Device.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Lag.Device.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.Lag.Device.Name))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Lag.Name))
				color.Cyan("\t  Cable: " + color.YellowString("%s", responseObject.Lag.Cable))
				color.Cyan("\t  Occupied: " + color.YellowString("%t", responseObject.Lag.Occupied))
			} else {
				color.Cyan("\tLink Aggregation Groups: " + color.RedString("No LAGs found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Mtu != 0 {
				color.Cyan("\tMTU: " + color.YellowString("%d", responseObject.Mtu))
			} else {
				color.Cyan("\tMTU: " + color.RedString("No MTU found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.MacAddress != "" {
				color.Cyan("\tMac Address: " + color.YellowString("%s", responseObject.MacAddress))
			} else {
				color.Cyan("\tMac Address: " + color.RedString("No mac address found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Speed != 0 {
				color.Cyan("\tSpeed: " + color.YellowString("%d", responseObject.Speed))
			} else {
				color.Cyan("\tSpeed: " + color.RedString("No Speed found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Duplex.Value != "" {
				color.Cyan("\tDuplex: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Duplex.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Duplex.Label))
			} else {
				color.Cyan("\tDuplex: " + color.RedString("No duplex found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Wwn != "" {
				color.Cyan("\tWorld Wide Name: " + color.YellowString("%s", responseObject.Wwn))
			} else {
				color.Cyan("\tWorld Wide Name: " + color.RedString("No world wide name found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			color.Cyan("\tMgmt Only: " + color.YellowString("%t", responseObject.MgmtOnly))

			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Mode.Value != "" {
				color.Cyan("\tMode: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Mode.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Mode.Label))
			} else {
				color.Cyan("\tMode: " + color.RedString("No mode found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.RfRole.Value != "" {
				color.Cyan("\tRadio Frequency (RF) Role: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.RfRole.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.RfRole.Label))
			} else {
				color.Cyan("\tRadio Frequency (RF) Role: " + color.RedString("No RF role found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.RfChannel.Value != "" {
				color.Cyan("\tRadio Frequency (RF) Channel: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.RfChannel.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.RfChannel.Label))
			} else {
				color.Cyan("\tRadio Frequency (RF) Channel: " + color.RedString("No RF channel found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.PoeMode.Value != "" {
				color.Cyan("\tPower Over Ethernet (PoE) Mode: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.PoeMode.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.PoeMode.Label))
			} else {
				color.Cyan("\tPower Over Ethernet (PoE) Mode: " + color.RedString("No PoE mode found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.PoeType.Value != "" {
				color.Cyan("\tPower Over Ethernet (PoE) Type: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.PoeType.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.PoeType.Label))
			} else {
				color.Cyan("\tPower Over Ethernet (PoE) Type: " + color.RedString("No PoE type found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.RfChannelFrequency != 0 {
				color.Cyan("\tRF Channel Frequency: " + color.YellowString("%d", responseObject.RfChannelFrequency))
			} else {
				color.Cyan("\tRF Channel Frequency: " + color.RedString("No RF channel frequency found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.RfChannelWidth != 0 {
				color.Cyan("\tRF Channel Width: " + color.YellowString("%d", responseObject.RfChannelWidth))
			} else {
				color.Cyan("\tRF Channel Width: " + color.RedString("No RF channel width found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.TxPower != 0 {
				color.Cyan("\tTx Power: " + color.YellowString("%d", responseObject.TxPower))
			} else {
				color.Cyan("\tTx Power: " + color.RedString("No Tx power found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.UntaggedVlan.Id != 0 {
				color.Cyan("\tUntagged Vlan: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.UntaggedVlan.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.UntaggedVlan.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.UntaggedVlan.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.UntaggedVlan.Name))
			} else {
				color.Cyan("\tUntagged Vlan: " + color.RedString("No untagged vlan found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			color.Cyan("\tMark Connected: " + color.YellowString("%t", responseObject.MarkConnected))
			if responseObject.Cable.Id != 0 {
				color.Cyan("\tCable: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Cable.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Cable.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Cable.Display))
				color.Cyan("\t    Label: " + color.YellowString("%s", responseObject.Cable.Label))
			} else {
				color.Cyan("\tCable: " + color.RedString("No cable found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.CableEnd != "" {
				color.Cyan("\tCable End: " + color.YellowString("%s", responseObject.CableEnd))
			} else {
				color.Cyan("\tCable End: " + color.RedString("No cable end found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.WirelessLink.Id != 0 {
				color.Cyan("\tWireless Link: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.WirelessLink.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.WirelessLink.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.WirelessLink.Display))
				color.Cyan("\t    SSID: " + color.YellowString("%s", responseObject.WirelessLink.Ssid))
			} else {
				color.Cyan("\tWireless Link: " + color.RedString("No wireless link found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.LinkPeers == nil || len(responseObject.LinkPeers) == 0 {
				color.Cyan("\tLink Peers: " + color.RedString("No link peers found for interface: %s", color.YellowString("%s", responseObject.Display)))
			} else {
				color.Cyan("\tLink Peers: " + color.YellowString("%d", responseObject.LinkPeers))
			}
			if responseObject.LinkPeersType != "" {
				color.Cyan("\tLink Peers Type: " + color.YellowString("%d", responseObject.LinkPeers))
			} else {
				color.Cyan("\tLink Peers Type: " + color.RedString("No link peers type found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.WirelessLans == nil || len(responseObject.WirelessLans) == 0 {
				color.Cyan("\tWireless LANs: " + color.RedString("No wireless lans found for interface: %s", color.YellowString("%s", responseObject.Display)))
			} else {
				color.Cyan("\tWireless LANs: " + color.YellowString("%d", responseObject.WirelessLans))
			}
			if responseObject.Vrf.Id != 0 {
				color.Cyan("\tVirtual Router Forwarding (vrf): ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Vrf.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Vrf.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Vrf.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.Vrf.Name))
				color.Cyan("\t    RD: " + color.YellowString("%s", responseObject.Vrf.Rd))
			} else {
				color.Cyan("\tVirtual Router Forwarding (vrf): " + color.RedString("No vrf found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.L2VpnTermination.Id != 0 {
				color.Cyan("\tL2VPN Termination: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.L2VpnTermination.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.L2VpnTermination.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.L2VpnTermination.Display))
				color.Cyan("\t  L2VPN: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.L2VpnTermination.L2Vpn.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.L2VpnTermination.L2Vpn.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.L2VpnTermination.L2Vpn.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.L2VpnTermination.L2Vpn.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", responseObject.L2VpnTermination.L2Vpn.Slug))
				color.Cyan("\t    Identifier: " + color.YellowString("%f", responseObject.L2VpnTermination.L2Vpn.Identifier))
				color.Cyan("\t    Type: " + color.YellowString("%s", responseObject.L2VpnTermination.L2Vpn.Type))
			} else {
				color.Cyan("\tL2VPN Termination: " + color.RedString("No l2vpn termination found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.ConnectedEndpoints == nil || len(responseObject.ConnectedEndpoints) == 0 {
				color.Cyan("\tConnected Enpoints: " + color.RedString("No connected endpoints found for interface: %s", color.YellowString("%s", responseObject.Display)))
			} else {
				color.Cyan("\tConnected Enpoints: " + color.YellowString("%d", responseObject.LinkPeers))
			}
			if responseObject.ConnectedEndpointsType != "" {
				color.Cyan("\tConnected Endpoints Type: " + color.YellowString("%d", responseObject.ConnectedEndpointsType))
			} else {
				color.Cyan("\tConnected Endpoints Type: " + color.RedString("No connected endpoints type found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			color.Cyan("\tConnected Endpoints Reachable: " + color.YellowString("%t", responseObject.ConnectedEndpointsReachable))
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
					color.Cyan("\tTags: " + color.RedString("No tags found for interface: %s", color.YellowString("%s", responseObject.Display)))
				}
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
			if responseObject.CountIpaddresses != 0 {
				color.Cyan("\tCount IP Addresses: " + color.YellowString("%d", responseObject.CountIpaddresses))
			} else {
				color.Cyan("\tCount IP Addresses: " + color.RedString("No count IP addresses found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.CountFhrpGroups != 0 {
				color.Cyan("\tCount FHRP Groups: " + color.YellowString("%d", responseObject.CountFhrpGroups))
			} else {
				color.Cyan("\tCount FHRP Groups: " + color.RedString("No count fhrp groups found for interface: %s", color.YellowString("%s", responseObject.Display)))
			}
			color.Cyan("\tOccupied: " + color.YellowString("%t\n", responseObject.Occupied))
		} else {
			color.Red("  Doh! No Interface object found on server for ID: "+color.YellowString("%d/n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimInterfacesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimInterfacesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	GetDcimInterfacesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the interface")
	err = GetDcimInterfacesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err.Error())
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimInterfacesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimInterfacesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

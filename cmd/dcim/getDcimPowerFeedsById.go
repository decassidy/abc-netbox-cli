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

type powerFeedsByID struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	Display    string `json:"display"`
	PowerPanel struct {
		CommonFieldsNoSlug
	} `json:"power_panel"`
	Rack struct {
		CommonFieldsNoSlug
	} `json:"rack"`
	Name   string `json:"name"`
	Status struct {
		ValueLabel
	} `json:"status"`
	Type struct {
		ValueLabel
	} `json:"type"`
	Supply struct {
		ValueLabel
	} `json:"supply"`
	Phase struct {
		ValueLabel
	} `json:"phase"`
	Voltage        int  `json:"voltage"`
	Amperage       int  `json:"amperage"`
	MaxUtilization int  `json:"max_utilization"`
	MarkConnected  bool `json:"mark_connected"`
	Cable          struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Label   string `json:"label"`
	} `json:"cable"`
	CableEnd                    string   `json:"cable_end"`
	LinkPeers                   []string `json:"link_peers"`
	LinkPeersType               string   `json:"link_peers_type"`
	ConnectedEndpoints          []string `json:"connected_endpoints"`
	ConnectedEndpointsType      string   `json:"connected_endpoints_type"`
	ConnectedEndpointsReachable bool     `json:"connected_endpoints_reachable"`
	Description                 string   `json:"description"`
	Tenant                      struct {
		CommonFieldsSlug
	} `json:"tenant"`
	Comments string `json:"comments"`
	Tags     []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
	Occupied    bool   `json:"_occupied"`
}

// GetDcimPowerFeedsByIdCmd represents the getDcimPowerFeedsById command
var GetDcimPowerFeedsByIdCmd = &cobra.Command{
	Use:   "getDcimPowerFeedsById",
	Short: "GET an power feed object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an power feed object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(powerFeedsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.power_feeds_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    ABC Power Feed: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.PowerPanel.Id != 0 {
				color.Cyan("\tPower Panel: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.PowerPanel.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.PowerPanel.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.PowerPanel.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.PowerPanel.Name))
			} else {
				color.Cyan("\tPower Panel: " + color.RedString("No power panel entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Rack.Id != 0 {
				color.Cyan("\tRack: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Rack.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Rack.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Rack.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Rack.Name))
			} else {
				color.Cyan("\tRack: " + color.RedString("No rack entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tName: " + color.RedString("No name entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Status.Value != "" {
				color.Cyan("\tStatus: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Status.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Status.Label))
			} else {
				color.Cyan("\tStatus: " + color.RedString("No Status entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Type.Value != "" {
				color.Cyan("\tType: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Type.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Type.Label))
			} else {
				color.Cyan("\tType: " + color.RedString("No type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Supply.Value != "" {
				color.Cyan("\tSupply: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Supply.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Supply.Label))
			} else {
				color.Cyan("\tSupply: " + color.RedString("No supply entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Phase.Value != "" {
				color.Cyan("\tPhase: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Phase.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Phase.Label))
			} else {
				color.Cyan("\tPhase: " + color.RedString("No phase entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Voltage != 0 {
				color.Cyan("\tVoltage: " + color.YellowString("%d", responseObject.Voltage))
			} else {
				color.Cyan("\tVoltage: " + color.RedString("No voltage entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Amperage != 0 {
				color.Cyan("\tAmperage: " + color.YellowString("%d", responseObject.Amperage))
			} else {
				color.Cyan("\tAmperage: " + color.RedString("No amperage entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.MaxUtilization != 0 {
				color.Cyan("\tMax Utilization: " + color.YellowString("%d", responseObject.MaxUtilization))
			} else {
				color.Cyan("\tMax Utilization: " + color.RedString("No max utilization entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tMark Connected: " + color.YellowString("%t", responseObject.MarkConnected))
			if responseObject.Cable.Id != 0 {
				color.Cyan("\tCable: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Cable.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Cable.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Cable.Display))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Cable.Label))
			} else {
				color.Cyan("\tCable: " + color.RedString("No cable entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.CableEnd != "" {
				color.Cyan("\tCable End: " + color.YellowString("%s", responseObject.CableEnd))
			} else {
				color.Cyan("\tCable End: " + color.RedString("No cable end entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			for _, link := range responseObject.LinkPeers {
				if link != "" {
					color.Cyan("\tLink Peer: " + color.YellowString("%s", link))
				} else {
					color.Cyan("\tLink Peer: " + color.RedString("No link peer entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			}
			if responseObject.LinkPeersType != "" {
				color.Cyan("\tLink Peers Type: " + color.YellowString("%s", responseObject.LinkPeersType))
			} else {
				color.Cyan("\tLink Peers Type: " + color.RedString("No link peers type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			for _, endpoint := range responseObject.ConnectedEndpoints {
				if endpoint != "" {
					color.Cyan("\tConnected Endpoint: " + color.YellowString("%s", endpoint))
				} else {
					color.Cyan("\tConnected Endpoint: " + color.RedString("No connected endpoint entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			}
			if responseObject.ConnectedEndpointsType != "" {
				color.Cyan("\tConnected Endpoint Type: " + color.YellowString("%s", responseObject.ConnectedEndpointsType))
			} else {
				color.Cyan("\tConnected Endpoint Type: " + color.RedString("No connected endpoint type entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			color.Cyan("\tConnected Endpoint Type: " + color.YellowString("%t", responseObject.ConnectedEndpointsReachable))

			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%t", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Tenant.Id != 0 {
				color.Cyan("\tTenant: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Tenant.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Tenant.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Tenant.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Tenant.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Tenant.Slug))
			} else {
				color.Cyan("\tTenant: " + color.RedString("No tenant entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Comments != "" {
				color.Cyan("\tComments: " + color.YellowString("%t", responseObject.Comments))
			} else {
				color.Cyan("\tComments: " + color.RedString("No comments entry found for ") + color.YellowString("%s", responseObject.Display))
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
			color.Cyan("\tOccupied: " + color.YellowString("%t\n", responseObject.Occupied))
		} else {
			color.Red("  Doh! No powerfeed object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimPowerFeedsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimPowerFeedsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPowerFeedsByIdCmd", err)
	}

	GetDcimPowerFeedsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the powerfeed object")
	err = GetDcimPowerFeedsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPowerFeedsByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimPowerFeedsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimPowerFeedsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

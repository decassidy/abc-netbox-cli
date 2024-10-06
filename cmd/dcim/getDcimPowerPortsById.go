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

type powerPortsByID struct {
	Id      uint   `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Device  struct {
		CommonFieldsNoSlug
	} `json:"device"`
	Module struct {
		Id        uint   `json:"id"`
		Url       string `json:"url"`
		Display   string `json:"display"`
		Device    uint   `json:"device"`
		ModuleBay struct {
			CommonFieldsNoSlug
		} `json:"module_bay"`
	} `json:"module"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Type  struct {
		ValueLabel
	} `json:"type"`
	MaximumDraw   uint   `json:"maximum_draw"`
	AllocatedDraw uint   `json:"allocated_draw"`
	Description   string `json:"description"`
	MarkConnected bool   `json:"mark_connected"`
	Cable         struct {
		Id      uint   `json:"id"`
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
	Tags                        []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
	Occupied    bool   `json:"_occupied"`
}

// GetDcimPowerPortsByIdCmd represents the getDcimPowerPortsById command
var GetDcimPowerPortsByIdCmd = &cobra.Command{
	Use:   "getDcimPowerPortsById",
	Short: "GET an power port object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an power outlet port object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(powerPortsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.power_ports_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    ABC Power Port: %s\n", color.YellowString(responseObject.Display))
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
				color.Cyan("\tDevice: " + color.RedString("No device entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Module.Id != 0 {
				color.Cyan("\tModule: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Module.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Module.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Module.Display))
				color.Cyan("\t  Device: " + color.YellowString("%d", responseObject.Module.Device))
			} else {
				color.Cyan("\tModule: " + color.RedString("No module entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tName: " + color.RedString("No name entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Label != "" {
				color.Cyan("\tLabel: " + color.YellowString("%s", responseObject.Label))
			} else {
				color.Cyan("\tLabel: " + color.RedString("No label entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Type.Value != "" {
				color.Cyan("\tType: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Type.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Type.Label))
			} else {
				color.Cyan("\tType: " + color.RedString("No type entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.MaximumDraw != 0 {
				color.Cyan("\tMaximum Draw: " + color.YellowString("%d", responseObject.MaximumDraw))
			} else {
				color.Cyan("\tMaximum Draw: " + color.RedString("No maximum draw entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.AllocatedDraw != 0 {
				color.Cyan("\tAllocated Draw: " + color.YellowString("%d", responseObject.AllocatedDraw))
			} else {
				color.Cyan("\tAllocated Draw: " + color.RedString("No maximum draw entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription" + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
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
					color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			}

			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
			color.Cyan("\tOccupied: " + color.YellowString("%t\n", responseObject.Occupied))
		} else {
			color.Red("  Doh! No power port object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimPowerPortsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimPowerPortsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPowerPortsByIdCmd", err)
	}

	GetDcimPowerPortsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "Id of the power port object")
	err = GetDcimPowerPortsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimPowerPortsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimPowerPortsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

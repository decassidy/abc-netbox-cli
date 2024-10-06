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

type powerOutlets struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Device  struct {
			CommonFieldsNoSlug
		} `json:"device"`
		Module struct {
			Id        int    `json:"id"`
			Url       string `json:"url"`
			Display   string `json:"display"`
			Device    int    `json:"device"`
			ModuleBay struct {
				CommonFieldsNoSlug
			} `json:"module_bay"`
		} `json:"module"`
		Name  string `json:"name"`
		Label string `json:"label"`
		Type  struct {
			ValueLabel
		} `json:"type"`
		PowerPort struct {
			Id      int    `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Device  struct {
				CommonFieldsNoSlug
			} `json:"device"`
			Name     string `json:"name"`
			Cable    int    `json:"cable"`
			Occupied bool   `json:"_occupied"`
		} `json:"power_port"`
		FeedLeg struct {
			ValueLabel
		} `json:"feed_leg"`
		Description   string `json:"description"`
		MarkConnected bool   `json:"mark_connected"`
		Cable         struct {
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
		Tags                        []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created     string `json:"created"`
		LastUpdated string `json:"last_updated"`
		Occupied    bool   `json:"_occupied"`
	} `json:"results"`
}

// GetDcimPowerOutletsCmd represents the getDcimPowerOutlets command
var GetDcimPowerOutletsCmd = &cobra.Command{
	Use:   "getDcimPowerOutlets",
	Short: "GET a list of power outlet objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of power outlet objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(powerOutlets)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.power_outlets")

		if responseObject.Count != 0 {
			color.Cyan("\n  ABC Power Outlets: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    ABC Power Outlet: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.Device.Id != 0 {
					color.Cyan("\tDevice: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Device.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Device.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Device.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Device.Name))
				} else {
					color.Cyan("\tDevice: " + color.RedString("No device entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Module.Id != 0 {
					color.Cyan("\tModule: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Module.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Module.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Module.Display))
					color.Cyan("\t  Device: " + color.YellowString("%d", result.Module.Device))
				} else {
					color.Cyan("\tModule: " + color.RedString("No module entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Module.ModuleBay.Id != 0 {
					color.Cyan("\t  Module Bay: " + color.YellowString("%d", result.Module.ModuleBay))
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Module.ModuleBay.Id))
					color.Cyan("\t    URL: " + color.YellowString("%d", result.Module.ModuleBay.Url))
					color.Cyan("\t    Display: " + color.YellowString("%d", result.Module.ModuleBay.Display))
					color.Cyan("\t    Name: " + color.YellowString("%d", result.Module.ModuleBay.Name))
				} else {
					color.Cyan("\t  Module Bay" + color.RedString("No module bay entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Name != "" {
					color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				} else {
					color.Cyan("\tName: " + color.RedString("No name entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Label != "" {
					color.Cyan("\tLabel: " + color.YellowString("%s", result.Label))
				} else {
					color.Cyan("\tLabel: " + color.RedString("No label entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Type.Value != "" {
					color.Cyan("\tType: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Type.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Type.Label))
				} else {
					color.Cyan("\tType: " + color.RedString("No type entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.PowerPort.Id != 0 {
					color.Cyan("\tPower Port: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.PowerPort.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.PowerPort.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.PowerPort.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.PowerPort.Name))
				} else {
					color.Cyan("\t  Cable: " + color.RedString("No cable entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.PowerPort.Cable != 0 {
					color.Cyan("\t  Cable: " + color.YellowString("%d", result.PowerPort.Cable))
				} else {
					color.Cyan("\t  Cable: " + color.RedString("No cable entry found for ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\t  Occupied: " + color.YellowString("%t", result.PowerPort.Occupied))
				if result.FeedLeg.Value != "" {
					color.Cyan("\tFeed Leg: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.FeedLeg.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.FeedLeg.Label))
				} else {
					color.Cyan("\tFeed Leg: " + color.RedString("No feed leg entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No Description entry found for ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tMark Connected: " + color.YellowString("%t", result.MarkConnected))
				if result.Cable.Id != 0 {
					color.Cyan("\tCable: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Cable.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Cable.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Cable.Display))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Cable.Label))
				} else {
					color.Cyan("\tCable: " + color.RedString("No cable entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.CableEnd != "" {
					color.Cyan("\tCable End: " + color.YellowString("%s", result.CableEnd))
				} else {
					color.Cyan("\tCable End: " + color.RedString("No cable end entry found for ") + color.YellowString("%s", result.Display))
				}
				for _, link := range result.LinkPeers {
					if link != "" {
						color.Cyan("\tLink Peer: " + color.YellowString("%s", link))
					} else {
						color.Cyan("\tLink Peer: " + color.RedString("No link peer entry found for ") + color.YellowString("%s", result.Display))
					}
				}
				if result.LinkPeersType != "" {
					color.Cyan("\tLink Peers Type: " + color.YellowString("%s", result.LinkPeersType))
				} else {
					color.Cyan("\tLink Peers Type: " + color.RedString("No link peers type entry found for ") + color.YellowString("%s", result.Display))
				}
				for _, endpoint := range result.ConnectedEndpoints {
					if endpoint != "" {
						color.Cyan("\tConnected Endpoint: " + color.YellowString("%s", endpoint))
					} else {
						color.Cyan("\tConnected Endpoint: " + color.RedString("No connected endpoint entry found for ") + color.YellowString("%s", result.Display))
					}
				}
				if result.ConnectedEndpointsType != "" {
					color.Cyan("\tConnected Endpoint Type: " + color.YellowString("%s", result.ConnectedEndpointsType))
				} else {
					color.Cyan("\tConnected Endpoint Type: " + color.RedString("No connected endpoint type entry found for ") + color.YellowString("%s", result.Display))
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
						color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s", result.Display))
					}
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
				color.Cyan("\tOccupied: " + color.YellowString("%t\n", result.Occupied))
			}
		} else {
			color.Cyan("  ABC Power Outlets: " + color.RedString("No power outlets found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimPowerOutletsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimPowerOutletsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPowerOutletsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimPowerOutletsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimPowerOutletsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

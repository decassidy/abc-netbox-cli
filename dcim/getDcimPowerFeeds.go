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

type powerFeeds struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
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
	} `json:"results"`
}

// GetDcimPowerFeedsCmd represents the getDcimPowerFeeds command
var GetDcimPowerFeedsCmd = &cobra.Command{
	Use:   "getDcimPowerFeeds",
	Short: "GET a list of powerfeed objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of powerfeed objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(powerFeeds)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.power_feeds")

		if responseObject.Count != 0 {
			color.Cyan("\n  ABC Power Feeds: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    ABC Power Feed: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.PowerPanel.Id != 0 {
					color.Cyan("\tPower Panel: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.PowerPanel.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.PowerPanel.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.PowerPanel.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.PowerPanel.Name))
				} else {
					color.Cyan("\tPower Panel: " + color.RedString("No power panel entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Rack.Id != 0 {
					color.Cyan("\tRack: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Rack.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Rack.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Rack.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Rack.Name))
				} else {
					color.Cyan("\tRack: " + color.RedString("No rack entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Name != "" {
					color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				} else {
					color.Cyan("\tName: " + color.RedString("No name entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Status.Value != "" {
					color.Cyan("\tStatus: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Status.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Status.Label))
				} else {
					color.Cyan("\tStatus: " + color.RedString("No Status entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Type.Value != "" {
					color.Cyan("\tType: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Type.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Type.Label))
				} else {
					color.Cyan("\tType: " + color.RedString("No type entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Supply.Value != "" {
					color.Cyan("\tSupply: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Supply.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Supply.Label))
				} else {
					color.Cyan("\tSupply: " + color.RedString("No supply entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Phase.Value != "" {
					color.Cyan("\tPhase: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Phase.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Phase.Label))
				} else {
					color.Cyan("\tPhase: " + color.RedString("No phase entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Voltage != 0 {
					color.Cyan("\tVoltage: " + color.YellowString("%d", result.Voltage))
				} else {
					color.Cyan("\tVoltage: " + color.RedString("No voltage entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Amperage != 0 {
					color.Cyan("\tAmperage: " + color.YellowString("%d", result.Amperage))
				} else {
					color.Cyan("\tAmperage: " + color.RedString("No amperage entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.MaxUtilization != 0 {
					color.Cyan("\tMax Utilization: " + color.YellowString("%d", result.MaxUtilization))
				} else {
					color.Cyan("\tMax Utilization: " + color.RedString("No max utilization entry found for ") + color.YellowString("%s", result.Display))
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

				color.Cyan("\tConnected Endpoint Type: " + color.YellowString("%t", result.ConnectedEndpointsReachable))

				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%t", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Tenant.Id != 0 {
					color.Cyan("\tTenant: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Tenant.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Tenant.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Tenant.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Tenant.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Tenant.Slug))
				} else {
					color.Cyan("\tTenant: " + color.RedString("No tenant entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Comments != "" {
					color.Cyan("\tComments: " + color.YellowString("%t", result.Comments))
				} else {
					color.Cyan("\tComments: " + color.RedString("No comments entry found for ") + color.YellowString("%s", result.Display))
				}
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
			color.Cyan("  ABC Power Feeds: " + color.RedString("No power feeds found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimPowerFeedsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimPowerFeedsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPowerFeedsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimPowerFeedsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimPowerFeedsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

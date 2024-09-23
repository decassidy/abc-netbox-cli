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
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type consoleServerPortsByID struct {
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
	Speed struct {
		Value int    `json:"value"`
		Label string `json:"label"`
	} `json:"speed"`
	Description   string `json:"description"`
	MarkConnected bool   `json:"mark_connected"`
	Cable         struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Label   string `json:"label"`
	} `json:"cable"`
	CableEnd                    string           `json:"cable_end"`
	LinkPeers                   []string         `json:"link_peers"`
	LinkPeersType               string           `json:"link_peers_type"`
	ConnectedEndpoints          []string         `json:"connected_endpoints"`
	ConnectedEndpointsType      string           `json:"connected_endpoints_type"`
	ConnectedEndpointsReachable bool             `json:"connected_endpoints_reachable"`
	Tags                        CommonFieldsSlug `json:"tags"`
	Created                     string           `json:"created"`
	LastUpdated                 string           `json:"last_updated"`
	Occupied                    bool             `json:"_occupied"`
}

// GetDcimConsoleServerPortsByIdCmd represents the getConsoleServerPortsById command
var GetDcimConsoleServerPortsByIdCmd = &cobra.Command{
	Use:   "getDcimConsoleServerPortsById",
	Short: "GET an console server port object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an console server port object by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(consoleServerPortsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.console_server_ports_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    ABC Console Server Port: %s\n", color.YellowString(responseObject.Display))
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
				color.Cyan("\tDevice: " + color.RedString("No device entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Module.Id != 0 {
				color.Cyan("\tModule: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Module.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Module.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Module.Display))
				color.Cyan("\t  Module Bay: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Module.ModuleBay.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Module.ModuleBay.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Module.ModuleBay.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.Module.ModuleBay.Name))
			} else {
				color.Cyan("\tModule: " + color.RedString("No module entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString(responseObject.Name))
			} else {
				color.Cyan("\tName: " + color.RedString("No label entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Label != "" {
				color.Cyan("\tLabel: " + color.YellowString(responseObject.Label))
			} else {
				color.Cyan("\tLabel: " + color.RedString("No label entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Type.Value != "" {
				color.Cyan("\tType: " + color.YellowString(responseObject.Label))
				color.Cyan("\t  Value: " + color.YellowString(responseObject.Type.Value))
				color.Cyan("\t  Label: " + color.YellowString(responseObject.Type.Label))
			} else {
				color.Cyan("\tType: " + color.RedString("No type entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Speed.Value != 0 {
				color.Cyan("\tSpeed: " + color.YellowString(responseObject.Label))
				color.Cyan("\t  Value: " + color.YellowString(strconv.Itoa(responseObject.Speed.Value)))
				color.Cyan("\t  Label: " + color.YellowString(responseObject.Speed.Label))
			} else {
				color.Cyan("\tSpeed: " + color.RedString("No speed entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString(responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tMarked Connected: " + color.YellowString("%v", responseObject.MarkConnected))
			if responseObject.Cable.Id != 0 {
				color.Cyan("\tCable: ")
				color.Cyan("\t  ID: " + color.RedString("%d", responseObject.Cable.Id))
				color.Cyan("\t  URL: " + color.RedString("%s", responseObject.Cable.Url))
				color.Cyan("\t  Display: " + color.RedString("%s", responseObject.Cable.Display))
				color.Cyan("\t  Label: " + color.RedString("%s", responseObject.Cable.Label))
			} else {
				color.Cyan("\tCable: " + color.RedString("No cable entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.CableEnd != "" {
				color.Cyan("\tCable End: " + color.YellowString(responseObject.CableEnd))
			} else {
				color.Cyan("\tCable End: " + color.RedString("No cable end entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			for _, link := range responseObject.LinkPeers {
				if link != "" {
					color.Cyan("\tLink Peer: " + color.YellowString(link))
				} else {
					color.Cyan("\tLink Peer: " + color.RedString("No link peers entry for console server port: ") + color.YellowString("%s", responseObject.Display))
				}
			}
			if responseObject.LinkPeersType != "" {
				color.Cyan("\tLink Peers Type: " + color.YellowString(responseObject.LinkPeersType))
			} else {
				color.Cyan("\tLink Peers Type: " + color.RedString("No link peers type entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			for _, endpoint := range responseObject.ConnectedEndpoints {
				if endpoint != "" {
					color.Cyan("\tConnected Endpoints: " + color.YellowString(endpoint))
				} else {
					color.Cyan("\tConnected Endpoints: " + color.RedString("No connected endpoints entry for console server port: ") + color.YellowString("%s", responseObject.Display))
				}
			}
			if responseObject.ConnectedEndpointsType != "" {
				color.Cyan("\tConnected Endpoints Type: " + color.YellowString(responseObject.ConnectedEndpointsType))
			} else {
				color.Cyan("\tConnected Endpoints Type: " + color.RedString("No connected endpoints type entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tConnected Endpoints Reachable: " + color.YellowString("%v", responseObject.ConnectedEndpointsReachable))
			if responseObject.Tags.Id != 0 {
				color.Cyan("\tTags: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Tags.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Tags.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Tags.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Tags.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Tags.Slug))
			} else {
				color.Cyan("\tTags: " + color.RedString("No tags entry for console server port: ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
			color.Cyan("\tOccupied: " + color.YellowString("%t\n", responseObject.Occupied))
		} else {
			color.Cyan("  ABC Console Server Port: " + color.RedString("No console server port entries found on server for ID: %d. Exiting...\n", id))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimConsoleServerPortsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimConsoleServerPortsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimConsoleServerPortsByIdCmd", err)
	}

	GetDcimConsoleServerPortsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the console server port")
	err = GetDcimConsoleServerPortsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getConsoleServerPortsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getConsoleServerPortsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

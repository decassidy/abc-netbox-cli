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

type frontPortsByID struct {
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
	Color    string `json:"color"`
	RearPort struct {
		CommonFieldsNoSlug
		Label       string `json:"label"`
		Description string `json:"description"`
	} `json:"rear_port"`
	RearPortPosition int    `json:"rear_port_position"`
	Description      string `json:"description"`
	MarkConnected    bool   `json:"mark_connected"`
	Cable            struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Label   string `json:"label"`
	} `json:"cable"`
	CableEnd      string   `json:"cable_end"`
	LinkPeers     []string `json:"link_peers"`
	LinkPeersType string   `json:"link_peers_type"`
	Tags          []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
	Occupied    bool   `json:"_occupied"`
}

// GetDcimFrontPortsByIdCmd represents the getDcimFrontPortsById command
var GetDcimFrontPortsByIdCmd = &cobra.Command{
	Use:   "getDcimFrontPortsById",
	Short: "GET an front port object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an front port object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(frontPortsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.front_ports_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    Metropolis Front Port Name: %s\n", color.YellowString(responseObject.Display))
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
			color.Cyan("\tModule: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Module.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Module.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Module.Display))
			color.Cyan("\t  Device: " + color.YellowString("%d", responseObject.Module.Device))
			color.Cyan("\t  Module Bay: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Module.ModuleBay.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Module.ModuleBay.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Module.ModuleBay.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.Module.ModuleBay.Name))
			color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			color.Cyan("\tLabel: " + color.YellowString("%s", responseObject.Label))
			color.Cyan("\tType: ")
			color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Type.Value))
			color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Type.Label))
			color.Cyan("\tColor: " + color.YellowString("%s", responseObject.Color))
			color.Cyan("\tRear Port: " + color.YellowString("%s", responseObject.RearPort))
			color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.RearPort.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.RearPort.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.RearPort.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.RearPort.Name))
			color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.RearPort.Label))
			color.Cyan("\t  Description: " + color.YellowString("%s", responseObject.RearPort.Description))
			color.Cyan("\tRear Port Position: " + color.YellowString("%d", responseObject.RearPortPosition))
			color.Cyan("\tDescription: " + color.YellowString("%d", responseObject.Description))
			color.Cyan("\tMarked Connected: " + color.YellowString("%v", responseObject.MarkConnected))
			color.Cyan("\tCable: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Cable.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Cable.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Cable.Display))
			color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Cable.Label))
			color.Cyan("\tCable End: " + color.YellowString("%s", responseObject.CableEnd))
			for _, link := range responseObject.LinkPeers {
				color.Cyan("\tLink Peer: " + color.YellowString("%d", link))
			}
			color.Cyan("\tLink Peers Type: " + color.YellowString("%s", responseObject.LinkPeersType))
			for _, tag := range responseObject.Tags {
				color.Cyan("\tTags: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
				color.Cyan("\t  URL: " + color.YellowString("%d", tag.Url))
				color.Cyan("\t  Display: " + color.YellowString("%d", tag.Display))
				color.Cyan("\t  Name: " + color.YellowString("%d", tag.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%d", tag.Slug))
				color.Cyan("\t  Color: " + color.YellowString("%d", tag.Color))
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
			color.Cyan("\tOccupied: " + color.YellowString("%s\n", responseObject.Occupied))
		} else {
			color.Red("  Doh! Front Port object not found with ID: " + color.YellowString("%s\n", id))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimFrontPortsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimFrontPortsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	GetDcimFrontPortsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of front port object")
	err = GetDcimFrontPortsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimFrontPortsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimFrontPortsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

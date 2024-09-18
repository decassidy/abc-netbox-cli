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

type rearPortsByID struct {
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
	Color         string `json:"color"`
	Positions     uint   `json:"positions"`
	Description   string `json:"description"`
	MarkConnected bool   `json:"mark_connected"`
	Cable         struct {
		Id      uint   `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Label   string `json:"label"`
	} `json:"cable"`
	CableEnd  string `json:"cable_end"`
	LinkPeers []struct {
		Id      uint   `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Device  struct {
			Id      uint   `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Name    string `json:"name"`
		}
		Name     string `json:"name"`
		Cable    uint   `json:"cable"`
		Occupied bool   `json:"_occupied"`
	} `json:"link_peers"`
	LinkPeersType string `json:"link_peers_type"`
	Tags          []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
	Occupied    bool   `json:"_occupied"`
}

// GetDcimRearPortsByIdCmd represents the getDcimRearPortsById command
var GetDcimRearPortsByIdCmd = &cobra.Command{
	Use:   "getDcimRearPortsById",
	Short: "GET an rear port object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an rear port object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(rearPortsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.rear_ports_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    Metropolis Rear Port: %s\n", color.YellowString(responseObject.Display))
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

			if responseObject.Module.ModuleBay.Id != 0 {
				color.Cyan("\t  Module Bay: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Module.ModuleBay.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Module.ModuleBay.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Module.ModuleBay.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.Module.ModuleBay.Name))
			} else {
				color.Cyan("\t  Module Bay: " + color.RedString("No module bay entry found for ") + color.YellowString("%s", responseObject.Display))
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

			if responseObject.Color != "" {
				color.Cyan("\tColor: " + color.YellowString("%s", responseObject.Color))
			} else {
				color.Cyan("\tColor: " + color.RedString("No color entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Positions != 0 {
				color.Cyan("\tPositions: " + color.YellowString("%d", responseObject.Positions))
			} else {
				color.Cyan("\tPositions: " + color.RedString("No positions entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
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
				if link.Id != 0 {
					color.Cyan("\tLink Peer: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", link.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", link.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", link.Display))
					color.Cyan("\t  Device: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", link.Device.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", link.Device.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", link.Device.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", link.Device.Name))
					color.Cyan("\t  Name: " + color.YellowString("%s", link.Name))
					color.Cyan("\t  Cable: " + color.YellowString("%d", link.Cable))
					color.Cyan("\t  Occupied: " + color.YellowString("%t", link.Occupied))
				} else {
					color.Cyan("\tLink Peer: " + color.RedString("No link peer entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			}

			if responseObject.LinkPeersType != "" {
				color.Cyan("\tLink Peers Type: " + color.YellowString("%s", responseObject.LinkPeersType))
			} else {
				color.Cyan("\tLink Peers Type: " + color.RedString("No link peers type entry found for ") + color.YellowString("%s", responseObject.Display))
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
			color.Red("  Doh! No rear port object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimRearPortsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimRearPortsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimRearPortsByIdCmd", err)
	}

	GetDcimRearPortsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the rear port object")
	err = GetDcimRearPortsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimRearPortsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimRearPortsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

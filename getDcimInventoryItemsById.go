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

type inventoryItemsByID struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Device  struct {
		CommonFieldsNoSlug
	} `json:"device"`
	Parent int    `json:"parent"`
	Name   string `json:"name"`
	Label  string `json:"label"`
	Role   struct {
		CommonFieldsSlug
	} `json:"role"`
	Manufacturer struct {
		CommonFieldsSlug
	} `json:"manufacturer"`
	PartId        string  `json:"part_id"`
	Serial        string  `json:"serial"`
	AssetTag      string  `json:"asset_tag"`
	Discovered    bool    `json:"discovered"`
	Description   string  `json:"description"`
	ComponentType string  `json:"component_type"`
	ComponentId   float64 `json:"component_id"`
	Component     string  `json:"component"`
	Tags          []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
	Depth       int    `json:"_depth"`
}

// GetDcimInventoryItemsByIdCmd represents the getDcimInventoryItemsById command
var GetDcimInventoryItemsByIdCmd = &cobra.Command{
	Use:   "getDcimInventoryItemsById",
	Short: "GET an inventory item object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an inventory item object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(inventoryItemsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.inventory_items_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    Metropolis Inventory Items: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))

			if responseObject.Device.Id != 0 {
				color.Cyan("\tDevice: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tDevice: " + color.RedString("No device entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Parent != 0 {
				color.Cyan("\tParent: " + color.YellowString("%d", responseObject.Parent))
			} else {
				color.Cyan("\tParent: " + color.RedString("No parent entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tName: " + color.RedString("No name entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Label != "" {
				color.Cyan("\tLabel: " + color.YellowString("%s", responseObject.Label))
			} else {
				color.Cyan("\tLabel: " + color.RedString("No Label entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Role.Id != 0 {
				color.Cyan("\tRole: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Role.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Role.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Role.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Role.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Role.Slug))
			} else {
				color.Cyan("\tRole: " + color.RedString("No role entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Manufacturer.Id != 0 {
				color.Cyan("\tManufacturer: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Manufacturer.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Manufacturer.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Manufacturer.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Manufacturer.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Manufacturer.Slug))
			} else {
				color.Cyan("\tManufacturer: " + color.RedString("No manufacturer entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.PartId != "" {
				color.Cyan("\tPart ID: " + color.YellowString("%s", responseObject.PartId))
			} else {
				color.Cyan("\tPart ID: " + color.RedString("No part ID entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Serial != "" {
				color.Cyan("\tSerial: " + color.YellowString("%s", responseObject.Serial))
			} else {
				color.Cyan("\tSerial: " + color.RedString("No serial entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.AssetTag != "" {
				color.Cyan("\tAsset Tag: " + color.YellowString("%s", responseObject.AssetTag))
			} else {
				color.Cyan("\tAsset Tag: " + color.RedString("No asset tag entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			color.Cyan("\tDiscovered: " + color.YellowString("%t", responseObject.Discovered))

			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No asset tag entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.ComponentType != "" {
				color.Cyan("\tComponent Type: " + color.YellowString("%s", responseObject.ComponentType))
			} else {
				color.Cyan("\tComponent Type: " + color.RedString("No component type entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.ComponentId != 0 {
				color.Cyan("\tComponent ID: " + color.YellowString("%d", responseObject.ComponentId))
			} else {
				color.Cyan("\tComponent ID: " + color.RedString("No component type entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Component != "" {
				color.Cyan("\tComponent: " + color.YellowString("%s", responseObject.Component))
			} else {
				color.Cyan("\tComponent: " + color.RedString("No component entry found for ") + color.YellowString("%s", responseObject.Display))
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
			color.Cyan("\tDepth: " + color.YellowString("%d\n", responseObject.Depth))
		} else {
			color.Red("  Doh! No Inventory Item object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimInventoryItemsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimInventoryItemsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	GetDcimInventoryItemsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "Inventory Item object by ID")
	err = GetDcimInventoryItemsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimInventoryItemsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimInventoryItemsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

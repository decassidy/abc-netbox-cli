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

type inventoryItemTemplateByID struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	Display    string `json:"display"`
	DeviceType struct {
		Id           int    `json:"id"`
		Url          string `json:"url"`
		Display      string `json:"display"`
		Manufacturer struct {
			CommonFieldsSlug
		} `json:"manufacturer"`
		Model string `json:"model"`
		Slug  string `json:"slug"`
	} `json:"device_type"`
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
	Description   string  `json:"description"`
	ComponentType string  `json:"component_type"`
	ComponentId   float32 `json:"component_id"`
	Component     string  `json:"component"`
	Created       string  `json:"created"`
	LastUpdated   string  `json:"last_updated"`
	Depth         int     `json:"_depth"`
}

// GetDcimInventoryItemTemplatesByIdCmd represents the getDcimInventoryItemTemplatesById command
var GetDcimInventoryItemTemplatesByIdCmd = &cobra.Command{
	Use:   "getDcimInventoryItemTemplatesById",
	Short: "GET an inventory item template object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an inventory item template object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(inventoryItemTemplateByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.inventory_item_templates_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    ABC Inventory Item Templates: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.DeviceType.Id != 0 {
				color.Cyan("\tDevice Type: " + color.YellowString("%s", responseObject.DeviceType))
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.DeviceType.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.DeviceType.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.DeviceType.Display))
				color.Cyan("\t  Manufacturer: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer))
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Manufacturer.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Manufacturer.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Manufacturer.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.Manufacturer.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", responseObject.Manufacturer.Slug))
				color.Cyan("\t  Model: " + color.YellowString("%s", responseObject.DeviceType.Model))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.DeviceType.Slug))
			} else {
				color.Cyan("\tDevice Type: " + color.RedString("No device type entry found for ") + color.YellowString("%s", responseObject.Display))
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
				color.Cyan("\tLabel: " + color.YellowString("%d", responseObject.Label))
			} else {
				color.Cyan("\tLabel: " + color.RedString("No label entry found for ") + color.YellowString("%s", responseObject.Display))
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
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.ComponentType != "" {
				color.Cyan("\tComponent Type: " + color.YellowString("%s", responseObject.ComponentType))
			} else {
				color.Cyan("\tComponent Type: " + color.RedString("No component type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.ComponentId != 0 {
				color.Cyan("\tComponent ID: " + color.YellowString("%f", responseObject.ComponentId))
			} else {
				color.Cyan("\tComponent ID: " + color.RedString("No component ID entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Component != "" {
				color.Cyan("\tComponent: " + color.YellowString("%s", responseObject.Component))
			} else {
				color.Cyan("\tComponent: " + color.RedString("No component entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
			if responseObject.Depth != 0 {
				color.Cyan("\tDepth: " + color.YellowString("%d", responseObject.Depth))
			} else {
				color.Cyan("\tDepth: " + color.RedString("No depth entry found for ") + color.YellowString("%s\n", responseObject.Display))
			}

		} else {
			color.Red("  Doh! No Inventory Item template object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimInventoryItemTemplatesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimInventoryItemTemplatesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	GetDcimInventoryItemTemplatesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the inventory item template object")
	err = GetDcimInventoryItemTemplatesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimInventoryItemTemplatesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimInventoryItemTemplatesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

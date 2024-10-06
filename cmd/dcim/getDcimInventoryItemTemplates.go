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

type inventoryItemTemplates struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
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
	} `json:"results"`
}

// GetDcimInventoryItemTemplatesCmd represents the getDcimInventoryItemTemplates command
var GetDcimInventoryItemTemplatesCmd = &cobra.Command{
	Use:   "getDcimInventoryItemTemplates",
	Short: "GET a list of inventory item template objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of inventory item template objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(inventoryItemTemplates)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.inventory_item_templates")

		if responseObject.Count != 0 {
			color.Cyan("\n  ABC Inventory Item Templates: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    ABC Inventory Item Templates: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.DeviceType.Id != 0 {
					color.Cyan("\tDevice Type: " + color.YellowString("%s", result.DeviceType))
					color.Cyan("\t  ID: " + color.YellowString("%d", result.DeviceType.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.DeviceType.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.DeviceType.Display))
					color.Cyan("\t  Manufacturer: " + color.YellowString("%s", result.DeviceType.Manufacturer))
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Manufacturer.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Manufacturer.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Manufacturer.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Manufacturer.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.Manufacturer.Slug))
					color.Cyan("\t  Model: " + color.YellowString("%s", result.DeviceType.Model))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.DeviceType.Slug))
				} else {
					color.Cyan("\tDevice Type: " + color.RedString("No device type entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Parent != 0 {
					color.Cyan("\tParent: " + color.YellowString("%d", result.Parent))
				} else {
					color.Cyan("\tParent: " + color.RedString("No parent entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Name != "" {
					color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				} else {
					color.Cyan("\tName: " + color.RedString("No name entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Label != "" {
					color.Cyan("\tLabel: " + color.YellowString("%d", result.Label))
				} else {
					color.Cyan("\tLabel: " + color.RedString("No label entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Role.Id != 0 {
					color.Cyan("\tRole: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Role.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Role.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Role.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Role.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Role.Slug))
				} else {
					color.Cyan("\tRole: " + color.RedString("No role entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Manufacturer.Id != 0 {
					color.Cyan("\tManufacturer: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Manufacturer.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Manufacturer.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Manufacturer.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Manufacturer.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Manufacturer.Slug))
				} else {
					color.Cyan("\tManufacturer: " + color.RedString("No manufacturer entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.PartId != "" {
					color.Cyan("\tPart ID: " + color.YellowString("%s", result.PartId))
				} else {
					color.Cyan("\tPart ID: " + color.RedString("No part ID entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.ComponentType != "" {
					color.Cyan("\tComponent Type: " + color.YellowString("%s", result.ComponentType))
				} else {
					color.Cyan("\tComponent Type: " + color.RedString("No component type entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.ComponentId != 0 {
					color.Cyan("\tComponent ID: " + color.YellowString("%f", result.ComponentId))
				} else {
					color.Cyan("\tComponent ID: " + color.RedString("No component ID entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Component != "" {
					color.Cyan("\tComponent: " + color.YellowString("%s", result.Component))
				} else {
					color.Cyan("\tComponent: " + color.RedString("No component entry found for ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
				if result.Depth != 0 {
					color.Cyan("\tDepth: " + color.YellowString("%d\n", result.Depth))
				} else {
					color.Cyan("\tDepth: " + color.RedString("No depth entry found for ") + color.YellowString("%s\n", result.Display))
				}
			}
		} else {
			color.Cyan("  ABC Inventory Item Templates: " + color.RedString("No inventory item templates found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimInventoryItemTemplatesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimInventoryItemTemplatesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimInventoryItemTemplatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimInventoryItemTemplatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

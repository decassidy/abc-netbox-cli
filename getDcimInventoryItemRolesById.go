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

type inventoryItemRolesByID struct {
	CommonFieldsSlug
	Color       string `json:"color"`
	Description string `json:"description"`
	Tags        []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created            string `json:"created"`
	LastUpdated        string `json:"last_updated"`
	InventoryitemCount int    `json:"inventoryitem_count"`
}

// GetDcimInventoryItemRolesByIdCmd represents the getDcimInventoryItemRolesById command
var GetDcimInventoryItemRolesByIdCmd = &cobra.Command{
	Use:   "getDcimInventoryItemRolesById",
	Short: "GET an inventory item role object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an inventory item role object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(inventoryItemRolesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.inventory_item_roles_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    Metropolis Inventory Item Role Name: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			color.Cyan("\tSlug: " + color.YellowString("%s", responseObject.Slug))
			color.Cyan("\tColor: " + color.YellowString("%s", responseObject.Color))
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s\n", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s\n", responseObject.Display))
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
				color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))
				color.Cyan("\tInventory Item Count: " + color.YellowString("%d\n", responseObject.InventoryitemCount))
			}
		} else {
			color.Red("  Doh! No Inventory Item Role object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimInventoryItemRolesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimInventoryItemRolesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	GetDcimInventoryItemRolesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of inventory item role object")
	err = GetDcimInventoryItemRolesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimInventoryItemRolesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimInventoryItemRolesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

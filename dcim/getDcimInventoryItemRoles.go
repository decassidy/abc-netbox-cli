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

type inventoryItemRoles struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
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
	} `json:"results"`
}

// GetDcimInventoryItemRolesCmd represents the getDcimInventoryItemRoles command
var GetDcimInventoryItemRolesCmd = &cobra.Command{
	Use:   "getDcimInventoryItemRoles",
	Short: "GET a list of inventory item role objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of inventory item role objects`,
	Run: func(cmd *cobra.Command, args []string) {

		responseObject := new(inventoryItemRoles)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.inventory_item_roles")

		if responseObject.Count != 0 {
			color.Cyan("\n  ABC Inventory Item Roles: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    ABC Inventory Item Role Name: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				color.Cyan("\tSlug: " + color.YellowString("%s", result.Slug))
				color.Cyan("\tColor: " + color.YellowString("%s", result.Color))
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s\n", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s\n", result.Display))
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
					color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
					color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
					color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
					color.Cyan("\tInventory Item Count: " + color.YellowString("%d\n", result.InventoryitemCount))
				}
			}
		} else {
			color.Cyan("  ABC Inventory Item Roles: " + color.RedString("No interface template found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimInventoryItemRolesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimInventoryItemRolesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking dcim inventory item roles flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimInventoryItemRolesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimInventoryItemRolesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

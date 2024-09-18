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

type moduleTypesByID struct {
	Id           int    `json:"id"`
	Url          string `json:"url"`
	Display      string `json:"display"`
	Manufacturer struct {
		CommonFieldsSlug
	} `json:"manufacturer"`
	Model      string  `json:"model"`
	PartNumber string  `json:"part_number"`
	Weight     float32 `json:"weight"`
	WeightUnit struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"weight_unit"`
	Description string `json:"description"`
	Comments    string `json:"comments"`
	Tags        []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

// GetDcimModuleTypesByIdCmd represents the getDcimModuleTypesById command
var GetDcimModuleTypesByIdCmd = &cobra.Command{
	Use:   "getDcimModuleTypesById",
	Short: "GET an module type object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an module type object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(moduleTypesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.module_types_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    Metropolis Module Type: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
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
			if responseObject.Model != "" {
				color.Cyan("\tModel: " + color.YellowString("%s", responseObject.Model))
			} else {
				color.Cyan("\tModel: " + color.RedString("No model entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.PartNumber != "" {
				color.Cyan("\tPart Number: " + color.YellowString("%s", responseObject.PartNumber))
			} else {
				color.Cyan("\tPart Number: " + color.RedString("No part number entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Weight > 0.00 {
				color.Cyan("\tWeight: " + color.YellowString("%2f", responseObject.Weight))
			} else {
				color.Cyan("\tWeight: " + color.RedString("No weight entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.WeightUnit.Value != "" {
				color.Cyan("\tWeight Unit: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.WeightUnit.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.WeightUnit.Label))
			} else {
				color.Cyan("\tWeight Unit: " + color.RedString("No weight unit entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Comments != "" {
				color.Cyan("\tComments: " + color.YellowString("%s", responseObject.Comments))
			} else {
				color.Cyan("\tComments: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
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
			color.Cyan("\tLast Updated: " + color.YellowString("%s\n", responseObject.LastUpdated))
		} else {
			color.Red("  Doh! No module type object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimModuleTypesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimModuleTypesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking id flag as required flag: %s - for GetDcimModuleTypesByIdCmd", err)
	}

	GetDcimModuleTypesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the module type object")
	err = GetDcimModuleTypesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking id flag as required flag: %s - for GetDcimModuleTypesByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimModuleTypesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimModuleTypesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

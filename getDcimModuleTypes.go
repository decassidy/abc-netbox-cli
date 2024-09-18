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

type moduleTypes struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Id           int    `json:"id"`
		Url          string `json:"url"`
		Display      string `json:"display"`
		Manufacturer struct {
			CommonFieldsSlug
		} `json:"manufacturer"`
		Model      string `json:"model"`
		PartNumber string `json:"part_number"`
		Weight     int    `json:"weight"`
		WeightUnit struct {
			ValueLabel
		} `json:"weight_unit"`
		Description string `json:"description"`
		Comments    string `json:"comments"`
		Tags        []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created     string `json:"created"`
		LastUpdated string `json:"last_updated"`
	} `json:"results"`
}

// GetDcimModuleTypesCmd represents the getDcimModuleTypes command
var GetDcimModuleTypesCmd = &cobra.Command{
	Use:   "getDcimModuleTypes",
	Short: "GET a list of module type objects",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of module type objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(moduleTypes)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.module_types")

		if responseObject.Count != 0 {
			color.Cyan("\n  Metropolis Module Types: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    Metropolis Module Type: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				color.Cyan("\tManufacturer: " + color.YellowString("%s", result.Manufacturer))
				color.Cyan("\t  ID: " + color.YellowString("%d", result.Manufacturer.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.Manufacturer.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.Manufacturer.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", result.Manufacturer.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", result.Manufacturer.Slug))
				color.Cyan("\tModel: " + color.YellowString("%s", result.Model))
				color.Cyan("\tPart Number: " + color.YellowString("%s", result.PartNumber))
				color.Cyan("\tWeight: " + color.YellowString("%d", result.Weight))
				color.Cyan("\tWeight Unit: ")
				color.Cyan("\t  Value: " + color.YellowString("%d", result.WeightUnit.Value))
				color.Cyan("\t  Label: " + color.YellowString("%d", result.WeightUnit.Label))
				color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				color.Cyan("\tComments: " + color.YellowString("%s", result.Comments))
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
				color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
			}
		} else {
			color.Cyan("  Metropolis Module Types: " + color.RedString("No modules found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimModuleTypesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimModuleTypesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimModuleTypesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimModuleTypesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimModuleTypesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

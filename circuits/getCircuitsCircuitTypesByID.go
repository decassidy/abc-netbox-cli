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

package circuits

import (
	"github.com/decassidy/metropolis-netbox-cli/cmd/dcim"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
)

type circuitTypesByID struct {
	dcim.CommonFieldsSlug
	Color       string `json:"color"`
	Description string `json:"description"`
	Tags        []struct {
		dcim.CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created      string `json:"created"`
	LastUpdated  string `json:"last_updated"`
	CircuitCount int    `json:"circuit_count"`
}

// GetCircuitsCircuitTypesByIdCmd represents the getCircuitsCircuitTypesById command
var GetCircuitsCircuitTypesByIdCmd = &cobra.Command{
	Use:   "getCircuitsCircuitTypesById",
	Short: "Get a circuit type object.",
	Long: `
Metropolis Netbox Automation Tools:
  Get a circuit type object.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(circuitTypesByID)
		apiConnectionID(responseObject, "GET", "cmd.circuits.circuits_api_url.circuit_types_id")

		if responseObject.Id > 0 {
			color.Cyan("\n============================================================================")
			color.Cyan("\n\tMetropolis Circuit Type Name: "+color.YellowString("%s\n"), responseObject.Display)
			color.Cyan("============================================================================\n")
			color.Cyan("\tID: "+color.YellowString("%d"), responseObject.Id)
			color.Cyan("\tURL: "+color.YellowString("%s"), responseObject.Url)
			color.Cyan("\tDisplay: "+color.YellowString("%s"), responseObject.Display)
			color.Cyan("\tName: "+color.YellowString("%s"), responseObject.Name)
			color.Cyan("\tSlug: "+color.YellowString("%s"), responseObject.Slug)
			if responseObject.Color != "" {
				color.Cyan("\tColor: "+color.YellowString("%s"), responseObject.Color)
			} else {
				color.Cyan("\tColor: " + color.RedString("No color found for type: %s", color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: "+color.YellowString("%s"), responseObject.Description)
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description found for type: %s", color.YellowString("%s", responseObject.Display)))
			}
			for _, tag := range responseObject.Tags {
				if tag.Id != 0 {
					color.Cyan("\tTag: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), tag.Id)
					color.Cyan("\t  URL: "+color.YellowString("%d"), tag.Url)
					color.Cyan("\t  Display: "+color.YellowString("%d"), tag.Display)
					color.Cyan("\t  Name: "+color.YellowString("%d"), tag.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%d"), tag.Slug)
					color.Cyan("\t  Color: "+color.YellowString("%d"), tag.Color)
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags found for type: %s", color.YellowString("%s", responseObject.Display)))
				}
			}
			color.Cyan("\tCreated: "+color.YellowString("%s"), responseObject.Created)
			color.Cyan("\tLast Updated: "+color.YellowString("%s"), responseObject.LastUpdated)
		} else {
			color.Red("  Doh! No circuit type object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetCircuitsCircuitTypesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetCircuitsCircuitTypesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetCircuitsCircuitTypesByIdCmd", err)
	}

	GetCircuitsCircuitTypesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the circuit type object")
	err = GetCircuitsCircuitTypesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCircuitsCircuitTypesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCircuitsCircuitTypesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

package circuits

import (
	"github.com/decassidy/abc-netbox-cli/cmd/dcim"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
)

type circuitTypes struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
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
	} `json:"results"`
}

// GetCircuitsCircuitTypesCmd represents the circuitTypes command
var GetCircuitsCircuitTypesCmd = &cobra.Command{
	Use:   "getCircuitsCircuitTypes",
	Short: "GET a list of Circuit Type objects.",
	Long: `
ABC Netbox Automation Tools:
  GET a list of Circuit Type objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(circuitTypes)
		ApiConnectionNonID(responseObject, "GET", "cmd.circuits.circuits_api_url.circuit_types")

		if responseObject.Count != 0 {
			color.Cyan("\nABC Total Circuit Types in Netbox: "+color.YellowString("%d"), responseObject.Count)
			for _, types := range responseObject.Results {
				color.Cyan("\n============================================================================")
				color.Cyan("\n\tABC Circuit Type Name: "+color.YellowString("%s\n"), types.Display)
				color.Cyan("============================================================================\n")
				color.Cyan("\tID: "+color.YellowString("%d"), types.Id)
				color.Cyan("\tURL: "+color.YellowString("%s"), types.Url)
				color.Cyan("\tDisplay: "+color.YellowString("%s"), types.Display)
				color.Cyan("\tName: "+color.YellowString("%s"), types.Name)
				color.Cyan("\tSlug: "+color.YellowString("%s"), types.Slug)
				if types.Color != "" {
					color.Cyan("\tColor: "+color.YellowString("%s"), types.Color)
				} else {
					color.Cyan("\tColor: " + color.RedString("No color found for type: %s", color.YellowString("%s", types.Display)))
				}
				if types.Description != "" {
					color.Cyan("\tDescription: "+color.YellowString("%s"), types.Description)
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description found for type: %s", color.YellowString("%s", types.Display)))
				}
				for _, tag := range types.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTag: ")
						color.Cyan("\t  ID: "+color.YellowString("%d"), tag.Id)
						color.Cyan("\t  URL: "+color.YellowString("%d"), tag.Url)
						color.Cyan("\t  Display: "+color.YellowString("%d"), tag.Display)
						color.Cyan("\t  Name: "+color.YellowString("%d"), tag.Name)
						color.Cyan("\t  Slug: "+color.YellowString("%d"), tag.Slug)
						color.Cyan("\t  Color: "+color.YellowString("%d"), tag.Color)
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags found for type: %s", color.YellowString("%s", types.Display)))
					}
				}
				color.Cyan("\tCreated: "+color.YellowString("%s"), types.Created)
				color.Cyan("\tLast Updated: "+color.YellowString("%s"), types.LastUpdated)
			}
		} else {
			color.Cyan("\tCircuit Types: " + color.RedString("No circuit types found on the server: "))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetCircuitsCircuitTypesCmd.Flags().StringVarP(&serverEnv, "env", "", "", "Environment ('development' or 'production')")
	err := GetCircuitsCircuitTypesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env as required: %s - for GetCircuitsCircuitTerminationsByIDCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// circuitTypesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// circuitTypesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	GetCircuitsCircuitTypesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

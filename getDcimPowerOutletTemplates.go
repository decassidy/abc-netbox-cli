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

type powerOutletTemplates struct {
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
		ModuleType struct {
			Id           int    `json:"id"`
			Url          string `json:"url"`
			Display      string `json:"display"`
			Manufacturer struct {
				CommonFieldsSlug
			} `json:"manufacturer"`
			Model string `json:"model"`
		} `json:"module_type"`
		Name  string `json:"name"`
		Label string `json:"label"`
		Type  struct {
			ValueLabel
		} `json:"type"`
		PowerPort struct {
			CommonFieldsNoSlug
		} `json:"power_port"`
		FeedLeg struct {
			ValueLabel
		} `json:"feed_leg"`
		Description string `json:"description"`
		Created     string `json:"created"`
		LastUpdated string `json:"last_updated"`
	} `json:"results"`
}

// GetDcimPowerOutletTemplatesCmd represents the getDcimPowerOutletTemplates command
var GetDcimPowerOutletTemplatesCmd = &cobra.Command{
	Use:   "getDcimPowerOutletTemplates",
	Short: "GET a list of power outlet template objects",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of power outlet template objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(powerOutletTemplates)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.power_outlet_templates")

		if responseObject.Count != 0 {
			color.Cyan("\n  Metropolis Power Outlet Templates: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    Metropolis Power Outlet Template: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.DeviceType.Id != 0 {
					color.Cyan("\tDevice Type: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.DeviceType.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.DeviceType.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.DeviceType.Display))
					if result.DeviceType.Manufacturer.Id != 0 {
						color.Cyan("\t  Manufacturer: " + color.YellowString("%s", result.DeviceType.Manufacturer))
						color.Cyan("\t    ID: " + color.YellowString("%d", result.DeviceType.Manufacturer.Id))
						color.Cyan("\t    URL: " + color.YellowString("%d", result.DeviceType.Manufacturer.Url))
						color.Cyan("\t    Display: " + color.YellowString("%d", result.DeviceType.Manufacturer.Display))
						color.Cyan("\t    Name: " + color.YellowString("%d", result.DeviceType.Manufacturer.Name))
						color.Cyan("\t    Slug: " + color.YellowString("%d", result.DeviceType.Manufacturer.Slug))
					} else {
						color.Cyan("\t  Manufacturer" + color.RedString("No manufacturer entry found for ") + color.YellowString("%s", result.Display))
					}
					if result.DeviceType.Model != "" {
						color.Cyan("\t  Model: " + color.YellowString("%s", result.DeviceType.Model))
					} else {
						color.Cyan("\t  Model" + color.RedString("No model entry found for ") + color.YellowString("%s", result.Display))
					}
					if result.DeviceType.Slug != "" {
						color.Cyan("\t  Slug: " + color.YellowString("%s", result.DeviceType.Slug))
					} else {
						color.Cyan("\t  Slug" + color.RedString("No slug entry found for ") + color.YellowString("%s", result.Display))
					}
					if result.ModuleType.Id != 0 {
						color.Cyan("\tDevice Type: ")
						color.Cyan("\t  ID: " + color.YellowString("%d", result.ModuleType.Id))
						color.Cyan("\t  URL: " + color.YellowString("%s", result.ModuleType.Url))
						color.Cyan("\t  Display: " + color.YellowString("%s", result.ModuleType.Display))
						if result.ModuleType.Manufacturer.Id > 0 {
							color.Cyan("\t  Manufacturer: " + color.YellowString("%s", result.ModuleType.Manufacturer))
							color.Cyan("\t    ID: " + color.YellowString("%d", result.ModuleType.Manufacturer.Id))
							color.Cyan("\t    URL: " + color.YellowString("%d", result.ModuleType.Manufacturer.Url))
							color.Cyan("\t    Display: " + color.YellowString("%d", result.ModuleType.Manufacturer.Display))
							color.Cyan("\t    Name: " + color.YellowString("%d", result.ModuleType.Manufacturer.Name))
							color.Cyan("\t    Slug: " + color.YellowString("%d", result.ModuleType.Manufacturer.Slug))
						} else {
							color.Cyan("\t  Manufacturer" + color.RedString("No manufacturer entry found for ") + color.YellowString("%s", result.Display))
						}
						if result.DeviceType.Model != "" {
							color.Cyan("\t  Model: " + color.YellowString("%s", result.DeviceType.Model))
						} else {
							color.Cyan("\t  Model" + color.RedString("No model entry found for ") + color.YellowString("%s", result.Display))
						}
					}
				}
				if result.Name != "" {
					color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				} else {
					color.Cyan("\tName" + color.RedString("No name entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Label != "" {
					color.Cyan("\tLabel: " + color.YellowString("%s", result.Label))
				} else {
					color.Cyan("\tLabel" + color.RedString("No label entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Type.Value != "" {
					color.Cyan("\tType: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Type.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Type.Label))
				} else {
					color.Cyan("\tType" + color.RedString("No type entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.PowerPort.Id > 0 {
					color.Cyan("\tPower Port: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.PowerPort.Id))
					color.Cyan("\t  URL: " + color.YellowString("%d", result.PowerPort.Url))
					color.Cyan("\t  Display: " + color.YellowString("%d", result.PowerPort.Display))
					color.Cyan("\t  Name: " + color.YellowString("%d", result.PowerPort.Name))
				} else {
					color.Cyan("\tPower Port" + color.RedString("No power port entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.FeedLeg.Value != "" {
					color.Cyan("\tFeed Leg: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.FeedLeg.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.FeedLeg.Label))
				} else {
					color.Cyan("\tFeed Leg" + color.RedString("No feed leg entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription" + color.RedString("No description entry found for ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
			}
		} else {
			color.Cyan("  Metropolis Power Outlet Templates: " + color.RedString("No power outlet templates found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimPowerOutletTemplatesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimPowerOutletTemplatesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPowerOutletTemplatesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimPowerOutletTemplatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimPowerOutletTemplatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

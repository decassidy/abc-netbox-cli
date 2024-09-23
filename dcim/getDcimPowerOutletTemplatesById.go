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

type powerOutletTemplatesById struct {
	Id         uint   `json:"id"`
	Url        string `json:"url"`
	Display    string `json:"display"`
	DeviceType struct {
		Id           uint   `json:"id"`
		Url          string `json:"url"`
		Display      string `json:"display"`
		Manufacturer struct {
			CommonFieldsSlug
		} `json:"manufacturer"`
		Model string `json:"model"`
		Slug  string `json:"slug"`
	} `json:"device_type"`
	ModuleType struct {
		Id           uint   `json:"id"`
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
}

// GetDcimPowerOutletTemplatesByIdCmd represents the getDcimPowerOutletTemplatesById command
var GetDcimPowerOutletTemplatesByIdCmd = &cobra.Command{
	Use:   "getDcimPowerOutletTemplatesById",
	Short: "GET an power outlet template object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an power outlet object template by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(powerOutletTemplatesById)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.power_outlet_templates_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    ABC Power Outlet Template: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.DeviceType.Id != 0 {
				color.Cyan("\tDevice Type: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.DeviceType.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.DeviceType.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.DeviceType.Display))
				if responseObject.DeviceType.Manufacturer.Id != 0 {
					color.Cyan("\t  Manufacturer: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer))
					color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.DeviceType.Manufacturer.Id))
					color.Cyan("\t    URL: " + color.YellowString("%d", responseObject.DeviceType.Manufacturer.Url))
					color.Cyan("\t    Display: " + color.YellowString("%d", responseObject.DeviceType.Manufacturer.Display))
					color.Cyan("\t    Name: " + color.YellowString("%d", responseObject.DeviceType.Manufacturer.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%d", responseObject.DeviceType.Manufacturer.Slug))
				} else {
					color.Cyan("\t  Manufacturer" + color.RedString("No manufacturer entry found for ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.DeviceType.Model != "" {
					color.Cyan("\t  Model: " + color.YellowString("%s", responseObject.DeviceType.Model))
				} else {
					color.Cyan("\t  Model" + color.RedString("No model entry found for ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.DeviceType.Slug != "" {
					color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.DeviceType.Slug))
				} else {
					color.Cyan("\t  Slug" + color.RedString("No slug entry found for ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.ModuleType.Id != 0 {
					color.Cyan("\tDevice Type: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.ModuleType.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.ModuleType.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.ModuleType.Display))
					if responseObject.ModuleType.Manufacturer.Id > 0 {
						color.Cyan("\t  Manufacturer: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer))
						color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.ModuleType.Manufacturer.Id))
						color.Cyan("\t    URL: " + color.YellowString("%d", responseObject.ModuleType.Manufacturer.Url))
						color.Cyan("\t    Display: " + color.YellowString("%d", responseObject.ModuleType.Manufacturer.Display))
						color.Cyan("\t    Name: " + color.YellowString("%d", responseObject.ModuleType.Manufacturer.Name))
						color.Cyan("\t    Slug: " + color.YellowString("%d", responseObject.ModuleType.Manufacturer.Slug))
					} else {
						color.Cyan("\t  Manufacturer" + color.RedString("No manufacturer entry found for ") + color.YellowString("%s", responseObject.Display))
					}
					if responseObject.DeviceType.Model != "" {
						color.Cyan("\t  Model: " + color.YellowString("%s", responseObject.DeviceType.Model))
					} else {
						color.Cyan("\t  Model" + color.RedString("No model entry found for ") + color.YellowString("%s", responseObject.Display))
					}
				}
			}
			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tName" + color.RedString("No name entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Label != "" {
				color.Cyan("\tLabel: " + color.YellowString("%s", responseObject.Label))
			} else {
				color.Cyan("\tLabel" + color.RedString("No label entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Type.Value != "" {
				color.Cyan("\tType: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Type.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Type.Label))
			} else {
				color.Cyan("\tType" + color.RedString("No type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.PowerPort.Id != 0 {
				color.Cyan("\tPower Port: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.PowerPort.Id))
				color.Cyan("\t  URL: " + color.YellowString("%d", responseObject.PowerPort.Url))
				color.Cyan("\t  Display: " + color.YellowString("%d", responseObject.PowerPort.Display))
				color.Cyan("\t  Name: " + color.YellowString("%d", responseObject.PowerPort.Name))
			} else {
				color.Cyan("\tPower Port" + color.RedString("No power port entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.FeedLeg.Value != "" {
				color.Cyan("\tFeed Leg: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.FeedLeg.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.FeedLeg.Label))
			} else {
				color.Cyan("\tFeed Leg" + color.RedString("No feed leg entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription" + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s\n", responseObject.LastUpdated))
		} else {
			color.Red("  Doh! No power outlet template object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimPowerOutletTemplatesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimPowerOutletTemplatesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimPowerOutletTemplatesByIdCmd", err)
	}

	GetDcimPowerOutletTemplatesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the power outlet template")
	err = GetDcimPowerOutletTemplatesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimPowerOutletTemplatesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimPowerOutletTemplatesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

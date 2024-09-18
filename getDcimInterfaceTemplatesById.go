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

type interfaceTemplatesByID struct {
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
	Enabled     bool   `json:"enabled"`
	MgmtOnly    bool   `json:"mgmt_only"`
	Description string `json:"description"`
	Bridge      struct {
		CommonFieldsNoSlug
	} `json:"bridge"`
	PoeMode struct {
		ValueLabel
	} `json:"poe_mode"`
	PoeType struct {
		ValueLabel
	} `json:"poe_type"`
	RfRole struct {
		ValueLabel
	} `json:"rf_role"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

// GetDcimInterfaceTemplatesByIdCmd represents the getDcimInterfaceTemplatesById command
var GetDcimInterfaceTemplatesByIdCmd = &cobra.Command{
	Use:   "getDcimInterfaceTemplatesById",
	Short: "GET an interface template object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an interface template object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(interfaceTemplatesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.interface_templates_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    Metropolis Interface Template Name: %s\n", color.YellowString(responseObject.Display))
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
			} else {
				color.Cyan("\tDevice Type: " + color.RedString("No device type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.DeviceType.Manufacturer.Id != 0 {
				color.Cyan("\t  Manufacturer: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.DeviceType.Manufacturer.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Slug))
			} else {
				color.Cyan("\t  Manufacturer: " + color.RedString("No manufacturer entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.DeviceType.Model != "" {
				color.Cyan("\t  Model: " + color.YellowString("%s", responseObject.DeviceType.Model))
			} else {
				color.Cyan("\t  Model: " + color.RedString("No model entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.DeviceType.Slug != "" {
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.DeviceType.Slug))
			} else {
				color.Cyan("\t  Slug: " + color.RedString("No slug entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.ModuleType.Id != 0 {
				color.Cyan("\tModule Type: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.ModuleType.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.ModuleType.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.ModuleType.Display))
				color.Cyan("\t  Manufacturer: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.ModuleType.Manufacturer.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Slug))
				color.Cyan("\t  Model: " + color.YellowString("%s", responseObject.ModuleType.Model))

			} else {
				color.Cyan("\tModel Type: " + color.RedString("No model type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tName: " + color.RedString("No name entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Label != "" {
				color.Cyan("\tLabel: " + color.YellowString("%s", responseObject.Label))
			} else {
				color.Cyan("\tLabel: " + color.RedString("No label entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Type.Value != "" {
				color.Cyan("\tType: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Type.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Type.Label))
			} else {
				color.Cyan("\tType: " + color.RedString("No type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tEnabled: " + color.YellowString("%t", responseObject.Enabled))
			color.Cyan("\tMgmt Only: " + color.YellowString("%t", responseObject.MgmtOnly))
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Bridge.Id != 0 {
				color.Cyan("\tBridge: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Bridge.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Bridge.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Bridge.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Bridge.Name))
			} else {
				color.Cyan("\tBridge: " + color.RedString("No bridge entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.PoeMode.Value != "" {
				color.Cyan("\tPoE Mode: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.PoeMode.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.PoeMode.Label))
			} else {
				color.Cyan("\tPoE Mode: " + color.RedString("No PoE mode entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.PoeType.Value != "" {
				color.Cyan("\tPoE Type: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.PoeType.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.PoeType.Label))
			} else {
				color.Cyan("\tPoE Type: " + color.RedString("No PoE type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.RfRole.Value != "" {
				color.Cyan("\tRF Role: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.RfRole.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.RfRole.Label))
			} else {
				color.Cyan("\tRF Role: " + color.RedString("No RF role entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s\n", responseObject.LastUpdated))
		} else {
			color.Red("  Doh! No Interface template object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimInterfaceTemplatesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimInterfaceTemplatesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	GetDcimInterfaceTemplatesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the interface template")
	err = GetDcimInterfaceTemplatesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimInterfaceTemplatesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimInterfaceTemplatesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

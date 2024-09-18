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

type frontPortTemplatesByID struct {
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
	Color    string `json:"color"`
	RearPort struct {
		CommonFieldsNoSlug
	} `json:"rear_port"`
	RearPortPosition int    `json:"rear_port_position"`
	Description      string `json:"description"`
	Created          string `json:"created"`
	LastUpdated      string `json:"last_updated"`
}

// GetDcimFrontPortTemplatesByIdCmd represents the getDcimFrontPortTemplatesById command
var GetDcimFrontPortTemplatesByIdCmd = &cobra.Command{
	Use:   "getDcimFrontPortTemplatesById",
	Short: "GET an front port template object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an front port template object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(frontPortTemplatesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.front_port_templates_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    Metropolis Front Port Template Name: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			color.Cyan("\tDevice Type: ")
			color.Cyan("\t  ID: " + color.YellowString("%s", responseObject.DeviceType.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.DeviceType.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.DeviceType.Display))
			color.Cyan("\t  Manufacturer: ")
			color.Cyan("\t    ID: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Name))
			color.Cyan("\t    Slug: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Slug))
			color.Cyan("\t  Model: " + color.YellowString("%s", responseObject.DeviceType.Model))
			color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.DeviceType.Slug))
			color.Cyan("\tModule Type: ")
			color.Cyan("\t  ID: " + color.YellowString("%s", responseObject.ModuleType.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.ModuleType.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.ModuleType.Display))
			color.Cyan("\t  Manufacturer: ")
			color.Cyan("\t    ID: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Name))
			color.Cyan("\t    Slug: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Slug))
			color.Cyan("\t  Model: " + color.YellowString("%s", responseObject.ModuleType.Model))
			color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			color.Cyan("\tLabel: " + color.YellowString("%s", responseObject.Label))
			color.Cyan("\tType: ")
			color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Type.Value))
			color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Type.Label))
			color.Cyan("\tColor: " + color.YellowString("%s", responseObject.Color))
			color.Cyan("\tRear Port: " + color.YellowString("%s", responseObject.RearPort))
			color.Cyan("\t  ID: " + color.YellowString("%s", responseObject.RearPort.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.RearPort.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.RearPort.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.RearPort.Name))
			color.Cyan("\tRear Port Position: " + color.YellowString("%d", responseObject.RearPortPosition))
			color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s\n", responseObject.LastUpdated))
		} else {
			color.Red("  Doh! Front Port Template object not found with ID: " + color.YellowString("%s\n", id))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimFrontPortTemplatesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimFrontPortTemplatesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	GetDcimFrontPortTemplatesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of front port template object")
	err = GetDcimFrontPortTemplatesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimFrontPortTemplatesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimFrontPortTemplatesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

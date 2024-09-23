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

type moduleBayTemplatesByID struct {
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
	Name        string `json:"name"`
	Label       string `json:"label"`
	Position    string `json:"position"`
	Description string `json:"description"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

// GetDcimModuleBayTemplatesByIdCmd represents the getDcimModuleBayTemplatesById command
var GetDcimModuleBayTemplatesByIdCmd = &cobra.Command{
	Use:   "getDcimModuleBayTemplatesById",
	Short: "GET an module bay template object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an module bay template object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(moduleBayTemplatesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.module_bay_templates_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    ABC Module Bay Templates: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.DeviceType.Id != 0 {
				color.Cyan("\tDevice Type: " + color.YellowString("%s", responseObject.DeviceType))
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.DeviceType.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.DeviceType.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.DeviceType.Display))
				color.Cyan("\t  Manufacturer: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.DeviceType.Manufacturer.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", responseObject.DeviceType.Manufacturer.Slug))
				color.Cyan("\t  Model: " + color.YellowString("%s", responseObject.DeviceType.Model))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.DeviceType.Slug))
			} else {
				color.Cyan("\tDevice Type: " + color.RedString("No device type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Name != "" {
				color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			} else {
				color.Cyan("\tName: " + color.RedString("No name entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Label != "" {
				color.Cyan("\tLabel: " + color.YellowString("%d", responseObject.Label))
			} else {
				color.Cyan("\tLabel: " + color.RedString("No label entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Position != "" {
				color.Cyan("\tPosition: " + color.YellowString("%d", responseObject.Position))
			} else {
				color.Cyan("\tPosition: " + color.RedString("No position entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%d", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No position entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s\n", responseObject.LastUpdated))

		} else {
			color.Red("  Doh! No module bay template object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimModuleBayTemplatesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimModuleBayTemplatesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required. %s - for GetDcimModuleBayTemplatesByIdCmd", err)
	}

	GetDcimModuleBayTemplatesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the module bay template object")
	err = GetDcimModuleBayTemplatesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required. %s - for GetDcimModuleBayTemplatesByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimModuleBayTemplatesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimModuleBayTemplatesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

type consoleServerPortTemplatesByID struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	Display    string `json:"display"`
	DeviceType struct {
		Id           int              `json:"id"`
		Url          string           `json:"url"`
		Display      string           `json:"display"`
		Manufacturer CommonFieldsSlug `json:"manufacturer"`
		Model        string           `json:"model"`
		Slug         string           `json:"slug"`
	} `json:"device_type"`
	ModuleType struct {
		Id           int              `json:"id"`
		Url          string           `json:"url"`
		Display      string           `json:"display"`
		Manufacturer CommonFieldsSlug `json:"manufacturer"`
		Model        string           `json:"model"`
	} `json:"module_type"`
	Name        string     `json:"name"`
	Label       string     `json:"label"`
	Type        ValueLabel `json:"type"`
	Description string     `json:"description"`
	Created     string     `json:"created"`
	LastUpdated string     `json:"last_updated"`
}

// GetDcimConsoleServerPortTemplatesByIdCmd represents the getDcimConsoleServerPortTemplatesById command
var GetDcimConsoleServerPortTemplatesByIdCmd = &cobra.Command{
	Use:   "getDcimConsoleServerPortTemplatesById",
	Short: "GET an console server port template object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an console server port template object by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(consoleServerPortTemplatesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.console_server_port_templates_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    ABC Console Server Port Template: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: %d", responseObject.Id)
			color.Cyan("\tURL: %s", responseObject.Url)
			color.Cyan("\tDisplay: %s", responseObject.Display)
			if responseObject.DeviceType.Id != 0 {
				color.Cyan("\tDevice Type: ")
				color.Cyan("\t  ID: %d", responseObject.DeviceType.Id)
				color.Cyan("\t  URL: %s", responseObject.DeviceType.Url)
				color.Cyan("\t  Display: %s", responseObject.DeviceType.Display)
				color.Cyan("\t  Manufacturer: ")
				color.Cyan("\t    ID: %d", responseObject.DeviceType.Manufacturer.Id)
				color.Cyan("\t    URL: %s", responseObject.DeviceType.Manufacturer.Url)
				color.Cyan("\t    Display: %s", responseObject.DeviceType.Manufacturer.Display)
				color.Cyan("\t    Name: %s", responseObject.DeviceType.Manufacturer.Name)
				color.Cyan("\t    Slug: %s", responseObject.DeviceType.Manufacturer.Slug)
				color.Cyan("\t  Model: %s", responseObject.DeviceType.Model)
				color.Cyan("\t  Slug: %s", responseObject.DeviceType.Slug)
			} else {
				color.Cyan("\tDevice Type: " + color.RedString("No device type entry found for: ") + color.YellowString("%s", responseObject.DeviceType.Display))
			}
			if responseObject.ModuleType.Id != 0 {
				color.Cyan("\tModule Type: ")
				color.Cyan("\t  ID: %d", responseObject.ModuleType.Id)
				color.Cyan("\t  URL: %s", responseObject.ModuleType.Url)
				color.Cyan("\t  Display: %s", responseObject.ModuleType.Display)
				color.Cyan("\t  Manufacturer: ")
				color.Cyan("\t    ID: %d", responseObject.ModuleType.Manufacturer.Id)
				color.Cyan("\t    URL: %s", responseObject.ModuleType.Manufacturer.Url)
				color.Cyan("\t    Display: %s", responseObject.ModuleType.Manufacturer.Display)
				color.Cyan("\t    Name: %s", responseObject.ModuleType.Manufacturer.Name)
				color.Cyan("\t    Slug: %s", responseObject.ModuleType.Manufacturer.Slug)
				color.Cyan("\t  Model: %s", responseObject.ModuleType.Model)
			} else {
				color.Cyan("\tModule Type: " + color.RedString("No model type entry found for: ") + color.YellowString("%s", responseObject.DeviceType.Display))
			}
			if responseObject.Name != "" {
				color.Cyan("\tName: %s", responseObject.Name)
			} else {
				color.Cyan("\tName: " + color.RedString("No name entry found for: ") + color.YellowString("%s", responseObject.DeviceType.Display))
			}
			if responseObject.Label != "" {
				color.Cyan("\tLabel: %s", responseObject.Label)
			} else {
				color.Cyan("\tLabel: " + color.RedString("No label entry found for: ") + color.YellowString("%s", responseObject.DeviceType.Display))
			}
			if responseObject.Type.Value != "" {
				color.Cyan("\tType: %s", responseObject.Type.Value)
				color.Cyan("\t  Value: %s", responseObject.Type.Value)
				color.Cyan("\t  Label: %s", responseObject.Type.Label)
			} else {
				color.Cyan("\tType: " + color.RedString("No type entry found for: ") + color.YellowString("%s", responseObject.DeviceType.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: %s", responseObject.Description)
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for: ") + color.YellowString("%s", responseObject.DeviceType.Display))
			}
			color.Cyan("\tCreated: %s", responseObject.Created)
			color.Cyan("\tLast Updated: %s\n", responseObject.LastUpdated)
		} else {
			color.Cyan("  ABC Console Server Port Template: " + color.RedString("No console server port template entries found on server for ID: %d. Exiting...\n", id))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimConsoleServerPortTemplatesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimConsoleServerPortTemplatesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimConsoleServerPortTemplatesByIdCmd", err)
	}

	GetDcimConsoleServerPortTemplatesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the console server port template")
	err = GetDcimConsoleServerPortTemplatesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimConsoleServerPortTemplatesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimConsoleServerPortTemplatesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

type modulesByID struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Device  struct {
		CommonFieldsNoSlug
	} `json:"device"`
	ModuleBay struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Module  struct {
			Id      int    `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Device  struct {
				CommonFieldsNoSlug
			} `json:"device"`
			ModuleBay struct {
				CommonFieldsNoSlug
			} `json:"module_bay"`
			ModuleType struct {
				Id           int    `json:"id"`
				Url          string `json:"url"`
				Display      string `json:"display"`
				Manufacturer struct {
					CommonFieldsSlug
				} `json:"manufacturer"`
				Model string `json:"model"`
			} `json:"module_type"`
		} `json:"module"`
		Name string `json:"name"`
	} `json:"module_bay"`
	ModuleType struct {
		Id           int    `json:"id"`
		Url          string `json:"url"`
		Display      string `json:"display"`
		Manufacturer struct {
			CommonFieldsSlug
		} `json:"manufacturer"`
		Model string `json:"model"`
	} `json:"module_type"`
	Status struct {
		ValueLabel
	} `json:"status"`
	Serial      string `json:"serial"`
	AssetTag    string `json:"asset_tag"`
	Description string `json:"description"`
	Comments    string `json:"comments"`
	Tags        []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

// GetDcimModulesByIdCmd represents the getDcimModulesById command
var GetDcimModulesByIdCmd = &cobra.Command{
	Use:   "getDcimModulesById",
	Short: "GET an module object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an module object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(modulesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.modules_id")

		if responseObject.Id != 0 {
			display := fmt.Sprintf("    Metropolis Module: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.Device.Id != 0 {
				color.Cyan("\tDevice Type: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Device.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Device.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Device.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Device.Name))
			} else {
				color.Cyan("\tDevice: " + color.RedString("No device entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.ModuleBay.Id != 0 {
				color.Cyan("\tModule Bay: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.ModuleBay.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.ModuleBay.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.ModuleBay.Display))
				if responseObject.ModuleBay.Module.Id != 0 {
					color.Cyan("\t  Module: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.ModuleBay.Module.Id))
					color.Cyan("\t    URL: " + color.YellowString("%d", responseObject.ModuleBay.Module.Url))
					color.Cyan("\t    Display: " + color.YellowString("%d", responseObject.ModuleBay.Module.Display))
				} else {
					color.Cyan("\t  Module: " + color.RedString("No module entry found for ") + color.YellowString("%s", responseObject.Display))
				}
				color.Cyan("\t    Device: ")
				color.Cyan("\t     ID: " + color.YellowString("%d", responseObject.ModuleBay.Module.Device.Id))
				color.Cyan("\t     URL: " + color.YellowString("%s", responseObject.ModuleBay.Module.Device.Url))
				color.Cyan("\t     Display: " + color.YellowString("%s", responseObject.ModuleBay.Module.Device.Display))
				color.Cyan("\t     Name: " + color.YellowString("%s", responseObject.ModuleBay.Module.Device.Name))
				color.Cyan("\t    Module Bay: ")
				color.Cyan("\t     ID: " + color.YellowString("%d", responseObject.ModuleBay.Module.ModuleBay.Id))
				color.Cyan("\t     URL: " + color.YellowString("%s", responseObject.ModuleBay.Module.ModuleBay.Url))
				color.Cyan("\t     Display: " + color.YellowString("%s", responseObject.ModuleBay.Module.ModuleBay.Display))
				color.Cyan("\t     Name: " + color.YellowString("%s", responseObject.ModuleBay.Module.ModuleBay.Name))
				color.Cyan("\t    Module Type: ")
				color.Cyan("\t     ID: " + color.YellowString("%d", responseObject.ModuleBay.Module.ModuleType.Id))
				color.Cyan("\t     URL: " + color.YellowString("%s", responseObject.ModuleBay.Module.ModuleType.Url))
				color.Cyan("\t     Display: " + color.YellowString("%s", responseObject.ModuleBay.Module.ModuleType.Display))
				color.Cyan("\t     Manufacturer: ")
				color.Cyan("\t       ID: " + color.YellowString("%d", responseObject.ModuleBay.Module.ModuleType.Manufacturer.Id))
				color.Cyan("\t       URL: " + color.YellowString("%s", responseObject.ModuleBay.Module.ModuleType.Manufacturer.Url))
				color.Cyan("\t       Display: " + color.YellowString("%s", responseObject.ModuleBay.Module.ModuleType.Manufacturer.Display))
				color.Cyan("\t       Name: " + color.YellowString("%s", responseObject.ModuleBay.Module.ModuleType.Manufacturer.Name))
				color.Cyan("\t       Slug: " + color.YellowString("%s", responseObject.ModuleBay.Module.ModuleType.Manufacturer.Slug))
				color.Cyan("\t     Model: " + color.YellowString("%s", responseObject.ModuleBay.Module.ModuleType.Model))
				color.Cyan("\t   Name: " + color.YellowString("%s", responseObject.ModuleBay.Name))
			} else {
				color.Cyan("\tModule Bay: " + color.RedString("No module bay entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.ModuleType.Id != 0 {
				color.Cyan("\tModule Type: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.ModuleType.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.ModuleType.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.ModuleType.Display))
				color.Cyan("\t  Manufacturer: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer))
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.ModuleType.Manufacturer.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", responseObject.ModuleType.Manufacturer.Slug))
				color.Cyan("\t  Model: " + color.YellowString("%s", responseObject.ModuleType.Model))
			} else {
				color.Cyan("\tModule Type: " + color.RedString("No module type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			color.Cyan("\tStatus: ")
			color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Status.Value))
			color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Status.Label))
			color.Cyan("\tSerial: " + color.YellowString("%s", responseObject.Serial))
			color.Cyan("\tAsset Tag: " + color.YellowString("%s", responseObject.AssetTag))
			color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			color.Cyan("\tComments: " + color.YellowString("%s", responseObject.Comments))
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
			color.Red("  Doh! No module object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimModulesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimModulesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking id flag as required flag: %s - for GetDcimModulesByIdCmd", err)
	}

	GetDcimModulesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the module object")
	err = GetDcimModulesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking id flag as required flag: %s - for GetDcimModulesByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimModulesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimModulesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

type modules struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
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
	} `json:"results"`
}

// GetDcimModulesCmd represents the getDcimModules command
var GetDcimModulesCmd = &cobra.Command{
	Use:   "getDcimModules",
	Short: "GET a list of module objects",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of module objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(modules)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.modules")

		if responseObject.Count != 0 {
			color.Cyan("\n  Metropolis Modules: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    Metropolis Modules: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.Device.Id != 0 {
					color.Cyan("\tDevice Type: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Device.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Device.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Device.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Device.Name))
				} else {
					color.Cyan("\tDevice: " + color.RedString("No device entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.ModuleBay.Id != 0 {
					color.Cyan("\tModule Bay: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.ModuleBay.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.ModuleBay.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.ModuleBay.Display))
					if result.ModuleBay.Module.Id != 0 {
						color.Cyan("\t  Module: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", result.ModuleBay.Module.Id))
						color.Cyan("\t    URL: " + color.YellowString("%d", result.ModuleBay.Module.Url))
						color.Cyan("\t    Display: " + color.YellowString("%d", result.ModuleBay.Module.Display))
					} else {
						color.Cyan("\t  Module: " + color.RedString("No module entry found for ") + color.YellowString("%s", result.Display))
					}
					color.Cyan("\t    Device: ")
					color.Cyan("\t     ID: " + color.YellowString("%d", result.ModuleBay.Module.Device.Id))
					color.Cyan("\t     URL: " + color.YellowString("%s", result.ModuleBay.Module.Device.Url))
					color.Cyan("\t     Display: " + color.YellowString("%s", result.ModuleBay.Module.Device.Display))
					color.Cyan("\t     Name: " + color.YellowString("%s", result.ModuleBay.Module.Device.Name))
					color.Cyan("\t    Module Bay: ")
					color.Cyan("\t     ID: " + color.YellowString("%d", result.ModuleBay.Module.ModuleBay.Id))
					color.Cyan("\t     URL: " + color.YellowString("%s", result.ModuleBay.Module.ModuleBay.Url))
					color.Cyan("\t     Display: " + color.YellowString("%s", result.ModuleBay.Module.ModuleBay.Display))
					color.Cyan("\t     Name: " + color.YellowString("%s", result.ModuleBay.Module.ModuleBay.Name))
					color.Cyan("\t    Module Type: ")
					color.Cyan("\t     ID: " + color.YellowString("%d", result.ModuleBay.Module.ModuleType.Id))
					color.Cyan("\t     URL: " + color.YellowString("%s", result.ModuleBay.Module.ModuleType.Url))
					color.Cyan("\t     Display: " + color.YellowString("%s", result.ModuleBay.Module.ModuleType.Display))
					color.Cyan("\t     Manufacturer: ")
					color.Cyan("\t       ID: " + color.YellowString("%d", result.ModuleBay.Module.ModuleType.Manufacturer.Id))
					color.Cyan("\t       URL: " + color.YellowString("%s", result.ModuleBay.Module.ModuleType.Manufacturer.Url))
					color.Cyan("\t       Display: " + color.YellowString("%s", result.ModuleBay.Module.ModuleType.Manufacturer.Display))
					color.Cyan("\t       Name: " + color.YellowString("%s", result.ModuleBay.Module.ModuleType.Manufacturer.Name))
					color.Cyan("\t       Slug: " + color.YellowString("%s", result.ModuleBay.Module.ModuleType.Manufacturer.Slug))
					color.Cyan("\t     Model: " + color.YellowString("%s", result.ModuleBay.Module.ModuleType.Model))
					color.Cyan("\t   Name: " + color.YellowString("%s", result.ModuleBay.Name))
				} else {
					color.Cyan("\tModule Bay: " + color.RedString("No module bay entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.ModuleType.Id != 0 {
					color.Cyan("\tModule Type: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.ModuleType.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.ModuleType.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.ModuleType.Display))
					color.Cyan("\t  Manufacturer: " + color.YellowString("%s", result.ModuleType.Manufacturer))
					color.Cyan("\t    ID: " + color.YellowString("%d", result.ModuleType.Manufacturer.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.ModuleType.Manufacturer.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.ModuleType.Manufacturer.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.ModuleType.Manufacturer.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.ModuleType.Manufacturer.Slug))
					color.Cyan("\t  Model: " + color.YellowString("%s", result.ModuleType.Model))
				} else {
					color.Cyan("\tModule Type: " + color.RedString("No module type entry found for ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tStatus: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", result.Status.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", result.Status.Label))
				color.Cyan("\tSerial: " + color.YellowString("%s", result.Serial))
				color.Cyan("\tAsset Tag: " + color.YellowString("%s", result.AssetTag))
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
			color.Cyan("  Metropolis Modules: " + color.RedString("No modules found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimModulesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimModulesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimModulesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimModulesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimModulesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

type regionsByID struct {
	CommonFieldsSlug
	Parent struct {
		CommonFieldsSlug
		Depth int `json:"_depth"`
	} `json:"parent"`
	Description string `json:"description"`
	Tags        []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
	SiteCount   int    `json:"site_count"`
	Depth       int    `json:"_depth"`
}

// GetDcimRegionsByIdCmd represents the getDcimRegionsById command
var GetDcimRegionsByIdCmd = &cobra.Command{
	Use:   "getDcimRegionsById",
	Short: "GET an region object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an region object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(regionsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.regions_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    ABC Rear Port Template: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			color.Cyan("\tSlug: " + color.YellowString("%s", responseObject.Slug))
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			for _, tag := range responseObject.Tags {
				if tag.Id > 0 {
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
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))

			if responseObject.SiteCount > 0 {
				color.Cyan("\tSite Count: " + color.YellowString("%d", responseObject.SiteCount))
			} else {
				color.Cyan("\tSite Count: " + color.RedString("No site count entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Depth > 0 {
				color.Cyan("\tDepth: " + color.YellowString("%d\n", responseObject.Depth))
			} else {
				color.Cyan("\tDepth: " + color.RedString("No depth entry found for ") + color.YellowString("%s\n", responseObject.Display))
			}
		} else {
			color.Red("  Doh! No region object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimRegionsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimRegionsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimRegionsByIdCmd", err)
	}

	GetDcimRegionsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the region object")
	err = GetDcimRegionsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimRegionsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimRegionsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

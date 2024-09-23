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

type racks struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		CommonFieldsNoSlug
		FacilityId string `json:"facility_id"`
		Site       struct {
			CommonFieldsSlug
		} `json:"site"`
		Location struct {
			CommonFieldsSlug
			Depth int `json:"_depth"`
		} `json:"location"`
		Tenant struct {
			CommonFieldsSlug
		} `json:"tenant"`
		Status struct {
			ValueLabel
		} `json:"status"`
		Role struct {
			CommonFieldsSlug
		} `json:"role"`
		Serial   string `json:"serial"`
		AssetTag string `json:"asset_tag"`
		Type     struct {
			ValueLabel
		} `json:"type"`
		Width struct {
			Value uint   `json:"value"`
			Label string `json:"label"`
		} `json:"width"`
		UHeight      uint `json:"u_height"`
		StartingUnit uint `json:"starting_unit"`
		Weight       uint `json:"weight"`
		MaxWeight    uint `json:"max_weight"`
		WeightUnit   struct {
			ValueLabel
		} `json:"weight_unit"`
		DescUnits  bool `json:"desc_units"`
		OuterWidth uint `json:"outer_width"`
		OuterDepth uint `json:"outer_depth"`
		OuterUnit  struct {
			ValueLabel
		} `json:"outer_unit"`
		MountingDepth uint   `json:"mounting_depth"`
		Description   string `json:"description"`
		Comments      string `json:"comments"`
		Tags          []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created        string `json:"created"`
		LastUpdated    string `json:"last_updated"`
		DeviceCount    uint   `json:"device_count"`
		PowerfeedCount uint   `json:"powerfeed_count"`
	} `json:"results"`
}

// GetDcimRacksCmd represents the getDcimRacks command
var GetDcimRacksCmd = &cobra.Command{
	Use:   "getDcimRacks",
	Short: "GET a list of rack objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of rack objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(racks)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.racks")

		if responseObject.Count != 0 {
			color.Cyan("\n  ABC Rack Count: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    ABC Rack: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.FacilityId != "" {
					color.Cyan("\tFacility ID: " + color.YellowString("%d", result.FacilityId))
				} else {
					color.Cyan("\tFacility ID: " + color.RedString("No facility ID entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Site.Id != 0 {
					color.Cyan("\tSite: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Site.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Site.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Site.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Site.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Site.Slug))
				} else {
					color.Cyan("\tSite" + color.RedString("No site entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Location.Id != 0 {
					color.Cyan("\tSite: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Location.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Location.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Location.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Location.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Location.Slug))
					color.Cyan("\t  Depth: " + color.YellowString("%d", result.Location.Depth))
				} else {
					color.Cyan("\tSite: " + color.RedString("No site entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Tenant.Id != 0 {
					color.Cyan("\tTenant: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Tenant.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Tenant.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Tenant.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Tenant.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Tenant.Slug))
				} else {
					color.Cyan("\tTenant: " + color.RedString("No tenant entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Status.Value != "" {
					color.Cyan("\tStatus: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Status.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Status.Label))
				} else {
					color.Cyan("\tStatus: " + color.RedString("No tenant entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Role.Id != 0 {
					color.Cyan("\tRole: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Role.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Role.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Role.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Role.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Role.Slug))
				} else {
					color.Cyan("\tRole: " + color.RedString("No role entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Serial != "" {
					color.Cyan("\tSerial: " + color.YellowString("%d", result.Serial))
				} else {
					color.Cyan("\tSerial: " + color.RedString("No serial entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.AssetTag != "" {
					color.Cyan("\tAsset Tag: " + color.YellowString("%d", result.AssetTag))
				} else {
					color.Cyan("\tAsset Tag: " + color.RedString("No asset tag entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Type.Value != "" {
					color.Cyan("\tType: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.Type.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Type.Label))
				} else {
					color.Cyan("\tType: " + color.RedString("No type entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Width.Value != 0 {
					color.Cyan("\tWidth: ")
					color.Cyan("\t  Value: " + color.YellowString("%d", result.Width.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.Width.Label))
				} else {
					color.Cyan("\tWidth: " + color.RedString("No width entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.UHeight != 0 {
					color.Cyan("\tU-Height: " + color.YellowString("%d", result.UHeight))
				} else {
					color.Cyan("\tU-Height: " + color.RedString("No u-height entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.StartingUnit != 0 {
					color.Cyan("\tStarting Unit: " + color.YellowString("%d", result.StartingUnit))
				} else {
					color.Cyan("\tStarting Unit: " + color.RedString("No starting unit entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Weight != 0 {
					color.Cyan("\tWeight: " + color.YellowString("%d", result.Weight))
				} else {
					color.Cyan("\tWeight: " + color.RedString("No weight entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.MaxWeight != 0 {
					color.Cyan("\tMax Weight: " + color.YellowString("%d", result.MaxWeight))
				} else {
					color.Cyan("\tMax Weight: " + color.RedString("No max weight entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.WeightUnit.Value != "" {
					color.Cyan("\tWeight Unit: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.WeightUnit.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.WeightUnit.Label))
				} else {
					color.Cyan("\tWeight Unit: " + color.RedString("No weight unit entry found for ") + color.YellowString("%s", result.Display))
				}

				color.Cyan("\tDesc Units: " + color.YellowString("%t", result.DescUnits))

				if result.OuterWidth != 0 {
					color.Cyan("\tOuter Width: " + color.YellowString("%d", result.OuterWidth))
				} else {
					color.Cyan("\tOuter Width: " + color.RedString("No outer width entry found for ") + color.YellowString("%s", result.Display))
				}

				if result.OuterDepth != 0 {
					color.Cyan("\tOuter Depth: " + color.YellowString("%d", result.OuterDepth))
				} else {
					color.Cyan("\tOuter Depth: " + color.RedString("No outer depth entry found for ") + color.YellowString("%s", result.Display))
				}

				if result.OuterUnit.Value != "" {
					color.Cyan("\tOuter Unit: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", result.OuterUnit.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", result.OuterUnit.Label))
				} else {
					color.Cyan("\tOuter Unit: " + color.RedString("No outer unit entry found for ") + color.YellowString("%s", result.Display))
				}

				if result.MountingDepth != 0 {
					color.Cyan("\tMounting Depth: " + color.YellowString("%d", result.MountingDepth))
				} else {
					color.Cyan("\tMounting Depth: " + color.RedString("No mounting depth entry found for ") + color.YellowString("%s", result.Display))
				}

				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", result.Display))
				}

				if result.Comments != "" {
					color.Cyan("\tComments: " + color.YellowString("%s", result.Comments))
				} else {
					color.Cyan("\tComments: " + color.RedString("No comments entry found for ") + color.YellowString("%s", result.Display))
				}

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
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))

				if result.DeviceCount != 0 {
					color.Cyan("\tDevice Count: " + color.YellowString("%d", result.DeviceCount))
				} else {
					color.Cyan("\tDevice Count: " + color.RedString("No device count entry found for ") + color.YellowString("%s", result.Display))
				}

				if result.PowerfeedCount != 0 {
					color.Cyan("\tPowerfeed Count: " + color.YellowString("%d\n", result.PowerfeedCount))
				} else {
					color.Cyan("\tPowerfeed Count: " + color.RedString("No powerfeed count entry found for ") + color.YellowString("%s\n", result.Display))
				}
			}
		} else {
			color.Cyan("  ABC Racks: " + color.RedString("No racks found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimRacksCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimRacksCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimRacksCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimRacksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimRacksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

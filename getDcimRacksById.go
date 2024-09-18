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

type racksByID struct {
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
}

// GetDcimRacksByIdCmd represents the getDcimRacksById command
var GetDcimRacksByIdCmd = &cobra.Command{
	Use:   "getDcimRacksById",
	Short: "GET an rack object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an rack object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(racksByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.racks_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    Metropolis Rack: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.FacilityId != "" {
				color.Cyan("\tFacility ID: " + color.YellowString("%d", responseObject.FacilityId))
			} else {
				color.Cyan("\tFacility ID: " + color.RedString("No facility ID entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Site.Id != 0 {
				color.Cyan("\tSite: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Site.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Site.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Site.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Site.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Site.Slug))
			} else {
				color.Cyan("\tSite" + color.RedString("No site entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Location.Id != 0 {
				color.Cyan("\tSite: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Location.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Location.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Location.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Location.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Location.Slug))
				color.Cyan("\t  Depth: " + color.YellowString("%d", responseObject.Location.Depth))
			} else {
				color.Cyan("\tSite: " + color.RedString("No site entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Tenant.Id != 0 {
				color.Cyan("\tTenant: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Tenant.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Tenant.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Tenant.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Tenant.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Tenant.Slug))
			} else {
				color.Cyan("\tTenant: " + color.RedString("No tenant entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Status.Value != "" {
				color.Cyan("\tStatus: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Status.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Status.Label))
			} else {
				color.Cyan("\tStatus: " + color.RedString("No tenant entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Role.Id != 0 {
				color.Cyan("\tRole: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Role.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Role.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Role.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Role.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Role.Slug))
			} else {
				color.Cyan("\tRole: " + color.RedString("No role entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Serial != "" {
				color.Cyan("\tSerial: " + color.YellowString("%d", responseObject.Serial))
			} else {
				color.Cyan("\tSerial: " + color.RedString("No serial entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.AssetTag != "" {
				color.Cyan("\tAsset Tag: " + color.YellowString("%d", responseObject.AssetTag))
			} else {
				color.Cyan("\tAsset Tag: " + color.RedString("No asset tag entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Type.Value != "" {
				color.Cyan("\tType: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.Type.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Type.Label))
			} else {
				color.Cyan("\tType: " + color.RedString("No type entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Width.Value != 0 {
				color.Cyan("\tWidth: ")
				color.Cyan("\t  Value: " + color.YellowString("%d", responseObject.Width.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.Width.Label))
			} else {
				color.Cyan("\tWidth: " + color.RedString("No width entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.UHeight != 0 {
				color.Cyan("\tU-Height: " + color.YellowString("%d", responseObject.UHeight))
			} else {
				color.Cyan("\tU-Height: " + color.RedString("No u-height entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.StartingUnit != 0 {
				color.Cyan("\tStarting Unit: " + color.YellowString("%d", responseObject.StartingUnit))
			} else {
				color.Cyan("\tStarting Unit: " + color.RedString("No starting unit entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Weight != 0 {
				color.Cyan("\tWeight: " + color.YellowString("%d", responseObject.Weight))
			} else {
				color.Cyan("\tWeight: " + color.RedString("No weight entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.MaxWeight != 0 {
				color.Cyan("\tMax Weight: " + color.YellowString("%d", responseObject.MaxWeight))
			} else {
				color.Cyan("\tMax Weight: " + color.RedString("No max weight entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.WeightUnit.Value != "" {
				color.Cyan("\tWeight Unit: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.WeightUnit.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.WeightUnit.Label))
			} else {
				color.Cyan("\tWeight Unit: " + color.RedString("No weight unit entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			color.Cyan("\tDesc Units: " + color.YellowString("%t", responseObject.DescUnits))

			if responseObject.OuterWidth != 0 {
				color.Cyan("\tOuter Width: " + color.YellowString("%d", responseObject.OuterWidth))
			} else {
				color.Cyan("\tOuter Width: " + color.RedString("No outer width entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.OuterDepth != 0 {
				color.Cyan("\tOuter Depth: " + color.YellowString("%d", responseObject.OuterDepth))
			} else {
				color.Cyan("\tOuter Depth: " + color.RedString("No outer depth entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.OuterUnit.Value != "" {
				color.Cyan("\tOuter Unit: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", responseObject.OuterUnit.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", responseObject.OuterUnit.Label))
			} else {
				color.Cyan("\tOuter Unit: " + color.RedString("No outer unit entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.MountingDepth != 0 {
				color.Cyan("\tMounting Depth: " + color.YellowString("%d", responseObject.MountingDepth))
			} else {
				color.Cyan("\tMounting Depth: " + color.RedString("No mounting depth entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Comments != "" {
				color.Cyan("\tComments: " + color.YellowString("%s", responseObject.Comments))
			} else {
				color.Cyan("\tComments: " + color.RedString("No comments entry found for ") + color.YellowString("%s", responseObject.Display))
			}

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
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))

			if responseObject.DeviceCount != 0 {
				color.Cyan("\tDevice Count: " + color.YellowString("%d", responseObject.DeviceCount))
			} else {
				color.Cyan("\tDevice Count: " + color.RedString("No device count entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.PowerfeedCount != 0 {
				color.Cyan("\tPowerfeed Count: " + color.YellowString("%d\n", responseObject.PowerfeedCount))
			} else {
				color.Cyan("\tPowerfeed Count: " + color.RedString("No powerfeed count entry found for ") + color.YellowString("%s\n", responseObject.Display))
			}
		} else {
			color.Red("  Doh! No rack object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimRacksByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimRacksByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimRacksByIdCmd", err)
	}

	GetDcimRacksByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the rack object")
	err = GetDcimRacksByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimRacksByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimRacksByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

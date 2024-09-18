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
	"log"
	"strings"

	"github.com/spf13/cobra"
)

type sitesByID struct {
	CommonFieldsSlug
	Status struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"status"`
	Region struct {
		CommonFieldsSlug
		Depth uint `json:"_depth"`
	} `json:"region"`
	Group struct {
		CommonFieldsSlug
		Depth uint `json:"_depth"`
	} `json:"group"`
	Tenant struct {
		CommonFieldsSlug
	} `json:"tenant"`
	Facility        string  `json:"facility"`
	TimeZone        string  `json:"time_zone"`
	Description     string  `json:"description"`
	PhysicalAddress string  `json:"physical_address"`
	ShippingAddress string  `json:"shipping_address"`
	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
	Comments        string  `json:"comments"`
	Asns            []int   `json:"asns"`
	Tags            []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created             string `json:"created"`
	LastUpdated         string `json:"last_updated"`
	CircuitCount        uint   `json:"circuit_count"`
	DeviceCount         uint   `json:"device_count"`
	PrefixCount         uint   `json:"prefix_count"`
	RackCount           uint   `json:"rack_count"`
	VirtualmachineCount uint   `json:"virtualmachine_count"`
	VlanCount           uint   `json:"vlan_count"`
}

// GetDcimSitesByIDCmd represents the responseObjectByID command
var GetDcimSitesByIDCmd = &cobra.Command{
	Use:   "getDcimSitesByID",
	Short: "GET an site object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an site object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(sitesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.sites_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    Metropolis Site: %s", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: "+color.YellowString("%d"), responseObject.Id)
			color.Cyan("\tURL: "+color.YellowString("%s"), responseObject.Url)
			color.Cyan("\tDisplay: "+color.YellowString("%s"), responseObject.Display)
			color.Cyan("\tSlug: "+color.YellowString("%s"), responseObject.Slug)
			if responseObject.Status.Value != "" {
				color.Cyan("\tStatus: ")
				color.Cyan("\t  Status Value: "+color.YellowString("%s"), responseObject.Status.Value)
				color.Cyan("\t  Status Label: "+color.YellowString("%s"), responseObject.Status.Label)
			} else {
				color.Cyan("\tStatus: " + color.RedString("No status entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Region.Id != 0 {
				color.Cyan("\tRegion:")
				color.Cyan("\t  Region ID: "+color.YellowString("%d"), responseObject.Region.Id)
				color.Cyan("\t  Region URL: "+color.YellowString("%s"), responseObject.Region.Url)
				color.Cyan("\t  Region Display: "+color.YellowString("%s"), responseObject.Region.Display)
				color.Cyan("\t  Region Name: "+color.YellowString("%s"), responseObject.Region.Name)
				color.Cyan("\t  Region Slug: "+color.YellowString("%s"), responseObject.Region.Slug)
				if responseObject.Region.Depth != 0 {
					color.Cyan("\t  Region Depth: "+color.YellowString("%v"), responseObject.Region.Depth)
				} else {
					color.Cyan("\t  Region Depth: " + color.RedString("No region depth entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			} else {
				color.Cyan("\tRegion: " + color.RedString("No region entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Group.Id != 0 {
				color.Cyan("\tGroup:\n")
				color.Cyan("\t  Group ID: "+color.YellowString("%d"), responseObject.Group.Id)
				color.Cyan("\t  Group URL: "+color.YellowString("%s"), responseObject.Group.Url)
				color.Cyan("\t  Group Display: "+color.YellowString("%s"), responseObject.Group.Display)
				color.Cyan("\t  Group Name: "+color.YellowString("%s"), responseObject.Group.Name)
				color.Cyan("\t  Group Slug: "+color.YellowString("%s"), responseObject.Group.Slug)
				if responseObject.Group.Depth != 0 {
					color.Cyan("\t  Group Depth: "+color.YellowString("%v"), responseObject.Group.Depth)
				} else {
					color.Cyan("\t  Group Depth: " + color.RedString("No group depth entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			} else {
				color.Cyan("\tGroup: " + color.RedString("No group entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Tenant.Id != 0 {
				color.Cyan("\tTenant:\n")
				color.Cyan("\t  Tenant ID: "+color.YellowString("%d"), responseObject.Tenant.Id)
				color.Cyan("\t  Tenant URL: "+color.YellowString("%s"), responseObject.Tenant.Url)
				color.Cyan("\t  Tenant Display: "+color.YellowString("%s"), responseObject.Tenant.Display)
				color.Cyan("\t  Tenant Name: "+color.YellowString("%s"), responseObject.Tenant.Name)
				color.Cyan("\t  Tenant Slug: "+color.YellowString("%s"), responseObject.Tenant.Slug)
			} else {
				color.Cyan("\tTenant: " + color.RedString("No tenant entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Facility != "" {
				color.Cyan("\tFacility: "+color.YellowString("%s"), responseObject.Facility)
			} else {
				color.Cyan("\tFacility: " + color.RedString("No facility entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.TimeZone != "" {
				color.Cyan("\tTimezone: "+color.YellowString("%s"), responseObject.TimeZone)
			} else {
				color.Cyan("\tTimezone: " + color.RedString("No timezone entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: "+color.YellowString("\n\t%s"), responseObject.Description)
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.PhysicalAddress != "" {
				color.Cyan("\tPhysical Address: "+color.YellowString("%s"), responseObject.PhysicalAddress)
			} else {
				color.Cyan("\tPhysical Address: " + color.RedString("No physical address entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.ShippingAddress != "" {
				color.Cyan("\tShipping Address: "+color.YellowString("%s"), responseObject.ShippingAddress)
			} else {
				color.Cyan("\tShipping Address: " + color.RedString("No shipping address entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Latitude != 0 {
				color.Cyan("\tLatitude: "+color.YellowString("%v"), responseObject.Latitude)
			} else {
				color.Cyan("\tLatitude: " + color.RedString("No latitude entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Longitude != 0 {
				color.Cyan("\tLongitude: "+color.YellowString("%v"), responseObject.Longitude)
			} else {
				color.Cyan("\tLongitude: " + color.RedString("No longitude entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Comments != "" {
				color.Cyan("\tComments: "+color.YellowString("%s"), responseObject.Comments)
			} else {
				color.Cyan("\tComments: " + color.RedString("No longitude entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			for _, asn := range responseObject.Asns {
				if asn > 0 {
					color.Cyan("\tASN: "+color.YellowString("%v"), asn)
				} else {
					color.Cyan("\tASN: " + color.RedString("No asn entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			}
			for _, tag := range responseObject.Tags {
				if tag.Id > 0 {
					color.Cyan("\tTags:")
					color.Cyan("\t  Tag ID: "+color.YellowString("%v"), tag.Id)
					color.Cyan("\t  Tag URL: "+color.YellowString("%v"), tag.Url)
					color.Cyan("\t  Tag Display: "+color.YellowString("%v"), tag.Display)
					color.Cyan("\t  Tag Name: "+color.YellowString("%v"), tag.Name)
					color.Cyan("\t  Tag Slug: "+color.YellowString("%v"), tag.Slug)
					color.Cyan("\t  Tag Color: "+color.YellowString("%v"), tag.Color)
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			}
			color.Cyan("\tCreated: "+color.YellowString("%s"), responseObject.Created)
			color.Cyan("\tLast Updated: "+color.YellowString("%s"), responseObject.LastUpdated)
			if responseObject.CircuitCount != 0 {
				color.Cyan("\tCircuit Count: "+color.YellowString("%d"), responseObject.CircuitCount)
			} else {
				color.Cyan("\tCircuit Count: " + color.RedString("No circuit count entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.DeviceCount != 0 {
				color.Cyan("\tDevice Count: "+color.YellowString("%d"), responseObject.DeviceCount)
			} else {
				color.Cyan("\tDevice Count: " + color.RedString("No device count entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.PrefixCount != 0 {
				color.Cyan("\tPrefix Count: "+color.YellowString("%d"), responseObject.PrefixCount)
			} else {
				color.Cyan("\tPrefix Count: " + color.RedString("No prefix count entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.RackCount != 0 {
				color.Cyan("\tRack Count: "+color.YellowString("%d"), responseObject.RackCount)
			} else {
				color.Cyan("\tRack Count: " + color.RedString("No rack count entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.VirtualmachineCount != 0 {
				color.Cyan("\tVirtual Machine Count: "+color.YellowString("%d"), responseObject.VirtualmachineCount)
			} else {
				color.Cyan("\tVirtual Machine Count: " + color.RedString("No virtual machine count entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.VlanCount != 0 {
				color.Cyan("\tVLAN Count: "+color.YellowString("%d\n"), responseObject.VlanCount)
			} else {
				color.Cyan("\tVLAN Count: " + color.RedString("No vlan count entry found for ") + color.YellowString("%s\n", responseObject.Display))
			}
		} else {
			color.Red("  Doh! No site object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

// init is a function that is automatically executed at program startup.
// It initializes the 'responseObjectByIDCmd' command by adding it as a subcommand to 'responseObjectByIDCmd'.
// It also defines flags and configuration settings for the command.
func init() {

	// Here you will define your flags and configuration settings.
	GetDcimSitesByIDCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimSitesByIDCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimSitesByIDCmd", err)
	}

	GetDcimSitesByIDCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the site object")
	err = GetDcimSitesByIDCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// responseObjectByIDCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// responseObjectByIDCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

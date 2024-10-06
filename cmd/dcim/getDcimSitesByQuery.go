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

// GetDcimSitesByQueryCmd represents the siteByQuery command
var GetDcimSitesByQueryCmd = &cobra.Command{
	Use:   "getDcimSitesByQuery",
	Short: "GET a site object by string query",
	Long: `
ABC Netbox Automation Tools:
  GET a site object by string query`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(sites)
		ApiConnectionQuery(responseObject, "GET", "cmd.dcim.dcim_api_url.sites_id")

		for _, site := range responseObject.Results {
			display := fmt.Sprintf("    ABC Site: %s", color.YellowString(site.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: "+color.YellowString("%d"), site.Id)
			color.Cyan("\tURL: "+color.YellowString("%s"), site.Url)
			color.Cyan("\tDisplay: "+color.YellowString("%s"), site.Display)
			color.Cyan("\tSlug: "+color.YellowString("%s"), site.Slug)
			if site.Status.Value != "" {
				color.Cyan("\tStatus: ")
				color.Cyan("\t  Status Value: "+color.YellowString("%s"), site.Status.Value)
				color.Cyan("\t  Status Label: "+color.YellowString("%s"), site.Status.Label)
			} else {
				color.Cyan("\tStatus: " + color.RedString("No status entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.Region.Id != 0 {
				color.Cyan("\tRegion:")
				color.Cyan("\t  Region ID: "+color.YellowString("%d"), site.Region.Id)
				color.Cyan("\t  Region URL: "+color.YellowString("%s"), site.Region.Url)
				color.Cyan("\t  Region Display: "+color.YellowString("%s"), site.Region.Display)
				color.Cyan("\t  Region Name: "+color.YellowString("%s"), site.Region.Name)
				color.Cyan("\t  Region Slug: "+color.YellowString("%s"), site.Region.Slug)
				if site.Region.Depth != 0 {
					color.Cyan("\t  Region Depth: "+color.YellowString("%v"), site.Region.Depth)
				} else {
					color.Cyan("\t  Region Depth: " + color.RedString("No region depth entry found for ") + color.YellowString("%s", site.Display))
				}
			} else {
				color.Cyan("\tRegion: " + color.RedString("No region entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.Group.Id != 0 {
				color.Cyan("\tGroup:\n")
				color.Cyan("\t  Group ID: "+color.YellowString("%d"), site.Group.Id)
				color.Cyan("\t  Group URL: "+color.YellowString("%s"), site.Group.Url)
				color.Cyan("\t  Group Display: "+color.YellowString("%s"), site.Group.Display)
				color.Cyan("\t  Group Name: "+color.YellowString("%s"), site.Group.Name)
				color.Cyan("\t  Group Slug: "+color.YellowString("%s"), site.Group.Slug)
				if site.Group.Depth != 0 {
					color.Cyan("\t  Group Depth: "+color.YellowString("%v"), site.Group.Depth)
				} else {
					color.Cyan("\t  Group Depth: " + color.RedString("No group depth entry found for ") + color.YellowString("%s", site.Display))
				}
			} else {
				color.Cyan("\tGroup: " + color.RedString("No group entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.Tenant.Id != 0 {
				color.Cyan("\tTenant:\n")
				color.Cyan("\t  Tenant ID: "+color.YellowString("%d"), site.Tenant.Id)
				color.Cyan("\t  Tenant URL: "+color.YellowString("%s"), site.Tenant.Url)
				color.Cyan("\t  Tenant Display: "+color.YellowString("%s"), site.Tenant.Display)
				color.Cyan("\t  Tenant Name: "+color.YellowString("%s"), site.Tenant.Name)
				color.Cyan("\t  Tenant Slug: "+color.YellowString("%s"), site.Tenant.Slug)
			} else {
				color.Cyan("\tTenant: " + color.RedString("No tenant entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.Facility != "" {
				color.Cyan("\tFacility: "+color.YellowString("%s"), site.Facility)
			} else {
				color.Cyan("\tFacility: " + color.RedString("No facility entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.TimeZone != "" {
				color.Cyan("\tTimezone: "+color.YellowString("%s"), site.TimeZone)
			} else {
				color.Cyan("\tTimezone: " + color.RedString("No timezone entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.Description != "" {
				color.Cyan("\tDescription: "+color.YellowString("\n\t%s"), site.Description)
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.PhysicalAddress != "" {
				color.Cyan("\tPhysical Address: "+color.YellowString("%s"), site.PhysicalAddress)
			} else {
				color.Cyan("\tPhysical Address: " + color.RedString("No physical address entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.ShippingAddress != "" {
				color.Cyan("\tShipping Address: "+color.YellowString("%s"), site.ShippingAddress)
			} else {
				color.Cyan("\tShipping Address: " + color.RedString("No shipping address entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.Latitude != 0 {
				color.Cyan("\tLatitude: "+color.YellowString("%v"), site.Latitude)
			} else {
				color.Cyan("\tLatitude: " + color.RedString("No latitude entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.Longitude != 0 {
				color.Cyan("\tLongitude: "+color.YellowString("%v"), site.Longitude)
			} else {
				color.Cyan("\tLongitude: " + color.RedString("No longitude entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.Comments != "" {
				color.Cyan("\tComments: "+color.YellowString("%s"), site.Comments)
			} else {
				color.Cyan("\tComments: " + color.RedString("No longitude entry found for ") + color.YellowString("%s", site.Display))
			}
			for _, asn := range site.Asns {
				if asn > 0 {
					color.Cyan("\tASN: "+color.YellowString("%v"), asn)
				} else {
					color.Cyan("\tASN: " + color.RedString("No asn entry found for ") + color.YellowString("%s", site.Display))
				}
			}
			for _, tag := range site.Tags {
				if tag.Id > 0 {
					color.Cyan("\tTags:")
					color.Cyan("\t  Tag ID: "+color.YellowString("%v"), tag.Id)
					color.Cyan("\t  Tag URL: "+color.YellowString("%v"), tag.Url)
					color.Cyan("\t  Tag Display: "+color.YellowString("%v"), tag.Display)
					color.Cyan("\t  Tag Name: "+color.YellowString("%v"), tag.Name)
					color.Cyan("\t  Tag Slug: "+color.YellowString("%v"), tag.Slug)
					color.Cyan("\t  Tag Color: "+color.YellowString("%v"), tag.Color)
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s", site.Display))
				}
			}
			color.Cyan("\tCreated: "+color.YellowString("%s"), site.Created)
			color.Cyan("\tLast Updated: "+color.YellowString("%s"), site.LastUpdated)
			if site.CircuitCount != 0 {
				color.Cyan("\tCircuit Count: "+color.YellowString("%d"), site.CircuitCount)
			} else {
				color.Cyan("\tCircuit Count: " + color.RedString("No circuit count entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.DeviceCount != 0 {
				color.Cyan("\tDevice Count: "+color.YellowString("%d"), site.DeviceCount)
			} else {
				color.Cyan("\tDevice Count: " + color.RedString("No device count entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.PrefixCount != 0 {
				color.Cyan("\tPrefix Count: "+color.YellowString("%d"), site.PrefixCount)
			} else {
				color.Cyan("\tPrefix Count: " + color.RedString("No prefix count entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.RackCount != 0 {
				color.Cyan("\tRack Count: "+color.YellowString("%d"), site.RackCount)
			} else {
				color.Cyan("\tRack Count: " + color.RedString("No rack count entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.VirtualmachineCount != 0 {
				color.Cyan("\tVirtual Machine Count: "+color.YellowString("%d"), site.VirtualmachineCount)
			} else {
				color.Cyan("\tVirtual Machine Count: " + color.RedString("No virtual machine count entry found for ") + color.YellowString("%s", site.Display))
			}
			if site.VlanCount != 0 {
				color.Cyan("\tVLAN Count: "+color.YellowString("%d\n"), site.VlanCount)
			} else {
				color.Cyan("\tVLAN Count: " + color.RedString("No vlan count entry found for ") + color.YellowString("%s\n", site.Display))
			}
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimSitesByQueryCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimSitesByQueryCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for GetDcimSitesByQueryCmd", err)
	}

	GetDcimSitesByQueryCmd.Flags().StringVarP(&query, "query", "q", "", "string of object search you want to get")
	err = GetDcimSitesByQueryCmd.MarkFlagRequired("query")
	if err != nil {
		log.Fatalf("Error marking query flag as required: %s - for GetDcimSitesByQueryCmd", err)
	}
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// siteByQueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// siteByQueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

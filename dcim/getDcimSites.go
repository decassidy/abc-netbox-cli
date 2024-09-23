/*
Copyright © 2024 Derrick Cassidy - .Technologies, Inc.

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
	"bufio"
	"fmt"
	_ "fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	_ "github.com/spf13/viper"
	_ "io"
	"log"
	_ "net/http"
	"os"
	_ "strconv"
	"strings"
)

type sites struct {
	Count    uint   `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		CommonFieldsSlug
		Status struct {
			ValueLabel
		} `json:"status"`
		Region struct {
			CommonFieldsSlug
			Description string `json:"description"`
			SiteCount   uint   `json:"site_count"`
			Depth       uint   `json:"_depth"`
		} `json:"region"`
		Group struct {
			CommonFieldsSlug
			Depth float32 `json:"_depth"`
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
		Asns            []uint  `json:"asns"`
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
	} `json:"results"`
}

var responseObjectSites = new(sites)

// GetDcimSitesCmd represents the allNetboxSites command
var GetDcimSitesCmd = &cobra.Command{
	Use:   "getDcimSites",
	Short: "GET a list of site objects",
	Long: `
.Netbox Automation Tools:
  GET a list of site objects`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectSites, "GET", "cmd.dcim.dcim_api_url.sites")

		if responseObjectSites.Count > 0 {
			color.Cyan("\n  Total ABC Sites: "+color.YellowString("%d"), responseObjectSites.Count)
			for _, site := range responseObjectSites.Results {
				display := fmt.Sprintf("    ABC Site: %s", color.YellowString(site.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: "+color.YellowString("%d"), site.Id)
				color.Cyan("\tURL: "+color.YellowString("%s"), site.Url)
				color.Cyan("\tDisplay: "+color.YellowString("%s"), site.Display)
				color.Cyan("\tCable: "+color.YellowString("%s"), site.Slug)
				if site.Status.Value != "" {
					color.Cyan("\tStatus: ")
					color.Cyan("\t  Value: "+color.YellowString("%s"), site.Status.Value)
					color.Cyan("\t  Label: "+color.YellowString("%s"), site.Status.Label)
				} else {
					color.Cyan("\tStatus: " + color.RedString("No status entry found for ") + color.YellowString("%s", site.Display))
				}
				if site.Region.Id != 0 {
					color.Cyan("\tRegion:")
					color.Cyan("\t  ID: "+color.YellowString("%d"), site.Region.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), site.Region.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), site.Region.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), site.Region.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%s"), site.Region.Slug)
					if site.Region.Depth != 0 {
						color.Cyan("\t  Depth: "+color.YellowString("%v"), site.Region.Depth)
					} else {
						color.Cyan("\t  Depth: " + color.RedString("No region depth entry found for ") + color.YellowString("%s", site.Display))
					}
				} else {
					color.Cyan("\tRegion: " + color.RedString("No region entry found for ") + color.YellowString("%s", site.Display))
				}
				if site.Group.Id != 0 {
					color.Cyan("\tGroup:\n")
					color.Cyan("\t  ID: "+color.YellowString("%d"), site.Group.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), site.Group.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), site.Group.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), site.Group.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%s"), site.Group.Slug)
					if site.Group.Depth != 0 {
						color.Cyan("\t  Depth: "+color.YellowString("%v"), site.Group.Depth)
					} else {
						color.Cyan("\t  Depth: " + color.RedString("No group depth entry found for ") + color.YellowString("%s", site.Display))
					}
				} else {
					color.Cyan("\tGroup: " + color.RedString("No group entry found for ") + color.YellowString("%s", site.Display))
				}
				if site.Tenant.Id != 0 {
					color.Cyan("\tTenant:\n")
					color.Cyan("\t  ID: "+color.YellowString("%d"), site.Tenant.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), site.Tenant.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), site.Tenant.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), site.Tenant.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%s"), site.Tenant.Slug)
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
						color.Cyan("\t  ID: "+color.YellowString("%v"), tag.Id)
						color.Cyan("\t  URL: "+color.YellowString("%v"), tag.Url)
						color.Cyan("\t  Display: "+color.YellowString("%v"), tag.Display)
						color.Cyan("\t  Name: "+color.YellowString("%v"), tag.Name)
						color.Cyan("\t  Slug: "+color.YellowString("%v"), tag.Slug)
						color.Cyan("\t  Color: "+color.YellowString("%v"), tag.Color)
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
			for responseObjectSites.Next != "" {
				nextPageSites()
			}
			if responseObjectSites.Next == "" {
				display := color.HiGreenString("\tAll Netbox site objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n")
			}

		} else {
			color.Cyan("  ABC Sites: " + color.RedString("No sites found on server. Exiting...\n"))
		}
	},
}

func displaySitesOutput() {
	for _, site := range responseObjectSites.Results {
		display := fmt.Sprintf("    ABC Site: %s", color.YellowString(site.Display))
		equals := strings.Repeat("=", len(display))
		color.Cyan("\n  " + equals + "\n")
		color.Cyan(display)
		color.Cyan("  " + equals + "\n")
		color.Cyan("\tID: "+color.YellowString("%d"), site.Id)
		color.Cyan("\tURL: "+color.YellowString("%s"), site.Url)
		color.Cyan("\tDisplay: "+color.YellowString("%s"), site.Display)
		color.Cyan("\tCable: "+color.YellowString("%s"), site.Slug)
		if site.Status.Value != "" {
			color.Cyan("\tStatus: ")
			color.Cyan("\t  Value: "+color.YellowString("%s"), site.Status.Value)
			color.Cyan("\t  Label: "+color.YellowString("%s"), site.Status.Label)
		} else {
			color.Cyan("\tStatus: " + color.RedString("No status entry found for ") + color.YellowString("%s", site.Display))
		}
		if site.Region.Id != 0 {
			color.Cyan("\tRegion:")
			color.Cyan("\t  ID: "+color.YellowString("%d"), site.Region.Id)
			color.Cyan("\t  URL: "+color.YellowString("%s"), site.Region.Url)
			color.Cyan("\t  Display: "+color.YellowString("%s"), site.Region.Display)
			color.Cyan("\t  Name: "+color.YellowString("%s"), site.Region.Name)
			color.Cyan("\t  Slug: "+color.YellowString("%s"), site.Region.Slug)
			if site.Region.Depth != 0 {
				color.Cyan("\t  Depth: "+color.YellowString("%v"), site.Region.Depth)
			} else {
				color.Cyan("\t  Depth: " + color.RedString("No region depth entry found for ") + color.YellowString("%s", site.Display))
			}
		} else {
			color.Cyan("\tRegion: " + color.RedString("No region entry found for ") + color.YellowString("%s", site.Display))
		}
		if site.Group.Id != 0 {
			color.Cyan("\tGroup:\n")
			color.Cyan("\t  ID: "+color.YellowString("%d"), site.Group.Id)
			color.Cyan("\t  URL: "+color.YellowString("%s"), site.Group.Url)
			color.Cyan("\t  Display: "+color.YellowString("%s"), site.Group.Display)
			color.Cyan("\t  Name: "+color.YellowString("%s"), site.Group.Name)
			color.Cyan("\t  Slug: "+color.YellowString("%s"), site.Group.Slug)
			if site.Group.Depth != 0 {
				color.Cyan("\t  Depth: "+color.YellowString("%v"), site.Group.Depth)
			} else {
				color.Cyan("\t  Depth: " + color.RedString("No group depth entry found for ") + color.YellowString("%s", site.Display))
			}
		} else {
			color.Cyan("\tGroup: " + color.RedString("No group entry found for ") + color.YellowString("%s", site.Display))
		}
		if site.Tenant.Id != 0 {
			color.Cyan("\tTenant:\n")
			color.Cyan("\t  ID: "+color.YellowString("%d"), site.Tenant.Id)
			color.Cyan("\t  URL: "+color.YellowString("%s"), site.Tenant.Url)
			color.Cyan("\t  Display: "+color.YellowString("%s"), site.Tenant.Display)
			color.Cyan("\t  Name: "+color.YellowString("%s"), site.Tenant.Name)
			color.Cyan("\t  Slug: "+color.YellowString("%s"), site.Tenant.Slug)
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
				color.Cyan("\t  ID: "+color.YellowString("%v"), tag.Id)
				color.Cyan("\t  URL: "+color.YellowString("%v"), tag.Url)
				color.Cyan("\t  Display: "+color.YellowString("%v"), tag.Display)
				color.Cyan("\t  Name: "+color.YellowString("%v"), tag.Name)
				color.Cyan("\t  Slug: "+color.YellowString("%v"), tag.Slug)
				color.Cyan("\t  Color: "+color.YellowString("%v"), tag.Color)
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
}

func nextPageSites() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\tDo you want to continue to the next page of site objects? [yes/no]: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageSites(responseObjectSites, "GET", responseObjectSites.Next)
		displaySitesOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("\tInvalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] \n")
	}
}

func ApiConnectionNextPageSites[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := responseObjectSites.Next

	color.Yellow("\n  Getting Netbox API objects from %s\n", suffix)
	token := config.GetString("cmd.token_key")

	err := CheckSSL(rootURL)
	if err != nil {
		fmt.Println("  SSL certificate is not valid: ", err)
	} else {
		color.Cyan("  SSL certificate is valid for: " + color.YellowString("%s", rootURL))
	}

	executeAPIRequest(httpMethod, fullAPIPath, token, r)
	if err != nil {
		log.Fatalf("Error getting Netbox API objects: %s\n", err)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimSitesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimSitesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for GetDcimSitesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allNetboxSitesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allNetboxSitesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

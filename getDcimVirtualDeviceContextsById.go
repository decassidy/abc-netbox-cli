/*
Copyright © 2024 Derrick Cassidy - Metropolis Technologies, Inc.
*/

package dcim

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type virtualDeviceContextsByID struct {
	CommonFieldsNoSlug
	Device struct {
		CommonFieldsNoSlug
	} `json:"device"`
	Identifier int `json:"identifier"`
	Tenant     struct {
		CommonFieldsSlug
	} `json:"tenant"`
	PrimaryIp struct {
		Id      uint32 `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Family  int    `json:"family"`
		Address string `json:"address"`
	} `json:"primary_ip"`
	PrimaryIp4 struct {
		Id      uint32 `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Family  int    `json:"family"`
		Address string `json:"address"`
	} `json:"primary_ip4"`
	PrimaryIp6 struct {
		Id      uint32 `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Family  uint32 `json:"family"`
		Address string `json:"address"`
	} `json:"primary_ip6"`
	Status struct {
		ValueLabel
	} `json:"status"`
	Description string `json:"description"`
	Comments    string `json:"comments"`
	Tags        []struct {
		CommonFieldsNoSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created        string `json:"created"`
	LastUpdated    string `json:"last_updated"`
	InterfaceCount uint32 `json:"interface_count"`
}

// GetDcimVirtualDeviceContextsByIdCmd represents the getDcimVirtualDeviceContextsById command
var GetDcimVirtualDeviceContextsByIdCmd = &cobra.Command{
	Use:   "getDcimVirtualDeviceContextsById",
	Short: "GET an device context object by ID",
	Long: `
Netbox Automation Tools:
  GET an device context object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(virtualDeviceContextsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.virtual_device_contexts_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    Metropolis Virtual Device Context: %s", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: "+color.YellowString("%d"), responseObject.Id)
			color.Cyan("\tURL: "+color.YellowString("%s"), responseObject.Url)
			color.Cyan("\tDisplay: "+color.YellowString("%s"), responseObject.Display)
			color.Cyan("\tName: "+color.YellowString("%s"), responseObject.Name)
			if responseObject.Identifier > 0 {
				color.Cyan("\tIdentifier: "+color.YellowString("%d"), responseObject.Identifier)
			} else {
				color.Cyan("\tIdentifier: " + color.RedString("No identifier entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Tenant.Id > 0 {
				color.Cyan("\tTenant: ")
				color.Cyan("\t  ID: "+color.YellowString("%d"), responseObject.Tenant.Id)
				color.Cyan("\t  URL: "+color.YellowString("%s"), responseObject.Tenant.Url)
				color.Cyan("\t  Display: "+color.YellowString("%s"), responseObject.Tenant.Display)
				color.Cyan("\t  Name: "+color.YellowString("%s"), responseObject.Tenant.Name)
				color.Cyan("\t  Slug: "+color.YellowString("%s"), responseObject.Tenant.Slug)
			} else {
				color.Cyan("\tTenant: " + color.RedString("No tenant entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.PrimaryIp.Id > 0 {
				color.Cyan("\tPrimary IP: ")
				color.Cyan("\t  ID: "+color.YellowString("%d"), responseObject.PrimaryIp.Id)
				color.Cyan("\t  URL: "+color.YellowString("%s"), responseObject.PrimaryIp.Url)
				color.Cyan("\t  Display: "+color.YellowString("%s"), responseObject.PrimaryIp.Display)
				color.Cyan("\t  Family: "+color.YellowString("%d"), responseObject.PrimaryIp.Family)
				color.Cyan("\t  Address: "+color.YellowString("%s"), responseObject.PrimaryIp.Address)
			} else {
				color.Cyan("\tPrimary IP: " + color.RedString("No primary IP entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.PrimaryIp4.Id > 0 {
				color.Cyan("\tPrimary IP: ")
				color.Cyan("\t  ID: "+color.YellowString("%d"), responseObject.PrimaryIp4.Id)
				color.Cyan("\t  URL: "+color.YellowString("%s"), responseObject.PrimaryIp4.Url)
				color.Cyan("\t  Display: "+color.YellowString("%s"), responseObject.PrimaryIp4.Display)
				color.Cyan("\t  Family: "+color.YellowString("%d"), responseObject.PrimaryIp4.Family)
				color.Cyan("\t  Address: "+color.YellowString("%s"), responseObject.PrimaryIp4.Address)
			} else {
				color.Cyan("\tPrimary IPv4: " + color.RedString("No primary IPv4 entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.PrimaryIp6.Id > 0 {
				color.Cyan("\tPrimary IP: ")
				color.Cyan("\t  ID: "+color.YellowString("%d"), responseObject.PrimaryIp6.Id)
				color.Cyan("\t  URL: "+color.YellowString("%s"), responseObject.PrimaryIp6.Url)
				color.Cyan("\t  Display: "+color.YellowString("%s"), responseObject.PrimaryIp6.Display)
				color.Cyan("\t  Family: "+color.YellowString("%d"), responseObject.PrimaryIp6.Family)
				color.Cyan("\t  Address: "+color.YellowString("%s"), responseObject.PrimaryIp6.Address)
			} else {
				color.Cyan("\tPrimary IPv6: " + color.RedString("No primary IPv6 entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Status.Value != "" {
				color.Cyan("\tStatus: ")
				color.Cyan("\t  Value: "+color.YellowString("%s"), responseObject.Status.Value)
				color.Cyan("\t  Label: "+color.YellowString("%s"), responseObject.Status.Label)
			} else {
				color.Cyan("\tStatus: " + color.RedString("No status entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: "+color.YellowString("%s"), responseObject.Description)
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Comments != "" {
				color.Cyan("\tComments: "+color.YellowString("%s"), responseObject.Comments)
			} else {
				color.Cyan("\tComments: " + color.RedString("No comments entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			for _, tag := range responseObject.Tags {
				if tag.Id != 0 {
					color.Cyan("\tTags:")
					color.Cyan("\t  Tag ID: "+color.YellowString("%v"), tag.Id)
					color.Cyan("\t  Tag URL: "+color.YellowString("%v"), tag.Url)
					color.Cyan("\t  Tag Display: "+color.YellowString("%v"), tag.Display)
					color.Cyan("\t  Tag Name: "+color.YellowString("%v"), tag.Name)
					color.Cyan("\t  Tag Color: "+color.YellowString("%v"), tag.Color)
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			}

			color.Cyan("\tCreated: "+color.YellowString("%s"), responseObject.Created)
			color.Cyan("\tLast Updated: "+color.YellowString("%s"), responseObject.LastUpdated)

			if responseObject.InterfaceCount != 0 {
				color.Cyan("\tInterface Count: "+color.YellowString("%d"), responseObject.InterfaceCount)
			} else {
				color.Cyan("\tInterface Count: " + color.RedString("No interface count entry found for ") + color.YellowString("%s", responseObject.Display))
			}
		} else {
			color.Red("  Doh! No virtual device context object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimVirtualDeviceContextsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimVirtualDeviceContextsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimVirtualDeviceContextsByIdCmd", err)
	}

	GetDcimVirtualDeviceContextsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of virtual device context object")
	err = GetDcimVirtualDeviceContextsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimVirtualDeviceContextsByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimVirtualDeviceContextsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimVirtualDeviceContextsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

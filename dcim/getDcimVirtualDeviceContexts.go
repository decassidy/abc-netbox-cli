/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type virtualDeviceContexts struct {
	Count    uint32 `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
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
	} `json:"results"`
}

// GetDcimVirtualDeviceContextsCmd represents the getDcimVirtualDeviceContexts command
var GetDcimVirtualDeviceContextsCmd = &cobra.Command{
	Use:   "getDcimVirtualDeviceContexts",
	Short: "GET a list of device context objects",
	Long: `
Netbox Automation Tools:
  GET a list of device context objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(virtualDeviceContexts)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.virtual_device_contexts")

		if responseObject.Count > 0 {
			color.Cyan("\n  Total ABC Virtual Device Contexts: "+color.YellowString("%d"), responseObject.Count)

			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    ABC Virtual Device Context: %s", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: "+color.YellowString("%d"), result.Id)
				color.Cyan("\tURL: "+color.YellowString("%s"), result.Url)
				color.Cyan("\tDisplay: "+color.YellowString("%s"), result.Display)
				color.Cyan("\tName: "+color.YellowString("%s"), result.Name)
				if result.Device.Id > 0 {
					color.Cyan("\tDevice: ")
					color.Cyan("\t  ID: "+color.YellowString("%s"), result.Device.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), result.Device.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), result.Device.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), result.Device.Name)
				} else {
					color.Cyan("\tDevice: " + color.RedString("No device entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Identifier > 0 {
					color.Cyan("\tIdentifier: "+color.YellowString("%d"), result.Identifier)
				} else {
					color.Cyan("\tIdentifier: " + color.RedString("No identifier entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.PrimaryIp.Id > 0 {
					color.Cyan("\tPrimary IP: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), result.PrimaryIp.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), result.PrimaryIp.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), result.PrimaryIp.Display)
					color.Cyan("\t  Family: "+color.YellowString("%d"), result.PrimaryIp.Family)
					color.Cyan("\t  Address: "+color.YellowString("%s"), result.PrimaryIp.Address)
				} else {
					color.Cyan("\tPrimary IP: " + color.RedString("No primary IP entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.PrimaryIp4.Id > 0 {
					color.Cyan("\tPrimary IPv4: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), result.PrimaryIp4.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), result.PrimaryIp4.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), result.PrimaryIp4.Display)
					color.Cyan("\t  Family: "+color.YellowString("%d"), result.PrimaryIp4.Family)
					color.Cyan("\t  Address: "+color.YellowString("%s"), result.PrimaryIp4.Address)
				} else {
					color.Cyan("\tPrimary IPv4: " + color.RedString("No primary IPv4 entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.PrimaryIp6.Id > 0 {
					color.Cyan("\tPrimary IPv6: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), result.PrimaryIp6.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), result.PrimaryIp6.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), result.PrimaryIp6.Display)
					color.Cyan("\t  Family: "+color.YellowString("%d"), result.PrimaryIp6.Family)
					color.Cyan("\t  Address: "+color.YellowString("%s"), result.PrimaryIp6.Address)
				} else {
					color.Cyan("\tPrimary IPv6: " + color.RedString("No primary IPv6 entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Status.Value != "" {
					color.Cyan("\tStatus: ")
					color.Cyan("\t  Value: "+color.YellowString("%s"), result.Status.Value)
					color.Cyan("\t  Label: "+color.YellowString("%s"), result.Status.Label)
				} else {
					color.Cyan("\tStatus: " + color.RedString("No Status entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: "+color.YellowString("%s"), result.Description)
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tComments: "+color.YellowString("%s"), result.Comments)
				for _, tag := range result.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTags:")
						color.Cyan("\t  Tag ID: "+color.YellowString("%v"), tag.Id)
						color.Cyan("\t  Tag URL: "+color.YellowString("%v"), tag.Url)
						color.Cyan("\t  Tag Display: "+color.YellowString("%v"), tag.Display)
						color.Cyan("\t  Tag Name: "+color.YellowString("%v"), tag.Name)
						color.Cyan("\t  Tag Color: "+color.YellowString("%v"), tag.Color)
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s", result.Display))
					}
				}
				color.Cyan("\tCreated: "+color.YellowString("%s"), result.Created)
				color.Cyan("\tLast Updated: "+color.YellowString("%s"), result.LastUpdated)
				if result.InterfaceCount > 0 {
					color.Cyan("\tInterface Count: "+color.YellowString("%d\n"), result.InterfaceCount)
				} else {
					color.Cyan("\tInterface Count: " + color.RedString("No interface count entry found for ") + color.YellowString("%s\n", result.Display))
				}
			}
		} else {
			color.Cyan("  ABC Virtual Device Contexts: " + color.RedString("No virtual device context found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimVirtualDeviceContextsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimVirtualDeviceContextsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimVirtualDeviceContextsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimVirtualDeviceContextsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimVirtualDeviceContextsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

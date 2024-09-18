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

type virtualChassisByID struct {
	CommonFieldsNoSlug
	Domain string `json:"domain"`
	Master struct {
		CommonFieldsNoSlug
	} `json:"master"`
	Description string `json:"description"`
	Comments    string `json:"comments"`
	Tags        []struct {
		CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
	MemberCount int    `json:"member_count"`
}

// GetDcimVirtualChassisByIdCmd represents the getDcimVirtualChassisById command
var GetDcimVirtualChassisByIdCmd = &cobra.Command{
	Use:   "getDcimVirtualChassisById",
	Short: "GET an virtual chassis object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an virtual chassis object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(virtualChassisByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.virtual_chassis_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    Metropolis Virtual Chassis: %s", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: "+color.YellowString("%d"), responseObject.Id)
			color.Cyan("\tURL: "+color.YellowString("%s"), responseObject.Url)
			color.Cyan("\tDisplay: "+color.YellowString("%s"), responseObject.Display)
			color.Cyan("\tName: "+color.YellowString("%s"), responseObject.Name)
			color.Cyan("\tDomain: "+color.YellowString("%s"), responseObject.Domain)

			if responseObject.Master.Id > 0 {
				color.Cyan("\tMaster: ")
				color.Cyan("\t  ID: "+color.YellowString("%s"), responseObject.Master.Id)
				color.Cyan("\t  URL: "+color.YellowString("%s"), responseObject.Master.Url)
				color.Cyan("\t  Display: "+color.YellowString("%s"), responseObject.Master.Display)
				color.Cyan("\t  Name: "+color.YellowString("%s"), responseObject.Master.Name)
			} else {
				color.Cyan("\tMaster: " + color.RedString("No master entry found for ") + color.YellowString("%s", responseObject.Display))
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
					color.Cyan("\t  Tag Slug: "+color.YellowString("%v"), tag.Slug)
					color.Cyan("\t  Tag Color: "+color.YellowString("%v"), tag.Color)
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s", responseObject.Display))
				}
			}
			color.Cyan("\tCreated: "+color.YellowString("%s"), responseObject.Created)
			color.Cyan("\tLast Updated: "+color.YellowString("%s"), responseObject.LastUpdated)

			if responseObject.MemberCount > 0 {
				color.Cyan("\tDevice Count: "+color.YellowString("%d\n"), responseObject.MemberCount)
			} else {
				color.Cyan("\tDevice Count: " + color.RedString("No device count entry found for ") + color.YellowString("%s\n", responseObject.Display))
			}
		} else {
			color.Red("  Doh! No virtual chassis object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimVirtualChassisByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimVirtualChassisByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimVirtualChassisByIdCmd", err)
	}

	GetDcimVirtualChassisByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of virtual chassis to retrieve")
	err = GetDcimVirtualChassisByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimVirtualChassByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimVirtualChassisByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimVirtualChassisByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

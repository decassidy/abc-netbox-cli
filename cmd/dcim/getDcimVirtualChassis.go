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

type virtualChassis struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
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
	} `json:"results"`
}

// GetDcimVirtualChassisCmd represents the getDcimVirtualChassis command
var GetDcimVirtualChassisCmd = &cobra.Command{
	Use:   "getDcimVirtualChassis",
	Short: "GET a list of virtual chassis objects",
	Long: `
Netbox Automation Tools:
  GET a list of virtual chassis objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(virtualChassis)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.virtual_chassis")

		if responseObject.Count > 0 {
			color.Cyan("\n  Total ABC Virtual Chassis: "+color.YellowString("%d"), responseObject.Count)

			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    ABC Virtual Chassis: %s", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: "+color.YellowString("%d"), result.Id)
				color.Cyan("\tURL: "+color.YellowString("%s"), result.Url)
				color.Cyan("\tDisplay: "+color.YellowString("%s"), result.Display)
				color.Cyan("\tName: "+color.YellowString("%s"), result.Name)
				color.Cyan("\tDomain: "+color.YellowString("%s"), result.Domain)

				if result.Master.Id > 0 {
					color.Cyan("\tMaster: ")
					color.Cyan("\t  ID: "+color.YellowString("%s"), result.Master.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), result.Master.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), result.Master.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), result.Master.Name)
				} else {
					color.Cyan("\tMaster: " + color.RedString("No master entry found for ") + color.YellowString("%s", result.Display))
				}

				if result.Description != "" {
					color.Cyan("\tDescription: "+color.YellowString("%s"), result.Description)
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", result.Display))
				}

				if result.Comments != "" {
					color.Cyan("\tComments: "+color.YellowString("%s"), result.Comments)
				} else {
					color.Cyan("\tComments: " + color.RedString("No comments entry found for ") + color.YellowString("%s", result.Display))
				}

				for _, tag := range result.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTags:")
						color.Cyan("\t  Tag ID: "+color.YellowString("%v"), tag.Id)
						color.Cyan("\t  Tag URL: "+color.YellowString("%v"), tag.Url)
						color.Cyan("\t  Tag Display: "+color.YellowString("%v"), tag.Display)
						color.Cyan("\t  Tag Name: "+color.YellowString("%v"), tag.Name)
						color.Cyan("\t  Tag Slug: "+color.YellowString("%v"), tag.Slug)
						color.Cyan("\t  Tag Color: "+color.YellowString("%v"), tag.Color)
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s", result.Display))
					}
				}
				color.Cyan("\tCreated: "+color.YellowString("%s"), result.Created)
				color.Cyan("\tLast Updated: "+color.YellowString("%s"), result.LastUpdated)

				if result.MemberCount > 0 {
					color.Cyan("\tDevice Count: "+color.YellowString("%d\n"), result.MemberCount)
				} else {
					color.Cyan("\tDevice Count: " + color.RedString("No device count entry found for ") + color.YellowString("%s\n", result.Display))
				}
			}
		} else {
			color.Cyan("  ABC Virtual Chassis: " + color.RedString("No virtual chassis found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimVirtualChassisCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimVirtualChassisCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking dcim virtual chassis flag as required: %s - for GetDcimVirtualChassisCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimVirtualChassisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimVirtualChassisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

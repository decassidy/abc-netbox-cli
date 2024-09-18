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

type siteGroups struct {
	Count    uint   `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
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
		SiteCount   uint   `json:"site_count"`
		Depth       uint   `json:"_depth"`
	} `json:"results"`
}

// GetDcimSiteGroupsCmd represents the getDcimSiteGroups command
var GetDcimSiteGroupsCmd = &cobra.Command{
	Use:   "getDcimSiteGroups",
	Short: "GET a list of site group objects",
	Long: `
Metropolis Netbox Automation Tools:
  GET a list of site group objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(siteGroups)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.site_groups")

		if responseObject.Count > 0 {
			color.Cyan("\n  Metropolis Site Groups: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    Metropolis Site Group: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				color.Cyan("\tSlug: " + color.YellowString("%s", result.Slug))
				if result.Parent.Id > 0 {
					color.Cyan("\tParent: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Parent.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Parent.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Parent.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Parent.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Parent.Slug))
					color.Cyan("\t  Depth: " + color.YellowString("%d", result.Parent.Depth))
				} else {
					color.Cyan("\tParent: " + color.RedString("No parent entry found for ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", result.Display))
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
				color.Cyan("\tSite Count: " + color.YellowString("%d", result.SiteCount))
				color.Cyan("\tDepth: " + color.YellowString("%d\n", result.Depth))
			}
		} else {
			color.Cyan("  Metropolis Site Groups: " + color.RedString("No site groups found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimSiteGroupsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimSiteGroupsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimSiteGroupsCmd", err)
	}
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimSiteGroupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimSiteGroupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

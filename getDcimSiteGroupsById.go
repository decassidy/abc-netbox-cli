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

type siteGroupsByID struct {
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
}

// GetDcimSiteGroupsByIdCmd represents the getDcimSiteGroupsById command
var GetDcimSiteGroupsByIdCmd = &cobra.Command{
	Use:   "getDcimSiteGroupsById",
	Short: "GET an site group object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an site group object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(siteGroupsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.site_groups_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    Metropolis Site Group: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			color.Cyan("\tName: " + color.YellowString("%s", responseObject.Name))
			color.Cyan("\tSlug: " + color.YellowString("%s", responseObject.Slug))
			if responseObject.Parent.Id > 0 {
				color.Cyan("\tParent: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Parent.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Parent.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Parent.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Parent.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Parent.Slug))
				color.Cyan("\t  Depth: " + color.YellowString("%d", responseObject.Parent.Depth))
			} else {
				color.Cyan("\tParent: " + color.RedString("No parent entry found for ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
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
			color.Cyan("\tSite Count: " + color.YellowString("%d", responseObject.SiteCount))
			color.Cyan("\tDepth: " + color.YellowString("%d\n", responseObject.Depth))
		} else {
			color.Red("  Doh! No site groups object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimSiteGroupsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimSiteGroupsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimSiteGroupsByIdCmd", err)
	}

	GetDcimSiteGroupsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the site group object")
	err = GetDcimSiteGroupsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimSiteGroupsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimSiteGroupsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

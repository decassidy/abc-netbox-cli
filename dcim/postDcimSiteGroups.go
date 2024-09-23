/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PostDcimSiteGroupsCmd represents the postDcimSiteGroups command
var PostDcimSiteGroupsCmd = &cobra.Command{
	Use:   "postDcimSiteGroups",
	Short: "POST a list of site group objects.",
	Long: `
ABC Netbox Automation Tools:
  POST a list of site group objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPost("cmd.dcim.dcim_api_url.site_groups_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PostDcimSiteGroupsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PostDcimSiteGroupsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PostDcimSiteGroupsCmd", err)
	}

	PostDcimSiteGroupsCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PostDcimSiteGroupsCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PostDcimSiteGroupsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postDcimSiteGroupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postDcimSiteGroupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

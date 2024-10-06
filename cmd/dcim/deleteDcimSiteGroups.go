/*
Copyright © 2024 Derrick Cassidy.

*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// DeleteDcimSiteGroupsCmd represents the deleteDcimSiteGroups command
var DeleteDcimSiteGroupsCmd = &cobra.Command{
	Use:   "deleteDcimSiteGroups",
	Short: "DELETE a list of site group objects",
	Long: `
ABC Netbox Automation Tools:
  DELETE a list of site group objects`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionDelete("cmd.dcim.dcim_api_url.site_groups_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteDcimSiteGroupsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := DeleteDcimSiteGroupsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for DeleteDcimSiteGroupsCmd", err)
	}

	DeleteDcimSiteGroupsCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be sent in delete request")
	err = DeleteDcimSiteGroupsCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for DeleteDcimSiteGroupsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDcimSiteGroupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteDcimSiteGroupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

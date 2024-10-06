/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// DeleteDcimSiteGroupsByIdCmd represents the deleteDcimSiteGroupsById command
var DeleteDcimSiteGroupsByIdCmd = &cobra.Command{
	Use:   "deleteDcimSiteGroupsById",
	Short: "DELETE an site group object by ID",
	Long: `
ABC Netbox Automation Tools:
  DELETE an site group object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionDeleteID("cmd.dcim.dcim_api_url.site_groups_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteDcimSiteGroupsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := DeleteDcimSiteGroupsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for DeleteDcimSiteGroupsByIdCmd", err)
	}

	DeleteDcimSiteGroupsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the site group object to be deleted")
	err = DeleteDcimSiteGroupsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for DeleteDcimSiteGroupsByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDcimSiteGroupsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteDcimSiteGroupsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

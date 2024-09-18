/*
Copyright © 2024 Derrick Cassidy - Metropolis Technologies, Inc.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// DeleteDcimSitesByIdCmd represents the deleteDcimSitesById command
var DeleteDcimSitesByIdCmd = &cobra.Command{
	Use:   "deleteDcimSitesById",
	Short: "DELETE an site object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  DELETE an site object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionDeleteID("cmd.dcim.dcim_api_url.sites_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteDcimSitesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := DeleteDcimSitesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for DeleteDcimSitesByIdCmd", err)
	}

	DeleteDcimSitesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the site object to be deleted")
	err = DeleteDcimSitesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for DeleteDcimSitesByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDcimSitesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteDcimSitesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

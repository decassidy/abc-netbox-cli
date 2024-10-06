/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PatchDcimSiteGroupsByIdCmd represents the patchDcimSiteGroupsById command
var PatchDcimSiteGroupsByIdCmd = &cobra.Command{
	Use:   "patchDcimSiteGroupsById",
	Short: "PATCH an site group object by ID",
	Long: `
ABC Netbox Automation Tools:
  PATCH an site group object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPatchID("cmd.dcim.dcim_api_url.site_groups_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PatchDcimSiteGroupsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PatchDcimSiteGroupsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PatchDcimSiteGroupsByIdCmd", err)
	}

	PatchDcimSiteGroupsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of device bay template to patch")
	err = PatchDcimSiteGroupsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking id flag as required: %s - for PatchDcimSiteGroupsByIdCmd", err)
	}

	PatchDcimSiteGroupsByIdCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PatchDcimSiteGroupsByIdCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PatchDcimSiteGroupsByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchDcimSiteGroupsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchDcimSiteGroupsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PatchDcimSitesByIdCmd represents the patchDcimSitesById command
var PatchDcimSitesByIdCmd = &cobra.Command{
	Use:   "patchDcimSitesById",
	Short: "PATCH an site object by ID",
	Long: `
ABC Netbox Automation Tools:
  PATCH an site object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPatchID("cmd.dcim.dcim_api_url.sites_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PatchDcimSitesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PatchDcimSitesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PatchDcimSitesByIdCmd", err)
	}

	PatchDcimSitesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of device bay template to patch")
	err = PatchDcimSitesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking id flag as required: %s - for PatchDcimSitesByIdCmd", err)
	}

	PatchDcimSitesByIdCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PatchDcimSitesByIdCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PatchDcimSitesByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchDcimSitesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchDcimSitesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

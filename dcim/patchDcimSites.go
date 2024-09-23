/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PatchDcimSitesCmd represents the patchDcimSites command
var PatchDcimSitesCmd = &cobra.Command{
	Use:   "patchDcimSites",
	Short: "PATCH a list of site objects.",
	Long: `
ABC Netbox Automation Tools:
  PATCH a list of site objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPatch("cmd.dcim.dcim_api_url.sites_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PatchDcimSitesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PatchDcimSitesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PatchDcimSitesCmd", err)
	}

	PatchDcimSitesCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PatchDcimSitesCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PatchDcimSitesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchDcimSitesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchDcimSitesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

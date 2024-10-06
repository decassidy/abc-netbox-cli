/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PatchDcimVirtualChassisByIdCmd represents the patchDcimVirtualChassisById command
var PatchDcimVirtualChassisByIdCmd = &cobra.Command{
	Use:   "patchDcimVirtualChassisById",
	Short: "PATCH an virtual chassis object by ID",
	Long: `
ABC Netbox Automation Tools:
  PATCH an virtual chassis object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPatchID("cmd.dcim.dcim_api_url.virtual_chassis_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PatchDcimVirtualChassisByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PatchDcimVirtualChassisByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PatchDcimVirtualChassisByIdCmd", err)
	}

	PatchDcimVirtualChassisByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of device bay template to patch")
	err = PatchDcimVirtualChassisByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking id flag as required: %s - for PatchDcimVirtualChassisByIdCmd", err)
	}

	PatchDcimVirtualChassisByIdCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PatchDcimVirtualChassisByIdCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PatchDcimVirtualChassisByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchDcimVirtualChassisByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchDcimVirtualChassisByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

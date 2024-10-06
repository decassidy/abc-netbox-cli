/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PatchDcimVirtualChassisCmd represents the patchDcimVirtualChassis command
var PatchDcimVirtualChassisCmd = &cobra.Command{
	Use:   "patchDcimVirtualChassis",
	Short: "PATCH a list of virtual chassis objects.",
	Long: `
ABC Netbox Automation Tools:
  PATCH a list of virtual chassis objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPatch("cmd.dcim.dcim_api_url.virtual_chassis_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PatchDcimVirtualChassisCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PatchDcimVirtualChassisCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PatchDcimVirtualChassisCmd", err)
	}

	PatchDcimVirtualChassisCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PatchDcimVirtualChassisCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PatchDcimVirtualChassisCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchDcimVirtualChassisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchDcimVirtualChassisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

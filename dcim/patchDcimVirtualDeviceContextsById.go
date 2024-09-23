/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PatchDcimVirtualDeviceContextsByIdCmd represents the patchDcimVirtualDeviceContextsById command
var PatchDcimVirtualDeviceContextsByIdCmd = &cobra.Command{
	Use:   "patchDcimVirtualDeviceContextsById",
	Short: "PATCH an virtual device context object by ID",
	Long: `
ABC Netbox Automation Tools:
  PATCH an virtual device context object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPatchID("cmd.dcim.dcim_api_url.virtual_device_contexts_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PatchDcimVirtualDeviceContextsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PatchDcimVirtualDeviceContextsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PatchDcimVirtualDeviceContextsByIdCmd", err)
	}

	PatchDcimVirtualDeviceContextsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of device bay template to patch")
	err = PatchDcimVirtualDeviceContextsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking id flag as required: %s - for PatchDcimVirtualDeviceContextsByIdCmd", err)
	}

	PatchDcimVirtualDeviceContextsByIdCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PatchDcimVirtualDeviceContextsByIdCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PatchDcimVirtualDeviceContextsByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchDcimVirtualDeviceContextsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchDcimVirtualDeviceContextsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

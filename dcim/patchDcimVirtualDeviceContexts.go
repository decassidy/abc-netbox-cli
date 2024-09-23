/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PatchDcimVirtualDeviceContextsCmd represents the patchDcimVirtualDeviceContexts command
var PatchDcimVirtualDeviceContextsCmd = &cobra.Command{
	Use:   "patchDcimVirtualDeviceContexts",
	Short: "PATCH a list of virtual device context objects.",
	Long: `
ABC Netbox Automation Tools:
  PATCH a list of virtual device context objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPatch("cmd.dcim.dcim_api_url.virtual_device_contexts_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PatchDcimVirtualDeviceContextsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PatchDcimVirtualDeviceContextsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PatchDcimVirtualDeviceContextsCmd", err)
	}

	PatchDcimVirtualDeviceContextsCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PatchDcimVirtualDeviceContextsCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PatchDcimVirtualDeviceContextsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchDcimVirtualDeviceContextsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchDcimVirtualDeviceContextsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

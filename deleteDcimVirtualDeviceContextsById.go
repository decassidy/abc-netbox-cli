/*
Copyright © 2024 Derrick Cassidy - Metropolis Technologies, Inc.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// DeleteDcimVirtualDeviceContextsByIdCmd represents the deleteDcimVirtualDeviceContextsById command
var DeleteDcimVirtualDeviceContextsByIdCmd = &cobra.Command{
	Use:   "deleteDcimVirtualDeviceContextsById",
	Short: "DELETE an virtual device context object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  DELETE an virtual device context object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionDeleteID("cmd.dcim.dcim_api_url.virtual_device_contexts_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteDcimVirtualDeviceContextsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := DeleteDcimVirtualDeviceContextsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for DeleteDcimVirtualDeviceContextsByIdCmd", err)
	}

	DeleteDcimVirtualDeviceContextsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the virtual device context object to be deleted")
	err = DeleteDcimVirtualDeviceContextsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for DeleteDcimVirtualDeviceContextsByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDcimVirtualDeviceContextsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteDcimVirtualDeviceContextsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

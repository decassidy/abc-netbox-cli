/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// DeleteDcimVirtualDeviceContextsCmd represents the deleteDcimVirtualDeviceContexts command
var DeleteDcimVirtualDeviceContextsCmd = &cobra.Command{
	Use:   "deleteDcimVirtualDeviceContexts",
	Short: "DELETE a list of device context objects",
	Long: `
ABC Netbox Automation Tools:
  DELETE a list of device context objects`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionDelete("cmd.dcim.dcim_api_url.virtual_device_contexts_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteDcimVirtualDeviceContextsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := DeleteDcimVirtualDeviceContextsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for DeleteDcimVirtualDeviceContextsCmd", err)
	}

	DeleteDcimVirtualDeviceContextsCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be sent in delete request")
	err = DeleteDcimVirtualDeviceContextsCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for DeleteDcimVirtualDeviceContextsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDcimVirtualDeviceContextsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteDcimVirtualDeviceContextsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2024 Derrick Cassidy.

*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PostDcimVirtualDeviceContextsCmd represents the postDcimVirtualDeviceContexts command
var PostDcimVirtualDeviceContextsCmd = &cobra.Command{
	Use:   "postDcimVirtualDeviceContexts",
	Short: "POST a list of virtual device context objects.",
	Long: `
ABC Netbox Automation Tools:
  POST a list of virtual device context objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPost("cmd.dcim.dcim_api_url.virtual_device_contexts_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PostDcimVirtualDeviceContextsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PostDcimVirtualDeviceContextsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PostDcimVirtualDeviceContextsCmd", err)
	}

	PostDcimVirtualDeviceContextsCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PostDcimVirtualDeviceContextsCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PostDcimVirtualDeviceContextsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postDcimVirtualDeviceContextsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postDcimVirtualDeviceContextsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

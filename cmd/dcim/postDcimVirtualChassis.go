/*
Copyright © 2024 Derrick Cassidy.

*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PostDcimVirtualChassisCmd represents the postDcimVirtualChassis command
var PostDcimVirtualChassisCmd = &cobra.Command{
	Use:   "postDcimVirtualChassis",
	Short: "POST a list of virtual chassis objects.",
	Long: `
ABC Netbox Automation Tools:
  POST a list of virtual chassis objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPost("cmd.dcim.dcim_api_url.virtual_chassis_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PostDcimVirtualChassisCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PostDcimVirtualChassisCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PostDcimVirtualChassisCmd", err)
	}

	PostDcimVirtualChassisCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PostDcimVirtualChassisCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PostDcimVirtualChassisCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postDcimVirtualChassisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postDcimVirtualChassisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

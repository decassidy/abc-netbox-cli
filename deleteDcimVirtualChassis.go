/*
Copyright © 2024 Derrick Cassidy - Metropolis Technologies, Inc.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// DeleteDcimVirtualChassisCmd represents the deleteDcimVirtualChassis command
var DeleteDcimVirtualChassisCmd = &cobra.Command{
	Use:   "deleteDcimVirtualChassis",
	Short: "DELETE a list of virtual chassis objects",
	Long: `
Metropolis Netbox Automation Tools:
  DELETE a list of virtual chassis objects`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionDelete("cmd.dcim.dcim_api_url.virtual_chassis_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteDcimVirtualChassisCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := DeleteDcimVirtualChassisCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for DeleteDcimVirtualChassisCmd", err)
	}

	DeleteDcimVirtualChassisCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be sent in delete request")
	err = DeleteDcimVirtualChassisCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for DeleteDcimVirtualChassisCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDcimVirtualChassisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteDcimVirtualChassisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

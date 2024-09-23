/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// DeleteDcimVirtualChassisByIdCmd represents the deleteDcimVirtualChassisById command
var DeleteDcimVirtualChassisByIdCmd = &cobra.Command{
	Use:   "deleteDcimVirtualChassisById",
	Short: "DELETE an virtual chassis object by ID",
	Long: `
ABC Netbox Automation Tools:
  DELETE an virtual chassis object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionDeleteID("cmd.dcim.dcim_api_url.virtual_chassis_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteDcimVirtualChassisByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := DeleteDcimVirtualChassisByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for DeleteDcimVirtualChassisByIdCmd", err)
	}

	DeleteDcimVirtualChassisByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of virtual chassis object to be deleted")
	err = DeleteDcimVirtualChassisByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for DeleteDcimVirtualChassisByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDcimVirtualChassisByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteDcimVirtualChassisByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

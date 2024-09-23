/*
Copyright © 2024 Derrick Cassidy.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// DeleteDcimSitesCmd represents the deleteDcimSites command
var DeleteDcimSitesCmd = &cobra.Command{
	Use:   "deleteDcimSites",
	Short: "DELETE a list of site objects",
	Long: `
ABC Netbox Automation Tools:
  DELETE a list of site objects`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionDelete("cmd.dcim.dcim_api_url.sites_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteDcimSitesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := DeleteDcimSitesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for DeleteDcimSitesCmd", err)
	}

	DeleteDcimSitesCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be sent in delete request")
	err = DeleteDcimSitesCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for DeleteDcimSitesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDcimSitesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteDcimSitesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

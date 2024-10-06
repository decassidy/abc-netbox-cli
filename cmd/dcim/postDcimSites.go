/*
Copyright © 2024 Derrick Cassidy.

*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// PostDcimSitesCmd represents the postDcimSites command
var PostDcimSitesCmd = &cobra.Command{
	Use:   "postDcimSites",
	Short: "POST a list of site objects.",
	Long: `
ABC Netbox Automation Tools:
  POST a list of site objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPost("cmd.dcim.dcim_api_url.sites_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PostDcimSitesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PostDcimSitesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PostDcimSitesCmd", err)
	}

	PostDcimSitesCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PostDcimSitesCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PostDcimSitesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postDcimSitesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postDcimSitesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2024 Derrick Cassidy - Metropolis Technologies, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package dcim

import (
	"github.com/spf13/cobra"
	"log"
)

// DeleteDcimInterfacesCmd represents the deleteDcimInterfaces command
var DeleteDcimInterfacesCmd = &cobra.Command{
	Use:   "deleteDcimInterfaces",
	Short: "DELETE a list of interface objects",
	Long: `
Metropolis Netbox Automation Tools:
  DELETE a list of interface objects`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionDelete("cmd.dcim.dcim_api_url.interfaces_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteDcimInterfacesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := DeleteDcimInterfacesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for DeleteDcimInterfacesCmd", err)
	}

	DeleteDcimInterfacesCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be sent in delete request")
	err = DeleteDcimInterfacesCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for DeleteDcimInterfacesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDcimInterfacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteDcimInterfacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

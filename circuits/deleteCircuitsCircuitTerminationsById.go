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

package circuits

import (
	"github.com/spf13/cobra"
	"log"
)

// DeleteCircuitsCircuitTerminationsByIdCmd represents the deleteCircuitsCircuitTerminationsById command
var DeleteCircuitsCircuitTerminationsByIdCmd = &cobra.Command{
	Use:   "deleteCircuitsCircuitTerminationsById",
	Short: "Delete a circuit termination object.",
	Long: `
Metropolis Netbox Automation Tools:
  Delete a circuit termination object.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionDeleteID("cmd.circuits.circuits_api_url.circuits_terminations_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteCircuitsCircuitTerminationsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := DeleteCircuitsCircuitTerminationsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for DeleteCircuitsCircuitTerminationsByIdCmd", err)
	}

	DeleteCircuitsCircuitTerminationsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the circuit termination object to be deleted")
	err = DeleteCircuitsCircuitTerminationsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: id: %s - for DeleteCircuitsCircuitTerminationsByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCircuitsCircuitTerminationsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCircuitsCircuitTerminationsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

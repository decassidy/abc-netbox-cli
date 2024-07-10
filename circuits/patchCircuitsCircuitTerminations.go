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

// PatchCircuitsCircuitTerminationsCmd represents the patchCircuitsCircuitTerminations command
var PatchCircuitsCircuitTerminationsCmd = &cobra.Command{
	Use:   "patchCircuitsCircuitTerminations",
	Short: "PATCH a list of circuit termination objects.",
	Long: `
Metropolis Netbox Automation Tools:
  PATCH a list of circuit termination objects.

  Example: "[{\"id\": 65, \"circuits\": {\"cid\": \"MyFakeCircuit654\"}, \"term_side\": \"Z\", \"site\": {\"name\": \"Your Site Name]\"}}]"`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPatch("cmd.circuits.circuits_api_url.circuits_terminations_id")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	PatchCircuitsCircuitTerminationsCmd.PersistentFlags().StringVarP(&serverEnv, "env", "e", "development", "Environment ('production' or 'development')")
	err := PatchCircuitsCircuitTerminationsCmd.MarkPersistentFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PatchCircuitsCircuitTerminationsCmd", err)
	}

	PatchCircuitsCircuitTerminationsCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be sent in PATCH request")
	err = PatchCircuitsCircuitTerminationsCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PatchCircuitsCircuitTerminationsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchCircuitsCircuitTerminationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchCircuitsCircuitTerminationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

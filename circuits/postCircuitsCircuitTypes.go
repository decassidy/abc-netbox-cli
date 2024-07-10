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

// PostCircuitsCircuitTypeCmd PostCircuitsCircuitTypesCmd represents the postCircuitsCircuitTypes command
var PostCircuitsCircuitTypeCmd = &cobra.Command{
	Use:   "postCircuitType",
	Short: "POST a list of Metropolis Circuit Type objects.",
	Long:  `POST a list of Metropolis Circuit Type objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPost("cmd.circuits.circuits_api_url.circuit_types_id")
	},
}

// init initializes the flags and configuration settings for the 'postCircuitType' and 'postCircuitTermination' commands.
// It defines the flags for each command, setting the corresponding variables and specifying the descriptions.
// It also marks the required flags as required.
// The function takes no parameters and does not return any values.
func init() {

	// Here you will define your flags and configuration settings.
	PostCircuitsCircuitTypeCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := PostCircuitsCircuitTypeCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PostCircuitsCircuitTypeCmd", err)
	}

	PostCircuitsCircuitTypeCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be posted (required)")
	err = PostCircuitsCircuitTypeCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PostCircuitsCircuitTypeCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postCircuitTypeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postCircuitTypeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

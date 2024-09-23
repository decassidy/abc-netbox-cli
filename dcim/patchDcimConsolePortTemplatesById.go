/*
Copyright © 2024 Derrick Cassidy.

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

// PatchDcimConsolePortTemplatesByIdCmd represents the patchDcimConsolePortTemplatesById command
var PatchDcimConsolePortTemplatesByIdCmd = &cobra.Command{
	Use:   "patchDcimConsolePortTemplatesById",
	Short: "PATCH an console port templates object by ID",
	Long: `
ABC Netbox Automation Tools:
  PATCH an console port templates object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		apiConnectionPatchID("cmd.dcim.dcim_api_url.console_port_templates_id")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	PatchDcimConsolePortTemplatesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "", "Environment ('development' or 'production')")
	err := PatchDcimConsolePortTemplatesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for PatchDcimConsolePortTemplatesByIdCmd", err)
	}

	PatchDcimConsolePortTemplatesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the console port template object")
	err = PatchDcimConsolePortTemplatesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking id flag as required: %s - for PatchDcimConsolePortTemplatesByIdCmd", err)
	}

	PatchDcimConsolePortTemplatesByIdCmd.Flags().StringVarP(&data, "data", "", "", "JSON data to be patched (required)")
	err = PatchDcimConsolePortTemplatesByIdCmd.MarkFlagRequired("data")
	if err != nil {
		log.Fatalf("Error marking data flag as required: %s - for PatchDcimConsolePortTemplatesByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchDcimConsolePortTemplatesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchDcimConsolePortTemplatesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

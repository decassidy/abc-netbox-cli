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

package core

import (
	"fmt"
	"github.com/spf13/cobra"
)

// PostCoreDataSourcesCmd represents the postCoreDataSources command. It is used to post a list of Metropolis Netbox data source objects.
var PostCoreDataSourcesCmd = &cobra.Command{
	Use:   "postCoreDataSources",
	Short: "Post a list of Metropolis Netbox data source objects.",
	Long:  `Post a list of Metropolis Netbox data source objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("postCoreDataSources called")
	},
}

// init initializes the flags and configuration settings for the PostCoreDataSourcesCmd command.
// It defines the flags for the command including Name, Type, SourceUrl, Enabled, Description, and Comments.
// The flags are set as mandatory or optional fields depending on their purpose.
// The command also supports Persistent Flags and local flags.
// When the command is called directly, it runs the postCoreDataSources function with the specified flag values.
// The postCoreDataSources function uses the flag values to perform an action or operation.
// Note: Refer to the PostCoreDataSourcesCmd declaration for more details on how the flag values are used.
func init() {

	// Here you will define your flags and configuration settings.
	//if len(os.Args) < 3 {
	//	log.Fatalf("Insufficient arguments. Please make sure to provide all mandatory fields: Name, Type, and SourceURL")
	//}
	////
	//PostCoreDataSourcesCmd.Flags().StringP("Name", "n", "", "DataSource Name (*Mandatory Field*)")
	//if _, err := PostCoreDataSourcesCmd.Flags().GetString("Name"); err != nil {
	//	log.Fatalf("Name not provided")
	//}
	//
	//PostCoreDataSourcesCmd.Flags().StringP("Type", "t", "", "DataSource Type (*Mandatory Field*)")
	//if _, err := PostCoreDataSourcesCmd.Flags().GetString("Type"); err != nil {
	//	log.Fatalf("Type not provided")
	//}
	//
	//PostCoreDataSourcesCmd.Flags().StringP("SourceUrl", "s", "", "DataSource URL (*Mandatory Field*)")
	//if _, err := PostCoreDataSourcesCmd.Flags().GetString("SourceUrl"); err != nil {
	//	log.Fatalf("SourceUrl not provided")
	//}
	//PostCoreDataSourcesCmd.Flags().BoolP("Enabled", "e", false, "DataSource Enabled (*Optional Field*)")
	//PostCoreDataSourcesCmd.Flags().StringP("Description", "d", "", "DataSource Description (*Optional Field*)")
	//PostCoreDataSourcesCmd.Flags().StringP("Comments", "c", "", "DataSource Comments (*Optional Field*)")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postCoreDataSourcesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postCoreDataSourcesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

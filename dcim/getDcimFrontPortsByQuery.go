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
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// GetDcimFrontPortsByQueryCmd represents the getDcimFrontPortsByQuery command
var GetDcimFrontPortsByQueryCmd = &cobra.Command{
	Use:   "getDcimFrontPortsByQuery",
	Short: "GET a front port object(s) by string query",
	Long: `
Netbox Automation Tools:
  GET a front port object(s) by string query`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(frontPorts)
		ApiConnectionQuery(responseObject, "GET", "cmd.dcim.dcim_api_url.front_ports_id")

		if responseObject.Count != 0 {
			color.Cyan("Total ABC Front Ports: "+color.YellowString("%v"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    ABC Front Port Name: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				color.Cyan("\tDevice: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", result.Device.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.Device.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.Device.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", result.Device.Name))
				color.Cyan("\tModule: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", result.Module.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.Module.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.Module.Display))
				color.Cyan("\t  Device: " + color.YellowString("%d", result.Module.Device))
				color.Cyan("\t  Module Bay: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", result.Module.ModuleBay.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", result.Module.ModuleBay.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", result.Module.ModuleBay.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", result.Module.ModuleBay.Name))
				color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				color.Cyan("\tLabel: " + color.YellowString("%s", result.Label))
				color.Cyan("\tType: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", result.Type.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", result.Type.Label))
				color.Cyan("\tColor: " + color.YellowString("%s", result.Color))
				color.Cyan("\tRear Port: " + color.YellowString("%s", result.RearPort))
				color.Cyan("\t  ID: " + color.YellowString("%d", result.RearPort.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.RearPort.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.RearPort.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", result.RearPort.Name))
				color.Cyan("\t  Label: " + color.YellowString("%s", result.RearPort.Label))
				color.Cyan("\t  Description: " + color.YellowString("%s", result.RearPort.Description))
				color.Cyan("\tRear Port Position: " + color.YellowString("%d", result.RearPortPosition))
				color.Cyan("\tDescription: " + color.YellowString("%d", result.Description))
				color.Cyan("\tMarked Connected: " + color.YellowString("%v", result.MarkConnected))
				color.Cyan("\tCable: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", result.Cable.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.Cable.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.Cable.Display))
				color.Cyan("\t  Label: " + color.YellowString("%s", result.Cable.Label))
				color.Cyan("\tCable End: " + color.YellowString("%s", result.CableEnd))
				for _, link := range result.LinkPeers {
					color.Cyan("\tLink Peer: " + color.YellowString("%d", link))
				}
				color.Cyan("\tLink Peers Type: " + color.YellowString("%s", result.LinkPeersType))
				for _, tag := range result.Tags {
					color.Cyan("\tTags: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
					color.Cyan("\t  URL: " + color.YellowString("%d", tag.Url))
					color.Cyan("\t  Display: " + color.YellowString("%d", tag.Display))
					color.Cyan("\t  Name: " + color.YellowString("%d", tag.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%d", tag.Slug))
					color.Cyan("\t  Color: " + color.YellowString("%d", tag.Color))
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
				color.Cyan("\tOccupied: " + color.YellowString("%s\n", result.Occupied))
			}

		} else {
			color.Cyan("  Total ABC Front Ports: " + color.RedString("No front ports found on the server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimFrontPortsByQueryCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimFrontPortsByQueryCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for GetDcimFrontPortsByQueryCmd", err)
	}

	GetDcimFrontPortsByQueryCmd.Flags().StringVarP(&query, "query", "q", "", "string query of object you want to get")
	err = GetDcimFrontPortsByQueryCmd.MarkFlagRequired("query")
	if err != nil {
		log.Fatalf("Error marking query flag as required: %s - for GetDcimFrontPortsByQueryCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimFrontPortsByQueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimFrontPortsByQueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

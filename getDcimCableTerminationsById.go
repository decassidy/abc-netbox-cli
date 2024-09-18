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
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type cableTerminationsByID struct {
	Id              uint   `json:"id"`
	Url             string `json:"url"`
	Display         string `json:"display"`
	Cable           uint   `json:"cable"`
	CableEnd        string `json:"cable_end"`
	TerminationType string `json:"termination_type"`
	TerminationId   uint   `json:"termination_id"`
	Termination     struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Circuit struct {
			Id          int    `json:"id"`
			Url         string `json:"url"`
			Display     string `json:"display"`
			Cid         string `json:"cid"`
			Description string `json:"description"`
		} `json:"circuit,omitempty"`
		TermSide    string `json:"term_side,omitempty"`
		Description string `json:"description"`
		Cable       struct {
			Id          uint   `json:"id"`
			Url         string `json:"url"`
			Display     string `json:"display"`
			Label       string `json:"label,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"cable"`
		Occupied bool `json:"_occupied"`
		Device   struct {
			Id          uint   `json:"id"`
			Url         string `json:"url"`
			Display     string `json:"display"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"device,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"termination"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

// GetDcimCableTerminationsByIdCmd represents the getDcimCableTerminationsById command
var GetDcimCableTerminationsByIdCmd = &cobra.Command{
	Use:   "getDcimCableTerminationsById",
	Short: "GET an cable termination object by ID",
	Long: `
Metropolis Netbox Automation Tools:
  GET an cable termination object by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(cableTerminationsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.cable_terminations_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    Metropolis Cable Termination: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			if responseObject.Cable > 0 {
				color.Cyan("\tCable: " + color.YellowString("%d", responseObject.Cable))
			} else {
				color.Cyan("\tCable: " + color.RedString("No cable entry found for termination: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.CableEnd != "" {
				color.Cyan("\tCable End: " + color.YellowString("%s", responseObject.CableEnd))
			} else {
				color.Cyan("\tCable End: " + color.RedString("No cable end entry found for termination: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.TerminationType != "" {
				color.Cyan("\tTermination Type: " + color.YellowString("%s", responseObject.TerminationType))
			} else {
				color.Cyan("\tTermination Type: " + color.RedString("No termination type entry found for termination: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Termination.Id > 0 {
				color.Cyan("\tTermination ID: " + color.YellowString("%v", responseObject.TerminationId))
			} else {
				color.Cyan("\tTermination ID: " + color.RedString("No termination ID entry found for termination: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Termination.Id > 0 {
				color.Cyan("\tTermination: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Termination.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Termination.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Termination.Display))
			} else {
				color.Cyan("\tTermination: " + color.RedString("No termination entry found for termination: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Termination.Circuit.Id > 0 {
				color.Cyan("\t  Circuit:")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Termination.Circuit.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Termination.Circuit.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Termination.Circuit.Display))
				color.Cyan("\t    CID: " + color.YellowString("%s", responseObject.Termination.Circuit.Cid))
				color.Cyan("\t    Description: " + color.YellowString("%s", responseObject.Termination.Circuit.Description))
			} else {
				color.Cyan("\t  Circuit: " + color.RedString("No circuit found for: "+color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Termination.TermSide != "" {
				color.Cyan("\t  Term Side: " + color.YellowString("%s", responseObject.Termination.TermSide))
			} else {
				color.Cyan("\t  Term Side: " + color.RedString("No term side found for: "+color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Termination.Description != "" {
				color.Cyan("\t  Description: " + color.YellowString("%s", responseObject.Termination.Description))
			} else {
				color.Cyan("\t  Description: " + color.RedString("No term side found for: "+color.YellowString("%s", responseObject.Display)))
			}
			if responseObject.Termination.Cable.Id > 0 {
				color.Cyan("\t  Cable: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Termination.Cable.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Termination.Cable.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Termination.Cable.Display))
				if responseObject.Termination.Cable.Label != "" {
					color.Cyan("\t    Label: " + color.YellowString("%s", responseObject.Termination.Cable.Label))
				} else {
					color.Cyan("\t    Label: " + color.RedString("No label found for: "+color.YellowString("%s", responseObject.Display)))
				}
				if responseObject.Termination.Cable.Description != "" {
					color.Cyan("\t    Description: " + color.YellowString("%s", responseObject.Termination.Cable.Description))
				} else {
					color.Cyan("\t    Description: " + color.RedString("No description found for: "+color.YellowString("%s", responseObject.Display)))
				}
			} else {
				color.Cyan("\t  Cable: " + color.RedString("No cable found for: "+color.YellowString("%s", responseObject.Display)))
			}

			color.Cyan("\t  Occupied: " + color.YellowString("%v", responseObject.Termination.Occupied))

			if responseObject.Termination.Device.Id > 0 {
				color.Cyan("\t  Device: ")
				color.Cyan("\t    ID: " + color.YellowString("%d", responseObject.Termination.Device.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", responseObject.Termination.Device.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", responseObject.Termination.Device.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", responseObject.Termination.Device.Name))
				color.Cyan("\t    Description: " + color.YellowString("%s", responseObject.Termination.Device.Description))
			} else {
				color.Cyan("\t  Device: " + color.RedString("No device found for: "+color.YellowString("%s", responseObject.Display)))
			}

			if responseObject.Termination.Name != "" {
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Termination.Name))
			} else {
				color.Cyan("\t  Name: " + color.RedString("No name found for: "+color.YellowString("%s", responseObject.Display)))
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s\n", responseObject.LastUpdated))
		} else {
			color.Cyan("\tMetropolis Cable Termination: " + color.RedString("No cable termination entries found for on server for ID: %d. Exiting...\n", id))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimCableTerminationsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimCableTerminationsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - GetDcimCableTerminationsByIdCmd", err)
	}

	GetDcimCableTerminationsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of Cable Termination")
	err = GetDcimCableTerminationsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimCableTerminationsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimCableTerminationsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

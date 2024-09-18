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
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

// cableTerminations represents the structure of the response containing cable terminations information.
// It is used to deserialize JSON into Go objects.
type cableTerminations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous string  `json:"previous"`
	Results  []struct {
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
	} `json:"results"`
}

var responseObject = new(cableTerminations)

// GetDcimCableTerminationsCmd represents the getDcimAllCableTeminations command.
var GetDcimCableTerminationsCmd = &cobra.Command{
	Use:   "getDcimCableTerminations",
	Short: "GET a list of cable termination objects.",
	Long: `
Metropolis Netbox Automation Tools: 
  GET a list of cable termination objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.cable_terminations")
		if responseObject.Count > 0 {
			color.Cyan("\n  Total Metropolis Cable Terminations: "+color.YellowString("%v"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    Metropolis Cable Termination: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.Cable > 0 {
					color.Cyan("\tCable: " + color.YellowString("%d", result.Cable))
				} else {
					color.Cyan("\tCable: " + color.RedString("No cable entry found for termination: ") + color.YellowString("%s", result.Display))
				}
				if result.CableEnd != "" {
					color.Cyan("\tCable End: " + color.YellowString("%s", result.CableEnd))
				} else {
					color.Cyan("\tCable End: " + color.RedString("No cable end entry found for termination: ") + color.YellowString("%s", result.Display))
				}
				if result.TerminationType != "" {
					color.Cyan("\tTermination Type: " + color.YellowString("%s", result.TerminationType))
				} else {
					color.Cyan("\tTermination Type: " + color.RedString("No termination type entry found for termination: ") + color.YellowString("%s", result.Display))
				}
				if result.Termination.Id > 0 {
					color.Cyan("\tTermination ID: " + color.YellowString("%v", result.TerminationId))
				} else {
					color.Cyan("\tTermination ID: " + color.RedString("No termination ID entry found for termination: ") + color.YellowString("%s", result.Display))
				}
				if result.Termination.Id > 0 {
					color.Cyan("\tTermination: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Termination.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Termination.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Termination.Display))
				} else {
					color.Cyan("\tTermination: " + color.RedString("No termination entry found for termination: ") + color.YellowString("%s", result.Display))
				}
				if result.Termination.Circuit.Id > 0 {
					color.Cyan("\t  Circuit:")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Termination.Circuit.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Termination.Circuit.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Termination.Circuit.Display))
					color.Cyan("\t    CID: " + color.YellowString("%s", result.Termination.Circuit.Cid))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Termination.Circuit.Description))
				} else {
					color.Cyan("\t  Circuit: " + color.RedString("No circuit found for: "+color.YellowString("%s", result.Display)))
				}
				if result.Termination.TermSide != "" {
					color.Cyan("\t  Term Side: " + color.YellowString("%s", result.Termination.TermSide))
				} else {
					color.Cyan("\t  Term Side: " + color.RedString("No term side found for: "+color.YellowString("%s", result.Display)))
				}
				if result.Termination.Description != "" {
					color.Cyan("\t  Description: " + color.YellowString("%s", result.Termination.Description))
				} else {
					color.Cyan("\t  Description: " + color.RedString("No term side found for: "+color.YellowString("%s", result.Display)))
				}
				if result.Termination.Cable.Id > 0 {
					color.Cyan("\t  Cable: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Termination.Cable.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Termination.Cable.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Termination.Cable.Display))
					if result.Termination.Cable.Label != "" {
						color.Cyan("\t    Label: " + color.YellowString("%s", result.Termination.Cable.Label))
					} else {
						color.Cyan("\t    Label: " + color.RedString("No label found for: "+color.YellowString("%s", result.Display)))
					}
					if result.Termination.Cable.Description != "" {
						color.Cyan("\t    Description: " + color.YellowString("%s", result.Termination.Cable.Description))
					} else {
						color.Cyan("\t    Description: " + color.RedString("No description found for: "+color.YellowString("%s", result.Display)))
					}
				} else {
					color.Cyan("\t  Cable: " + color.RedString("No cable found for: "+color.YellowString("%s", result.Display)))
				}

				color.Cyan("\t  Occupied: " + color.YellowString("%v", result.Termination.Occupied))

				if result.Termination.Device.Id > 0 {
					color.Cyan("\t  Device: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Termination.Device.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Termination.Device.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Termination.Device.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Termination.Device.Name))
					color.Cyan("\t    Description: " + color.YellowString("%s", result.Termination.Device.Description))
				} else {
					color.Cyan("\t  Device: " + color.RedString("No device found for: "+color.YellowString("%s", result.Display)))
				}

				if result.Termination.Name != "" {
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Termination.Name))
				} else {
					color.Cyan("\t  Name: " + color.RedString("No name found for: "+color.YellowString("%s", result.Display)))
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
			}
			for responseObject.Next != nil {
				nextPage()
			}
			if responseObject.Next == nil {
				display := color.HiGreenString("\tAll Netbox cable termination objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n")
			}
		} else {
			color.Cyan("  Total Metropolis Cable Terminations: " + color.RedString("No cable termination entries found on server.\n"))
		}
	},
}

func displayCableTerminationsOutput() {
	for _, result := range responseObject.Results {
		display := fmt.Sprintf("    Metropolis Cable Termination: %s\n", color.YellowString(result.Display))
		equals := strings.Repeat("=", len(display))
		color.Cyan("\n  " + equals + "\n")
		color.Cyan(display)
		color.Cyan("  " + equals + "\n")
		color.Cyan("\tID: " + color.YellowString("%d", result.Id))
		color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
		color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
		if result.Cable > 0 {
			color.Cyan("\tCable: " + color.YellowString("%d", result.Cable))
		} else {
			color.Cyan("\tCable: " + color.RedString("No cable entry found for termination: ") + color.YellowString("%s", result.Display))
		}
		if result.CableEnd != "" {
			color.Cyan("\tCable End: " + color.YellowString("%s", result.CableEnd))
		} else {
			color.Cyan("\tCable End: " + color.RedString("No cable end entry found for termination: ") + color.YellowString("%s", result.Display))
		}
		if result.TerminationType != "" {
			color.Cyan("\tTermination Type: " + color.YellowString("%s", result.TerminationType))
		} else {
			color.Cyan("\tTermination Type: " + color.RedString("No termination type entry found for termination: ") + color.YellowString("%s", result.Display))
		}
		if result.Termination.Id > 0 {
			color.Cyan("\tTermination ID: " + color.YellowString("%v", result.TerminationId))
		} else {
			color.Cyan("\tTermination ID: " + color.RedString("No termination ID entry found for termination: ") + color.YellowString("%s", result.Display))
		}
		if result.Termination.Id > 0 {
			color.Cyan("\tTermination: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", result.Termination.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", result.Termination.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", result.Termination.Display))
		} else {
			color.Cyan("\tTermination: " + color.RedString("No termination entry found for termination: ") + color.YellowString("%s", result.Display))
		}
		if result.Termination.Circuit.Id > 0 {
			color.Cyan("\t  Circuit:")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Termination.Circuit.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Termination.Circuit.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Termination.Circuit.Display))
			color.Cyan("\t    CID: " + color.YellowString("%s", result.Termination.Circuit.Cid))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Termination.Circuit.Description))
		} else {
			color.Cyan("\t  Circuit: " + color.RedString("No circuit found for: "+color.YellowString("%s", result.Display)))
		}
		if result.Termination.TermSide != "" {
			color.Cyan("\t  Term Side: " + color.YellowString("%s", result.Termination.TermSide))
		} else {
			color.Cyan("\t  Term Side: " + color.RedString("No term side found for: "+color.YellowString("%s", result.Display)))
		}
		if result.Termination.Description != "" {
			color.Cyan("\t  Description: " + color.YellowString("%s", result.Termination.Description))
		} else {
			color.Cyan("\t  Description: " + color.RedString("No term side found for: "+color.YellowString("%s", result.Display)))
		}
		if result.Termination.Cable.Id > 0 {
			color.Cyan("\t  Cable: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Termination.Cable.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Termination.Cable.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Termination.Cable.Display))
			if result.Termination.Cable.Label != "" {
				color.Cyan("\t    Label: " + color.YellowString("%s", result.Termination.Cable.Label))
			} else {
				color.Cyan("\t    Label: " + color.RedString("No label found for: "+color.YellowString("%s", result.Display)))
			}
			if result.Termination.Cable.Description != "" {
				color.Cyan("\t    Description: " + color.YellowString("%s", result.Termination.Cable.Description))
			} else {
				color.Cyan("\t    Description: " + color.RedString("No description found for: "+color.YellowString("%s", result.Display)))
			}
		} else {
			color.Cyan("\t  Cable: " + color.RedString("No cable found for: "+color.YellowString("%s", result.Display)))
		}

		color.Cyan("\t  Occupied: " + color.YellowString("%v", result.Termination.Occupied))

		if result.Termination.Device.Id > 0 {
			color.Cyan("\t  Device: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Termination.Device.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Termination.Device.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Termination.Device.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Termination.Device.Name))
			color.Cyan("\t    Description: " + color.YellowString("%s", result.Termination.Device.Description))
		} else {
			color.Cyan("\t  Device: " + color.RedString("No device found for: "+color.YellowString("%s", result.Display)))
		}

		if result.Termination.Name != "" {
			color.Cyan("\t  Name: " + color.YellowString("%s", result.Termination.Name))
		} else {
			color.Cyan("\t  Name: " + color.RedString("No name found for: "+color.YellowString("%s", result.Display)))
		}
		color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
		color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
	}
}

func nextPage() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\tDo you want to continue to the next page of cable termination objects? ['Y' or 'yes'] or ['n' or 'no']: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPage(responseObject, "GET", *responseObject.Next)
		displayCableTerminationsOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the metropolis-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("Invalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimCableTerminationsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimCableTerminationsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimCableTerminationsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimAllCableTeminationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimAllCableTeminationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

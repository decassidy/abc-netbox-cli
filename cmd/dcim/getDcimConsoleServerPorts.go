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
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
)

type consoleServerPorts struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Device  struct {
			CommonFieldsNoSlug
		} `json:"device"`
		Module struct {
			Id        int    `json:"id"`
			Url       string `json:"url"`
			Display   string `json:"display"`
			Device    int    `json:"device"`
			ModuleBay struct {
				CommonFieldsNoSlug
			} `json:"module_bay"`
		} `json:"module"`
		Name  string `json:"name"`
		Label string `json:"label"`
		Type  struct {
			ValueLabel
		} `json:"type"`
		Speed struct {
			Value int    `json:"value"`
			Label string `json:"label"`
		} `json:"speed"`
		Description   string `json:"description"`
		MarkConnected bool   `json:"mark_connected"`
		Cable         struct {
			Id      int    `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Label   string `json:"label"`
		} `json:"cable"`
		CableEnd                    string           `json:"cable_end"`
		LinkPeers                   []string         `json:"link_peers"`
		LinkPeersType               string           `json:"link_peers_type"`
		ConnectedEndpoints          []string         `json:"connected_endpoints"`
		ConnectedEndpointsType      string           `json:"connected_endpoints_type"`
		ConnectedEndpointsReachable bool             `json:"connected_endpoints_reachable"`
		Tags                        CommonFieldsSlug `json:"tags"`
		Created                     string           `json:"created"`
		LastUpdated                 string           `json:"last_updated"`
		Occupied                    bool             `json:"_occupied"`
	} `json:"results"`
}

var responseObjectConsoleServerPorts = new(consoleServerPorts)

// GetDcimConsoleServerPortsCmd represents the getConsoleServerPorts command
var GetDcimConsoleServerPortsCmd = &cobra.Command{
	Use:   "getDcimConsoleServerPorts",
	Short: "GET a list of console server port objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of console server port objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectConsoleServerPorts, "GET", "cmd.dcim.dcim_api_url.console_server_ports")

		if len(responseObjectConsoleServerPorts.Results) == 0 {
			color.Cyan("  Total ABC Console Server Ports: " + color.RedString("No console server port entries found on server. Exiting...\n"))
		} else {
			color.Cyan("  Total ABC Console Server Ports: "+color.YellowString("%d"), responseObjectConsoleServerPorts.Count)

			for _, result := range responseObjectConsoleServerPorts.Results {
				display := fmt.Sprintf("    ABC Console Server Port: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.Device.Id != 0 {
					color.Cyan("\tDevice: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Device.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Device.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Device.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Device.Name))
				} else {
					color.Cyan("\tDevice: " + color.RedString("No device entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				if result.Module.Id != 0 {
					color.Cyan("\tModule: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Module.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Module.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Module.Display))
					color.Cyan("\t  Module Bay: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.Module.ModuleBay.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.Module.ModuleBay.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.Module.ModuleBay.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.Module.ModuleBay.Name))
				} else {
					color.Cyan("\tModule: " + color.RedString("No module entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				if result.Name != "" {
					color.Cyan("\tName: " + color.YellowString(result.Name))
				} else {
					color.Cyan("\tName: " + color.RedString("No label entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				if result.Label != "" {
					color.Cyan("\tLabel: " + color.YellowString(result.Label))
				} else {
					color.Cyan("\tLabel: " + color.RedString("No label entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				if result.Type.Value != "" {
					color.Cyan("\tType: " + color.YellowString(result.Label))
					color.Cyan("\t  Value: " + color.YellowString(result.Type.Value))
					color.Cyan("\t  Label: " + color.YellowString(result.Type.Label))
				} else {
					color.Cyan("\tType: " + color.RedString("No type entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				if result.Speed.Value != 0 {
					color.Cyan("\tSpeed: ")
					color.Cyan("\t  Value: " + color.YellowString(strconv.Itoa(result.Speed.Value)))
					color.Cyan("\t  Label: " + color.YellowString(result.Speed.Label))
				} else {
					color.Cyan("\tSpeed: " + color.RedString("No speed entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString(result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tMarked Connected: " + color.YellowString("%v", result.MarkConnected))
				if result.Cable.Id != 0 {
					color.Cyan("\tCable: ")
					color.Cyan("\t  ID: " + color.RedString("%d", result.Cable.Id))
					color.Cyan("\t  URL: " + color.RedString("%s", result.Cable.Url))
					color.Cyan("\t  Display: " + color.RedString("%s", result.Cable.Display))
					color.Cyan("\t  Label: " + color.RedString("%s", result.Cable.Label))
				} else {
					color.Cyan("\tCable: " + color.RedString("No cable entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				if result.CableEnd != "" {
					color.Cyan("\tCable End: " + color.YellowString(result.CableEnd))
				} else {
					color.Cyan("\tCable End: " + color.RedString("No cable end entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				for _, link := range result.LinkPeers {
					if link != "" {
						color.Cyan("\tLink Peer: " + color.YellowString(link))
					} else {
						color.Cyan("\tLink Peer: " + color.RedString("No link peers entry for console server port: ") + color.YellowString("%s", result.Display))
					}
				}
				if result.LinkPeersType != "" {
					color.Cyan("\tLink Peers Type: " + color.YellowString(result.LinkPeersType))
				} else {
					color.Cyan("\tLink Peers Type: " + color.RedString("No link peers type entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				for _, endpoint := range result.ConnectedEndpoints {
					if endpoint != "" {
						color.Cyan("\tConnected Endpoints: " + color.YellowString(endpoint))
					} else {
						color.Cyan("\tConnected Endpoints: " + color.RedString("No connected endpoints entry for console server port: ") + color.YellowString("%s", result.Display))
					}
				}
				if result.ConnectedEndpointsType != "" {
					color.Cyan("\tConnected Endpoints Type: " + color.YellowString(result.ConnectedEndpointsType))
				} else {
					color.Cyan("\tConnected Endpoints Type: " + color.RedString("No connected endpoints type entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tConnected Endpoints Reachable: " + color.YellowString("%v", result.ConnectedEndpointsReachable))
				if result.Tags.Id != 0 {
					color.Cyan("\tTags: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.Tags.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.Tags.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.Tags.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.Tags.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", result.Tags.Slug))
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags entry for console server port: ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
				color.Cyan("\tOccupied: " + color.YellowString("%v", result.Occupied))
			}
			for responseObjectConsoleServerPorts.Next != nil {
				nextPageConsoleServerPorts()
			}
			if responseObjectConsoleServerPorts.Next == nil {
				display := color.HiGreenString("\tAll Netbox console server port objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n")
			}
		}
	},
}

func ApiConnectionNextPageConsoleServerPorts[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := *responseObjectConsolePorts.Next

	color.Yellow("\n  Getting Netbox API objects from %s\n", suffix)
	token := config.GetString("cmd.token_key")

	err := CheckSSL(rootURL)
	if err != nil {
		fmt.Println("  SSL certificate is not valid: ", err)
	} else {
		color.Cyan("  SSL certificate is valid for: " + color.YellowString("%s", rootURL))
	}

	executeAPIRequest(httpMethod, fullAPIPath, token, r)
	if err != nil {
		log.Fatalf("Error getting Netbox API objects: %s\n", err)
	}
}

func nextPageConsoleServerPorts() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\tDo you want to continue to the next page of console server port objects? ['Y' or 'yes'] or ['n' or 'no']: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageConsoleServerPorts(responseObjectConsoleServerPorts, "GET", *responseObjectConsoleServerPorts.Next)
		displayConsoleServerPortsOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("Invalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func displayConsoleServerPortsOutput() {
	for _, result := range responseObjectConsoleServerPorts.Results {
		display := fmt.Sprintf("    ABC Console Server Port: %s\n", color.YellowString(result.Display))
		equals := strings.Repeat("=", len(display))
		color.Cyan("\n  " + equals + "\n")
		color.Cyan(display)
		color.Cyan("  " + equals + "\n")
		color.Cyan("\tID: " + color.YellowString("%d", result.Id))
		color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
		color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
		if result.Device.Id != 0 {
			color.Cyan("\tDevice: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", result.Device.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", result.Device.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", result.Device.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", result.Device.Name))
		} else {
			color.Cyan("\tDevice: " + color.RedString("No device entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		if result.Module.Id != 0 {
			color.Cyan("\tModule: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", result.Module.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", result.Module.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", result.Module.Display))
			color.Cyan("\t  Module Bay: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.Module.ModuleBay.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.Module.ModuleBay.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.Module.ModuleBay.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.Module.ModuleBay.Name))
		} else {
			color.Cyan("\tModule: " + color.RedString("No module entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		if result.Name != "" {
			color.Cyan("\tName: " + color.YellowString(result.Name))
		} else {
			color.Cyan("\tName: " + color.RedString("No label entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		if result.Label != "" {
			color.Cyan("\tLabel: " + color.YellowString(result.Label))
		} else {
			color.Cyan("\tLabel: " + color.RedString("No label entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		if result.Type.Value != "" {
			color.Cyan("\tType: " + color.YellowString(result.Label))
			color.Cyan("\t  Value: " + color.YellowString(result.Type.Value))
			color.Cyan("\t  Label: " + color.YellowString(result.Type.Label))
		} else {
			color.Cyan("\tType: " + color.RedString("No type entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		if result.Speed.Value != 0 {
			color.Cyan("\tSpeed: ")
			color.Cyan("\t  Value: " + color.YellowString(strconv.Itoa(result.Speed.Value)))
			color.Cyan("\t  Label: " + color.YellowString(result.Speed.Label))
		} else {
			color.Cyan("\tSpeed: " + color.RedString("No speed entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		if result.Description != "" {
			color.Cyan("\tDescription: " + color.YellowString(result.Description))
		} else {
			color.Cyan("\tDescription: " + color.RedString("No description entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		color.Cyan("\tMarked Connected: " + color.YellowString("%v", result.MarkConnected))
		if result.Cable.Id != 0 {
			color.Cyan("\tCable: ")
			color.Cyan("\t  ID: " + color.RedString("%d", result.Cable.Id))
			color.Cyan("\t  URL: " + color.RedString("%s", result.Cable.Url))
			color.Cyan("\t  Display: " + color.RedString("%s", result.Cable.Display))
			color.Cyan("\t  Label: " + color.RedString("%s", result.Cable.Label))
		} else {
			color.Cyan("\tCable: " + color.RedString("No cable entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		if result.CableEnd != "" {
			color.Cyan("\tCable End: " + color.YellowString(result.CableEnd))
		} else {
			color.Cyan("\tCable End: " + color.RedString("No cable end entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		for _, link := range result.LinkPeers {
			if link != "" {
				color.Cyan("\tLink Peer: " + color.YellowString(link))
			} else {
				color.Cyan("\tLink Peer: " + color.RedString("No link peers entry for console server port: ") + color.YellowString("%s", result.Display))
			}
		}
		if result.LinkPeersType != "" {
			color.Cyan("\tLink Peers Type: " + color.YellowString(result.LinkPeersType))
		} else {
			color.Cyan("\tLink Peers Type: " + color.RedString("No link peers type entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		for _, endpoint := range result.ConnectedEndpoints {
			if endpoint != "" {
				color.Cyan("\tConnected Endpoints: " + color.YellowString(endpoint))
			} else {
				color.Cyan("\tConnected Endpoints: " + color.RedString("No connected endpoints entry for console server port: ") + color.YellowString("%s", result.Display))
			}
		}
		if result.ConnectedEndpointsType != "" {
			color.Cyan("\tConnected Endpoints Type: " + color.YellowString(result.ConnectedEndpointsType))
		} else {
			color.Cyan("\tConnected Endpoints Type: " + color.RedString("No connected endpoints type entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		color.Cyan("\tConnected Endpoints Reachable: " + color.YellowString("%v", result.ConnectedEndpointsReachable))
		if result.Tags.Id != 0 {
			color.Cyan("\tTags: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", result.Tags.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", result.Tags.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", result.Tags.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", result.Tags.Name))
			color.Cyan("\t  Slug: " + color.YellowString("%s", result.Tags.Slug))
		} else {
			color.Cyan("\tTags: " + color.RedString("No tags entry for console server port: ") + color.YellowString("%s", result.Display))
		}
		color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
		color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
		color.Cyan("\tOccupied: " + color.YellowString("%v", result.Occupied))
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimConsoleServerPortsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimConsoleServerPortsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimConsoleServerPortsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getConsoleServerPortsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getConsoleServerPortsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

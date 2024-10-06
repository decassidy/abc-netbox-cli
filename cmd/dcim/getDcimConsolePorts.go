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
	"strings"
)

type consolePorts struct {
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
			ValueLabel
		} `json:"speed"`
		Description   string `json:"description"`
		MarkConnected bool   `json:"mark_connected"`
		Cable         struct {
			Id      int    `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Label   string `json:"label"`
		} `json:"cable"`
		CableEnd                    string   `json:"cable_end"`
		LinkPeers                   []string `json:"link_peers"`
		LinkPeersType               string   `json:"link_peers_type"`
		ConnectedEndpoints          []string `json:"connected_endpoints"`
		ConnectedEndpointsType      string   `json:"connected_endpoints_type"`
		ConnectedEndpointsReachable bool     `json:"connected_endpoints_reachable"`
		Tags                        []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created     string `json:"created"`
		LastUpdated string `json:"last_updated"`
		Occupied    bool   `json:"_occupied"`
	} `json:"results"`
}

var responseObjectConsolePorts = new(consolePorts)

// GetDcimConsolePortsCmd represents the getDcimConsolePorts command
var GetDcimConsolePortsCmd = &cobra.Command{
	Use:   "getDcimConsolePorts",
	Short: "GET a list of console port objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of console port objects`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectConsolePorts, "GET", "cmd.dcim.dcim_api_url.console_ports")

		if responseObjectConsolePorts.Count > 0 {
			color.Cyan("\n  Total ABC Console Ports: "+color.YellowString("%v"), responseObjectConsolePorts.Count)
			for _, result := range responseObjectConsolePorts.Results {
				display := fmt.Sprintf("    ABC Console Port Name: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: %d\n", result.Id)
				color.Cyan("\tURL: %s\n", result.Url)
				color.Cyan("\tDisplay: %s\n", result.Display)
				if result.Device.Id > 0 {
					color.Cyan("\tDevice: ")
					color.Cyan("\t  ID: %d\n", result.Device.Id)
					color.Cyan("\t  URL: %s\n", result.Device.Url)
					color.Cyan("\t  Display: %s\n", result.Device.Display)
					color.Cyan("\t  Name: %s\n", result.Device.Name)
				} else {
					color.Cyan("\tDevice: " + color.RedString("No device entry found for console port name: "+color.YellowString(result.Display)))
				}
				if result.Module.Id > 0 {
					color.Cyan("\tModule: ")
					color.Cyan("\t  ID: %d\n", result.Module.Id)
					color.Cyan("\t  URL: %s\n", result.Module.Url)
					color.Cyan("\t  Display: %s\n", result.Module.Display)
					color.Cyan("\t  Device: %d\n", result.Module.Device)
					if result.Module.ModuleBay.Id > 0 {
						color.Cyan("\t  Module Bay: ")
						color.Cyan("\t    ID: %d\n", result.Module.ModuleBay.Id)
						color.Cyan("\t    Url: %s\n", result.Module.ModuleBay.Url)
						color.Cyan("\t    Display: %s\n", result.Module.ModuleBay.Display)
						color.Cyan("\t    Name: %s\n", result.Module.ModuleBay.Name)
					} else {
						color.Cyan("\t  Module Bay: " + color.RedString("No module bay entry found for console port name: "+color.YellowString(result.Display)))
					}
				} else {
					color.Cyan("\tModule: " + color.RedString("No module entry found for console port name: "+color.YellowString(result.Display)))
				}
				if result.Name != "" {
					color.Cyan("\tName: %s\n", result.Name)
				} else {
					color.Cyan("\tName: " + color.RedString("No name entry found for console port name: "+color.YellowString(result.Display)))
				}
				if result.Label != "" {
					color.Cyan("\tLabel: %s\n", result.Label)
				} else {
					color.Cyan("\tLabel: " + color.RedString("No label entry found for console port name: "+color.YellowString(result.Display)))
				}
				if result.Type.Value != "" {
					color.Cyan("\tType: ")
					color.Cyan("\t  Value: %d\n", result.Type.Value)
					color.Cyan("\t  Label: %d\n", result.Type.Label)
				} else {
					color.Cyan("\tType: " + color.RedString("No type entry found for console port name: "+color.YellowString(result.Display)))
				}
				if result.Speed.Value != "" {
					color.Cyan("\tSpeed: ")
					color.Cyan("\t  Value: %d\n", result.Speed.Value)
					color.Cyan("\t  Label: %d\n", result.Speed.Label)
				} else {
					color.Cyan("\tSpeed: " + color.RedString("No speed entry found for console port name: "+color.YellowString(result.Display)))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: %s\n", result.Description)
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for console port name: "+color.YellowString(result.Display)))
				}
				color.Cyan("\tMarked Connected: %t\n", result.MarkConnected)
				if result.Cable.Id > 0 {
					color.Cyan("\tCable: ")
					color.Cyan("\t  ID: %d\n", result.Cable.Id)
					color.Cyan("\t  URL: %s\n", result.Cable.Url)
					color.Cyan("\t  Display: %s\n", result.Cable.Display)
					color.Cyan("\t  Label: %s\n", result.Cable.Label)
				} else {
					color.Cyan("\tCable: " + color.RedString("No cable entry found for console port name: "+color.YellowString(result.Display)))
				}
				if result.CableEnd != "" {
					color.Cyan("\tCable End: %s\n", result.CableEnd)
				} else {
					color.Cyan("\tCable End: " + color.RedString("No cable end entry found for console port name: "+color.YellowString(result.Display)))
				}
				for _, link := range result.LinkPeers {
					if link != "" {
						color.Cyan("\tLink Peer: %s\n", link)
					} else {
						color.Cyan("\tLink Peers: " + color.RedString("No link peer entry found for console port name: "+color.YellowString(result.Display)))
					}
				}
				if result.LinkPeersType != "" {
					color.Cyan("\tLink Peers Type: %s\n", result.LinkPeersType)
				} else {
					color.Cyan("\tLink Peers Type: " + color.RedString("No link peers type entry found for console port name: "+color.YellowString(result.Display)))
				}
				for _, endpoint := range result.ConnectedEndpoints {
					if endpoint != "" {
						color.Cyan("\tConnected Endpoints: %s\n", endpoint)
					} else {
						color.Cyan("\tConnected Endpoints: " + color.RedString("No connected endpoint entry found for console port name: "+color.YellowString(result.Display)))
					}
				}
				if result.ConnectedEndpointsType != "" {
					color.Cyan("\tConnected Endpoint Type: %s\n", result.ConnectedEndpointsType)
				} else {
					color.Cyan("\tConnected Endpoint Type: " + color.RedString("No connected endpoint type entry found for console port name: "+color.YellowString(result.Display)))
				}
				color.Cyan("\tConnected Endpoints Reachable: %t\n", result.ConnectedEndpointsReachable)

				for _, tag := range result.Tags {
					if tag.Id > 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
						color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
						color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
						color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
						color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
						color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
					} else {
						color.Cyan("\tTags: " + color.RedString("No Tags entry found for console port name: ") + color.YellowString("%s", result.Display))
					}
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
				color.Cyan("\tOccupied: " + color.YellowString("%t\n", result.Occupied))
			}
			for responseObjectConsolePorts.Next != nil {
				nextPageConsolePorts()
			}
			if responseObjectConsolePorts.Next == nil {
				display := color.HiGreenString("\tAll Netbox console port objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n")
			}
		} else {
			color.Cyan("  ABC Console Ports: " + color.RedString("No console ports found on server. Exiting...\n"))
		}
	},
}

func ApiConnectionNextPageConsolePorts[T anyStruct](r T, httpMethod string, suffix string) {
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

func nextPageConsolePorts() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\tDo you want to continue to the next page of console poert objects? ['Y' or 'yes'] or ['n' or 'no']: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageConsolePorts(responseObjectConsolePorts, "GET", *responseObjectConsolePorts.Next)
		displayConsolePortsOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("Invalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func displayConsolePortsOutput() {
	for _, result := range responseObjectConsolePorts.Results {
		display := fmt.Sprintf("    ABC Console Port Name: %s\n", color.YellowString(result.Display))
		equals := strings.Repeat("=", len(display))
		color.Cyan("\n  " + equals + "\n")
		color.Cyan(display)
		color.Cyan("  " + equals + "\n")
		color.Cyan("\tID: %d\n", result.Id)
		color.Cyan("\tURL: %s\n", result.Url)
		color.Cyan("\tDisplay: %s\n", result.Display)
		if result.Device.Id > 0 {
			color.Cyan("\tDevice: ")
			color.Cyan("\t  ID: %d\n", result.Device.Id)
			color.Cyan("\t  URL: %s\n", result.Device.Url)
			color.Cyan("\t  Display: %s\n", result.Device.Display)
			color.Cyan("\t  Name: %s\n", result.Device.Name)
		} else {
			color.Cyan("\tDevice: " + color.RedString("No device entry found for console port name: "+color.YellowString(result.Display)))
		}
		if result.Module.Id > 0 {
			color.Cyan("\tModule: ")
			color.Cyan("\t  ID: %d\n", result.Module.Id)
			color.Cyan("\t  URL: %s\n", result.Module.Url)
			color.Cyan("\t  Display: %s\n", result.Module.Display)
			color.Cyan("\t  Device: %d\n", result.Module.Device)
			if result.Module.ModuleBay.Id > 0 {
				color.Cyan("\t  Module Bay: ")
				color.Cyan("\t    ID: %d\n", result.Module.ModuleBay.Id)
				color.Cyan("\t    Url: %s\n", result.Module.ModuleBay.Url)
				color.Cyan("\t    Display: %s\n", result.Module.ModuleBay.Display)
				color.Cyan("\t    Name: %s\n", result.Module.ModuleBay.Name)
			} else {
				color.Cyan("\t  Module Bay: " + color.RedString("No module bay entry found for console port name: "+color.YellowString(result.Display)))
			}
		} else {
			color.Cyan("\tModule: " + color.RedString("No module entry found for console port name: "+color.YellowString(result.Display)))
		}
		if result.Name != "" {
			color.Cyan("\tName: %s\n", result.Name)
		} else {
			color.Cyan("\tName: " + color.RedString("No name entry found for console port name: "+color.YellowString(result.Display)))
		}
		if result.Label != "" {
			color.Cyan("\tLabel: %s\n", result.Label)
		} else {
			color.Cyan("\tLabel: " + color.RedString("No label entry found for console port name: "+color.YellowString(result.Display)))
		}
		if result.Type.Value != "" {
			color.Cyan("\tType: ")
			color.Cyan("\t  Value: %d\n", result.Type.Value)
			color.Cyan("\t  Label: %d\n", result.Type.Label)
		} else {
			color.Cyan("\tType: " + color.RedString("No type entry found for console port name: "+color.YellowString(result.Display)))
		}
		if result.Speed.Value != "" {
			color.Cyan("\tSpeed: ")
			color.Cyan("\t  Value: %d\n", result.Speed.Value)
			color.Cyan("\t  Label: %d\n", result.Speed.Label)
		} else {
			color.Cyan("\tSpeed: " + color.RedString("No speed entry found for console port name: "+color.YellowString(result.Display)))
		}
		if result.Description != "" {
			color.Cyan("\tDescription: %s\n", result.Description)
		} else {
			color.Cyan("\tDescription: " + color.RedString("No description entry found for console port name: "+color.YellowString(result.Display)))
		}
		color.Cyan("\tMarked Connected: %t\n", result.MarkConnected)
		if result.Cable.Id > 0 {
			color.Cyan("\tCable: ")
			color.Cyan("\t  ID: %d\n", result.Cable.Id)
			color.Cyan("\t  URL: %s\n", result.Cable.Url)
			color.Cyan("\t  Display: %s\n", result.Cable.Display)
			color.Cyan("\t  Label: %s\n", result.Cable.Label)
		} else {
			color.Cyan("\tCable: " + color.RedString("No cable entry found for console port name: "+color.YellowString(result.Display)))
		}
		if result.CableEnd != "" {
			color.Cyan("\tCable End: %s\n", result.CableEnd)
		} else {
			color.Cyan("\tCable End: " + color.RedString("No cable end entry found for console port name: "+color.YellowString(result.Display)))
		}
		for _, link := range result.LinkPeers {
			if link != "" {
				color.Cyan("\tLink Peer: %s\n", link)
			} else {
				color.Cyan("\tLink Peers: " + color.RedString("No link peer entry found for console port name: "+color.YellowString(result.Display)))
			}
		}
		if result.LinkPeersType != "" {
			color.Cyan("\tLink Peers Type: %s\n", result.LinkPeersType)
		} else {
			color.Cyan("\tLink Peers Type: " + color.RedString("No link peers type entry found for console port name: "+color.YellowString(result.Display)))
		}
		for _, endpoint := range result.ConnectedEndpoints {
			if endpoint != "" {
				color.Cyan("\tConnected Endpoints: %s\n", endpoint)
			} else {
				color.Cyan("\tConnected Endpoints: " + color.RedString("No connected endpoint entry found for console port name: "+color.YellowString(result.Display)))
			}
		}
		if result.ConnectedEndpointsType != "" {
			color.Cyan("\tConnected Endpoint Type: %s\n", result.ConnectedEndpointsType)
		} else {
			color.Cyan("\tConnected Endpoint Type: " + color.RedString("No connected endpoint type entry found for console port name: "+color.YellowString(result.Display)))
		}
		color.Cyan("\tConnected Endpoints Reachable: %t\n", result.ConnectedEndpointsReachable)

		for _, tag := range result.Tags {
			if tag.Id > 0 {
				color.Cyan("\tTags: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
				color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
			} else {
				color.Cyan("\tTags: " + color.RedString("No Tags entry found for console port name: ") + color.YellowString("%s", result.Display))
			}
		}
		color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
		color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
		color.Cyan("\tOccupied: " + color.YellowString("%t\n", result.Occupied))
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimConsolePortsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimConsolePortsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimConsolePortsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimConsolePortsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimConsolePortsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

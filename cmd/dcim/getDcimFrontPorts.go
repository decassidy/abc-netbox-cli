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

type frontPorts struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous string  `json:"previous"`
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
		Color    string `json:"color"`
		RearPort struct {
			CommonFieldsNoSlug
			Label       string `json:"label"`
			Description string `json:"description"`
		} `json:"rear_port"`
		RearPortPosition int    `json:"rear_port_position"`
		Description      string `json:"description"`
		MarkConnected    bool   `json:"mark_connected"`
		Cable            struct {
			Id      int    `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Label   string `json:"label"`
		} `json:"cable"`
		CableEnd      string   `json:"cable_end"`
		LinkPeers     []string `json:"link_peers"`
		LinkPeersType string   `json:"link_peers_type"`
		Tags          []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created     string `json:"created"`
		LastUpdated string `json:"last_updated"`
		Occupied    bool   `json:"_occupied"`
	} `json:"results"`
}

var responseObjectFrontPorts = new(frontPorts)

// GetDcimFrontPortsCmd represents the getDcimFrontPorts command
var GetDcimFrontPortsCmd = &cobra.Command{
	Use:   "getDcimFrontPorts",
	Short: "GET a list of front port objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of front port objects`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectFrontPorts, "GET", "cmd.dcim.dcim_api_url.front_ports")

		if responseObjectFrontPorts.Count != 0 {
			color.Cyan("Total ABC Front Ports: "+color.YellowString("%v"), responseObjectFrontPorts.Count)
			for _, result := range responseObjectFrontPorts.Results {
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
			for responseObjectFrontPorts.Next != nil {
				nextPageFrontPorts()
			}
			if responseObjectFrontPorts.Next == nil {
				display := color.HiGreenString("\tAll Netbox front port objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n")
			}
		} else {
			color.Cyan("Total ABC Front Ports: " + color.RedString("No front ports found on the server. Exiting...\n"))
		}
	},
}

func nextPageFrontPorts() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\tDo you want to continue to the next page of device objects? [yes/no]: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageFrontPorts(responseObjectFrontPorts, "GET", *responseObjectFrontPorts.Next)
		displayFrontPortsOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("\nInvalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func ApiConnectionNextPageFrontPorts[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := *responseObjectFrontPorts.Next

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

func displayFrontPortsOutput() {
	for _, result := range responseObjectFrontPorts.Results {
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
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimFrontPortsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimFrontPortsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimFrontPortsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimFrontPortsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

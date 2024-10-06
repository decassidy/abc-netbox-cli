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

type frontPortTemplates struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Id         int    `json:"id"`
		Url        string `json:"url"`
		Display    string `json:"display"`
		DeviceType struct {
			Id           int    `json:"id"`
			Url          string `json:"url"`
			Display      string `json:"display"`
			Manufacturer struct {
				CommonFieldsSlug
			} `json:"manufacturer"`
			Model string `json:"model"`
			Slug  string `json:"slug"`
		} `json:"device_type"`
		ModuleType struct {
			Id           int    `json:"id"`
			Url          string `json:"url"`
			Display      string `json:"display"`
			Manufacturer struct {
				CommonFieldsSlug
			} `json:"manufacturer"`
			Model string `json:"model"`
		} `json:"module_type"`
		Name  string `json:"name"`
		Label string `json:"label"`
		Type  struct {
			ValueLabel
		} `json:"type"`
		Color    string `json:"color"`
		RearPort struct {
			CommonFieldsNoSlug
		} `json:"rear_port"`
		RearPortPosition int    `json:"rear_port_position"`
		Description      string `json:"description"`
		Created          string `json:"created"`
		LastUpdated      string `json:"last_updated"`
	} `json:"results"`
}

var responseObjectFrontPortTemplates = new(frontPortTemplates)

// GetDcimFrontPortTemplatesCmd represents the getDcimFrontPortTemplates command
var GetDcimFrontPortTemplatesCmd = &cobra.Command{
	Use:   "getDcimFrontPortTemplates",
	Short: "GET a list of front port template objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of front port template objects`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectFrontPortTemplates, "GET", "cmd.dcim.dcim_api_url.front_port_templates")

		if responseObjectFrontPortTemplates.Count > 0 {
			color.Cyan("Total ABC Front Port Templates: "+color.YellowString("%d"), responseObjectFrontPortTemplates.Count)
			for _, result := range responseObjectFrontPortTemplates.Results {
				display := fmt.Sprintf("    ABC Front Port Template Name: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				color.Cyan("\tDevice Type: ")
				color.Cyan("\t  ID: " + color.YellowString("%s", result.DeviceType.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.DeviceType.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.DeviceType.Display))
				color.Cyan("\t  Manufacturer: ")
				color.Cyan("\t    ID: " + color.YellowString("%s", result.DeviceType.Manufacturer.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", result.DeviceType.Manufacturer.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", result.DeviceType.Manufacturer.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", result.DeviceType.Manufacturer.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", result.DeviceType.Manufacturer.Slug))
				color.Cyan("\t  Model: " + color.YellowString("%s", result.DeviceType.Model))
				color.Cyan("\t  Slug: " + color.YellowString("%s", result.DeviceType.Slug))
				color.Cyan("\tModule Type: ")
				color.Cyan("\t  ID: " + color.YellowString("%s", result.ModuleType.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.ModuleType.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.ModuleType.Display))
				color.Cyan("\t  Manufacturer: ")
				color.Cyan("\t    ID: " + color.YellowString("%s", result.ModuleType.Manufacturer.Id))
				color.Cyan("\t    URL: " + color.YellowString("%s", result.ModuleType.Manufacturer.Url))
				color.Cyan("\t    Display: " + color.YellowString("%s", result.ModuleType.Manufacturer.Display))
				color.Cyan("\t    Name: " + color.YellowString("%s", result.ModuleType.Manufacturer.Name))
				color.Cyan("\t    Slug: " + color.YellowString("%s", result.ModuleType.Manufacturer.Slug))
				color.Cyan("\t  Model: " + color.YellowString("%s", result.ModuleType.Model))
				color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				color.Cyan("\tLabel: " + color.YellowString("%s", result.Label))
				color.Cyan("\tType: ")
				color.Cyan("\t  Value: " + color.YellowString("%s", result.Type.Value))
				color.Cyan("\t  Label: " + color.YellowString("%s", result.Type.Label))
				color.Cyan("\tColor: " + color.YellowString("%s", result.Color))
				color.Cyan("\tRear Port: " + color.YellowString("%s", result.RearPort))
				color.Cyan("\t  ID: " + color.YellowString("%s", result.RearPort.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.RearPort.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.RearPort.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", result.RearPort.Name))
				color.Cyan("\tRear Port Position: " + color.YellowString("%d", result.RearPortPosition))
				color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
			}
			for responseObjectFrontPortTemplates.Next != nil {
				nextPageFrontPortTemplates()
			}
			if responseObjectFrontPorts.Next == nil {
				display := color.HiGreenString("\tAll Netbox front port template objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n")
			}
		} else {
			color.Cyan("  Total ABC Front Port Templates: " + color.RedString("No front port templates found on the server. Exiting...\n"))
		}
	},
}

func nextPageFrontPortTemplates() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\tDo you want to continue to the next page of front port template objects? [yes/no]: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageFrontPortTemplates(responseObjectFrontPortTemplates, "GET", *responseObjectFrontPortTemplates.Next)
		displayFrontPortTemplatesOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("\nInvalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func ApiConnectionNextPageFrontPortTemplates[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := *responseObjectFrontPortTemplates.Next

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

func displayFrontPortTemplatesOutput() {
	for _, result := range responseObjectFrontPortTemplates.Results {
		display := fmt.Sprintf("    ABC Front Port Template Name: %s\n", color.YellowString(result.Display))
		equals := strings.Repeat("=", len(display))
		color.Cyan("\n  " + equals + "\n")
		color.Cyan(display)
		color.Cyan("  " + equals + "\n")
		color.Cyan("\tID: " + color.YellowString("%d", result.Id))
		color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
		color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
		color.Cyan("\tDevice Type: ")
		color.Cyan("\t  ID: " + color.YellowString("%s", result.DeviceType.Id))
		color.Cyan("\t  URL: " + color.YellowString("%s", result.DeviceType.Url))
		color.Cyan("\t  Display: " + color.YellowString("%s", result.DeviceType.Display))
		color.Cyan("\t  Manufacturer: ")
		color.Cyan("\t    ID: " + color.YellowString("%s", result.DeviceType.Manufacturer.Id))
		color.Cyan("\t    URL: " + color.YellowString("%s", result.DeviceType.Manufacturer.Url))
		color.Cyan("\t    Display: " + color.YellowString("%s", result.DeviceType.Manufacturer.Display))
		color.Cyan("\t    Name: " + color.YellowString("%s", result.DeviceType.Manufacturer.Name))
		color.Cyan("\t    Slug: " + color.YellowString("%s", result.DeviceType.Manufacturer.Slug))
		color.Cyan("\t  Model: " + color.YellowString("%s", result.DeviceType.Model))
		color.Cyan("\t  Slug: " + color.YellowString("%s", result.DeviceType.Slug))
		color.Cyan("\tModule Type: ")
		color.Cyan("\t  ID: " + color.YellowString("%s", result.ModuleType.Id))
		color.Cyan("\t  URL: " + color.YellowString("%s", result.ModuleType.Url))
		color.Cyan("\t  Display: " + color.YellowString("%s", result.ModuleType.Display))
		color.Cyan("\t  Manufacturer: ")
		color.Cyan("\t    ID: " + color.YellowString("%s", result.ModuleType.Manufacturer.Id))
		color.Cyan("\t    URL: " + color.YellowString("%s", result.ModuleType.Manufacturer.Url))
		color.Cyan("\t    Display: " + color.YellowString("%s", result.ModuleType.Manufacturer.Display))
		color.Cyan("\t    Name: " + color.YellowString("%s", result.ModuleType.Manufacturer.Name))
		color.Cyan("\t    Slug: " + color.YellowString("%s", result.ModuleType.Manufacturer.Slug))
		color.Cyan("\t  Model: " + color.YellowString("%s", result.ModuleType.Model))
		color.Cyan("\tName: " + color.YellowString("%s", result.Name))
		color.Cyan("\tLabel: " + color.YellowString("%s", result.Label))
		color.Cyan("\tType: ")
		color.Cyan("\t  Value: " + color.YellowString("%s", result.Type.Value))
		color.Cyan("\t  Label: " + color.YellowString("%s", result.Type.Label))
		color.Cyan("\tColor: " + color.YellowString("%s", result.Color))
		color.Cyan("\tRear Port: " + color.YellowString("%s", result.RearPort))
		color.Cyan("\t  ID: " + color.YellowString("%s", result.RearPort.Id))
		color.Cyan("\t  URL: " + color.YellowString("%s", result.RearPort.Url))
		color.Cyan("\t  Display: " + color.YellowString("%s", result.RearPort.Display))
		color.Cyan("\t  Name: " + color.YellowString("%s", result.RearPort.Name))
		color.Cyan("\tRear Port Position: " + color.YellowString("%d", result.RearPortPosition))
		color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
		color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
		color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimFrontPortTemplatesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimFrontPortTemplatesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimFrontPortTemplatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimFrontPortTemplatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

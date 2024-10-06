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

type deviceBayTemplates struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Id         int    `json:"id"`
		Url        string `json:"url"`
		Display    string `json:"display"`
		DeviceType struct {
			Id           int              `json:"id"`
			Url          string           `json:"url"`
			Display      string           `json:"display"`
			Manufacturer CommonFieldsSlug `json:"manufacturer"`
			Model        string           `json:"model"`
			Slug         string           `json:"slug"`
		} `json:"device_type"`
		Name        string `json:"name"`
		Label       string `json:"label"`
		Description string `json:"description"`
		Created     string `json:"created"`
		LastUpdated string `json:"last_updated"`
	} `json:"results"`
}

var responseObjectDeviceBayTemplates = new(deviceBayTemplates)

// GetDcimDeviceBayTemplatesCmd represents the getDcimDeviceBayTemplates command
var GetDcimDeviceBayTemplatesCmd = &cobra.Command{
	Use:   "getDcimDeviceBayTemplates",
	Short: "GET a list of device bay template objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of device bay template objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectDeviceBayTemplates, "GET", "cmd.dcim.dcim_api_url.device_bay_templates")

		if len(responseObjectDeviceBayTemplates.Results) == 0 {
			color.Cyan("  Total ABC Device Bay Templates: " + color.RedString("No device bay templates entries found on server. Exiting...\n"))
		} else {
			color.Cyan("\n  Total ABC Device Bay Templates: "+color.YellowString("%d\n"), responseObjectDeviceBayTemplates.Count)

			for _, result := range responseObjectDeviceBayTemplates.Results {
				display := fmt.Sprintf("    ABC Device Bay Template: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				if result.DeviceType.Id != 0 {
					color.Cyan("\tDevice Type: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.DeviceType.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.DeviceType.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.DeviceType.Display))
					color.Cyan("\t  Manufacturer: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", result.DeviceType.Manufacturer.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", result.DeviceType.Manufacturer.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", result.DeviceType.Manufacturer.Display))
					color.Cyan("\t    Name: " + color.YellowString("%s", result.DeviceType.Manufacturer.Name))
					color.Cyan("\t    Slug: " + color.YellowString("%s", result.DeviceType.Manufacturer.Slug))
				} else {
					color.Cyan("\t  Manufacturer: " + color.RedString("No manufacturer entry for device bay template: ") + color.YellowString("%s", result.Display))
				}
				if result.Name != "" {
					color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				} else {
					color.Cyan("\tName: " + color.RedString("No name entry for device bay template: ") + color.YellowString("%s", result.Display))
				}
				if result.Label != "" {
					color.Cyan("\tLabel: " + color.YellowString("%s", result.Label))
				} else {
					color.Cyan("\tLabel: " + color.RedString("No label entry for device bay template: ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry for device bay template: ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
			}
			for responseObjectDeviceBayTemplates.Next != nil {
				nextPageDeviceBayTemplates()
			}
			if responseObjectDeviceBayTemplates.Next == nil {
				display := color.HiGreenString("\tAll Netbox device bay template objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n\n")
			}
		}
	},
}

func ApiConnectionNextPageDeviceBayTemplates[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := *responseObjectDeviceBayTemplates.Next

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

func nextPageDeviceBayTemplates() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\tDo you want to continue to the next page of device bay objects? ['Y' or 'yes'] or ['n' or 'no']: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageDeviceBayTemplates(responseObjectDeviceBayTemplates, "GET", *responseObjectDeviceBayTemplates.Next)
		displayDeviceBayTemplatesOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("Invalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func displayDeviceBayTemplatesOutput() {
	for _, result := range responseObjectDeviceBayTemplates.Results {
		display := fmt.Sprintf("    ABC Device Bay Template: %s\n", color.YellowString(result.Display))
		equals := strings.Repeat("=", len(display))
		color.Cyan("\n  " + equals + "\n")
		color.Cyan(display)
		color.Cyan("  " + equals + "\n")
		color.Cyan("\tID: " + color.YellowString("%d", result.Id))
		color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
		color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
		if result.DeviceType.Id != 0 {
			color.Cyan("\tDevice Type: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", result.DeviceType.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", result.DeviceType.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", result.DeviceType.Display))
			color.Cyan("\t  Manufacturer: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", result.DeviceType.Manufacturer.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", result.DeviceType.Manufacturer.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", result.DeviceType.Manufacturer.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", result.DeviceType.Manufacturer.Name))
			color.Cyan("\t    Slug: " + color.YellowString("%s", result.DeviceType.Manufacturer.Slug))
		} else {
			color.Cyan("\t  Manufacturer: " + color.RedString("No manufacturer entry for device bay template: ") + color.YellowString("%s", result.Display))
		}
		if result.Name != "" {
			color.Cyan("\tName: " + color.YellowString("%s", result.Name))
		} else {
			color.Cyan("\tName: " + color.RedString("No name entry for device bay template: ") + color.YellowString("%s", result.Display))
		}
		if result.Label != "" {
			color.Cyan("\tLabel: " + color.YellowString("%s", result.Label))
		} else {
			color.Cyan("\tLabel: " + color.RedString("No label entry for device bay template: ") + color.YellowString("%s", result.Display))
		}
		if result.Description != "" {
			color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
		} else {
			color.Cyan("\tDescription: " + color.RedString("No description entry for device bay template: ") + color.YellowString("%s", result.Display))
		}
		color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
		color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimDeviceBayTemplatesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimDeviceBayTemplatesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimDeviceBayTemplatesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimDeviceBayTemplatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimDeviceBayTemplatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

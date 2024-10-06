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

type consolePortTemplates struct {
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
		ModuleType struct {
			Id           int              `json:"id"`
			Url          string           `json:"url"`
			Display      string           `json:"display"`
			Manufacturer CommonFieldsSlug `json:"manufacturer"`
			Model        string           `json:"model"`
		} `json:"module_type"`
		Name        string     `json:"name"`
		Label       string     `json:"label"`
		Type        ValueLabel `json:"type"`
		Description string     `json:"description"`
		Created     string     `json:"created"`
		LastUpdated string     `json:"last_updated"`
	} `json:"results"`
}

var responseObjectConsolePortTemplates = new(consolePortTemplates)

// GetDcimConsolePortTemplatesCmd represents the getDcimConsolePortTemplates command
var GetDcimConsolePortTemplatesCmd = &cobra.Command{
	Use:   "getDcimConsolePortTemplates",
	Short: "GET a list of console port template objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of console port template objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectConsolePortTemplates, "GET", "cmd.dcim.dcim_api_url.console_port_templates")

		if responseObject.Count > 0 {
			color.Cyan("\nTotal ABC Console Port Templates: "+color.YellowString("%v"), responseObjectConsolePortTemplates.Count)
			for _, result := range responseObjectConsolePortTemplates.Results {
				display := fmt.Sprintf("    ABC Console Port Template: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: %d", result.Id)
				color.Cyan("\tURL: %s", result.Url)
				color.Cyan("\tDisplay: %s", result.Display)
				if result.DeviceType.Id != 0 {
					color.Cyan("\tDevice Type: ")
					color.Cyan("\t  ID: %d", result.DeviceType.Id)
					color.Cyan("\t  URL: %s", result.DeviceType.Url)
					color.Cyan("\t  Display: %s", result.DeviceType.Display)
					color.Cyan("\t  Manufacturer: ")
					color.Cyan("\t    ID: %d", result.DeviceType.Manufacturer.Id)
					color.Cyan("\t    URL: %s", result.DeviceType.Manufacturer.Url)
					color.Cyan("\t    Display: %s", result.DeviceType.Manufacturer.Display)
					color.Cyan("\t    Name: %s", result.DeviceType.Manufacturer.Name)
					color.Cyan("\t    Slug: %s", result.DeviceType.Manufacturer.Slug)
					color.Cyan("\t  Model: %s", result.DeviceType.Model)
					color.Cyan("\t  Slug: %s", result.DeviceType.Slug)
				} else {
					color.Cyan("\tDevice Type: " + color.RedString("No device type entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
				}
				if result.ModuleType.Id != 0 {
					color.Cyan("\tModule Type: ")
					color.Cyan("\t  ID: %d", result.ModuleType.Id)
					color.Cyan("\t  URL: %s", result.ModuleType.Url)
					color.Cyan("\t  Display: %s", result.ModuleType.Display)
					color.Cyan("\t  Manufacturer: ")
					color.Cyan("\t    ID: %d", result.ModuleType.Manufacturer.Id)
					color.Cyan("\t    URL: %s", result.ModuleType.Manufacturer.Url)
					color.Cyan("\t    Display: %s", result.ModuleType.Manufacturer.Display)
					color.Cyan("\t    Name: %s", result.ModuleType.Manufacturer.Name)
					color.Cyan("\t    Slug: %s", result.ModuleType.Manufacturer.Slug)
					color.Cyan("\t  Model: %s", result.ModuleType.Model)
				} else {
					color.Cyan("\tModule Type: " + color.RedString("No model type entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
				}
				if result.Name != "" {
					color.Cyan("\tName: %s", result.Name)
				} else {
					color.Cyan("\tName: " + color.RedString("No name entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
				}
				if result.Label != "" {
					color.Cyan("\tLabel: %s", result.Label)
				} else {
					color.Cyan("\tLabel: " + color.RedString("No label entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
				}
				if result.Type.Value != "" {
					color.Cyan("\tType: %s", result.Type.Value)
					color.Cyan("\t  Value: %s", result.Type.Value)
					color.Cyan("\t  Label: %s", result.Type.Label)
				} else {
					color.Cyan("\tType: " + color.RedString("No type entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: %s", result.Description)
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
				}
				color.Cyan("\tCreated: %s", result.Created)
				color.Cyan("\tCreated: %s\n", result.LastUpdated)
			}
			for responseObjectConsolePortTemplates.Next != nil {
				nextPageConsolePortTemplates()
			}
			if responseObjectConsolePortTemplates.Next == nil {
				display := color.HiGreenString("\tAll Netbox console port template objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n")
			}
		} else {
			color.Cyan("  Total ABC Console Port Templates: " + color.RedString("No console port templates entries found on server. Exiting...\n"))
		}
	},
}

func ApiConnectionNextPageConsolePortTemplates[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := *responseObjectConsolePortTemplates.Next

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

func nextPageConsolePortTemplates() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\tDo you want to continue to the next page of console port template objects? ['Y' or 'yes'] or ['n' or 'no']: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageConsolePortTemplates(responseObjectConsolePortTemplates, "GET", *responseObjectConsolePortTemplates.Next)
		displayConsolePortTemplatesOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("Invalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func displayConsolePortTemplatesOutput() {
	for _, result := range responseObjectConsolePortTemplates.Results {
		display := fmt.Sprintf("    ABC Console Port Template: %s\n", color.YellowString(result.Display))
		equals := strings.Repeat("=", len(display))
		color.Cyan("\n  " + equals + "\n")
		color.Cyan(display)
		color.Cyan("  " + equals + "\n")
		color.Cyan("\tID: %d", result.Id)
		color.Cyan("\tURL: %s", result.Url)
		color.Cyan("\tDisplay: %s", result.Display)
		if result.DeviceType.Id != 0 {
			color.Cyan("\tDevice Type: ")
			color.Cyan("\t  ID: %d", result.DeviceType.Id)
			color.Cyan("\t  URL: %s", result.DeviceType.Url)
			color.Cyan("\t  Display: %s", result.DeviceType.Display)
			color.Cyan("\t  Manufacturer: ")
			color.Cyan("\t    ID: %d", result.DeviceType.Manufacturer.Id)
			color.Cyan("\t    URL: %s", result.DeviceType.Manufacturer.Url)
			color.Cyan("\t    Display: %s", result.DeviceType.Manufacturer.Display)
			color.Cyan("\t    Name: %s", result.DeviceType.Manufacturer.Name)
			color.Cyan("\t    Slug: %s", result.DeviceType.Manufacturer.Slug)
			color.Cyan("\t  Model: %s", result.DeviceType.Model)
			color.Cyan("\t  Slug: %s", result.DeviceType.Slug)
		} else {
			color.Cyan("\tDevice Type: " + color.RedString("No device type entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
		}
		if result.ModuleType.Id != 0 {
			color.Cyan("\tModule Type: ")
			color.Cyan("\t  ID: %d", result.ModuleType.Id)
			color.Cyan("\t  URL: %s", result.ModuleType.Url)
			color.Cyan("\t  Display: %s", result.ModuleType.Display)
			color.Cyan("\t  Manufacturer: ")
			color.Cyan("\t    ID: %d", result.ModuleType.Manufacturer.Id)
			color.Cyan("\t    URL: %s", result.ModuleType.Manufacturer.Url)
			color.Cyan("\t    Display: %s", result.ModuleType.Manufacturer.Display)
			color.Cyan("\t    Name: %s", result.ModuleType.Manufacturer.Name)
			color.Cyan("\t    Slug: %s", result.ModuleType.Manufacturer.Slug)
			color.Cyan("\t  Model: %s", result.ModuleType.Model)
		} else {
			color.Cyan("\tModule Type: " + color.RedString("No model type entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
		}
		if result.Name != "" {
			color.Cyan("\tName: %s", result.Name)
		} else {
			color.Cyan("\tName: " + color.RedString("No name entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
		}
		if result.Label != "" {
			color.Cyan("\tLabel: %s", result.Label)
		} else {
			color.Cyan("\tLabel: " + color.RedString("No label entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
		}
		if result.Type.Value != "" {
			color.Cyan("\tType: %s", result.Type.Value)
			color.Cyan("\t  Value: %s", result.Type.Value)
			color.Cyan("\t  Label: %s", result.Type.Label)
		} else {
			color.Cyan("\tType: " + color.RedString("No type entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
		}
		if result.Description != "" {
			color.Cyan("\tDescription: %s", result.Description)
		} else {
			color.Cyan("\tDescription: " + color.RedString("No description entry found for: ") + color.YellowString("%s", result.DeviceType.Display))
		}
		color.Cyan("\tCreated: %s", result.Created)
		color.Cyan("\tCreated: %s\n", result.LastUpdated)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimConsolePortTemplatesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimConsolePortTemplatesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimConsolePortTemplatesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimConsolePortTemplatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimConsolePortTemplatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

type deviceRoles struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		CommonFieldsSlug
		Color          string `json:"color"`
		VmRole         bool   `json:"vm_role"`
		ConfigTemplate struct {
			CommonFieldsNoSlug
		} `json:"config_template"`
		Description         string             `json:"description"`
		Tags                []CommonFieldsSlug `json:"tags"`
		Created             string             `json:"created"`
		LastUpdated         string             `json:"last_updated"`
		DeviceCount         int                `json:"device_count"`
		VirtualmachineCount int                `json:"virtualmachine_count"`
	} `json:"results"`
}

var responseObjectDeviceRoles = new(deviceRoles)

// GetDcimDeviceRolesCmd represents the getDcimDeviceRoles command
var GetDcimDeviceRolesCmd = &cobra.Command{
	Use:   "getDcimDeviceRoles",
	Short: "GET a list of device role objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of device role objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectDeviceRoles, "GET", "cmd.dcim.dcim_api_url.device_roles")

		if len(responseObjectDeviceRoles.Results) == 0 {
			color.Cyan("\n  Total ABC Device Roles: " + color.RedString("No device roles entries found on server. Exiting...\n"))
		} else {
			color.Cyan("\n  Total ABC Device Roles: "+color.YellowString("%d\n"), responseObjectDeviceRoles.Count)

			for _, result := range responseObjectDeviceRoles.Results {
				display := fmt.Sprintf("    ABC Device Role: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				color.Cyan("\tSlug: " + color.YellowString("%s", result.Slug))
				if result.Color != "" {
					color.Cyan("\tColor: " + color.YellowString("%s", result.Color))
				} else {
					color.Cyan("\tColor: " + color.RedString("No color entry for device role: ") + color.YellowString("%s", result.Display))
				}
				color.Cyan("\tVM Role: " + color.YellowString("%v", result.VmRole))
				if result.ConfigTemplate.Id != 0 {
					color.Cyan("\tConfig Template: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.ConfigTemplate.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.ConfigTemplate.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.ConfigTemplate.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.ConfigTemplate.Name))
				} else {
					color.Cyan("\tConfig Template: " + color.RedString("No config template entry for device role: ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry for device role: ") + color.YellowString("%s", result.Display))
				}
				for _, tag := range result.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
						color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
						color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
						color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
						color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags entry for device bay: ") + color.YellowString("%s", result.Display))
					}
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
				if result.DeviceCount != 0 {
					color.Cyan("\tDevice Count: " + color.YellowString("%d", result.DeviceCount))
				} else {
					color.Cyan("\tDevice Count: " + color.RedString("No device count entry for device bay: ") + color.YellowString("%s", result.Display))
				}
				if result.VirtualmachineCount != 0 {
					color.Cyan("\tVirtual Machine Count: " + color.YellowString("%d\n", result.VirtualmachineCount))
				} else {
					color.Cyan("\tVirtual Machine Count: " + color.RedString("No virtual machine count entry for device bay: ") + color.YellowString("%s\n", result.Display))
				}
			}
			for responseObjectDeviceRoles.Next != nil {
				nextPageDeviceRoles()
			}
			if responseObjectDeviceRoles.Next == nil {
				display := color.HiGreenString("\tAll Netbox device role objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n\n")
			}
		}
	},
}

func ApiConnectionNextPageDeviceRoles[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := *responseObjectDeviceRoles.Next

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

func nextPageDeviceRoles() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\tDo you want to continue to the next page of device roles objects? ['Y' or 'yes'] or ['n' or 'no']: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageDeviceRoles(responseObjectDeviceRoles, "GET", *responseObjectDeviceRoles.Next)
		displayDeviceRolesOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("Invalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func displayDeviceRolesOutput() {
	for _, result := range responseObjectDeviceRoles.Results {
		display := fmt.Sprintf("    ABC Device Role: %s\n", color.YellowString(result.Display))
		equals := strings.Repeat("=", len(display))
		color.Cyan("\n  " + equals + "\n")
		color.Cyan(display)
		color.Cyan("  " + equals + "\n")
		color.Cyan("\tID: " + color.YellowString("%d", result.Id))
		color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
		color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
		color.Cyan("\tName: " + color.YellowString("%s", result.Name))
		color.Cyan("\tSlug: " + color.YellowString("%s", result.Slug))
		if result.Color != "" {
			color.Cyan("\tColor: " + color.YellowString("%s", result.Color))
		} else {
			color.Cyan("\tColor: " + color.RedString("No color entry for device role: ") + color.YellowString("%s", result.Display))
		}
		color.Cyan("\tVM Role: " + color.YellowString("%v", result.VmRole))
		if result.ConfigTemplate.Id != 0 {
			color.Cyan("\tConfig Template: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", result.ConfigTemplate.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", result.ConfigTemplate.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", result.ConfigTemplate.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", result.ConfigTemplate.Name))
		} else {
			color.Cyan("\tConfig Template: " + color.RedString("No config template entry for device role: ") + color.YellowString("%s", result.Display))
		}
		if result.Description != "" {
			color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
		} else {
			color.Cyan("\tDescription: " + color.RedString("No description entry for device role: ") + color.YellowString("%s", result.Display))
		}
		for _, tag := range result.Tags {
			if tag.Id != 0 {
				color.Cyan("\tTags: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
			} else {
				color.Cyan("\tTags: " + color.RedString("No tags entry for device bay: ") + color.YellowString("%s", result.Display))
			}
		}
		color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
		color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))
		if result.DeviceCount != 0 {
			color.Cyan("\tDevice Count: " + color.YellowString("%d", result.DeviceCount))
		} else {
			color.Cyan("\tDevice Count: " + color.RedString("No device count entry for device bay: ") + color.YellowString("%s", result.Display))
		}
		if result.VirtualmachineCount != 0 {
			color.Cyan("\tVirtual Machine Count: " + color.YellowString("%d\n", result.VirtualmachineCount))
		} else {
			color.Cyan("\tVirtual Machine Count: " + color.RedString("No virtual machine count entry for device bay: ") + color.YellowString("%s\n", result.Display))
		}
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimDeviceRolesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimDeviceRolesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimDeviceRolesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimDeviceRolesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

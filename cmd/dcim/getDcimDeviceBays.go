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

type deviceBays struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Device  struct {
			Id      int    `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Name    string `json:"name"`
		} `json:"device"`
		Name            string `json:"name"`
		Label           string `json:"label"`
		Description     string `json:"description"`
		InstalledDevice struct {
			Id      int    `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Name    string `json:"name"`
		} `json:"installed_device"`
		Tags        []CommonFieldsSlug `json:"tags"`
		Created     string             `json:"created"`
		LastUpdated string             `json:"last_updated"`
	} `json:"results"`
}

var responseObjectDeviceBays = new(deviceBays)

// GetDcimDeviceBaysCmd represents the getDcimDeviceBays command
var GetDcimDeviceBaysCmd = &cobra.Command{
	Use:   "getDcimDeviceBays",
	Short: "GET a list of device bay objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of device bay objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectDeviceBays, "GET", "cmd.dcim.dcim_api_url.device_bays")

		if len(responseObjectDeviceBays.Results) == 0 {
			color.Cyan("  Total ABC Device Bays: " + color.RedString("No device bays entries found on server. Exiting...\n"))
		} else {
			color.Cyan("\nTotal ABC Device Bays: "+color.YellowString("%d"), responseObjectDeviceBays.Count)

			for _, result := range responseObjectDeviceBays.Results {
				display := fmt.Sprintf("    ABC Device Bay: %s\n", color.YellowString(result.Display))
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
					color.Cyan("\tDevice: " + color.RedString("No device entry for device bay: ") + color.YellowString("%s", result.Display))
				}
				if result.Name != "" {
					color.Cyan("\tName: " + color.YellowString("%s", result.Name))
				} else {
					color.Cyan("\tName: " + color.RedString("No name entry for device bay: ") + color.YellowString("%s", result.Display))
				}
				if result.Label != "" {
					color.Cyan("\tLabel: " + color.YellowString("%s", result.Label))
				} else {
					color.Cyan("\tLabel: " + color.RedString("No label entry for device bay: ") + color.YellowString("%s", result.Display))
				}
				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description entry for device bay: ") + color.YellowString("%s", result.Display))
				}
				if result.InstalledDevice.Id != 0 {
					color.Cyan("\tInstalled Device: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", result.InstalledDevice.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", result.InstalledDevice.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", result.InstalledDevice.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", result.InstalledDevice.Name))
				} else {
					color.Cyan("\tInstalled Device: " + color.RedString("No installed device entry for device bay: ") + color.YellowString("%s", result.Display))
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
				color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
			}
			for responseObjectDeviceBays.Next != nil {
				nextPageDeviceBays()
			}
			if responseObjectDeviceBays.Next == nil {
				display := color.HiGreenString("\tAll Netbox device bay objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n")
			}
		}
	},
}

func ApiConnectionNextPageDeviceBays[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := *responseObjectDeviceBays.Next

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

func nextPageDeviceBays() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\tDo you want to continue to the next page of device bay objects? ['Y' or 'yes'] or ['n' or 'no']: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageDeviceBays(responseObjectDeviceBays, "GET", *responseObjectDeviceBays.Next)
		displayDeviceBaysOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("Invalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func displayDeviceBaysOutput() {
	for _, result := range responseObjectDeviceBays.Results {
		display := fmt.Sprintf("    ABC Device Bay: %s\n", color.YellowString(result.Display))
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
			color.Cyan("\tDevice: " + color.RedString("No device entry for device bay: ") + color.YellowString("%s", result.Display))
		}
		if result.Name != "" {
			color.Cyan("\tName: " + color.YellowString("%s", result.Name))
		} else {
			color.Cyan("\tName: " + color.RedString("No name entry for device bay: ") + color.YellowString("%s", result.Display))
		}
		if result.Label != "" {
			color.Cyan("\tLabel: " + color.YellowString("%s", result.Label))
		} else {
			color.Cyan("\tLabel: " + color.RedString("No label entry for device bay: ") + color.YellowString("%s", result.Display))
		}
		if result.Description != "" {
			color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
		} else {
			color.Cyan("\tDescription: " + color.RedString("No description entry for device bay: ") + color.YellowString("%s", result.Display))
		}
		if result.InstalledDevice.Id != 0 {
			color.Cyan("\tInstalled Device: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", result.InstalledDevice.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", result.InstalledDevice.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", result.InstalledDevice.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", result.InstalledDevice.Name))
		} else {
			color.Cyan("\tInstalled Device: " + color.RedString("No installed device entry for device bay: ") + color.YellowString("%s", result.Display))
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
		color.Cyan("\tLast Updated: " + color.YellowString("%s\n", result.LastUpdated))
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimDeviceBaysCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimDeviceBaysCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimDeviceBaysCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimDeviceBaysCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimDeviceBaysCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

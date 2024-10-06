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

type deviceTypes struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous string  `json:"previous"`
	Results  []struct {
		Id           int    `json:"id"`
		Url          string `json:"url"`
		Display      string `json:"display"`
		Manufacturer struct {
			CommonFieldsSlug
		} `json:"manufacturer"`
		DefaultPlatform struct {
			CommonFieldsSlug
		} `json:"default_platform"`
		Model                  string  `json:"model"`
		Slug                   string  `json:"slug"`
		PartNumber             string  `json:"part_number"`
		UHeight                float32 `json:"u_height"`
		ExcludeFromUtilization bool    `json:"exclude_from_utilization"`
		IsFullDepth            bool    `json:"is_full_depth"`
		SubdeviceRole          struct {
			ValueLabel
		} `json:"subdevice_role"`
		Airflow struct {
			ValueLabel
		} `json:"airflow"`
		Weight     float32 `json:"weight"`
		WeightUnit struct {
			ValueLabel
		} `json:"weight_unit"`
		FrontImage  string `json:"front_image"`
		RearImage   string `json:"rear_image"`
		Description string `json:"description"`
		Comments    string `json:"comments"`
		Tags        []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
		Created                        string `json:"created"`
		LastUpdated                    string `json:"last_updated"`
		DeviceCount                    int    `json:"device_count"`
		ConsolePortTemplateCount       int    `json:"console_port_template_count"`
		ConsoleServerPortTemplateCount int    `json:"console_server_port_template_count"`
		PowerPortTemplateCount         int    `json:"power_port_template_count"`
		PowerOutletTemplateCount       int    `json:"power_outlet_template_count"`
		InterfaceTemplateCount         int    `json:"interface_template_count"`
		FrontPortTemplateCount         int    `json:"front_port_template_count"`
		RearPortTemplateCount          int    `json:"rear_port_template_count"`
		DeviceBayTemplateCount         int    `json:"device_bay_template_count"`
		ModuleBayTemplateCount         int    `json:"module_bay_template_count"`
		InventoryItemTemplateCount     int    `json:"inventory_item_template_count"`
	} `json:"results"`
}

var responseObjectDeviceTypes = new(deviceTypes)

// GetDcimDeviceTypesCmd represents the getDcimDeviceTypes command
var GetDcimDeviceTypesCmd = &cobra.Command{
	Use:   "getDcimDeviceTypes",
	Short: "GET a list of device type objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of device type objects`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectDeviceTypes, "GET", "cmd.dcim.dcim_api_url.device_types")

		if responseObjectDeviceTypes.Count != 0 {
			color.Cyan("\n  Total ABC Device Types: "+color.YellowString("%v"), responseObjectDeviceTypes.Count)

			for _, deviceType := range responseObjectDeviceTypes.Results {
				display := fmt.Sprintf("    ABC Device Type: %s\n", color.YellowString(deviceType.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", deviceType.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", deviceType.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", deviceType.Display))
				if deviceType.Manufacturer.Id != 0 {
					color.Cyan("\tManufacturer: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", deviceType.Manufacturer.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", deviceType.Manufacturer.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", deviceType.Manufacturer.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", deviceType.Manufacturer.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", deviceType.Manufacturer.Slug))
				} else {
					color.Cyan("\tManufacturer: " + color.RedString("No manufacturer found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.DefaultPlatform.Id != 0 {
					color.Cyan("\tDefault Platform: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", deviceType.DefaultPlatform.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", deviceType.DefaultPlatform.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", deviceType.DefaultPlatform.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", deviceType.DefaultPlatform.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", deviceType.DefaultPlatform.Slug))
				} else {
					color.Cyan("\tDefault Platform: " + color.RedString("No default platform found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.Model != "" {
					color.Cyan("\tModel: " + color.YellowString("%s", deviceType.Model))
				} else {
					color.Cyan("\tModel: " + color.RedString("No model found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.Slug != "" {
					color.Cyan("\tSlug: " + color.YellowString("%s", deviceType.Slug))
				} else {
					color.Cyan("\tSlug: " + color.RedString("No slug found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.PartNumber != "" {
					color.Cyan("\tPart Number: " + color.YellowString("%s", deviceType.PartNumber))
				} else {
					color.Cyan("\tPart Number: " + color.RedString("No part number found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.UHeight != 0.0 {
					color.Cyan("\tUnit Height: " + color.YellowString("%v", deviceType.UHeight))
				} else {
					color.Cyan("\tUnit Height: " + color.RedString("No unit height found for device: ") + color.YellowString(deviceType.Display))
				}
				color.Cyan("\tExcluded From Utilization: " + color.YellowString("%v", deviceType.ExcludeFromUtilization))
				color.Cyan("\tIs Full Depth: " + color.YellowString("%v", deviceType.IsFullDepth))
				if deviceType.SubdeviceRole.Value != "" {
					color.Cyan("\tSubdevice Role: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", deviceType.SubdeviceRole.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", deviceType.SubdeviceRole.Label))
				} else {
					color.Cyan("\tSubdevice Role: " + color.RedString("No subdevice role found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.Airflow.Value != "" {
					color.Cyan("\tAirflow: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", deviceType.Airflow.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", deviceType.Airflow.Label))
				} else {
					color.Cyan("\tAirflow: " + color.RedString("No airflow found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.Weight != 0.0 {
					color.Cyan("\tWeight: " + color.YellowString("%d", deviceType.Weight))
				} else {
					color.Cyan("\tWeight: " + color.RedString("No weight found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.WeightUnit.Value != "" {
					color.Cyan("\tWeight Unit: ")
					color.Cyan("\t  Value: " + color.YellowString("%s", deviceType.WeightUnit.Value))
					color.Cyan("\t  Label: " + color.YellowString("%s", deviceType.WeightUnit.Label))
				} else {
					color.Cyan("\tWeight Unit: " + color.RedString("No Weight unit found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.FrontImage != "" {
					color.Cyan("\tFront Image: " + color.YellowString("%s", deviceType.FrontImage))
				} else {
					color.Cyan("\tFront Image: " + color.RedString("No front image found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.RearImage != "" {
					color.Cyan("\tRear Image: " + color.YellowString("%s", deviceType.RearImage))
				} else {
					color.Cyan("\tRear Image: " + color.RedString("No rear image found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", deviceType.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description found for device: ") + color.YellowString(deviceType.Description))
				}
				if deviceType.Comments != "" {
					color.Cyan("\tComments: " + color.YellowString("%s", deviceType.Comments))
				} else {
					color.Cyan("\tComments: " + color.RedString("No comments found for device: ") + color.YellowString(deviceType.Comments))
				}
				for _, tag := range deviceType.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
						color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
						color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
						color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
						color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
						color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags found for device: ") + color.YellowString(deviceType.Display))
					}
				}

				color.Cyan("\tCreated: " + color.YellowString("%s", deviceType.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", deviceType.LastUpdated))
				if deviceType.DeviceCount != 0 {
					color.Cyan("\tDevice Count: " + color.YellowString("%d", deviceType.DeviceCount))
				} else {
					color.Cyan("\tDevice Count: " + color.RedString("No device count found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.ConsolePortTemplateCount != 0 {
					color.Cyan("\tConsole Port Template Count: " + color.YellowString("%d", deviceType.ConsolePortTemplateCount))
				} else {
					color.Cyan("\tConsole Port Template Count: " + color.RedString("No console port template count found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.ConsoleServerPortTemplateCount != 0 {
					color.Cyan("\tConsole Server Port Template Count: " + color.YellowString("%d", deviceType.ConsoleServerPortTemplateCount))
				} else {
					color.Cyan("\tConsole Server Port Template Count: " + color.RedString("No console server port template count found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.PowerPortTemplateCount != 0 {
					color.Cyan("\tPower Port Template Count: " + color.YellowString("%d", deviceType.PowerPortTemplateCount))
				} else {
					color.Cyan("\tPower Port Template Count: " + color.RedString("No power port template count found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.PowerOutletTemplateCount != 0 {
					color.Cyan("\tPower Outlet Template Count: " + color.YellowString("%d", deviceType.PowerOutletTemplateCount))
				} else {
					color.Cyan("\tPower Outlet Template Count: " + color.RedString("No power outlet template count found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.InterfaceTemplateCount != 0 {
					color.Cyan("\tInterface Template Count: " + color.YellowString("%d", deviceType.InterfaceTemplateCount))
				} else {
					color.Cyan("\tInterface Template Count: " + color.RedString("No interface template count found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.FrontPortTemplateCount != 0 {
					color.Cyan("\tFront Port Template Count: " + color.YellowString("%d", deviceType.FrontPortTemplateCount))
				} else {
					color.Cyan("\tFront Port Template Count: " + color.RedString("No front port template count found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.RearPortTemplateCount != 0 {
					color.Cyan("\tRear Port Template Count: " + color.YellowString("%d", deviceType.RearPortTemplateCount))
				} else {
					color.Cyan("\tRear Port Template Count: " + color.RedString("No rear port template count found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.DeviceBayTemplateCount != 0 {
					color.Cyan("\tDevice Bay Template Count: " + color.YellowString("%d", deviceType.DeviceBayTemplateCount))
				} else {
					color.Cyan("\tDevice Bay Template Count: " + color.RedString("No device bay template count found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.ModuleBayTemplateCount != 0 {
					color.Cyan("\tModule Bay Template Count: " + color.YellowString("%d", deviceType.ModuleBayTemplateCount))
				} else {
					color.Cyan("\tModule Bay Template Count: " + color.RedString("No module bay template count found for device: ") + color.YellowString(deviceType.Display))
				}
				if deviceType.InventoryItemTemplateCount != 0 {
					color.Cyan("\tInventory Item Template Count: " + color.YellowString("%d\n", deviceType.InventoryItemTemplateCount))
				} else {
					color.Cyan("\tInventory Item Template Count: " + color.RedString("No inventory template count found for device: ") + color.YellowString(deviceType.Display) + "\n")
				}
			}
			for responseObjectDeviceTypes.Next != nil {
				nextPageDeviceTypes()
			}
			if responseObjectDeviceTypes.Next == nil {
				display := color.HiGreenString("\tAll Netbox device type objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n\n")
			}
		} else {
			color.Cyan("  Total ABC Device Types: " + color.RedString("No device types found on the server. Exiting...\n"))
		}
	},
}

func nextPageDeviceTypes() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\tDo you want to continue to the next page of device objects? [yes/no]: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageDeviceTypes(responseObjectDeviceTypes, "GET", *responseObjectDeviceTypes.Next)
		displayDeviceTypesOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("\nInvalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func ApiConnectionNextPageDeviceTypes[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := *responseObjectDeviceTypes.Next

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
		log.Fatalf("Error patching Netbox API objects: %s\n", err)
	}
}

func displayDeviceTypesOutput() {
	for _, device := range responseObjectDevices.Results {
		display := fmt.Sprintf("    ABC Devices: %s\n", color.YellowString(device.Display))
		equals := strings.Repeat("=", len(display))
		color.Cyan("\n  " + equals + "\n")
		color.Cyan(display)
		color.Cyan("  " + equals + "\n")
		color.Cyan("\tID: " + color.YellowString("%d", device.Id))
		color.Cyan("\tURL: " + color.YellowString("%s", device.Url))
		color.Cyan("\tDisplay: " + color.YellowString("%s", device.Display))
		color.Cyan("\tName: " + color.YellowString("%s", device.Name))
		if device.DeviceType.Id != 0 {
			color.Cyan("\tDevice Type: \n")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.DeviceType.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.DeviceType.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.DeviceType.Display))
		} else {
			color.Cyan("\tDevice Type: No device type entry found for device: " + color.YellowString("%s", device.Name))
		}
		if device.DeviceType.Manufacturer.Id > 0 {
			color.Cyan("\t  Manufacturer: \n")
			color.Cyan("\t    ID: " + color.YellowString("%d", device.DeviceType.Manufacturer.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", device.DeviceType.Manufacturer.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", device.DeviceType.Manufacturer.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", device.DeviceType.Manufacturer.Name))
			color.Cyan("\t    Slug: " + color.YellowString("%s", device.DeviceType.Manufacturer.Slug))
			color.Cyan("\t    Description: " + color.YellowString("%s", device.DeviceType.Manufacturer.Description))
			color.Cyan("\t    Device Type Count: " + color.YellowString("%d", device.DeviceType.Manufacturer.DevicetypeCount))
			color.Cyan("\t  Model: " + color.YellowString("%s", device.DeviceType.Model))
			color.Cyan("\t  Slug: " + color.YellowString("%s", device.DeviceType.Slug))
			color.Cyan("\t  Description: " + color.YellowString("%s", device.DeviceType.Description))
			color.Cyan("\t  Device Count: " + color.YellowString("%d", device.DeviceType.DeviceCount))
		} else {
			color.Cyan("\t  Manufacturer: No manufacturer entry found for device: " + color.YellowString("%s", device.Name))
		}
		if device.Role.Id != 0 {
			color.Cyan("\tRole: \n")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.Role.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.Role.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.Role.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", device.Role.Name))
			color.Cyan("\t  Slug: " + color.YellowString("%s", device.Role.Slug))
			color.Cyan("\t  Description: " + color.YellowString("%s", device.Role.Description))
			color.Cyan("\t  Device Count: " + color.YellowString("%d", device.Role.DeviceCount))
			color.Cyan("\t  Virtual Machine Count: " + color.YellowString("%d", device.Role.VirtualmachineCount))
		} else {
			color.Cyan("\tRole: No role entry found for device: " + color.YellowString("%s", device.Name))
		}
		if device.Tenant.Id != 0 {
			color.Cyan("\tTenant: \n")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.Tenant.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.Tenant.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.Tenant.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", device.Tenant.Name))
			color.Cyan("\t  Slug: " + color.YellowString("%s", device.Tenant.Slug))
		} else {
			color.Cyan("\tTenant: " + color.RedString("No tenant entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Platform.Id != 0 {
			color.Cyan("\tPlatform: \n")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.Platform.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.Platform.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.Platform.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", device.Platform.Name))
			color.Cyan("\t  Slug: " + color.YellowString("%s", device.Platform.Slug))
		} else {
			color.Cyan("\tPlatform: " + color.RedString("No platform entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Location.Id != 0 {
			color.Cyan("\tLocation: \n")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.Location.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.Location.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.Location.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", device.Location.Name))
			color.Cyan("\t  Slug: " + color.YellowString("%s", device.Location.Slug))
			color.Cyan("\t  Depth: " + color.YellowString("%d", device.Location.Depth))
		} else {
			color.Cyan("\tLocation: " + color.RedString("No location entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Rack.Id != 0 {
			color.Cyan("\tRack: \n")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.Rack.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.Rack.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.Rack.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", device.Rack.Name))
		} else {
			color.Cyan("\tRack: " + color.RedString("No rack entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Position != 0 {
			color.Cyan("\tPosition: " + color.YellowString("%d", device.Position))
		} else {
			color.Cyan("\tPosition: " + color.RedString("No position entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Face.Value != "" {
			color.Cyan("\tFace: ")
			color.Cyan("\t  Value: " + color.YellowString("%s", device.Face.Value))
			color.Cyan("\t  Label: " + color.YellowString("%s", device.Face.Label))
		} else {
			color.Cyan("\tFace: " + color.RedString("No face entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Latitude != 0 {
			color.Cyan("\tLatitude: " + color.YellowString("%d", device.Latitude))
		} else {
			color.Cyan("\tLatitude: " + color.RedString("No latitude entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Longitude != 0 {
			color.Cyan("\tLongitude: " + color.YellowString("%d", device.Longitude))
		} else {
			color.Cyan("\tLongitude: " + color.RedString("No longitude entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.ParentDevice.Id != 0 {
			color.Cyan("\tParent Device: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.ParentDevice.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.ParentDevice.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.ParentDevice.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", device.ParentDevice.Name))
		} else {
			color.Cyan("\tParent Device: " + color.RedString("No parent device entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Status.Value != "" {
			color.Cyan("\tStatus: ")
			color.Cyan("\t  Value: " + color.YellowString("%s", device.Status.Value))
			color.Cyan("\t  Label: " + color.YellowString("%s", device.Status.Label))
		} else {
			color.Cyan("\tStatus: " + color.RedString("No status entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Airflow.Value != "" {
			color.Cyan("\tAirflow: ")
			color.Cyan("\t  Value: " + color.YellowString("%s", device.Airflow.Value))
			color.Cyan("\t  Label: " + color.YellowString("%s", device.Airflow.Label))
		} else {
			color.Cyan("\tAirflow: " + color.RedString("No airflow entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.PrimaryIp.Id != 0 {
			color.Cyan("\tPrimary IP: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.PrimaryIp.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.PrimaryIp.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.PrimaryIp.Display))
			color.Cyan("\t  Family: ")
			color.Cyan("\t    Value: " + color.YellowString("%d", device.PrimaryIp.Family.Value))
			color.Cyan("\t    Label: " + color.YellowString("%s", device.PrimaryIp.Family.Label))
			color.Cyan("\t  Address: " + color.YellowString("%s", device.PrimaryIp.Address))
		} else {
			color.Cyan("\tPrimary IP: " + color.RedString("No primary ip entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.PrimaryIp6.Id != 0 {
			color.Cyan("\tPrimary IPv6: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.PrimaryIp6.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.PrimaryIp6.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.PrimaryIp6.Display))
			color.Cyan("\t  Family: ")
			color.Cyan("\t    Value: " + color.YellowString("%d", device.PrimaryIp6.Family.Value))
			color.Cyan("\t    Label: " + color.YellowString("%s", device.PrimaryIp6.Family.Label))
			color.Cyan("\t  Address: " + color.YellowString("%s", device.PrimaryIp6.Address))
		} else {
			color.Cyan("\tPrimary IPv6: " + color.RedString("No primary ipv6 entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.OobIp.Id != 0 {
			color.Cyan("\tOOB IP: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.OobIp.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.OobIp.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.OobIp.Display))
			color.Cyan("\t  Family: ")
			color.Cyan("\t    Value: " + color.YellowString("%d", device.OobIp.Family.Value))
			color.Cyan("\t    Label: " + color.YellowString("%s", device.OobIp.Family.Label))
			color.Cyan("\t  Address: " + color.YellowString("%s", device.OobIp.Address))
		} else {
			color.Cyan("\tOOB IP: " + color.RedString("No oop ip entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.VirtualChassis.Id != 0 {
			color.Cyan("\tVirtual Chassis: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.VirtualChassis.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.VirtualChassis.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.VirtualChassis.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", device.VirtualChassis.Name))
			color.Cyan("\t  Master: ")
			color.Cyan("\t    ID: " + color.YellowString("%d", device.VirtualChassis.Master.Id))
			color.Cyan("\t    URL: " + color.YellowString("%s", device.VirtualChassis.Master.Url))
			color.Cyan("\t    Display: " + color.YellowString("%s", device.VirtualChassis.Master.Display))
			color.Cyan("\t    Name: " + color.YellowString("%s", device.VirtualChassis.Master.Name))
		} else {
			color.Cyan("\tVirtual Chassis: " + color.RedString("No virtual chassis entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.VcPosition != 0 {
			color.Cyan("\tVC Position: " + color.YellowString("%d", device.VcPosition))
		} else {
			color.Cyan("\tVC Position: " + color.RedString("No vc position entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.VcPriority != 0 {
			color.Cyan("\tVC Priority: " + color.YellowString("%d", device.VcPosition))
		} else {
			color.Cyan("\tVC Priority: " + color.RedString("No vc priority entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Description != "" {
			color.Cyan("\tDescription: " + color.YellowString("%s", device.Description))
		} else {
			color.Cyan("\tDescription: " + color.RedString("No description entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.Comments != "" {
			color.Cyan("\tComments: " + color.YellowString("%s", device.Description))
		} else {
			color.Cyan("\tComments: " + color.RedString("No comments entry found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.ConfigTemplate.Id != 0 {
			color.Cyan("\tConfig Template: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", device.ConfigTemplate.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", device.ConfigTemplate.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", device.ConfigTemplate.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", device.ConfigTemplate.Name))
		} else {
			color.Cyan("\tConfig Template: " + color.RedString("No config template entry found for device: ") + color.YellowString("%s", device.Name))
		}
		color.Cyan("\tConfig Context: " + color.YellowString("%v", device.ConfigContext))
		if device.LocalContextData != "" {
			color.Cyan("\tLocal Context Data: " + color.YellowString("%s", device.LocalContextData))
		} else {
			color.Cyan("\tLocal Context Data: " + color.RedString("No local context data entry found for device: ") + color.YellowString("%s", device.Name))
		}
		for _, tag := range device.Tags {
			if tag.Id != 0 {
				color.Cyan("\tTags: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
				color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
			} else {
				color.Cyan("\tTags: " + color.RedString("No tags entry found for device: ") + color.YellowString("%s", device.Name))
			}
		}

		color.Cyan("\tCreated: " + color.YellowString("%s", device.Created))
		color.Cyan("\tLast Updated: " + color.YellowString("%s", device.LastUpdated))
		if device.ConsolePortCount != 0 {
			color.Cyan("\tConsole Port Count: " + color.YellowString("%d", device.ConsolePortCount))
		} else {
			color.Cyan("\tConsole Port Count: " + color.RedString("No console port count found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.ConsoleServerPortCount != 0 {
			color.Cyan("\tConsole Server Port Count: " + color.YellowString("%d", device.ConsolePortCount))
		} else {
			color.Cyan("\tConsole Server Port Count: " + color.RedString("No console server port count found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.PowerPortCount != 0 {
			color.Cyan("\tPower Port Count: " + color.YellowString("%d", device.PowerPortCount))
		} else {
			color.Cyan("\tPower Port Count: " + color.RedString("No power port count found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.PowerOutletCount != 0 {
			color.Cyan("\tPower Outlet Count: " + color.YellowString("%d", device.PowerOutletCount))
		} else {
			color.Cyan("\tPower Outlet Count: " + color.RedString("No power outlet count found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.InterfaceCount != 0 {
			color.Cyan("\tInterface Count: " + color.YellowString("%d", device.InterfaceCount))
		} else {
			color.Cyan("\tPower Outlet Count: " + color.RedString("No interface count found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.FrontPortCount != 0 {
			color.Cyan("\tFront Port Count: " + color.YellowString("%d", device.FrontPortCount))
		} else {
			color.Cyan("\tFront Port Count: " + color.RedString("No front port count found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.RearPortCount != 0 {
			color.Cyan("\tRear Port Count: " + color.YellowString("%d", device.RearPortCount))
		} else {
			color.Cyan("\tRear Port Count: " + color.RedString("No rear port count found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.DeviceBayCount != 0 {
			color.Cyan("\tDevice Bay Count: " + color.YellowString("%d", device.DeviceBayCount))
		} else {
			color.Cyan("\tDevice Bay Count: " + color.RedString("No device bay count found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.ModuleBayCount != 0 {
			color.Cyan("\tModule Bay Count: " + color.YellowString("%d", device.ModuleBayCount))
		} else {
			color.Cyan("\tModule Bay Count: " + color.RedString("No module bay count found for device: ") + color.YellowString("%s", device.Name))
		}
		if device.InventoryItemCount != 0 {
			color.Cyan("\tInventory Item Count: " + color.YellowString("%d\n", device.InventoryItemCount))
		} else {
			color.Cyan("\tInventory Item Count: " + color.RedString("No inventory item count found for device: ") + color.YellowString("%s\n", device.Name))
		}
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimDeviceTypesCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimDeviceTypesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag required flag: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimDeviceTypesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimDeviceTypesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

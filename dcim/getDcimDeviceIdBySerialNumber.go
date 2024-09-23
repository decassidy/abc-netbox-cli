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
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// GetDcimDeviceIdBySerialNumberCmd represents the getDcimDeviceBySerialNumber command
var GetDcimDeviceIdBySerialNumberCmd = &cobra.Command{
	Use:   "getDcimDeviceIdBySerialNumber",
	Short: "GET a device's ID by serial number query",
	Long: `
ABC Netbox Automation Tools:
  GET a device's ID by serial number query`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(devices)
		ApiConnectionSerialNumber(responseObject, "GET", "cmd.dcim.dcim_api_url.devices_serial_number")

		if responseObject.Count > 0 {
			for _, device := range responseObject.Results {
				display := fmt.Sprintf("    ABC Device Name: %s\n", color.YellowString(device.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tDevice ID: " + color.YellowString("%d", device.Id))
				color.Cyan("\tSite ID: " + color.YellowString("%d", device.Site.Id))
				color.Cyan("\tSite Display: " + color.YellowString("%s", device.Site.Display))
				color.Cyan("\tPrimary IP: " + color.YellowString("%s\n", device.PrimaryIp.Address))
			}
		} else {
			color.Cyan("  ABC Device: " + color.RedString("No device entries found on server with serial number: "+color.YellowString("%s", serial)+color.RedString(" Exiting...\n")))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimDeviceIdBySerialNumberCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimDeviceIdBySerialNumberCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for GetDcimDeviceIdBySerialNumberCmd", err)
	}

	GetDcimDeviceIdBySerialNumberCmd.Flags().StringVarP(&serial, "serial", "s", "", "serial number of object you want to get")
	err = GetDcimDeviceIdBySerialNumberCmd.MarkFlagRequired("serial")
	if err != nil {
		log.Fatalf("Error marking serial flag as required: %s - for GetDcimDeviceIdBySerialNumberCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimDeviceBySerialNumberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimDeviceBySerialNumberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

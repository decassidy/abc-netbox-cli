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
	"time"
)

type rackReservations struct {
	Count    uint   `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Id      uint   `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Rack    struct {
			CommonFieldsNoSlug
		} `json:"rack"`
		Units       []int     `json:"units"`
		Created     time.Time `json:"created"`
		LastUpdated time.Time `json:"last_updated"`
		User        struct {
			Id       uint   `json:"id"`
			Url      string `json:"url"`
			Display  string `json:"display"`
			Username string `json:"username"`
		} `json:"user"`
		Tenant struct {
			CommonFieldsSlug
		} `json:"tenant"`
		Description string `json:"description"`
		Comments    string `json:"comments"`
		Tags        []struct {
			CommonFieldsSlug
			Color string `json:"color"`
		} `json:"tags"`
	} `json:"results"`
}

// GetDcimRackReservationsCmd represents the getDcimRackReservations command
var GetDcimRackReservationsCmd = &cobra.Command{
	Use:   "getDcimRackReservations",
	Short: "GET a list of rack reservation objects",
	Long: `
ABC Netbox Automation Tools:
  GET a list of rack reservation objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(rackReservations)
		ApiConnectionNonID(responseObject, "GET", "cmd.dcim.dcim_api_url.rack_reservations")

		if responseObject.Count != 0 {
			color.Cyan("\n  ABC Rack Reservations: "+color.YellowString("%d"), responseObject.Count)
			for _, result := range responseObject.Results {
				display := fmt.Sprintf("    ABC Rack Reservation: %s\n", color.YellowString(result.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", result.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", result.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", result.Display))
				color.Cyan("\tRack: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", result.Rack.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.Rack.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.Rack.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", result.Rack.Name))
				for _, unit := range result.Units {
					color.Cyan("\tUnits: " + color.YellowString("%d", unit))
				}
				color.Cyan("\tCreated: " + color.YellowString("%s", result.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%s", result.LastUpdated))

				color.Cyan("\tUser: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", result.User.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.User.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.User.Display))
				color.Cyan("\t  Username: " + color.YellowString("%s", result.User.Username))

				color.Cyan("\tTenant: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", result.Tenant.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", result.Tenant.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", result.Tenant.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", result.Tenant.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", result.Tenant.Slug))

				if result.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", result.Description))
				} else {
					color.Cyan("\tDescription" + color.RedString("No description entry found for ") + color.YellowString("%s", result.Display))
				}

				if result.Comments != "" {
					color.Cyan("\tComments: " + color.YellowString("%s", result.Comments))
				} else {
					color.Cyan("\tComments" + color.RedString("No comments entry found for ") + color.YellowString("%s", result.Display))
				}

				for _, tag := range result.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
						color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
						color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
						color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
						color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
						color.Cyan("\t  Color: " + color.YellowString("%s\n", tag.Color))
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s\n", result.Display))
					}
				}
			}
		} else {
			color.Cyan("  ABC Rack Reservations: " + color.RedString("No rack reservations found on server. Exiting...\n"))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimRackReservationsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimRackReservationsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimRackReservationsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimRackReservationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimRackReservationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

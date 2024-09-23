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

type rackReservationsByID struct {
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
}

// GetDcimRackReservationsByIdCmd represents the getDcimRackReservationsById command
var GetDcimRackReservationsByIdCmd = &cobra.Command{
	Use:   "getDcimRackReservationsById",
	Short: "GET an rack reservation object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an rack reservation object by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(rackReservationsByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.rack_reservations_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    ABC Rack Reservation: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			color.Cyan("\tRack: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Rack.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Rack.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Rack.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Rack.Name))
			for _, unit := range responseObject.Units {
				color.Cyan("\tUnits: " + color.YellowString("%d", unit))
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s", responseObject.LastUpdated))

			color.Cyan("\tUser: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.User.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.User.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.User.Display))
			color.Cyan("\t  Username: " + color.YellowString("%s", responseObject.User.Username))

			color.Cyan("\tTenant: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Tenant.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Tenant.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Tenant.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Tenant.Name))
			color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Tenant.Slug))

			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription" + color.RedString("No description entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			if responseObject.Comments != "" {
				color.Cyan("\tComments: " + color.YellowString("%s", responseObject.Comments))
			} else {
				color.Cyan("\tComments" + color.RedString("No comments entry found for ") + color.YellowString("%s", responseObject.Display))
			}

			for _, tag := range responseObject.Tags {
				if tag.Id != 0 {
					color.Cyan("\tTags: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
					color.Cyan("\t  Color: " + color.YellowString("%s\n", tag.Color))
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags entry found for ") + color.YellowString("%s\n", responseObject.Display))
				}
			}
		} else {
			color.Red("  Doh! No rack reservation object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimRackReservationsByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimRackReservationsByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimRackReservationsByIdCmd", err)
	}

	GetDcimRackReservationsByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the rack reservation object")
	err = GetDcimRackReservationsByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimRackReservationsByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimRackReservationsByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

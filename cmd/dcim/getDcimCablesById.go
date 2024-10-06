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

// cablesByID represents a JSON object containing detailed information about a specific cable.
// It includes information such as the cable ID, URL, display name, type, and termination details.
// The termination details include objectType, objectId, and object information such as ID, URL,
// display name, circuit, termSide, cable, occupied status, device details, and name.
type cablesByID struct {
	Id            uint   `json:"id"`
	Url           string `json:"url"`
	Display       string `json:"display"`
	Type          string `json:"type"`
	ATerminations []struct {
		ObjectType string `json:"object_type"`
		ObjectId   uint   `json:"object_id"`
		Object     struct {
			Id      uint   `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Circuit struct {
				Id          uint   `json:"id"`
				Url         string `json:"url"`
				Display     string `json:"display"`
				Cid         string `json:"cid"`
				Description string `json:"description"`
			} `json:"circuit,omitempty"`
			TermSide    string `json:"term_side,omitempty"`
			Description string `json:"description"`
			Cable       struct {
				Id          uint   `json:"id"`
				Url         string `json:"url"`
				Display     string `json:"display"`
				Label       string `json:"label"`
				Description string `json:"description"`
			} `json:"cable"`
			Occupied bool `json:"_occupied"`
			Device   struct {
				Id          uint   `json:"id"`
				Url         string `json:"url"`
				Display     string `json:"display"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"device,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"object"`
	} `json:"a_terminations"`
	BTerminations []struct {
		ObjectType string `json:"object_type"`
		ObjectId   uint   `json:"object_id"`
		Object     struct {
			Id      uint   `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Device  struct {
				Id          uint   `json:"id"`
				Url         string `json:"url"`
				Display     string `json:"display"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"device"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Cable       struct {
				Id          uint   `json:"id"`
				Url         string `json:"url"`
				Display     string `json:"display"`
				Label       string `json:"label"`
				Description string `json:"description"`
			} `json:"cable"`
			Occupied bool `json:"_occupied"`
		} `json:"object"`
	} `json:"b_terminations"`
	Status struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"status"`
	Tenant struct {
		CommonFieldsSlug
	} `json:"tenant,omitempty"`
	Label      string  `json:"label"`
	Color      string  `json:"color"`
	Length     float32 `json:"length,omitempty"`
	LengthUnit struct {
		ValueLabel
	} `json:"length_unit,omitempty"`
	Description string `json:"description"`
	Comments    string `json:"comments"`
	Tags        []struct {
		CommonFieldsSlug
		Color string `json:"color,omitempty"`
	} `json:"tags,omitempty"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

// GetDcimCablesByIdCmd represents the getDcimCablesById command.
var GetDcimCablesByIdCmd = &cobra.Command{
	Use:   "getDcimCablesById",
	Short: "GET an cable object by ID",
	Long: `
ABC Netbox Automation Tools:
  GET an cable object by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(cablesByID)
		apiConnectionID(responseObject, "GET", "cmd.dcim.dcim_api_url.cables_id")

		if responseObject.Id > 0 {
			display := fmt.Sprintf("    ABC Cable: %s\n", color.YellowString(responseObject.Display))
			equals := strings.Repeat("=", len(display))
			color.Cyan("\n  " + equals + "\n")
			color.Cyan(display)
			color.Cyan("  " + equals + "\n")
			color.Cyan("\tID: " + color.YellowString("%d", responseObject.Id))
			color.Cyan("\tURL: " + color.YellowString("%s", responseObject.Url))
			color.Cyan("\tDisplay: " + color.YellowString("%s", responseObject.Display))
			color.Cyan("\tType: " + color.YellowString("%s", responseObject.Type))
			for _, term := range responseObject.ATerminations {
				if term.ObjectId > 0 {
					color.Cyan("\tA Terminations: ")
					color.Cyan("\t  Object Type: " + color.YellowString("%s", term.ObjectType))
					color.Cyan("\t  Object Id: " + color.YellowString("%d", term.ObjectId))
					if term.Object.Id != 0 {
						color.Cyan("\t  Object: ")
						color.Cyan("\t    ID: " + color.YellowString("%d", term.Object.Id))
						color.Cyan("\t    URL: " + color.YellowString("%s", term.Object.Url))
						color.Cyan("\t    Display: " + color.YellowString("%s", term.Object.Display))
					} else {
						color.Cyan("\t  Object: " + color.RedString("No Object entry found for cable: ") + color.YellowString("%s", responseObject.Display))
					}
					if term.Object.Circuit.Id > 0 {
						color.Cyan("\t    Circuit: ")
						color.Cyan("\t      ID: " + color.YellowString("%d", term.Object.Circuit.Id))
						color.Cyan("\t      URL: " + color.YellowString("%s", term.Object.Circuit.Url))
						color.Cyan("\t      Display: " + color.YellowString("%s", term.Object.Circuit.Display))
						color.Cyan("\t      CID: " + color.YellowString("%s", term.Object.Circuit.Cid))
					} else {
						color.Cyan("\t    Circuit: " + color.RedString("No circuit entry found for cable: ") + color.YellowString("%s", responseObject.Display))
					}
					if term.Object.TermSide != "" {
						color.Cyan("\t    Term Side: " + color.YellowString("%s", term.Object.TermSide))
					} else {
						color.Cyan("\t    Term Side: " + color.RedString("No term side entry found for cable: ") + color.YellowString("%s", responseObject.Display))
					}
					if term.Object.Cable.Id > 0 {
						color.Cyan("\t    Cable: ")
						color.Cyan("\t      ID: " + color.YellowString("%d", term.Object.Cable.Id))
						color.Cyan("\t      URL: " + color.YellowString("%s", term.Object.Cable.Url))
						color.Cyan("\t      Display: " + color.YellowString("%s", term.Object.Cable.Display))
						if term.Object.Cable.Label != "" {
							color.Cyan("\t      Label: " + color.YellowString("%s", term.Object.Cable.Label))
						} else {
							color.Cyan("\t      Label: " + color.RedString("No label entry found for cable: ") + color.YellowString("%s", responseObject.Display))
						}
						if term.Object.Cable.Description != "" {
							color.Cyan("\t      Description: " + color.YellowString("%s", term.Object.Cable.Description))
						} else {
							color.Cyan("\t      Description: " + color.RedString("No description entry found for cable: ") + color.YellowString("%s", responseObject.Display))
						}
					} else {
						color.Cyan("\t    Cable: " + color.RedString("No cable entry found for cable: ") + color.YellowString("%s", responseObject.Display))
					}
					color.Cyan("\t    Occupied: " + color.YellowString("%v", term.Object.Occupied))
				} else {
					color.Cyan("\tA Terminations: " + color.RedString("No A Terminations entry found for cable: ") + color.YellowString("%s", responseObject.Display))
				}
			}
			for _, bTermination := range responseObject.BTerminations {
				if bTermination.ObjectId > 0 {
					color.Cyan("\tB Terminations: ")
					color.Cyan("\t  Object Type: " + color.YellowString("%s", bTermination.ObjectType))
					color.Cyan("\t  Object Id: " + color.YellowString("%d", bTermination.ObjectId))
					color.Cyan("\t  Object: ")
					color.Cyan("\t    ID: " + color.YellowString("%d", bTermination.Object.Id))
					color.Cyan("\t    URL: " + color.YellowString("%s", bTermination.Object.Url))
					color.Cyan("\t    Display: " + color.YellowString("%s", bTermination.Object.Display))
					color.Cyan("\t    Device: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", bTermination.Object.Device.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", bTermination.Object.Device.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", bTermination.Object.Device.Display))
					color.Cyan("\t      Name: " + color.YellowString("%s", bTermination.Object.Device.Name))
					color.Cyan("\t    Name: " + color.YellowString("%s", bTermination.Object.Name))
					if bTermination.Object.Cable.Id > 0 {
						color.Cyan("\t    Cable: ")
						color.Cyan("\t      ID: " + color.YellowString("%d", bTermination.Object.Cable.Id))
						color.Cyan("\t      URL: " + color.YellowString("%s", bTermination.Object.Cable.Url))
						color.Cyan("\t      Display: " + color.YellowString("%s", bTermination.Object.Cable.Display))
						if bTermination.Object.Cable.Label != "" {
							color.Cyan("\t      Label: " + color.YellowString("%s", bTermination.Object.Cable.Label))
						} else {
							color.Cyan("\t      Label: " + color.RedString("No label entry found for cable: ") + color.YellowString("%s", responseObject.Display))
						}
						if bTermination.Object.Cable.Description != "" {
							color.Cyan("\t      Description: " + color.YellowString("%s", bTermination.Object.Cable.Description))
						} else {
							color.Cyan("\t      Description: " + color.RedString("No description entry found for cable: ") + color.YellowString("%s", responseObject.Display))
						}
					} else {
						color.Cyan("\t    Cable: " + color.RedString("No cable entry found for cable: ") + color.YellowString("%s", responseObject.Display))
					}
					color.Cyan("\t    Occupied: " + color.YellowString("%v", bTermination.Object.Occupied))
				} else {
					color.Cyan("\tB Terminations: " + color.RedString("No B Terminations entry found for device: ") + color.YellowString("%s", responseObject.Display))
				}
			}
			if responseObject.Status.Value != "" {
				color.Cyan("\tStatus: ")
				color.Cyan("\t  Value: " + color.YellowString("%v", responseObject.Status.Value))
				color.Cyan("\t  Label: " + color.YellowString("%v", responseObject.Status.Label))
			} else {
				color.Cyan("\tStatus: " + color.RedString("No Status entry found for device: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Tenant.Id > 0 {
				color.Cyan("\tTenant: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", responseObject.Tenant.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", responseObject.Tenant.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", responseObject.Tenant.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", responseObject.Tenant.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", responseObject.Tenant.Slug))
			} else {
				color.Cyan("\tTenant: " + color.RedString("No Tenant entry found for device: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Label != "" {
				color.Cyan("\tLabel: " + color.YellowString("%s", responseObject.Label))
			} else {
				color.Cyan("\tLabel: " + color.RedString("No Label entry found for device: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Color != "" {
				color.Cyan("\tColor: " + color.YellowString("%s", responseObject.Color))
			} else {
				color.Cyan("\tColor: " + color.RedString("No Color entry found for device: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Length != 0 {
				color.Cyan("\tLength: " + color.YellowString("%s", responseObject.Length))
			} else {
				color.Cyan("\tLength: " + color.RedString("No Length entry found for device: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.LengthUnit.Value != "" {
				color.Cyan("\tLength Unit: ")
				color.Cyan("\t  Length Unit: " + color.YellowString("%s", responseObject.LengthUnit.Value))
				color.Cyan("\t  Length Unit: " + color.YellowString("%s", responseObject.LengthUnit.Label))
			} else {
				color.Cyan("\tLength Unit: " + color.RedString("No Length Unit entry found for device: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: " + color.YellowString("%s", responseObject.Description))
			} else {
				color.Cyan("\tDescription: " + color.RedString("No Description entry found for device: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Comments != "" {
				color.Cyan("\tComments:" + color.YellowString("%s", responseObject.Comments))
			} else {
				color.Cyan("\tComments: " + color.RedString("No Comments entry found for device: ") + color.YellowString("%s", responseObject.Display))
			}
			for _, tag := range responseObject.Tags {
				if tag.Id != 0 {
					color.Cyan("\tTags: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
					color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
				} else {
					color.Cyan("\tTags: " + color.RedString("No Tags entry found for device: ") + color.YellowString("%s", responseObject.Display))
				}
			}
			color.Cyan("\tCreated: " + color.YellowString("%s", responseObject.Created))
			color.Cyan("\tLast Updated: " + color.YellowString("%s\n", responseObject.LastUpdated))
		} else {
			color.Cyan("\n  ABC Cables: " + color.RedString("No cables entries found for on server for ID: %d. Exiting...\n", id))
		}
	},
}

// Define environment flag
func init() {

	// Here you will define your flags and configuration settings.
	GetDcimCablesByIdCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment ('development' or 'production')")
	err := GetDcimCablesByIdCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking env flag as required: %s - for GetDcimCablesByIdCmd", err)
	}

	GetDcimCablesByIdCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the cable")
	err = GetDcimCablesByIdCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking id flag as required: %s - for GetDcimCablesByIdCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimCablesByIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimCablesByIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

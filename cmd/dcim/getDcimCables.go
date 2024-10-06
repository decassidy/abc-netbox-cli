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

// cables represents the structure of the response for the cables API endpoint.
type cables struct {
	Count    uint    `json:"count"`
	Next     *string `json:"next"`
	Previous string  `json:"previous"`
	Results  []struct {
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
	} `json:"results"`
}

var responseObjectCables = new(cables)

// GetDcimCablesCmd represents the getDcimCables command
var GetDcimCablesCmd = &cobra.Command{
	Use:   "getDcimCables",
	Short: "Get a list of cable objects.",
	Long: `
ABC Netbox Automation Tools:
  Get a list of cable objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiConnectionNonID(responseObjectCables, "GET", "cmd.dcim.dcim_api_url.cables")

		if responseObjectCables.Count > 0 {
			color.Cyan("\n  Total ABC Cables: "+color.YellowString("%v"), responseObjectCables.Count)

			for _, cable := range responseObjectCables.Results {
				display := fmt.Sprintf("    ABC Cable: %s\n", color.YellowString(cable.Display))
				equals := strings.Repeat("=", len(display))
				color.Cyan("\n  " + equals + "\n")
				color.Cyan(display)
				color.Cyan("  " + equals + "\n")
				color.Cyan("\tID: " + color.YellowString("%d", cable.Id))
				color.Cyan("\tURL: " + color.YellowString("%s", cable.Url))
				color.Cyan("\tDisplay: " + color.YellowString("%s", cable.Display))
				color.Cyan("\tType: " + color.YellowString("%s", cable.Type))
				for _, term := range cable.ATerminations {
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
							color.Cyan("\t  Object: " + color.RedString("No Object entry found for cable: ") + color.YellowString("%s", cable.Display))
						}
						if term.Object.Circuit.Id != 0 {
							color.Cyan("\t    Circuit: ")
							color.Cyan("\t      ID: " + color.YellowString("%d", term.Object.Circuit.Id))
							color.Cyan("\t      URL: " + color.YellowString("%s", term.Object.Circuit.Url))
							color.Cyan("\t      Display: " + color.YellowString("%s", term.Object.Circuit.Display))
							color.Cyan("\t      CID: " + color.YellowString("%s", term.Object.Circuit.Cid))
						} else {
							color.Cyan("\t    Circuit: " + color.RedString("No circuit entry found for cable: ") + color.YellowString("%s", cable.Display))
						}
						if term.Object.TermSide != "" {
							color.Cyan("\t    Term Side: " + color.YellowString("%s", term.Object.TermSide))
						} else {
							color.Cyan("\t    Term Side: " + color.RedString("No term side entry found for cable: ") + color.YellowString("%s", cable.Display))
						}
						if term.Object.Cable.Id > 0 {
							color.Cyan("\t    Cable: ")
							color.Cyan("\t      ID: " + color.YellowString("%d", term.Object.Cable.Id))
							color.Cyan("\t      URL: " + color.YellowString("%s", term.Object.Cable.Url))
							color.Cyan("\t      Display: " + color.YellowString("%s", term.Object.Cable.Display))
							if term.Object.Cable.Label != "" {
								color.Cyan("\t      Label: " + color.YellowString("%s", term.Object.Cable.Label))
							} else {
								color.Cyan("\t      Label: " + color.RedString("No label entry found for cable: ") + color.YellowString("%s", cable.Display))
							}
							if term.Object.Cable.Description != "" {
								color.Cyan("\t      Description: " + color.YellowString("%s", term.Object.Cable.Description))
							} else {
								color.Cyan("\t      Description: " + color.RedString("No description entry found for cable: ") + color.YellowString("%s", cable.Display))
							}
						} else {
							color.Cyan("\t    Cable: " + color.RedString("No cable entry found for cable: ") + color.YellowString("%s", cable.Display))
						}
						color.Cyan("\t    Occupied: " + color.YellowString("%t", term.Object.Occupied))
					} else {
						color.Cyan("\tA Terminations: " + color.RedString("No A Terminations entry found for cable: ") + color.YellowString("%s", cable.Display))
					}
				}
				for _, bTermination := range cable.BTerminations {
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
								color.Cyan("\t      Label: " + color.RedString("No label entry found for cable: ") + color.YellowString("%s", cable.Display))
							}
							if bTermination.Object.Cable.Description != "" {
								color.Cyan("\t      Description: " + color.YellowString("%s", bTermination.Object.Cable.Description))
							} else {
								color.Cyan("\t      Description: " + color.RedString("No description entry found for cable: ") + color.YellowString("%s", cable.Display))
							}
						} else {
							color.Cyan("\t    Cable: " + color.RedString("No cable entry found for cable: ") + color.YellowString("%s", cable.Display))
						}
						color.Cyan("\t    Occupied: " + color.YellowString("%v", bTermination.Object.Occupied))
					} else {
						color.Cyan("\tB Terminations: " + color.RedString("No B Terminations entry found for device: ") + color.YellowString("%s", cable.Display))
					}
				}
				if cable.Status.Value != "" {
					color.Cyan("\tStatus: ")
					color.Cyan("\t  Value: " + color.YellowString("%v", cable.Status.Value))
					color.Cyan("\t  Label: " + color.YellowString("%v", cable.Status.Label))
				} else {
					color.Cyan("\tStatus: " + color.RedString("No Status entry found for device: ") + color.YellowString("%s", cable.Display))
				}
				if cable.Tenant.Id != 0 {
					color.Cyan("\tTenant: ")
					color.Cyan("\t  ID: " + color.YellowString("%d", cable.Tenant.Id))
					color.Cyan("\t  URL: " + color.YellowString("%s", cable.Tenant.Url))
					color.Cyan("\t  Display: " + color.YellowString("%s", cable.Tenant.Display))
					color.Cyan("\t  Name: " + color.YellowString("%s", cable.Tenant.Name))
					color.Cyan("\t  Slug: " + color.YellowString("%s", cable.Tenant.Slug))
				} else {
					color.Cyan("\tTenant: " + color.RedString("No Tenant entry found for device: ") + color.YellowString("%s", cable.Display))
				}
				if cable.Label != "" {
					color.Cyan("\tLabel: " + color.YellowString("%s", cable.Label))
				} else {
					color.Cyan("\tLabel: " + color.RedString("No Label entry found for device: ") + color.YellowString("%s", cable.Display))
				}
				if cable.Color != "" {
					color.Cyan("\tColor: " + color.YellowString("%s", cable.Color))
				} else {
					color.Cyan("\tColor: " + color.RedString("No Color entry found for device: ") + color.YellowString("%s", cable.Display))
				}
				if cable.Length != 0 {
					color.Cyan("\tLength: " + color.YellowString("%.2f", cable.Length))
				} else {
					color.Cyan("\tLength: " + color.RedString("No Length entry found for device: ") + color.YellowString("%s", cable.Display))
				}
				if cable.LengthUnit.Value != "" {
					color.Cyan("\tLength Unit: ")
					color.Cyan("\t  Length Unit: " + color.YellowString("%s", cable.LengthUnit.Value))
					color.Cyan("\t  Length Unit: " + color.YellowString("%s", cable.LengthUnit.Label))
				} else {
					color.Cyan("\tLength Unit: " + color.RedString("No Length Unit entry found for device: ") + color.YellowString("%s", cable.Display))
				}
				if cable.Description != "" {
					color.Cyan("\tDescription: " + color.YellowString("%s", cable.Description))
				} else {
					color.Cyan("\tDescription: " + color.RedString("No Description entry found for device: ") + color.YellowString("%s", cable.Display))
				}
				if cable.Comments != "" {
					color.Cyan("\tComments:" + color.YellowString("%s", cable.Comments))
				} else {
					color.Cyan("\tComments: " + color.RedString("No Comments entry found for device: ") + color.YellowString("%s", cable.Display))
				}
				for _, tag := range cable.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
						color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
						color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
						color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
						color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
						color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
					} else {
						color.Cyan("\tTags: " + color.RedString("No Tags entry found for device: ") + color.YellowString("%s", cable.Display))
					}
				}
				color.Cyan("\tCreated: " + color.YellowString("%v", cable.Created))
				color.Cyan("\tLast Updated: " + color.YellowString("%v\n", cable.LastUpdated))
			}
			for responseObjectCables.Next != nil {
				nextPageCables()
			}
			if responseObjectCables.Next == nil {
				display := color.HiGreenString("\tAll Netbox cable objects have been successfully displayed...")
				equals := strings.Repeat("*", len(display))
				color.HiGreen("\n  " + equals)
				color.Cyan(display)
				color.HiGreen("  " + equals + "\n")
			}
		} else {
			color.Cyan("  Total ABC Cables: " + color.RedString("No cables entries found for on server. Exiting...\n"))
		}
	},
}

func ApiConnectionNextPageCables[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	fullAPIPath := *responseObjectCables.Next

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

func nextPageCables() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\tDo you want to continue to the next page of cable objects? ['Y' or 'yes'] or ['n' or 'no']: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "yes":
		ApiConnectionNextPageCables(responseObjectCables, "GET", *responseObjectCables.Next)
		displayCablesOutput()
	case "n", "no":
		color.HiMagenta("\tExiting the ABC-netbox-cli application...\n")
		os.Exit(0)
	default:
		color.Cyan("Invalid input, Please type ['Y' or 'yes'] or ['n' or 'no'] ")
	}
}

func displayCablesOutput() {
	for _, cable := range responseObjectCables.Results {
		display := fmt.Sprintf("    ABC Cable: %s\n", color.YellowString(cable.Display))
		equals := strings.Repeat("=", len(display))
		color.Cyan("\n  " + equals + "\n")
		color.Cyan(display)
		color.Cyan("  " + equals + "\n")
		color.Cyan("\tID: " + color.YellowString("%d", cable.Id))
		color.Cyan("\tURL: " + color.YellowString("%s", cable.Url))
		color.Cyan("\tDisplay: " + color.YellowString("%s", cable.Display))
		color.Cyan("\tType: " + color.YellowString("%s", cable.Type))
		for _, term := range cable.ATerminations {
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
					color.Cyan("\t  Object: " + color.RedString("No Object entry found for cable: ") + color.YellowString("%s", cable.Display))
				}
				if term.Object.Circuit.Id != 0 {
					color.Cyan("\t    Circuit: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", term.Object.Circuit.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", term.Object.Circuit.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", term.Object.Circuit.Display))
					color.Cyan("\t      CID: " + color.YellowString("%s", term.Object.Circuit.Cid))
				} else {
					color.Cyan("\t    Circuit: " + color.RedString("No circuit entry found for cable: ") + color.YellowString("%s", cable.Display))
				}
				if term.Object.TermSide != "" {
					color.Cyan("\t    Term Side: " + color.YellowString("%s", term.Object.TermSide))
				} else {
					color.Cyan("\t    Term Side: " + color.RedString("No term side entry found for cable: ") + color.YellowString("%s", cable.Display))
				}
				if term.Object.Cable.Id > 0 {
					color.Cyan("\t    Cable: ")
					color.Cyan("\t      ID: " + color.YellowString("%d", term.Object.Cable.Id))
					color.Cyan("\t      URL: " + color.YellowString("%s", term.Object.Cable.Url))
					color.Cyan("\t      Display: " + color.YellowString("%s", term.Object.Cable.Display))
					if term.Object.Cable.Label != "" {
						color.Cyan("\t      Label: " + color.YellowString("%s", term.Object.Cable.Label))
					} else {
						color.Cyan("\t      Label: " + color.RedString("No label entry found for cable: ") + color.YellowString("%s", cable.Display))
					}
					if term.Object.Cable.Description != "" {
						color.Cyan("\t      Description: " + color.YellowString("%s", term.Object.Cable.Description))
					} else {
						color.Cyan("\t      Description: " + color.RedString("No description entry found for cable: ") + color.YellowString("%s", cable.Display))
					}
				} else {
					color.Cyan("\t    Cable: " + color.RedString("No cable entry found for cable: ") + color.YellowString("%s", cable.Display))
				}
				color.Cyan("\t    Occupied: " + color.YellowString("%t", term.Object.Occupied))
			} else {
				color.Cyan("\tA Terminations: " + color.RedString("No A Terminations entry found for cable: ") + color.YellowString("%s", cable.Display))
			}
		}
		for _, bTermination := range cable.BTerminations {
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
						color.Cyan("\t      Label: " + color.RedString("No label entry found for cable: ") + color.YellowString("%s", cable.Display))
					}
					if bTermination.Object.Cable.Description != "" {
						color.Cyan("\t      Description: " + color.YellowString("%s", bTermination.Object.Cable.Description))
					} else {
						color.Cyan("\t      Description: " + color.RedString("No description entry found for cable: ") + color.YellowString("%s", cable.Display))
					}
				} else {
					color.Cyan("\t    Cable: " + color.RedString("No cable entry found for cable: ") + color.YellowString("%s", cable.Display))
				}
				color.Cyan("\t    Occupied: " + color.YellowString("%v", bTermination.Object.Occupied))
			} else {
				color.Cyan("\tB Terminations: " + color.RedString("No B Terminations entry found for device: ") + color.YellowString("%s", cable.Display))
			}
		}
		if cable.Status.Value != "" {
			color.Cyan("\tStatus: ")
			color.Cyan("\t  Value: " + color.YellowString("%v", cable.Status.Value))
			color.Cyan("\t  Label: " + color.YellowString("%v", cable.Status.Label))
		} else {
			color.Cyan("\tStatus: " + color.RedString("No Status entry found for device: ") + color.YellowString("%s", cable.Display))
		}
		if cable.Tenant.Id != 0 {
			color.Cyan("\tTenant: ")
			color.Cyan("\t  ID: " + color.YellowString("%d", cable.Tenant.Id))
			color.Cyan("\t  URL: " + color.YellowString("%s", cable.Tenant.Url))
			color.Cyan("\t  Display: " + color.YellowString("%s", cable.Tenant.Display))
			color.Cyan("\t  Name: " + color.YellowString("%s", cable.Tenant.Name))
			color.Cyan("\t  Slug: " + color.YellowString("%s", cable.Tenant.Slug))
		} else {
			color.Cyan("\tTenant: " + color.RedString("No Tenant entry found for device: ") + color.YellowString("%s", cable.Display))
		}
		if cable.Label != "" {
			color.Cyan("\tLabel: " + color.YellowString("%s", cable.Label))
		} else {
			color.Cyan("\tLabel: " + color.RedString("No Label entry found for device: ") + color.YellowString("%s", cable.Display))
		}
		if cable.Color != "" {
			color.Cyan("\tColor: " + color.YellowString("%s", cable.Color))
		} else {
			color.Cyan("\tColor: " + color.RedString("No Color entry found for device: ") + color.YellowString("%s", cable.Display))
		}
		if cable.Length != 0 {
			color.Cyan("\tLength: " + color.YellowString("%.2f", cable.Length))
		} else {
			color.Cyan("\tLength: " + color.RedString("No Length entry found for device: ") + color.YellowString("%s", cable.Display))
		}
		if cable.LengthUnit.Value != "" {
			color.Cyan("\tLength Unit: ")
			color.Cyan("\t  Length Unit: " + color.YellowString("%s", cable.LengthUnit.Value))
			color.Cyan("\t  Length Unit: " + color.YellowString("%s", cable.LengthUnit.Label))
		} else {
			color.Cyan("\tLength Unit: " + color.RedString("No Length Unit entry found for device: ") + color.YellowString("%s", cable.Display))
		}
		if cable.Description != "" {
			color.Cyan("\tDescription: " + color.YellowString("%s", cable.Description))
		} else {
			color.Cyan("\tDescription: " + color.RedString("No Description entry found for device: ") + color.YellowString("%s", cable.Display))
		}
		if cable.Comments != "" {
			color.Cyan("\tComments:" + color.YellowString("%s", cable.Comments))
		} else {
			color.Cyan("\tComments: " + color.RedString("No Comments entry found for device: ") + color.YellowString("%s", cable.Display))
		}
		for _, tag := range cable.Tags {
			if tag.Id != 0 {
				color.Cyan("\tTags: ")
				color.Cyan("\t  ID: " + color.YellowString("%d", tag.Id))
				color.Cyan("\t  URL: " + color.YellowString("%s", tag.Url))
				color.Cyan("\t  Display: " + color.YellowString("%s", tag.Display))
				color.Cyan("\t  Name: " + color.YellowString("%s", tag.Name))
				color.Cyan("\t  Slug: " + color.YellowString("%s", tag.Slug))
				color.Cyan("\t  Color: " + color.YellowString("%s", tag.Color))
			} else {
				color.Cyan("\tTags: " + color.RedString("No Tags entry found for device: ") + color.YellowString("%s", cable.Display))
			}
		}
		color.Cyan("\tCreated: " + color.YellowString("%v", cable.Created))
		color.Cyan("\tLast Updated: " + color.YellowString("%v", cable.LastUpdated))
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	GetDcimCablesCmd.Flags().StringVarP(&serverEnv, "env", "", "", "Environment to get DcimCables ('development', 'production').")
	err := GetDcimCablesCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetDcimCablesCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDcimCablesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDcimCablesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

package circuits

import (
	"github.com/decassidy/abc-netbox-cli/cmd/dcim"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
)

type circuitsByID struct {
	Id       uint   `json:"id"`
	Url      string `json:"url"`
	Display  string `json:"display"`
	Cid      string `json:"cid"`
	Provider struct {
		dcim.CommonFieldsSlug
	} `json:"provider"`
	ProviderAccount struct {
		dcim.CommonFieldsNoSlug
		Account string `json:"account"`
	} `json:"provider_account"`
	Type struct {
		dcim.CommonFieldsSlug
	} `json:"type"`
	Status struct {
		dcim.ValueLabel
	} `json:"status"`
	Tenant struct {
		dcim.CommonFieldsSlug
	} `json:"tenant"`
	InstallDate     string `json:"install_date"`
	TerminationDate string `json:"termination_date"`
	CommitRate      uint   `json:"commit_rate"`
	Description     string `json:"description"`
	TerminationA    struct {
		Id              uint                  `json:"id"`
		Url             string                `json:"url"`
		Display         string                `json:"display"`
		Site            dcim.CommonFieldsSlug `json:"site"`
		ProviderNetwork struct {
			dcim.CommonFieldsNoSlug
		} `json:"provider_network"`
		PortSpeed     uint   `json:"port_speed"`
		UpstreamSpeed uint   `json:"upstream_speed"`
		XconnectId    string `json:"xconnect_id"`
		Description   string `json:"description"`
	} `json:"termination_a"`
	TerminationZ struct {
		Id              uint                  `json:"id"`
		Url             string                `json:"url"`
		Display         string                `json:"display"`
		Site            dcim.CommonFieldsSlug `json:"site"`
		ProviderNetwork struct {
			dcim.CommonFieldsNoSlug
		} `json:"provider_network"`
		PortSpeed     uint   `json:"port_speed"`
		UpstreamSpeed uint   `json:"upstream_speed"`
		XconnectId    string `json:"xconnect_id"`
		Description   string `json:"description"`
	} `json:"termination_z"`
	Comments string `json:"comments"`
	Tags     []struct {
		dcim.CommonFieldsSlug
		Color string `json:"color"`
	} `json:"tags"`
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

// GetCircuitsCircuitsByIDCmd represents the circuitByID command
var GetCircuitsCircuitsByIDCmd = &cobra.Command{
	Use:   "getCircuitsCircuitsById",
	Short: "Get a ABC Circuit object by ID.",
	Long: `
ABC Netbox Automation Tools:
  Get a ABC Circuit object by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(circuitsByID)
		apiConnectionID(responseObject, "GET", "cmd.circuits.circuits_api_url.circuits_id")

		if responseObject.Id != 0 {
			color.Cyan("\n============================================================================")
			color.Cyan("\n\tABC Circuit Name: "+color.YellowString("%s\n"), responseObject.Display)
			color.Cyan("============================================================================\n")
			color.Cyan("\tID: "+color.YellowString("%d"), responseObject.Id)
			color.Cyan("\tURL: "+color.YellowString("%s"), responseObject.Url)
			color.Cyan("\tDisplay: "+color.YellowString("%s"), responseObject.Display)
			color.Cyan("\tCID: "+color.YellowString("%s"), responseObject.Cid)
			if responseObject.Provider.Id != 0 {
				color.Cyan("\tProvider: ")
				color.Cyan("\t  ID: "+color.YellowString("%d"), responseObject.Provider.Id)
				color.Cyan("\t  URL: "+color.YellowString("%s"), responseObject.Provider.Url)
				color.Cyan("\t  Display: "+color.YellowString("%s"), responseObject.Provider.Display)
				color.Cyan("\t  Name: "+color.YellowString("%s"), responseObject.Provider.Name)
				color.Cyan("\t  Slug: "+color.YellowString("%s"), responseObject.Provider.Slug)
			} else {
				color.Cyan("\tProvider: " + color.RedString("No provider found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.ProviderAccount.Id != 0 {
				color.Cyan("\tProvider Account: ")
				color.Cyan("\t  Provider Account ID: "+color.YellowString("%d"), responseObject.ProviderAccount.Id)
				color.Cyan("\t  Provider Account URL: "+color.YellowString("%s"), responseObject.ProviderAccount.Url)
				color.Cyan("\t  Provider Account Display: "+color.YellowString("%s"), responseObject.ProviderAccount.Display)
				color.Cyan("\t  Provider Account Name: "+color.YellowString("%s"), responseObject.ProviderAccount.Name)
				color.Cyan("\t  Provider Account Account: "+color.YellowString("%s"), responseObject.ProviderAccount.Account)
			} else {
				color.Cyan("\tProvider Account: " + color.RedString("No provider account found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Type.Id != 0 {
				color.Cyan("\tType: ")
				color.Cyan("\t  ID: "+color.YellowString("%d"), responseObject.Type.Id)
				color.Cyan("\t  URL: "+color.YellowString("%s"), responseObject.Type.Url)
				color.Cyan("\t  Display: "+color.YellowString("%s"), responseObject.Type.Display)
				color.Cyan("\t  Name: "+color.YellowString("%s"), responseObject.Type.Name)
				color.Cyan("\t  Slug: "+color.YellowString("%s"), responseObject.Type.Slug)
			} else {
				color.Cyan("\tType: " + color.RedString("No provider type found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Status.Value != "" {
				color.Cyan("\tStatus: ")
				color.Cyan("\t  Status Value: "+color.YellowString("%s"), responseObject.Status.Value)
				color.Cyan("\t  Status Label: "+color.YellowString("%s"), responseObject.Status.Label)
			} else {
				color.Cyan("\tStatus: " + color.RedString("No status found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Tenant.Id != 0 {
				color.Cyan("\tTenant: ")
				color.Cyan("\t  ID: "+color.YellowString("%d"), responseObject.Tenant.Id)
				color.Cyan("\t  URL: "+color.YellowString("%s"), responseObject.Tenant.Url)
				color.Cyan("\t  Display: "+color.YellowString("%s"), responseObject.Tenant.Display)
				color.Cyan("\t  Name: "+color.YellowString("%s"), responseObject.Tenant.Name)
				color.Cyan("\t  Slug: "+color.YellowString("%s"), responseObject.Tenant.Slug)
			} else {
				color.Cyan("\tTenant: " + color.RedString("No tenant found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.InstallDate != "" {
				color.Cyan("\tInstall Date: "+color.YellowString("%s"), responseObject.InstallDate)
			} else {
				color.Cyan("\tInstall Date: " + color.RedString("No install date found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.TerminationDate != "" {
				color.Cyan("\tTermination Date: "+color.YellowString("%s"), responseObject.TerminationDate)
			} else {
				color.Cyan("\tTermination Date: " + color.RedString("No termination date found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.CommitRate != 0 {
				color.Cyan("\tCommit Rate: "+color.YellowString("%d"), responseObject.CommitRate)
			} else {
				color.Cyan("\tCommit Rate: " + color.RedString("No commit rate found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Description != "" {
				color.Cyan("\tDescription: "+color.YellowString("%s"), responseObject.Description)
			} else {
				color.Cyan("\tDescription: " + color.RedString("No description found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.TerminationA.Id != 0 {
				color.Cyan("\tTermination A: ")
				color.Cyan("\t  ID: "+color.YellowString("%d"), responseObject.TerminationA.Id)
				color.Cyan("\t  URL: "+color.YellowString("%s"), responseObject.TerminationA.Url)
				color.Cyan("\t  Display: "+color.YellowString("%s"), responseObject.TerminationA.Display)
				if responseObject.TerminationA.Site.Id != 0 {
					color.Cyan("\t  Site: ")
					color.Cyan("\t    ID: "+color.YellowString("%d"), responseObject.TerminationA.Site.Id)
					color.Cyan("\t    URL: "+color.YellowString("%s"), responseObject.TerminationA.Site.Url)
					color.Cyan("\t    Display: "+color.YellowString("%s"), responseObject.TerminationA.Site.Display)
					color.Cyan("\t    Name: "+color.YellowString("%s"), responseObject.TerminationA.Site.Name)
					color.Cyan("\t    Slug: "+color.YellowString("%s"), responseObject.TerminationA.Site.Slug)
				} else {
					color.Cyan("\t  Site: " + color.RedString("No termination A site found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.TerminationA.ProviderNetwork.Id != 0 {
					color.Cyan("\t  Provider Network: ")
					color.Cyan("\t    ID: "+color.YellowString("%d"), responseObject.TerminationA.ProviderNetwork.Id)
					color.Cyan("\t    URL: "+color.YellowString("%s"), responseObject.TerminationA.ProviderNetwork.Url)
					color.Cyan("\t    Display: "+color.YellowString("%s"), responseObject.TerminationA.ProviderNetwork.Display)
					color.Cyan("\t    Name: "+color.YellowString("%s"), responseObject.TerminationA.ProviderNetwork.Name)
				} else {
					color.Cyan("\t  Provider Network: " + color.RedString("No termination A provider network found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.TerminationA.PortSpeed != 0 {
					color.Cyan("\t  Port Speed: "+color.YellowString("%d"), responseObject.TerminationA.PortSpeed)
				} else {
					color.Cyan("\t  Port Speed: " + color.RedString("No termination A port speed found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.TerminationA.UpstreamSpeed != 0 {
					color.Cyan("\t  Upstream Speed: "+color.YellowString("%d"), responseObject.TerminationA.UpstreamSpeed)
				} else {
					color.Cyan("\t  Upstream Speed: " + color.RedString("No termination A upstream speed found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.TerminationA.XconnectId != "" {
					color.Cyan("\t  XConnect ID: "+color.YellowString("%s"), responseObject.TerminationA.XconnectId)
				} else {
					color.Cyan("\t  XConnect ID: " + color.RedString("No termination A xconnect ID found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.TerminationA.Description != "" {
					color.Cyan("\t  Description: "+color.YellowString("%s"), responseObject.TerminationA.Description)
				} else {
					color.Cyan("\t  Description: " + color.RedString("No termination A description found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
			} else {
				color.Cyan("\tTermination A: " + color.RedString("No termination A found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.TerminationZ.Id != 0 {
				color.Cyan("\tTermination Z: ")
				color.Cyan("\t  ID: "+color.YellowString("%d"), responseObject.TerminationZ.Id)
				color.Cyan("\t  URL: "+color.YellowString("%s"), responseObject.TerminationZ.Url)
				color.Cyan("\t  Display: "+color.YellowString("%s"), responseObject.TerminationZ.Display)
				if responseObject.TerminationZ.Site.Id != 0 {
					color.Cyan("\t  Site: ")
					color.Cyan("\t    ID: "+color.YellowString("%d"), responseObject.TerminationZ.Site.Id)
					color.Cyan("\t    URL: "+color.YellowString("%s"), responseObject.TerminationZ.Site.Url)
					color.Cyan("\t    Display: "+color.YellowString("%s"), responseObject.TerminationZ.Site.Display)
					color.Cyan("\t    Name: "+color.YellowString("%s"), responseObject.TerminationZ.Site.Name)
					color.Cyan("\t    Slug: "+color.YellowString("%s"), responseObject.TerminationZ.Site.Slug)
				} else {
					color.Cyan("\t  Site: " + color.RedString("No termination Z site found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.TerminationZ.ProviderNetwork.Id != 0 {
					color.Cyan("\t  Provider Network: ")
					color.Cyan("\t    ID: "+color.YellowString("%d"), responseObject.TerminationZ.ProviderNetwork.Id)
					color.Cyan("\t    URL: "+color.YellowString("%s"), responseObject.TerminationZ.ProviderNetwork.Url)
					color.Cyan("\t    Display: "+color.YellowString("%s"), responseObject.TerminationZ.ProviderNetwork.Display)
					color.Cyan("\t    Name: "+color.YellowString("%s"), responseObject.TerminationZ.ProviderNetwork.Name)
				} else {
					color.Cyan("\t  Provider Network: " + color.RedString("No termination Z provider network found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.TerminationZ.PortSpeed != 0 {
					color.Cyan("\t  Port Speed: "+color.YellowString("%d"), responseObject.TerminationZ.PortSpeed)
				} else {
					color.Cyan("\t  Port Speed: " + color.RedString("No termination Z port speed found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.TerminationZ.UpstreamSpeed != 0 {
					color.Cyan("\t  Upstream Speed: "+color.YellowString("%d"), responseObject.TerminationZ.UpstreamSpeed)
				} else {
					color.Cyan("\t  Upstream Speed: " + color.RedString("No termination Z upstream speed found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.TerminationZ.XconnectId != "" {
					color.Cyan("\t  XConnect ID: "+color.YellowString("%s"), responseObject.TerminationZ.XconnectId)
				} else {
					color.Cyan("\t  XConnect ID: " + color.RedString("No termination Z xconnect ID found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
				if responseObject.TerminationZ.Description != "" {
					color.Cyan("\t  Description: "+color.YellowString("%s"), responseObject.TerminationZ.Description)
				} else {
					color.Cyan("\t  Description: " + color.RedString("No termination Z description found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
			} else {
				color.Cyan("\tTermination Z: " + color.RedString("No termination Z found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			if responseObject.Comments != "" {
				color.Cyan("\tComments: "+color.YellowString("%s"), responseObject.Comments)
			} else {
				color.Cyan("\tComments: " + color.RedString("No comments found for circuit: ") + color.YellowString("%s", responseObject.Display))
			}
			for _, tag := range responseObject.Tags {
				if tag.Id != 0 {
					color.Cyan("\tTags: ")
					color.Cyan("\t  ID: "+color.YellowString("%v"), tag.Id)
					color.Cyan("\t  URL: "+color.YellowString("%v"), tag.Url)
					color.Cyan("\t  Display: "+color.YellowString("%v"), tag.Display)
					color.Cyan("\t  Name: "+color.YellowString("%v"), tag.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%v"), tag.Slug)
					color.Cyan("\t  Color: "+color.YellowString("%v"), tag.Color)
				} else {
					color.Cyan("\tTags: " + color.RedString("No tags found for circuit: ") + color.YellowString("%s", responseObject.Display))
				}
			}
			color.Cyan("\tABC Circuit Provider Created: "+color.YellowString("%s"), responseObject.Created)
			color.Cyan("\tABC Circuit Provider Last Updated: "+color.YellowString("%s"), responseObject.LastUpdated)
		} else {
			color.Red("  Doh! No circuit object found on server for ID: "+color.YellowString("%d\n"), id)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetCircuitsCircuitsByIDCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment from which the server is running")
	err := GetCircuitsCircuitsByIDCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Could not mark 'env' flag as required %s - for GetCircuitsCircuitsByIDCmd", err)
	}

	GetCircuitsCircuitsByIDCmd.Flags().IntVarP(&id, "id", "", 0, "ID of the circuit object to get")
	err = GetCircuitsCircuitsByIDCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Could not mark 'id' flag as required %s - for GetCircuitsCircuitsByIDCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// circuitByIDCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// circuitByIDCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

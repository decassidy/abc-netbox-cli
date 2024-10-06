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

type circuit struct {
	Count    uint   `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
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
	} `json:"results"`
}

// GetCircuitsCircuitsCmd GetCircuitsCircuitsCmd represents the allCircuits command
var GetCircuitsCircuitsCmd = &cobra.Command{
	Use:   "getCircuitsCircuits",
	Short: "Get a list of Circuits objects.",
	Long: `
ABC Netbox Automation Tools:
  Get a list of Circuits objects`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(circuit)
		ApiConnectionNonID(responseObject, "GET", "cmd.circuits.circuits_api_url.circuits")

		if responseObject.Count != 0 {
			color.Cyan("\nABC Total Circuits in Netbox: "+color.YellowString("%d"), responseObject.Count)
			for _, circuit := range responseObject.Results {
				color.Cyan("\n============================================================================")
				color.Cyan("\n\tABC Circuit Name: "+color.YellowString("%s\n"), circuit.Display)
				color.Cyan("============================================================================\n")
				color.Cyan("\tID: "+color.YellowString("%d"), circuit.Id)
				color.Cyan("\tURL: "+color.YellowString("%s"), circuit.Url)
				color.Cyan("\tDisplay: "+color.YellowString("%s"), circuit.Display)
				color.Cyan("\tCID: "+color.YellowString("%s"), circuit.Cid)
				if circuit.Provider.Id != 0 {
					color.Cyan("\tProvider: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), circuit.Provider.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), circuit.Provider.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), circuit.Provider.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), circuit.Provider.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%s"), circuit.Provider.Slug)
				} else {
					color.Cyan("\tProvider: " + color.RedString("No provider found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.ProviderAccount.Id != 0 {
					color.Cyan("\tProvider Account: ")
					color.Cyan("\t  Provider Account ID: "+color.YellowString("%d"), circuit.ProviderAccount.Id)
					color.Cyan("\t  Provider Account URL: "+color.YellowString("%s"), circuit.ProviderAccount.Url)
					color.Cyan("\t  Provider Account Display: "+color.YellowString("%s"), circuit.ProviderAccount.Display)
					color.Cyan("\t  Provider Account Name: "+color.YellowString("%s"), circuit.ProviderAccount.Name)
					color.Cyan("\t  Provider Account Account: "+color.YellowString("%s"), circuit.ProviderAccount.Account)
				} else {
					color.Cyan("\tProvider Account: " + color.RedString("No provider account found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.Type.Id != 0 {
					color.Cyan("\tType: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), circuit.Type.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), circuit.Type.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), circuit.Type.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), circuit.Type.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%s"), circuit.Type.Slug)
				} else {
					color.Cyan("\tType: " + color.RedString("No provider type found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.Status.Value != "" {
					color.Cyan("\tStatus: ")
					color.Cyan("\t  Status Value: "+color.YellowString("%s"), circuit.Status.Value)
					color.Cyan("\t  Status Label: "+color.YellowString("%s"), circuit.Status.Label)
				} else {
					color.Cyan("\tStatus: " + color.RedString("No status found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.Tenant.Id != 0 {
					color.Cyan("\tTenant: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), circuit.Tenant.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), circuit.Tenant.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), circuit.Tenant.Display)
					color.Cyan("\t  Name: "+color.YellowString("%s"), circuit.Tenant.Name)
					color.Cyan("\t  Slug: "+color.YellowString("%s"), circuit.Tenant.Slug)
				} else {
					color.Cyan("\tTenant: " + color.RedString("No tenant found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.InstallDate != "" {
					color.Cyan("\tInstall Date: "+color.YellowString("%s"), circuit.InstallDate)
				} else {
					color.Cyan("\tInstall Date: " + color.RedString("No install date found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.TerminationDate != "" {
					color.Cyan("\tTermination Date: "+color.YellowString("%s"), circuit.TerminationDate)
				} else {
					color.Cyan("\tTermination Date: " + color.RedString("No termination date found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.CommitRate != 0 {
					color.Cyan("\tCommit Rate: "+color.YellowString("%d"), circuit.CommitRate)
				} else {
					color.Cyan("\tCommit Rate: " + color.RedString("No commit rate found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.Description != "" {
					color.Cyan("\tDescription: "+color.YellowString("%s"), circuit.Description)
				} else {
					color.Cyan("\tDescription: " + color.RedString("No description found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.TerminationA.Id != 0 {
					color.Cyan("\tTermination A: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), circuit.TerminationA.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), circuit.TerminationA.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), circuit.TerminationA.Display)
					if circuit.TerminationA.Site.Id != 0 {
						color.Cyan("\t  Site: ")
						color.Cyan("\t    ID: "+color.YellowString("%d"), circuit.TerminationA.Site.Id)
						color.Cyan("\t    URL: "+color.YellowString("%s"), circuit.TerminationA.Site.Url)
						color.Cyan("\t    Display: "+color.YellowString("%s"), circuit.TerminationA.Site.Display)
						color.Cyan("\t    Name: "+color.YellowString("%s"), circuit.TerminationA.Site.Name)
						color.Cyan("\t    Slug: "+color.YellowString("%s"), circuit.TerminationA.Site.Slug)
					} else {
						color.Cyan("\t  Site: " + color.RedString("No termination A site found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
					if circuit.TerminationA.ProviderNetwork.Id != 0 {
						color.Cyan("\t  Provider Network: ")
						color.Cyan("\t    ID: "+color.YellowString("%d"), circuit.TerminationA.ProviderNetwork.Id)
						color.Cyan("\t    URL: "+color.YellowString("%s"), circuit.TerminationA.ProviderNetwork.Url)
						color.Cyan("\t    Display: "+color.YellowString("%s"), circuit.TerminationA.ProviderNetwork.Display)
						color.Cyan("\t    Name: "+color.YellowString("%s"), circuit.TerminationA.ProviderNetwork.Name)
					} else {
						color.Cyan("\t  Provider Network: " + color.RedString("No termination A provider network found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
					if circuit.TerminationA.PortSpeed != 0 {
						color.Cyan("\t  Port Speed: "+color.YellowString("%d"), circuit.TerminationA.PortSpeed)
					} else {
						color.Cyan("\t  Port Speed: " + color.RedString("No termination A port speed found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
					if circuit.TerminationA.UpstreamSpeed != 0 {
						color.Cyan("\t  Upstream Speed: "+color.YellowString("%d"), circuit.TerminationA.UpstreamSpeed)
					} else {
						color.Cyan("\t  Upstream Speed: " + color.RedString("No termination A upstream speed found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
					if circuit.TerminationA.XconnectId != "" {
						color.Cyan("\t  XConnect ID: "+color.YellowString("%s"), circuit.TerminationA.XconnectId)
					} else {
						color.Cyan("\t  XConnect ID: " + color.RedString("No termination A xconnect ID found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
					if circuit.TerminationA.Description != "" {
						color.Cyan("\t  Description: "+color.YellowString("%s"), circuit.TerminationA.Description)
					} else {
						color.Cyan("\t  Description: " + color.RedString("No termination A description found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
				} else {
					color.Cyan("\tTermination A: " + color.RedString("No termination A found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.TerminationZ.Id != 0 {
					color.Cyan("\tTermination Z: ")
					color.Cyan("\t  ID: "+color.YellowString("%d"), circuit.TerminationZ.Id)
					color.Cyan("\t  URL: "+color.YellowString("%s"), circuit.TerminationZ.Url)
					color.Cyan("\t  Display: "+color.YellowString("%s"), circuit.TerminationZ.Display)
					if circuit.TerminationZ.Site.Id != 0 {
						color.Cyan("\t  Site: ")
						color.Cyan("\t    ID: "+color.YellowString("%d"), circuit.TerminationZ.Site.Id)
						color.Cyan("\t    URL: "+color.YellowString("%s"), circuit.TerminationZ.Site.Url)
						color.Cyan("\t    Display: "+color.YellowString("%s"), circuit.TerminationZ.Site.Display)
						color.Cyan("\t    Name: "+color.YellowString("%s"), circuit.TerminationZ.Site.Name)
						color.Cyan("\t    Slug: "+color.YellowString("%s"), circuit.TerminationZ.Site.Slug)
					} else {
						color.Cyan("\t  Site: " + color.RedString("No termination Z site found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
					if circuit.TerminationZ.ProviderNetwork.Id != 0 {
						color.Cyan("\t  Provider Network: ")
						color.Cyan("\t    ID: "+color.YellowString("%d"), circuit.TerminationZ.ProviderNetwork.Id)
						color.Cyan("\t    URL: "+color.YellowString("%s"), circuit.TerminationZ.ProviderNetwork.Url)
						color.Cyan("\t    Display: "+color.YellowString("%s"), circuit.TerminationZ.ProviderNetwork.Display)
						color.Cyan("\t    Name: "+color.YellowString("%s"), circuit.TerminationZ.ProviderNetwork.Name)
					} else {
						color.Cyan("\t  Provider Network: " + color.RedString("No termination Z provider network found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
					if circuit.TerminationZ.PortSpeed != 0 {
						color.Cyan("\t  Port Speed: "+color.YellowString("%d"), circuit.TerminationZ.PortSpeed)
					} else {
						color.Cyan("\t  Port Speed: " + color.RedString("No termination Z port speed found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
					if circuit.TerminationZ.UpstreamSpeed != 0 {
						color.Cyan("\t  Upstream Speed: "+color.YellowString("%d"), circuit.TerminationZ.UpstreamSpeed)
					} else {
						color.Cyan("\t  Upstream Speed: " + color.RedString("No termination Z upstream speed found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
					if circuit.TerminationZ.XconnectId != "" {
						color.Cyan("\t  XConnect ID: "+color.YellowString("%s"), circuit.TerminationZ.XconnectId)
					} else {
						color.Cyan("\t  XConnect ID: " + color.RedString("No termination Z xconnect ID found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
					if circuit.TerminationZ.Description != "" {
						color.Cyan("\t  Description: "+color.YellowString("%s"), circuit.TerminationZ.Description)
					} else {
						color.Cyan("\t  Description: " + color.RedString("No termination Z description found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
				} else {
					color.Cyan("\tTermination Z: " + color.RedString("No termination Z found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				if circuit.Comments != "" {
					color.Cyan("\tComments: "+color.YellowString("%s"), circuit.Comments)
				} else {
					color.Cyan("\tComments: " + color.RedString("No comments found for circuit: ") + color.YellowString("%s", circuit.Display))
				}
				for _, tag := range circuit.Tags {
					if tag.Id != 0 {
						color.Cyan("\tTags: ")
						color.Cyan("\t  ID: "+color.YellowString("%v"), tag.Id)
						color.Cyan("\t  URL: "+color.YellowString("%v"), tag.Url)
						color.Cyan("\t  Display: "+color.YellowString("%v"), tag.Display)
						color.Cyan("\t  Name: "+color.YellowString("%v"), tag.Name)
						color.Cyan("\t  Slug: "+color.YellowString("%v"), tag.Slug)
						color.Cyan("\t  Color: "+color.YellowString("%v"), tag.Color)
					} else {
						color.Cyan("\tTags: " + color.RedString("No tags found for circuit: ") + color.YellowString("%s", circuit.Display))
					}
				}
				color.Cyan("\tABC Circuit Provider Created: "+color.YellowString("%s"), circuit.Created)
				color.Cyan("\tABC Circuit Provider Last Updated: "+color.YellowString("%s"), circuit.LastUpdated)
			}
		} else {
			color.Cyan("\tCircuits: " + color.RedString("No circuits found on the server: "))
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetCircuitsCircuitsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "environment ('development' or 'production')")
	err := GetCircuitsCircuitsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("%s - for GetCircuitsCircuitsCmd", err)
	}
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCircuitsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	GetCircuitsCircuitsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

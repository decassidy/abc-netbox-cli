/*
Copyright © 2024 Derrick Cassidy - Metropolis Technologies, Inc.

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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

type terminations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Circuit struct {
			Id      int    `json:"id,omitempty"`
			Url     string `json:"url,omitempty"`
			Display string `json:"display,omitempty"`
			Cid     string `json:"cid,omitempty"`
		} `json:"circuit,omitempty"`
		TermSide string `json:"term_side,omitempty"`
		Site     struct {
			Id      int    `json:"id,omitempty"`
			Url     string `json:"url,omitempty"`
			Display string `json:"display,omitempty"`
			Name    string `json:"name,omitempty"`
			Slug    string `json:"slug,omitempty"`
		} `json:"site,omitempty"`
		ProviderNetwork struct {
			Id      int    `json:"id,omitempty"`
			Url     string `json:"url,omitempty"`
			Display string `json:"display,omitempty"`
			Name    string `json:"name,omitempty"`
		} `json:"provider_network,omitempty"`
		PortSpeed     int    `json:"port_speed,omitempty"`
		UpstreamSpeed int    `json:"upstream_speed,omitempty"`
		XconnectId    string `json:"xconnect_id,omitempty"`
		PpInfo        string `json:"pp_info,omitempty"`
		Description   string `json:"description,omitempty"`
		MarkConnected bool   `json:"mark_connected,omitempty"`
		Cable         struct {
			Id      int    `json:"id,omitempty"`
			Url     string `json:"url,omitempty"`
			Display string `json:"display,omitempty"`
			Label   string `json:"label,omitempty"`
		} `json:"cable,omitempty"`
		CableEnd      string   `json:"cable_end,omitempty"`
		LinkPeers     []string `json:"link_peers,omitempty"`
		LinkPeersType string   `json:"link_peers_type,omitempty"`
		Tags          []struct {
			Id      int    `json:"id,omitempty"`
			Url     string `json:"url,omitempty"`
			Display string `json:"display,omitempty"`
			Name    string `json:"name,omitempty"`
			Slug    string `json:"slug,omitempty"`
			Color   string `json:"color,omitempty"`
		} `json:"tags,omitempty"`
		Created     string `json:"created,omitempty"`
		LastUpdated string `json:"last_updated,omitempty"`
		Occupied    bool   `json:"_occupied,omitempty"`
	} `json:"results,omitempty"`
}

// GetCircuitsCircuitTerminationsCmd represents the getCircuitsCircuitTerminations command
var GetCircuitsCircuitTerminationsCmd = &cobra.Command{
	Use:   "getCircuitsCircuitTerminations",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		responseObject := new(terminations)
		ApiConnectionNonID(responseObject, "GET", "cmd.circuits.circuits_api_url.circuits_terminations")

		color.Cyan("\nTerminations Count: %d\n", responseObject.Count)
		for _, result := range responseObject.Results {
			color.Cyan("==========================================================\n")
			color.Cyan("\tDisplay: %s\n", color.YellowString(result.Display))
			color.Cyan("==========================================================\n")
			color.Cyan("\tID: %v\n", color.YellowString(strconv.Itoa(result.Id)))
			color.Cyan("\tURL: %s\n", color.YellowString(result.Url))
			color.Cyan("\tDisplay: %s\n", color.YellowString(result.Display))
			if result.Circuit.Id > 0 {
				color.Cyan("\tCircuit: \n")
				color.Cyan("\t  ID: %s\n", color.YellowString(strconv.Itoa(result.Circuit.Id)))
				color.Cyan("\t  URL: %s\n", color.YellowString(result.Circuit.Url))
				color.Cyan("\t  Display: %s\n", color.YellowString(result.Circuit.Display))
				color.Cyan("\t  CID: %s\n", color.YellowString(result.Circuit.Cid))
			} else {
				color.Cyan("\tCircuit: %s\n", color.RedString("No circuit found"))
			}
			if result.Site.Id > 0 {
				color.Cyan("\tSite: \n")
				color.Cyan("\t  ID: %v\n", color.YellowString(strconv.Itoa(result.Site.Id)))
				color.Cyan("\t  URL: %s\n", color.YellowString(result.Site.Url))
				color.Cyan("\t  Display: %s\n", color.YellowString(result.Site.Display))
				color.Cyan("\t  Name: %s\n", color.YellowString(result.Site.Name))
				color.Cyan("\t  Slug: %s\n", color.YellowString(result.Site.Slug))
			} else {
				color.Cyan("\tSite: %s\n", color.RedString("No site found"))
			}
			if result.ProviderNetwork.Id > 0 {
				color.Cyan("\tProvider Network: \n")
				color.Cyan("\t  ID: %v\n", color.YellowString(strconv.Itoa(result.ProviderNetwork.Id)))
				color.Cyan("\t  URL: %s\n", color.YellowString(result.ProviderNetwork.Url))
				color.Cyan("\t  Display: %s\n", color.YellowString(result.ProviderNetwork.Display))
				color.Cyan("\t  Name: %s\n", color.YellowString(result.ProviderNetwork.Name))
			} else {
				color.Cyan("\tProvider Network: %s\n", color.RedString("%s", "No provider network"))
			}
			if result.PortSpeed > 0 {
				color.Cyan("\tPort Speed: %v\n", color.YellowString(strconv.Itoa(result.PortSpeed)))
			} else {
				color.Cyan("\tPort Speed: %s\n", color.RedString("No port speed found"))
			}
			if result.UpstreamSpeed > 0 {
				color.Cyan("\tUpstream Speed: %v\n", color.YellowString(strconv.Itoa(result.UpstreamSpeed)))
			} else {
				color.Cyan("\tUpstream Speed: %s\n", color.RedString("No upstream speed found"))
			}
			if result.XconnectId != "" {
				color.Cyan("\tXconnect ID: %s\n", color.YellowString(result.XconnectId))
			} else {
				color.Cyan("\tXconnect ID: %s\n", color.RedString("No xconnect ID found"))
			}
			if result.PpInfo != "" {
				color.Cyan("\tPP Info: %s\n", color.YellowString(result.PpInfo))
			} else {
				color.Cyan("\tPP Info: %s\n", color.RedString("No PP info found"))
			}
			if result.Description != "" {
				color.Cyan("\tDescription: %s\n", color.YellowString(result.Description))
			} else {
				color.Cyan("\tDescription: %s\n", color.RedString("No description found"))
			}
			color.Cyan("\tMark Connected: %t\n", result.MarkConnected)
			if result.Cable.Id > 0 {
				color.Cyan("\tCable: \n")
				color.Cyan("\t  ID: %v\n", color.YellowString(strconv.Itoa(result.Cable.Id)))
				color.Cyan("\t  URL: %s\n", color.YellowString(result.Cable.Url))
				color.Cyan("\t  Display: %s\n", color.YellowString(result.Cable.Display))
				color.Cyan("\t  Label: %s\n", color.YellowString(result.Cable.Label))
			} else {
				color.Cyan("\tCable: %s\n", color.RedString("No cable"))
			}
			if result.CableEnd != "" {
				color.Cyan("\tCableEnd: %s\n", color.YellowString(result.CableEnd))
			} else {
				color.Cyan("\tCableEnd: %s\n", color.RedString("No cable end found"))
			}
			color.Cyan("\tLink Peers:\n")
			for _, link := range result.LinkPeers {
				if link != "" {
					color.Cyan("\t  Link Peer: %s\n", color.YellowString(link))
				} else {
					color.Cyan("\t  Link Peer: %s\n", color.RedString("No link peer found"))
				}
			}
			if result.LinkPeersType != "" {
				color.Cyan("\tLink Peers Type: %s\n", color.YellowString(result.LinkPeersType))
			} else {
				color.Cyan("\tLink Peers Type: %s\n", color.RedString("No link peers type found"))
			}
			color.Cyan("\tTags:\n")
			for _, tag := range result.Tags {
				if tag.Id != 0 {
					color.Cyan("\t  Tag ID: %v\n", color.YellowString(strconv.Itoa(tag.Id)))
					color.Cyan("\t  Tag URL: %s\n", color.YellowString(tag.Url))
					color.Cyan("\t  Tag Display: %s\n", color.YellowString(tag.Display))
					color.Cyan("\t  Tag Name: %s\n", color.YellowString(tag.Name))
					color.Cyan("\t  Tag Slug: %s\n", color.YellowString(tag.Slug))
					color.Cyan("\t  Tag Color: %s\n", color.YellowString(tag.Color))
				} else {
					color.Cyan("\tTags: %s\n", color.RedString("No tags found"))
				}
			}
			color.Cyan("\tCreated: %s\n", color.YellowString(result.Created))
			color.Cyan("\tLast Updated: %s\n", color.YellowString(result.LastUpdated))
			color.Cyan("\tOccupied: %t\n\n", result.Occupied)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	GetCircuitsCircuitTerminationsCmd.Flags().StringVarP(&serverEnv, "env", "", "development", "Environment from which the server is running")
	err := GetCircuitsCircuitTerminationsCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatalf("Error marking flag as required: %s - for GetCircuitsCircuitTerminationsCmd", err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCircuitsCircuitTerminationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCircuitsCircuitTerminationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

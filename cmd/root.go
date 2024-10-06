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

package cmd

import (
	"fmt"
	"github.com/decassidy/abc-netbox-cli/cmd/circuits"
	"github.com/decassidy/abc-netbox-cli/cmd/core"
	"github.com/decassidy/abc-netbox-cli/cmd/dcim"
	"github.com/decassidy/abc-netbox-cli/cmd/extras"
	"github.com/decassidy/abc-netbox-cli/cmd/ipam"
	"github.com/decassidy/abc-netbox-cli/cmd/tenancy"
	"github.com/decassidy/abc-netbox-cli/cmd/users"
	"github.com/decassidy/abc-netbox-cli/cmd/virtualization"
	"github.com/decassidy/abc-netbox-cli/cmd/vpn"
	"github.com/decassidy/abc-netbox-cli/cmd/wireless"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "abc-netbox.cli",
	Short: "\nA command line tool to interact with ABC Netbox APIs",
	Long: `
ABC Netbox Automation Tools - ABC Netbox APIs.
Copyright (c) 2024 ABC Technologies, Inc. All Rights Reserved.`,
	Version: "v0.1.0",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.abc-netbox.cli")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version information for abc-netbox.cli",
	Long:  `ABC Netbox Automation Tools - ABC Netbox APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\nVersion: %s\nBuild Time: %s\n", rootCmd.Long, rootCmd.Version, "2024-05-20 21:53:30")
	},
}

var CompletionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	Long:                  "To load completions",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			err := cmd.Root().GenBashCompletion(os.Stdout)
			if err != nil {
				return
			}
		case "zsh":
			err := cmd.Root().GenZshCompletion(os.Stdout)
			if err != nil {
				return
			}
		case "fish":
			err := cmd.Root().GenFishCompletion(os.Stdout, true)
			if err != nil {
				return
			}
		case "powershell":
			err := cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			if err != nil {
				return
			}
		}
	},
}

var CircuitsCmd = &cobra.Command{
	Use:   "Circuits",
	Short: "ABC Netbox Circuit Management APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Circuit Management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var circuitsGetCmd = &cobra.Command{
	Use:   "CircuitsGet",
	Short: "ABC Netbox Circuit Management GET APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Circuit Management GET APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var circuitsPostCmd = &cobra.Command{
	Use:   "CircuitsPost",
	Short: "ABC Netbox Circuit Management POST APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Circuit Management POST APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var circuitsPatchCmd = &cobra.Command{
	Use:   "CircuitsPatch",
	Short: "ABC Netbox Circuit Management PATCH APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Circuit Management PATCH APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var circuitsDeleteCmd = &cobra.Command{
	Use:   "CircuitsDelete",
	Short: "ABC Netbox Circuit Management DELETE APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Circuit Management DELETE APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var CoreCmd = &cobra.Command{
	Use:   "Core",
	Short: "ABC Netbox Core Management APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Core Management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var coreGetCmd = &cobra.Command{
	Use:   "CoreGet",
	Short: "ABC Netbox Core Management GET APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Core Management GET APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var corePostCmd = &cobra.Command{
	Use:   "CorePost",
	Short: "ABC Netbox Core Management POST APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Core Management POST APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var corePatchCmd = &cobra.Command{
	Use:   "CorePatch",
	Short: "ABC Netbox Core Management PATCH APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Core Management PATCH APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var coreDeleteCmd = &cobra.Command{
	Use:   "CoreDelete",
	Short: "ABC Netbox Core Management Delete APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Core Management DELETE APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var DcimCmd = &cobra.Command{
	Use:   "DCIM",
	Short: "ABC Netbox DCIM Management APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox DCIM Management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var dcimGetCmd = &cobra.Command{
	Use:   "DcimGet",
	Short: "ABC Netbox DCIM Management GET APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox DCIM Management GET APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var dcimPostCmd = &cobra.Command{
	Use:   "DcimPost",
	Short: "ABC Netbox DCIM Management POST APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox DCIM Management POST APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var dcimPatchCmd = &cobra.Command{
	Use:   "DcimPatch",
	Short: "ABC Netbox DCIM Management PATCH APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox DCIM Management PATCH APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var dcimDeleteCmd = &cobra.Command{
	Use:   "DcimDelete",
	Short: "ABC Netbox DCIM Management DELETE APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox DCIM Management DELETE APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var ExtrasCmd = &cobra.Command{
	Use:   "Extras",
	Short: "ABC Netbox Extras Management APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Extras Management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var extrasGetCmd = &cobra.Command{
	Use:   "ExtrasGet",
	Short: "ABC Netbox Extras Management GET APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Extras Management GET APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var extrasPostCmd = &cobra.Command{
	Use:   "ExtrasPost",
	Short: "ABC Netbox Extras Management POST APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Extras Management POST APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var extrasPatchCmd = &cobra.Command{
	Use:   "ExtrasPatch",
	Short: "ABC Netbox Extras Management PATCH APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Extras Management PATCH APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var extrasDeleteCmd = &cobra.Command{
	Use:   "ExtrasDelete",
	Short: "ABC Netbox Extras Management DELETE APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Extras Management DELETE APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var IpamCmd = &cobra.Command{
	Use:   "IPAM",
	Short: "ABC Netbox IPAM Management APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox IPAM Management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var ipamGetCmd = &cobra.Command{
	Use:   "IpamGet",
	Short: "ABC Netbox IPAM Management GET APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox IPAM Management GET APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var ipamPostCmd = &cobra.Command{
	Use:   "IpamPost",
	Short: "ABC Netbox IPAM Management POST APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox IPAM Management POST APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var ipamPatchCmd = &cobra.Command{
	Use:   "IpamPatch",
	Short: "ABC Netbox IPAM Management PATCH APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox IPAM Management PATCH APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var ipamDeleteCmd = &cobra.Command{
	Use:   "IpamDelete",
	Short: "ABC Netbox IPAM Management DELETE APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox IPAM Management DELETE APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var TenancyCmd = &cobra.Command{
	Use:   "Tenancy",
	Short: "ABC Netbox Tenancy Management APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Tenancy Management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var tenancyGetCmd = &cobra.Command{
	Use:   "TenancyGet",
	Short: "ABC Netbox Tenancy Management GET APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Tenancy Management GET APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var tenancyPostCmd = &cobra.Command{
	Use:   "TenancyPost",
	Short: "ABC Netbox Tenancy Management POST APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Tenancy Management POST APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var tenancyPatchCmd = &cobra.Command{
	Use:   "TenancyPatch",
	Short: "ABC Netbox Tenancy Management PATCH APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Tenancy Management PATCH APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var tenancyDeleteCmd = &cobra.Command{
	Use:   "TenancyDelete",
	Short: "ABC Netbox Tenancy Management DELETE APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Tenancy Management DELETE APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var UsersCmd = &cobra.Command{
	Use:   "Users",
	Short: "ABC Netbox Users Management APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Users Management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var usersGetCmd = &cobra.Command{
	Use:   "UsersGet",
	Short: "ABC Netbox Users Management GET APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Users Management GET APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var usersPostCmd = &cobra.Command{
	Use:   "UsersPost",
	Short: "ABC Netbox Users Management POST APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Users Management POST APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var usersPatchCmd = &cobra.Command{
	Use:   "UsersPatch",
	Short: "ABC Netbox Users Management PATCH APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Users Management PATCH APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var usersDeleteCmd = &cobra.Command{
	Use:   "UsersDelete",
	Short: "ABC Netbox Users Management DELETE APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Users Management DELETE APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var VirtualizationCmd = &cobra.Command{
	Use:   "Virtualization",
	Short: "ABC Netbox Virtualization Management APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Virtualization Management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var virtualizationGetCmd = &cobra.Command{
	Use:   "VirtualizationGet",
	Short: "ABC Netbox Virtualization Management GET APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Virtualization Management GET APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var virtualizationPostCmd = &cobra.Command{
	Use:   "VirtualizationPost",
	Short: "ABC Netbox Virtualization Management POST APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Virtualization Management POST APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var virtualizationPatchCmd = &cobra.Command{
	Use:   "VirtualizationPatch",
	Short: "ABC Netbox Virtualization Management PATCH APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Virtualization Management PATCH APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var virtualizationDeleteCmd = &cobra.Command{
	Use:   "VirtualizationDelete",
	Short: "ABC Netbox Virtualization Management DELETE APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Virtualization Management DELETE APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var VpnCmd = &cobra.Command{
	Use:   "VPN",
	Short: "ABC Netbox VPN Management APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox VPN Management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var vpnGetCmd = &cobra.Command{
	Use:   "VpnGet",
	Short: "ABC Netbox VPN Management GET APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox VPN Management GET APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var vpnPostCmd = &cobra.Command{
	Use:   "VpnPost",
	Short: "ABC Netbox VPN Management POST APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox VPN Management POST APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var vpnPatchCmd = &cobra.Command{
	Use:   "VpnPatch",
	Short: "ABC Netbox VPN Management PATCH APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox VPN Management PATCH APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var vpnDeleteCmd = &cobra.Command{
	Use:   "VpnDelete",
	Short: "ABC Netbox VPN Management DELETE APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox VPN Management DELETE APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var WirelessCmd = &cobra.Command{
	Use:   "Wireless",
	Short: "ABC Netbox Wireless Management APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Wireless Management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var wirelessGetCmd = &cobra.Command{
	Use:   "WirelessGet",
	Short: "ABC Netbox Wireless Management GET APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Wireless Management GET APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var wirelessPostCmd = &cobra.Command{
	Use:   "WirelessPost",
	Short: "ABC Netbox Wireless Management POST APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Wireless Management POST APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var wirelessPatchCmd = &cobra.Command{
	Use:   "WirelessPatch",
	Short: "ABC Netbox Wireless Management PATCH APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Wireless Management PATCH APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var wirelessDeleteCmd = &cobra.Command{
	Use:   "WirelessDelete",
	Short: "ABC Netbox Wireless Management DELETE APIs.",
	Long: `
ABC Netbox Automation Tools:
  ABC Netbox Wireless Management DELETE APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addCircuitsSubcommandPalettes() {
	rootCmd.AddCommand(CircuitsCmd)
	CircuitsCmd.AddCommand(circuitsGetCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsCircuitTerminationsCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsCircuitTerminationsByIdCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsCircuitTypesCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsCircuitTypesByIdCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsCircuitsCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsCircuitsByIDCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsProviderAccountsCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsProviderAccountsByIDCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsProviderNetworksCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsProviderNetworksByIDCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsProvidersCmd)
	circuitsGetCmd.AddCommand(circuits.GetCircuitsProvidersByIdCmd)
	CircuitsCmd.AddCommand(circuitsPostCmd)
	circuitsPostCmd.AddCommand(circuits.PostCircuitsCircuitTypeCmd)
	circuitsPostCmd.AddCommand(circuits.PostCircuitsCircuitTerminationsCmd)
	circuitsPostCmd.AddCommand(circuits.PostCircuitsCircuitTypeCmd)
	circuitsPostCmd.AddCommand(circuits.PostCircuitsCircuitsCmd)
	circuitsPostCmd.AddCommand(circuits.PostCircuitsProviderAccountsCmd)
	circuitsPostCmd.AddCommand(circuits.PostCircuitsProviderNetworksCmd)
	circuitsPostCmd.AddCommand(circuits.PostCircuitsProvidersCmd)
	CircuitsCmd.AddCommand(circuitsPatchCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsCircuitTerminationsCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsCircuitTerminationsByIdCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsCircuitsCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsCircuitsByIdCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsCircuitTypesCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsCircuitTypesByIdCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsProviderAccountsCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsProviderAccountsByIdCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsProviderNetworksCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsProviderNetworksByIdCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsProvidersCmd)
	circuitsPatchCmd.AddCommand(circuits.PatchCircuitsProvidersByIdCmd)
	CircuitsCmd.AddCommand(circuitsDeleteCmd)
	circuitsDeleteCmd.AddCommand(circuits.DeleteCircuitsCircuitTerminationsCmd)
	circuitsDeleteCmd.AddCommand(circuits.DeleteCircuitsCircuitTerminationsByIdCmd)
	circuitsDeleteCmd.AddCommand(circuits.DeleteCircuitsCircuitTypesCmd)
	circuitsDeleteCmd.AddCommand(circuits.DeleteCircuitsCircuitTypesByIdCmd)
	circuitsDeleteCmd.AddCommand(circuits.DeleteCircuitsCircuitsCmd)
	circuitsDeleteCmd.AddCommand(circuits.DeleteCircuitsCircuitsByIdCmd)
	circuitsDeleteCmd.AddCommand(circuits.DeleteCircuitsProviderNetworksCmd)
	circuitsDeleteCmd.AddCommand(circuits.DeleteCircuitsProviderNetworksByIdCmd)
	circuitsDeleteCmd.AddCommand(circuits.DeleteCircuitsProvidersCmd)
	circuitsDeleteCmd.AddCommand(circuits.DeleteCircuitsProvidersByIdCmd)
}

func addCoreSubcommandPalettes() {
	rootCmd.AddCommand(CoreCmd)
	CoreCmd.AddCommand(coreGetCmd)
	coreGetCmd.AddCommand(core.GetCoreDataFilesCmd)
	coreGetCmd.AddCommand(core.GetCoreDataFileByIDCmd)
	coreGetCmd.AddCommand(core.GetCoreDataSourcesCmd)
	coreGetCmd.AddCommand(core.GetCoreDataFileByIDCmd)
	CoreCmd.AddCommand(corePostCmd)
	corePostCmd.AddCommand(core.PostCoreDataSourcesCmd)
	CoreCmd.AddCommand(corePatchCmd)
	corePatchCmd.AddCommand(core.PatchCoreDataSourcesCmd)
	corePatchCmd.AddCommand(core.PatchCoreDataSourcesByIDCmd)
	CoreCmd.AddCommand(coreDeleteCmd)
	coreDeleteCmd.AddCommand(core.DeleteCoreDataSourcesCmd)
	coreDeleteCmd.AddCommand(core.DeleteCoreDataSourcesByIDCmd)
}

func addDcimSubcommandPalettes() {
	rootCmd.AddCommand(DcimCmd)
	DcimCmd.AddCommand(dcimGetCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimCablesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimCablesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimCableTerminationsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimCableTerminationsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimConnectedDeviceCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimConsolePortTemplatesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimConsolePortTemplatesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimConsolePortsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimConsolePortsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimConsoleServerPortTemplatesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimConsoleServerPortTemplatesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimConsoleServerPortsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimConsoleServerPortsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDeviceBayTemplatesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDeviceBayTemplatesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDeviceBaysCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDeviceBaysByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDeviceRolesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDeviceRolesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDeviceTypesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDeviceTypesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDevicesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDevicesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDeviceIdBySerialNumberCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimDeviceByQueryCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimFrontPortTemplatesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimFrontPortTemplatesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimFrontPortsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimFrontPortsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimFrontPortsByQueryCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInterfaceTemplatesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInterfaceTemplatesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInterfacesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInterfacesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInterfacesByQueryCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInventoryItemRolesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInventoryItemRolesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInventoryItemTemplatesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInventoryItemTemplatesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInventoryItemsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimInventoryItemsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimLocationsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimLocationsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimManufacturersCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimManufacturersByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimModuleBayTemplatesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimModuleBayTemplatesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimModuleTypesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimModuleTypesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimModulesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimModulesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPlatformsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPlatformsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerFeedsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerFeedsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerOutletTemplatesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerOutletTemplatesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerOutletsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerOutletsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerPanelsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerPanelsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerPortTemplatesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerPortTemplatesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerPortsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimPowerPortsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRackReservationsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRackReservationsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRackRolesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRackRolesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRacksCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRacksByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRearPortTemplatesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRearPortTemplatesByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRearPortsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRearPortsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRegionsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimRegionsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimSiteGroupsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimSiteGroupsByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimSitesCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimSitesByIDCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimSitesByQueryCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimVirtualChassisCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimVirtualChassisByIdCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimVirtualDeviceContextsCmd)
	dcimGetCmd.AddCommand(dcim.GetDcimVirtualDeviceContextsByIdCmd)
	DcimCmd.AddCommand(dcimPatchCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimCablesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimCablesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimCableTerminationsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimCableTerminationsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimConsolePortTemplatesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimConsolePortTemplatesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimConsolePortsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimConsolePortsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimConsoleServerPortTemplatesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimConsoleServerPortTemplatesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimConsoleServerPortsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimConsoleServerPortsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimDeviceBaysCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimDeviceBaysByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimDeviceBayTemplatesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimDeviceBayTemplatesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimDeviceRolesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimDeviceRolesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimDeviceTypesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimDeviceTypesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimDevicesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimDevicesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimFrontPortTemplatesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimFrontPortTemplatesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimFrontPortsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimFrontPortsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimInterfaceTemplatesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimInterfaceTemplatesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimInterfacesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimInterfacesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimInventoryItemRolesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimInventoryItemRolesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimInventoryItemTemplatesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimInventoryItemTemplatesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimInventoryItemsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimInventoryItemsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimLocationsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimLocationsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimManufacturersCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimManufacturersByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimModuleBayTemplatesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimModuleBayTemplatesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimModuleTypesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimModuleTypesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimModulesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimModulesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPlatformsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPlatformsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerFeedsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerFeedsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerOutletTemplatesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerOutletTemplatesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerOutletsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerOutletsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerPanelsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerPanelsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerPortTemplatesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerPortTemplatesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerPortsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimPowerPortsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRackReservationsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRackReservationsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRackRolesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRackRolesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRacksCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRacksByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRearPortTemplatesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRearPortTemplatesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRearPortsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRearPortsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRegionsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimRegionsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimSiteGroupsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimSiteGroupsByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimSitesCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimSitesByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimVirtualChassisCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimVirtualChassisByIdCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimVirtualDeviceContextsCmd)
	dcimPatchCmd.AddCommand(dcim.PatchDcimVirtualDeviceContextsByIdCmd)
	DcimCmd.AddCommand(dcimPostCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimDeviceBayTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimDeviceTypesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimCablesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimCableTerminationsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimConsolePortTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimConsolePortsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimConsoleServerPortTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimConsoleServerPortsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimDeviceBayTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimDeviceBaysCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimDeviceRolesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimDeviceTypesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimDevicesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimFrontPortTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimFrontPortsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimInterfaceTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimInterfacesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimInventoryItemRolesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimInventoryItemTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimInventoryItemsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimLocationsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimManufacturersCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimModuleBayTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimModuleTypesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimModulesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimPlatformsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimPowerFeedsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimPowerOutletTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimPowerOutletsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimPowerPanelsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimPowerPortTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimPowerPortsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimRackReservationsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimRackRolesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimRacksCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimRearPortTemplatesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimRearPortsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimRegionsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimSiteGroupsCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimSitesCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimVirtualChassisCmd)
	dcimPostCmd.AddCommand(dcim.PostDcimVirtualDeviceContextsCmd)
	DcimCmd.AddCommand(dcimDeleteCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimCablesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimCablesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimCableTerminationCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimCableTerminationsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimConsolePortTemplatesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimConsolePortTemplatesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimConsolePortsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimConsolePortsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimConsoleServerPortTemplatesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimConsoleServerPortTemplatesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimConsoleServerPortsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimConsoleServerPortsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimConsolePortsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimDeviceBayTemplatesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimDeviceBayTemplatesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimDeviceBaysCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimDeviceBaysByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimDeviceRolesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimDeviceRolesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimDeviceTypesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimDeviceTypesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimDevicesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimDevicesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimFrontPortTemplatesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimFrontPortTemplatesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimFrontPortsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimInterfaceTemplatesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimInterfaceTemplatesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimInterfacesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimInterfacesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimInventoryItemRolesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimInventoryItemRolesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimInventoryItemTemplatesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimInventoryItemTemplatesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimInventoryItemsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimInventoryItemsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimLocationsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimLocationsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimManufacturersCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimManufacturersByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimModuleBayTemplatesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimModuleBayTemplatesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimModuleTypesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimModuleTypesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimModulesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimModulesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPlatformsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPlatformsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerFeedsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerFeedsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerOutletTemplatesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerOutletTemplatesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerOutletsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerOutletsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerPanelsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerPanelsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerPortTemplatesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerPortTemplatesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerPortsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimPowerPortsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRackReservationsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRackReservationsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRackRolesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRackRolesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRacksCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRacksByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRearPortTemplatesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRearPortTemplatesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRearPortsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRearPortsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRegionsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimRegionsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimSiteGroupsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimSiteGroupsByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimSitesCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimSitesByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimVirtualChassisCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimVirtualChassisByIdCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimVirtualDeviceContextsCmd)
	dcimDeleteCmd.AddCommand(dcim.DeleteDcimVirtualDeviceContextsByIdCmd)
}

func addExtrasSubcommandsPalettes() {
	rootCmd.AddCommand(ExtrasCmd)
	ExtrasCmd.AddCommand(extrasGetCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasBookmarksCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasBookmarksByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasConfigContextsCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasConfigContextsByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasConfigTemplatesCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasConfigTemplatesByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasContentTypesCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasContentTypesByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasCustomFieldChoiceSetsCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasCustomFieldChoiceSetsByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasCustomFieldsCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasCustomFieldsByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasCustomLinksCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasCustomLinksByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasDashboardCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasEventRulesCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasEventRulesByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasExportTemplatesCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasExportTemplatesByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasImageAttachmentsCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasImageAttachmentsByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasJournalEntriesCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasJournalEntriesByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasObjectChangesCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasObjectChangesByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasSavedFiltersCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasSavedFiltersByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasTagsCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasTagsByIdCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasWebhooksCmd)
	extrasGetCmd.AddCommand(extras.GetExtrasWebhooksByIdCmd)
	ExtrasCmd.AddCommand(extrasPostCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasBookmarksCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasConfigContextsCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasConfigTemplatesCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasCustomFieldChoiceSetsCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasCustomFieldsCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasCustomLinksCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasEventRulesCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasExportTemplatesCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasImageAttachmentsCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasJournalEntriesCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasSavedFiltersCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasTagsCmd)
	extrasPostCmd.AddCommand(extras.PostExtrasWebhooksCmd)
	ExtrasCmd.AddCommand(extrasPatchCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasBookmarksCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasBookmarksByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasConfigContextsCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasConfigContextsByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasConfigTemplatesCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasConfigTemplatesByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasCustomFieldChoiceSetsCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasCustomFieldChoiceSetsByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasCustomFieldsCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasCustomFieldsByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasCustomLinksCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasCustomLinksByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasDashboardCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasEventRulesCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasEventRulesByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasExportTemplatesCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasExportTemplatesByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasImageAttachmentsCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasImageAttachmentsByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasJournalEntriesCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasJournalEntriesByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasSavedFiltersCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasSavedFiltersByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasTagsCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasTagsByIdCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasWebhooksCmd)
	extrasPatchCmd.AddCommand(extras.PatchExtrasWebhooksByIdCmd)
	ExtrasCmd.AddCommand(extrasDeleteCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasBookmarksCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasBookmarksByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasConfigContextsCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasConfigContextsByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasConfigTemplatesCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasConfigTemplatesByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasCustomFieldChoiceSetsCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasCustomFieldChoiceSetsByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasCustomFieldsCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasCustomFieldsByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasCustomLinksCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasCustomLinksByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasDashboardCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasEventRulesCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasEventRulesByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasExportTemplatesCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasExportTemplatesByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasImageAttachmentsCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasImageAttachmentsByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasJournalEntriesCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasJournalEntriesByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasSavedFiltersCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasSavedFiltersByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasTagsCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasTagsByIdCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasWebhooksCmd)
	extrasDeleteCmd.AddCommand(extras.DeleteExtrasWebhooksByIdCmd)
}

func addIpamSubcommandPalettes() {
	rootCmd.AddCommand(IpamCmd)
	IpamCmd.AddCommand(ipamGetCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamAggregatesCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamAggregatesByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamAsnRangesCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamAsnRangesByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamAsnsCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamAsnsByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamFhrpGroupAssignmentsCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamFhrpGroupAssignmentsByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamFhrpGroupsCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamFhrpGroupsByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamIpAddressesCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamIpAddressesByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamIpAddressesByQueryCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamIpRangesCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamIpRangesByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamPrefixesCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamPrefixesByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamPrefixesAvailableIPsCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamPrefixesAvailablePrefixesCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamRirsCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamRirsByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamRolesCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamRolesByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamRouteTargetsCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamRouteTargetsByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamServicesCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamServicesByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamServiceTemplatesCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamServiceTemplatesByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamVlanGroupsCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamVlanGroupsByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamVlanGroupsAvailableVlansCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamVlansCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamVlansByIdCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamVrfsCmd)
	ipamGetCmd.AddCommand(ipam.GetIpamVrfsByIdCmd)
	IpamCmd.AddCommand(ipamPostCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamAggregatesCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamAsnRangesCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamAsnsCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamFhrpGroupAssignmentsCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamFhrpGroupsCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamIpAddressesCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamIpRangesCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamPrefixesCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamPrefixesAvailableIPsCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamPrefixesAvailablePrefixesCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamRirsCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamRolesCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamRouteTargetsCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamServicesCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamServiceTemplatesCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamVlanGroupsCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamVlanGroupsAvailableVlansCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamVlansCmd)
	ipamPostCmd.AddCommand(ipam.PostIpamVrfsCmd)
	IpamCmd.AddCommand(ipamPatchCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamAggregatesCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamAggregatesByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamAsnRangesCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamAsnRangesByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamAsnsCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamAsnsByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamFhrpGroupAssignmentsCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamFhrpGroupAssignmentsByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamFhrpGroupsCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamFhrpGroupsByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamIpAddressesCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamIpAddressesByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamIpRangesCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamIpRangesByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamPrefixesCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamPrefixesByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamRirsCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamRirsByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamRolesCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamRolesByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamRouteTargetsCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamRouteTargetsByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamServicesCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamServicesByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamServiceTemplatesCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamServiceTemplatesByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamVlanGroupsCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamVlanGroupsByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamVlansCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamVlansByIdCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamVrfsCmd)
	ipamPatchCmd.AddCommand(ipam.PatchIpamVrfsByIdCmd)
	IpamCmd.AddCommand(ipamDeleteCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamAggregatesCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamAggregatesByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamAsnRangesCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamAsnRangesByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamAsnsCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamAsnsByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamFhrpGroupAssignmentsCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamFhrpGroupAssignmentsByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamFhrpGroupsCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamFhrpGroupsByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamIpAddressesCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamIpAddressesByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamIpRangesCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamIpRangesByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamPrefixesCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamPrefixesByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamRirsCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamRirsByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamRolesCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamRolesByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamRouteTargetsCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamRouteTargetsByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamServicesCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamServicesByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamServiceTemplatesCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamServiceTemplatesByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamVlanGroupsCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamVlanGroupsByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamVlansCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamVlansByIdCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamVrfsCmd)
	ipamDeleteCmd.AddCommand(ipam.DeleteIpamVrfsByIdCmd)
}

func addTenancySubcommandPalettes() {
	rootCmd.AddCommand(TenancyCmd)
	TenancyCmd.AddCommand(tenancyGetCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyContactAssignmentsCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyContactAssignmentsByIdCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyContactGroupsCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyContactGroupsByIdCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyContactRolesCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyContactRolesByIdCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyContactsCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyContactsByIdCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyTenantGroupsCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyTenantGroupsByIdCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyTenantsCmd)
	tenancyGetCmd.AddCommand(tenancy.GetTenancyTenantsByIdCmd)
	TenancyCmd.AddCommand(tenancyPostCmd)
	tenancyPostCmd.AddCommand(tenancy.PostTenancyContactAssignmentsCmd)
	tenancyPostCmd.AddCommand(tenancy.PostTenancyContactGroupsCmd)
	tenancyPostCmd.AddCommand(tenancy.PostTenancyContactRolesCmd)
	tenancyPostCmd.AddCommand(tenancy.PostTenancyContactsCmd)
	tenancyPostCmd.AddCommand(tenancy.PostTenancyTenantGroupsCmd)
	tenancyPostCmd.AddCommand(tenancy.PostTenancyTenantsCmd)
	TenancyCmd.AddCommand(tenancyPatchCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyContactAssignmentsCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyContactAssignmentsByIdCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyContactGroupsCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyContactGroupsByIdCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyContactRolesCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyContactRolesByIdCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyContactsCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyContactsByIdCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyTenantGroupsCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyTenantGroupsByIdCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyTenantsCmd)
	tenancyPatchCmd.AddCommand(tenancy.PatchTenancyTenantsByIdCmd)
	TenancyCmd.AddCommand(tenancyDeleteCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyContactAssignmentsCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyContactAssignmentsByIdCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyContactGroupsCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyContactGroupsByIdCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyContactRolesCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyContactRolesByIdCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyContactsCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyContactsByIdCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyTenantGroupsCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyTenantGroupsByIdCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyTenantsCmd)
	tenancyDeleteCmd.AddCommand(tenancy.DeleteTenancyTenantsByIdCmd)
}

func addUsersSubcommandPalettes() {
	rootCmd.AddCommand(UsersCmd)
	UsersCmd.AddCommand(usersGetCmd)
	usersGetCmd.AddCommand(users.GetUsersGroupsCmd)
	usersGetCmd.AddCommand(users.GetUsersGroupsByIdCmd)
	usersGetCmd.AddCommand(users.GetUsersPermissionsCmd)
	usersGetCmd.AddCommand(users.GetUsersPermissionsByIdCmd)
	usersGetCmd.AddCommand(users.GetUsersTokensCmd)
	usersGetCmd.AddCommand(users.GetUsersTokensByIdCmd)
	usersGetCmd.AddCommand(users.GetUsersUsersCmd)
	usersGetCmd.AddCommand(users.GetUsersUsersByIdCmd)
	UsersCmd.AddCommand(usersPostCmd)
	usersPostCmd.AddCommand(users.PostUsersGroupsCmd)
	usersPostCmd.AddCommand(users.PostUsersPermissionsCmd)
	usersPostCmd.AddCommand(users.PostUsersTokensCmd)
	usersPostCmd.AddCommand(users.PostUsersUsersCmd)
	UsersCmd.AddCommand(usersPatchCmd)
	usersPatchCmd.AddCommand(users.PatchUsersGroupsCmd)
	usersPatchCmd.AddCommand(users.PatchUsersGroupsByIdCmd)
	usersPatchCmd.AddCommand(users.PatchUsersPermissionsCmd)
	usersPatchCmd.AddCommand(users.PatchUsersPermissionsByIdCmd)
	usersPatchCmd.AddCommand(users.PatchUsersTokensCmd)
	usersPatchCmd.AddCommand(users.PatchUsersTokensByIdCmd)
	usersPatchCmd.AddCommand(users.PatchUsersUsersCmd)
	usersPatchCmd.AddCommand(users.PatchUsersUsersByIdCmd)
	UsersCmd.AddCommand(usersDeleteCmd)
	usersDeleteCmd.AddCommand(users.DeleteUsersGroupsCmd)
	usersDeleteCmd.AddCommand(users.DeleteUsersGroupsByIdCmd)
	usersDeleteCmd.AddCommand(users.DeleteUsersPermissionsCmd)
	usersDeleteCmd.AddCommand(users.DeleteUsersPermissionsByIdCmd)
	usersDeleteCmd.AddCommand(users.DeleteUsersTokensCmd)
	usersDeleteCmd.AddCommand(users.DeleteUsersTokensByIdCmd)
	usersDeleteCmd.AddCommand(users.DeleteUsersUsersCmd)
	usersDeleteCmd.AddCommand(users.DeleteUsersUsersByIdCmd)
}

func addVirtualizationSubcommandPalettes() {
	rootCmd.AddCommand(VirtualizationCmd)
	VirtualizationCmd.AddCommand(virtualizationGetCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationClusterGroupsCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationClusterGroupsByIdCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationClustersCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationClustersByIdCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationClusterTypesCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationClusterTypesByIdCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationInterfacesCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationInterfacesByIdCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationVirtualDisksCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationVirtualDisksByIdCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationVirtualMachinesCmd)
	virtualizationGetCmd.AddCommand(virtualization.GetVirtualizationVirtualMachinesByIdCmd)
	VirtualizationCmd.AddCommand(virtualizationPostCmd)
	virtualizationPostCmd.AddCommand(virtualization.PostVirtualizationClusterGroupsCmd)
	virtualizationPostCmd.AddCommand(virtualization.PostVirtualizationClustersCmd)
	virtualizationPostCmd.AddCommand(virtualization.PostVirtualizationClusterTypesCmd)
	virtualizationPostCmd.AddCommand(virtualization.PostVirtualizationInterfacesCmd)
	virtualizationPostCmd.AddCommand(virtualization.PostVirtualizationVirtualDisksCmd)
	virtualizationPostCmd.AddCommand(virtualization.PostVirtualizationVirtualMachinesCmd)
	VirtualizationCmd.AddCommand(virtualizationPatchCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationClusterGroupsCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationClusterGroupsByIdCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationClustersCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationClustersByIdCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationClusterTypesCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationClusterTypesByIdCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationInterfacesCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationInterfacesByIdCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationVirtualDisksCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationVirtualDisksByIdCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationVirtualMachinesCmd)
	virtualizationPatchCmd.AddCommand(virtualization.PatchVirtualizationVirtualMachinesByIdCmd)
	VirtualizationCmd.AddCommand(virtualizationDeleteCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationClusterGroupsCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationClusterGroupsByIdCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationClustersCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationClustersByIdCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationClusterTypesCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationClusterTypesByIdCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationInterfacesCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationInterfacesByIdCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationVirtualDisksCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationVirtualDisksByIdCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationVirtualMachinesCmd)
	virtualizationDeleteCmd.AddCommand(virtualization.DeleteVirtualizationVirtualMachinesByIdCmd)
}

func addVpnSubcommandPalettes() {
	rootCmd.AddCommand(VpnCmd)
	VpnCmd.AddCommand(vpnGetCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnIkePoliciesCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnIkePoliciesByIdCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnIkeProposalsCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnIkeProposalsByIdCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnIpsecPoliciesCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnIpsecPoliciesByIdCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnIpsecProfilesCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnIpsecProfilesByIdCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnIpsecProposalsCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnIpsecProposalsByIdCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnL2vpnsCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnL2vpnsByIdCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnL2vpnTerminationsCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnL2vpnTerminationsByIdCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnTunnelGroupsCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnTunnelGroupsByIdCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnTunnelsCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnTunnelsByIdCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnTunnelTerminationsCmd)
	vpnGetCmd.AddCommand(vpn.GetVpnTunnelTerminationsByIdCmd)
	VpnCmd.AddCommand(vpnPostCmd)
	vpnPostCmd.AddCommand(vpn.PostVpnIkePoliciesCmd)
	vpnPostCmd.AddCommand(vpn.PostVpnIkeProposalsCmd)
	vpnPostCmd.AddCommand(vpn.PostVpnIpsecPoliciesCmd)
	vpnPostCmd.AddCommand(vpn.PostVpnIpsecProfilesCmd)
	vpnPostCmd.AddCommand(vpn.PostVpnIpsecProposalsCmd)
	vpnPostCmd.AddCommand(vpn.PostVpnL2vpnsCmd)
	vpnPostCmd.AddCommand(vpn.PostVpnL2vpnTerminationsCmd)
	vpnPostCmd.AddCommand(vpn.PostVpnTunnelGroupsCmd)
	vpnPostCmd.AddCommand(vpn.PostVpnTunnelsCmd)
	vpnPostCmd.AddCommand(vpn.PostVpnTunnelTerminationsCmd)
	VpnCmd.AddCommand(vpnPatchCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnIkePoliciesCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnIkePoliciesByIdCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnIkeProposalsCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnIkeProposalsByIdCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnIpsecPoliciesCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnIpsecPoliciesByIdCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnIpsecProfilesCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnIpsecProfilesByIdCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnIpsecProposalsCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnIpsecProposalsByIdCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnL2vpnsCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnL2vpnsByIdCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnL2vpnTerminationsCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnL2vpnTerminationsByIdCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnTunnelGroupsCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnTunnelGroupsByIdCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnTunnelsCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnTunnelsByIdCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnTunnelTerminationsCmd)
	vpnPatchCmd.AddCommand(vpn.PatchVpnTunnelTerminationsByIdCmd)
	VpnCmd.AddCommand(vpnDeleteCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnIkePoliciesCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnIkePoliciesByIdCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnIkeProposalsCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnIkeProposalsByIdCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnIpsecPoliciesCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnIpsecPoliciesByIdCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnIpsecProfilesCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnIpsecProfilesByIdCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnIpsecProposalsCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnIpsecProposalsByIdCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnL2vpnsCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnL2vpnsByIdCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnL2vpnTerminationsCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnL2vpnTerminationsByIdCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnTunnelGroupsCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnTunnelGroupsByIdCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnTunnelsCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnTunnelsByIdCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnTunnelTerminationsCmd)
	vpnDeleteCmd.AddCommand(vpn.DeleteVpnTunnelTerminationsByIdCmd)
}

func addWirelessSubcommandPalettes() {
	rootCmd.AddCommand(WirelessCmd)
	WirelessCmd.AddCommand(wirelessGetCmd)
	wirelessGetCmd.AddCommand(wireless.GetWirelessWirelessLanGroupsCmd)
	wirelessGetCmd.AddCommand(wireless.GetWirelessWirelessLanGroupsByIdCmd)
	wirelessGetCmd.AddCommand(wireless.GetWirelessWirelessLansCmd)
	wirelessGetCmd.AddCommand(wireless.GetWirelessWirelessLansByIdCmd)
	wirelessGetCmd.AddCommand(wireless.GetWirelessWirelessLinksCmd)
	wirelessGetCmd.AddCommand(wireless.GetWirelessWirelessLinksByIdCmd)
	WirelessCmd.AddCommand(wirelessPostCmd)
	wirelessPostCmd.AddCommand(wireless.PostWirelessWirelessLanGroupsCmd)
	wirelessPostCmd.AddCommand(wireless.PostWirelessWirelessLansCmd)
	wirelessPostCmd.AddCommand(wireless.PostWirelessWirelessLinksCmd)
	WirelessCmd.AddCommand(wirelessPatchCmd)
	wirelessPatchCmd.AddCommand(wireless.PatchWirelessWirelessLanGroupsCmd)
	wirelessPatchCmd.AddCommand(wireless.PatchWirelessWirelessLanGroupsByIdCmd)
	wirelessPatchCmd.AddCommand(wireless.PatchWirelessWirelessLansCmd)
	wirelessPatchCmd.AddCommand(wireless.PatchWirelessWirelessLansByIdCmd)
	wirelessPatchCmd.AddCommand(wireless.PatchWirelessWirelessLinksCmd)
	wirelessPatchCmd.AddCommand(wireless.PatchWirelessWirelessLinksByIdCmd)
	WirelessCmd.AddCommand(wirelessDeleteCmd)
	wirelessDeleteCmd.AddCommand(wireless.DeleteWirelessWirelessLanGroupsCmd)
	wirelessDeleteCmd.AddCommand(wireless.DeleteWirelessWirelessLanGroupsByIdCmd)
	wirelessDeleteCmd.AddCommand(wireless.DeleteWirelessWirelessLansCmd)
	wirelessDeleteCmd.AddCommand(wireless.DeleteWirelessWirelessLansByIdCmd)
	wirelessDeleteCmd.AddCommand(wireless.DeleteWirelessWirelessLinksCmd)
	wirelessDeleteCmd.AddCommand(wireless.DeleteWirelessWirelessLinksByIdCmd)
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.demo-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCircuitsSubcommandPalettes()
	addCoreSubcommandPalettes()
	addDcimSubcommandPalettes()
	addExtrasSubcommandsPalettes()
	addIpamSubcommandPalettes()
	addTenancySubcommandPalettes()
	addUsersSubcommandPalettes()
	addVirtualizationSubcommandPalettes()
	addVpnSubcommandPalettes()
	addWirelessSubcommandPalettes()
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(CompletionCmd)
}

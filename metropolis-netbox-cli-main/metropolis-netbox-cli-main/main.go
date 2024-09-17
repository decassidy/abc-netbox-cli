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

package main

import (
	"github.com/decassidy/metropolis-netbox-cli/cmd"
	"github.com/spf13/cobra/doc"
	"log"
)

func emptyStrFunction(_ string) string {
	return ""
}

func main() {
	cmd.Execute()

	err := doc.GenMarkdownTreeCustom(cmd.CircuitsCmd, "/Users/dcassidy/GolandProjects/metropolis-netbox-cli/Docs/Circuits_Documents", emptyStrFunction, emptyStrFunction)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTreeCustom(cmd.CoreCmd, "/Users/dcassidy/GolandProjects/metropolis-netbox-cli/Docs/Core_Documents", emptyStrFunction, emptyStrFunction)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTreeCustom(cmd.DcimCmd, "/Users/dcassidy/GolandProjects/metropolis-netbox-cli/Docs/DCIM_Documents", emptyStrFunction, emptyStrFunction)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTreeCustom(cmd.ExtrasCmd, "/Users/dcassidy/GolandProjects/metropolis-netbox-cli/Docs/Extras_Documents", emptyStrFunction, emptyStrFunction)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTreeCustom(cmd.IpamCmd, "/Users/dcassidy/GolandProjects/metropolis-netbox-cli/Docs/IPAM_Documents", emptyStrFunction, emptyStrFunction)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTreeCustom(cmd.TenancyCmd, "/Users/dcassidy/GolandProjects/metropolis-netbox-cli/Docs/Tenancy_Documents", emptyStrFunction, emptyStrFunction)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTreeCustom(cmd.UsersCmd, "/Users/dcassidy/GolandProjects/metropolis-netbox-cli/Docs/Users_Documents", emptyStrFunction, emptyStrFunction)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTreeCustom(cmd.VirtualizationCmd, "/Users/dcassidy/GolandProjects/metropolis-netbox-cli/Docs/Virtualization_Documents", emptyStrFunction, emptyStrFunction)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTreeCustom(cmd.VpnCmd, "/Users/dcassidy/GolandProjects/metropolis-netbox-cli/Docs/VPN_Documents", emptyStrFunction, emptyStrFunction)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTreeCustom(cmd.WirelessCmd, "/Users/dcassidy/GolandProjects/metropolis-netbox-cli/Docs/Wireless_Documents", emptyStrFunction, emptyStrFunction)
	if err != nil {
		log.Fatal(err)
	}

}

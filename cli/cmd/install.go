/*
Copyright © 2022 Om More <ommore501@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"npm3/cli/common"
	"os"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runInstallCmd,
}

func runInstallCmd(cmd *cobra.Command, args []string) {
	cmd.Println("Connecting To RPC Node")
	client := common.GetClient()
	cmd.Println("Connecting To Smart Contract")
	pkgMng := common.GetPackageManagerInstance(client)
	cmd.Println("Getting latest package version")
	res, err := pkgMng.NameToPackage(nil, args[0])
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	cmd.Println("Fetching latest package release")
	dataHash, err := pkgMng.GetRelease(nil, args[0], res.DefaultVersion)
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	fmt.Println("data hash is", dataHash)
}
func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

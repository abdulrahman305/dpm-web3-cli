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
	"archive/zip"
	"fmt"
	"io"
	"npm3/cli/common"
	"os"
	"path/filepath"
	"strings"

	ipfsShell "github.com/ipfs/go-ipfs-api"
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
	Run:     runInstallCmd,
	Aliases: []string{"i"},
}

func runInstallCmd(cmd *cobra.Command, args []string) {
	err := os.Mkdir("node_modules", 755)
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	cmd.Println("Connecting To RPC Node")
	client := common.GetClient()
	cmd.Println("Connecting To Smart Contract")
	pkgMng := common.GetPackageManagerInstance(client)
	cmd.Println("Getting latest package version")
	pkgName := args[0]
	res, err := pkgMng.NameToPackage(nil, pkgName)
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	cmd.Println("Fetching latest package release")
	dataHash, err := pkgMng.GetRelease(nil, pkgName, res.DefaultVersion)
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	// ipfs := ipfsShell.NewLocalShell()
	ipfs := ipfsShell.NewShell("http://localhost:5001")
	tempDir, err := os.MkdirTemp("", "npm3-*")
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	err = ipfs.Get(dataHash, tempDir)
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	archive, err := zip.OpenReader(filepath.Join(tempDir, dataHash))
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	defer archive.Close()
	os.Mkdir("node_modules", os.ModeDevice)
	dst := fmt.Sprintf("node_modules/%s", pkgName)
	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			return
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	cmd.Println("Installed")

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

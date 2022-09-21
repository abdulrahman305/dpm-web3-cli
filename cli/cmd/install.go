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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"npm3/cli/common"
	"npm3/cli/types"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install package",
	Long:    `Installs package, for e.g. dpm i packageABC`,
	Run:     runInstallCmd,
	Aliases: []string{"i"},
}

func runInstallCmd(cmd *cobra.Command, args []string) {
	os.Mkdir("node_modules", 755)
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
	dataHash, err := pkgMng.GetRelease(nil, pkgName, res.DefaultVersion)
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	tempDir, err := os.MkdirTemp("", "npm3-*")
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	out, err := os.Create(filepath.Join(tempDir, dataHash))
	defer out.Close()
	gateway := "https://ipfs.io/ipfs"
	url := fmt.Sprintf("%s/%s", gateway, dataHash)
	cmd.Println("Downloading release", res.DefaultVersion)
	resp, err := http.Get(url)
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	archive, err := zip.OpenReader(filepath.Join(tempDir, dataHash))
	if err != nil {
		cmd.PrintErrln(err.Error())
		os.Exit(1)
	}
	defer archive.Close()
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
	pkgJsonBytes, _ := os.ReadFile("package.json")
	var pkgJson types.PackageJSON
	json.Unmarshal(pkgJsonBytes, &pkgJson)
	if pkgJson.Dependencies == nil {
		pkgJson.Dependencies = make(map[string]string)
	}
	pkgJson.Dependencies["pkgName"] = res.DefaultVersion
	pkgJsonBytes, _ = json.MarshalIndent(pkgJson, "", "  ")
	os.WriteFile("package.json", pkgJsonBytes, 666)
	cmd.Println("Installed")

}
func init() {
	rootCmd.AddCommand(installCmd)
}

// Copyright skoved

package cmd

import (
	"fmt"
	"os"

	"github.com/skoved/kubearchive-config-cli/pkg/env"
	yFiles "github.com/skoved/kubearchive-config-cli/pkg/files"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const dataDirKey = "dataDir"

var (
	filesDir     string
	filesCommand = &cobra.Command{
		Use:   "files",
		Short: "Print the yaml files containing a KubeArchiveConfig",
		Run:   listFiles,
	}
)

func init() {
	filesCommand.Flags().StringVarP(&filesDir, "dir", "d", env.XdgDataHome(), "List KubeArchiveConfig yaml files")
	viper.BindPFlag(dataDirKey, filesCommand.Flags().Lookup("dir"))
	rootCmd.AddCommand(filesCommand)
}

func listFiles(cmd *cobra.Command, args []string) {
	dataDir := viper.GetString(dataDirKey)
	files, err := os.ReadDir(dataDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open directory:", dataDir)
	}
	for _, file := range files {
		if !file.IsDir() && yFiles.IsYaml(fmt.Sprintf("%s/%s", dataDir, file.Name())) {
			// TODO: Replace with logic to select a yaml file to apply
			fmt.Println(file.Name())
		}
	}
}

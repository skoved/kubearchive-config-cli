package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const appName = "kac"

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   appName,
		Short: "Kac is a cli to help manage different KubeArchiveConfigs",
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/kac/config or $HOME/.kac/config)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// initConfig tells viper where to look for the config file. The precedence is as follows:
// 1. file specified with the --config flag
// 2. $XDG_CONFIG_HOME/kac (if $XDG_CONFIG_HOME is not set, $HOME/.config is used)
// 3. $HOME/.kac
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		userCfgDir, err := os.UserConfigDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		viper.AddConfigPath(fmt.Sprintf("%s/.%s", homeDir, appName))
		viper.AddConfigPath(fmt.Sprintf("%s/%s", userCfgDir, appName))
		viper.SetConfigName("config")
		viper.SetEnvPrefix("kac")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		// don't exit if the config file is not found
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

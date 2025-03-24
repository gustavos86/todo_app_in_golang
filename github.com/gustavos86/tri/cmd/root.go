/*
Copyright © 2025 GoLang Cobra CLI
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo application",
	Long: `Tri will help you get more done in less time.
It's designed to be as simple as possible to help you accomplish your goals.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Read in config file and ENV variables if set
func initConfig() {
	viper.SetConfigName(".tri")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tri")

	// If a config file is found, read it in
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config YAML file not found")
		os.Exit(1)
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

var dataFile string
var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// home, err := homedir.Dir()
	// if err != nil {
	// 	log.Println("Unable to detect from home directory. Please set data file using --datafile.")
	// }

	// var default_storage_path = home + string(os.PathSeparator) + ".tridos.json"

	// // rootCmd.PersistentFlags().StringVar(&dataFile,
	// // 	"datafile",
	// // 	default_storage_path,
	// // 	"data file to store todos")

	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.tri.yaml)")
}

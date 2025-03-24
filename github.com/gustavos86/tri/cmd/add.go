/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"src/github.com/gustavos86/cobra/src/github.com/gustavos86/tri/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	datafile := viper.GetString("datafile")
	if datafile == "" {
		fmt.Print("Config YAML file does not contain the variable datafile\n")
		os.Exit(1)
	}
	fmt.Print("Storage JSON file: " + datafile + "\n")

	items, err := todo.ReadItems(datafile)
	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}

	err = todo.SaveItems(datafile, items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority,
		"priority",
		"p",
		2,
		"Priority:1,2,3")
}

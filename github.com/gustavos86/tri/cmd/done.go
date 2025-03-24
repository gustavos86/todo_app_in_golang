/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"src/github.com/gustavos86/cobra/src/github.com/gustavos86/tri/todo"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark Item as Done",
	Long:    `The done flag marks an Item as Done`,
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
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

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i <= len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done")

		sort.Sort(todo.ByPri(items))
		todo.SaveItems(datafile, items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

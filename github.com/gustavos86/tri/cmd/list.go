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
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all the tasks",
	Long:  `List will display all the tasks`,
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todos")
}

var (
	doneOpt bool
	allOpt  bool
)

func listRun(cmd *cobra.Command, args []string) {
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

	sort.Sort(todo.ByPri(items))
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Text+"\t")
		}
	}

	w.Flush()
}

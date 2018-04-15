package cmd

import (
	"fmt"

	"github.com/adityakeyal/gocli/command"
)

var newlist = &command.Command{
	Name:    "list",
	Use:     "Test Command",
	Short:   "Short Desc",
	Long:    `Long Desc`,
	Execute: listExecute,
}

func listExecute(args []string) {
	fmt.Println("In list_execute")
}

func init() {
	rootCmd.AddCommand(newlist)
}

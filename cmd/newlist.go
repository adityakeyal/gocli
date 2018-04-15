package cmd

import (
	"flag"
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
	flags := parseArguments(args)

	fmt.Println(*flags.all)
	fmt.Println(*flags.name)

}

type listFlag struct {
	all  *bool
	name *string
}

func parseArguments(args []string) listFlag {
	var flags listFlag

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	flags.name = listCommand.String("name", "", "Name of the command")
	flags.all = listCommand.Bool("all", false, "List All")
	listCommand.Parse(args[1:])

	return flags
}

func init() {
	rootCmd.AddCommand(newlist)
}

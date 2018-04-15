package cmd

import (
	"fmt"
	"os"

	"github.com/adityakeyal/gocli/command"
)

//var cfgFile string

var rootCmd = &command.Command{
	Use:     "Test Command",
	Short:   "Short Desc",
	Long:    `Long Desc`,
	Name:    "root",
	Execute: root_execute,
}

func root_execute(args []string) {

	fmt.Println("In root Command")

	//loop over all child commands and check if any of them satisfy the sub command or not
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	args := os.Args[1:]

	if len(args) > 0 {

		subCommand := args[0]

		isExecuted := false

		for _, x := range rootCmd.SubCommand {

			if subCommand == x.Name {
				x.Execute(nil)
				isExecuted = true
				break
			}

		}

		if isExecuted {
			//help command
		}

	} else {
		rootCmd.Execute(args)
	}

}

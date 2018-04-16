package command

import (
	"os"
)

//Command used to decide which child to call
type Command struct {
	Default    bool
	Name       string
	Use        string
	Short      string
	Long       string
	Execute    func(args []string)
	SubCommand []*Command
}

//AddCommand - Command
func (rc *Command) AddCommand(cmd *Command) {
	rc.SubCommand = append(rc.SubCommand, cmd)
}

//AddDefaultCommand - Used to add default commands
func (rc *Command) AddDefaultCommand(cmd *Command) {
	cmd.Default = true
	rc.SubCommand = append(rc.SubCommand, cmd)
}

//Execute - The default execute which should be called by the main
func Execute() {
	args := os.Args[1:]

	if len(args) > 0 {

		subCommand := args[0]

		isExecuted := false

		for _, x := range RootCmd.SubCommand {

			if subCommand == x.Name {
				x.Execute(args)
				isExecuted = true
				break
			}

		}

		if !isExecuted {
			panic("Invalid option")
		}

	} else {
		checkDefaultCommand(args)
	}
}

//RootCmd - Root Command
var RootCmd = &Command{
	Use:     "Root Command",
	Short:   "Holder for all commands",
	Long:    `Long Desc`,
	Name:    "root",
	Execute: rootExecute,
}

func rootExecute(args []string) {

	panic("No default command defined")

}

func checkDefaultCommand(args []string) {
	for _, x := range RootCmd.SubCommand {
		if x.Default {
			x.Execute(args)
			break
		}
	}
}

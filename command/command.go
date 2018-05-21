package command

import (
	"fmt"
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
	SubHelp    func() string
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
				checkIfHelp(args, x)
				x.Execute(args)
				isExecuted = true
				break
			}

		}

		if !isExecuted {
			//print Default command options
			printCommandOptions()
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

func checkIfHelp(args []string, command *Command) {
	for _, x := range args {
		if x == "--help" || x == "-h" {
			fmt.Println(command.Name)
			fmt.Println(command.Short)
			fmt.Println(command.Long)

		}
	}
}

func printCommandOptions() {

	fmt.Println("Usage:")
	fmt.Println("<appname> command")
	fmt.Println("")
	fmt.Println("Available Commands:")

	for _, x := range RootCmd.SubCommand {

		defaultHelp := ""
		if x.Default {
			defaultHelp = " (default)"
		}

		fmt.Println("  > " + x.Name + defaultHelp + " [" + x.Short + "] : " + x.Long)
		if x.SubHelp != nil {
			fmt.Println(x.SubHelp())
		}
	}
}

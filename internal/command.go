package internal

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Command interface {
	Name() string
	Example() string
	Help() string
	LongHelp() string
	Register(*flag.FlagSet)
	Run()
}

type CommandRoot struct {
	Name     string
	commands []Command
}

func CommandInit(name string) *CommandRoot {
	return &CommandRoot{
		Name: name,
	}
}

func (cr *CommandRoot) Start(commandList []Command) error {
	// existe comando registrado?
	// existe algum argumento?
	// o argumento é válido?
	//

	if len(commandList) == 0 {
		return errors.New("put some command!")
	}

	cr.commands = commandList

	if len(os.Args) < 2 {
		cr.showHelp()
		return nil
	}

	userCommand := ArgsFilter(os.Args[1:])

	// fmt.Println(userCommand)

	if userCommand.Command == "" {
		cr.showHelp()
		return nil
	}

	if userCommand.Command == "help" {
		cr.showHelp()
		return nil

	}

	for _, command := range cr.commands {
		if userCommand.Command == command.Name() {
			fs := flag.NewFlagSet(command.Name(), flag.ExitOnError)
			command.Register(fs)
			fs.Parse(os.Args[2:])
			command.Run()

			return nil
		}
	}

	cr.showHelp()

	return nil
}

func (cr *CommandRoot) showHelp() {
	fmt.Println("HELP!")
}

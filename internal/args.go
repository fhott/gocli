package internal

import "regexp"

type UserCommand struct {
	Command   string
	Arguments []string
}

func ArgsFilter(args []string) UserCommand {
	regexValidator := regexp.MustCompile("(?m)-")

	commandSet := false

	var finalUserCommand UserCommand

	for _, argument := range args {
		isFlag := regexValidator.MatchString(argument)

		if !isFlag && !commandSet {
			finalUserCommand.Command = argument
			commandSet = true
		} else if isFlag {
			finalUserCommand.Arguments = append(finalUserCommand.Arguments, argument)
		}
	}

	return finalUserCommand
}

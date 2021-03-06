package completer

import (
	"github.com/c-bata/go-prompt"
	"strings"
	"github.com/prazd/go-cli/options"
)

func Completer(d prompt.Document) []prompt.Suggest{
	args := strings.Split(d.TextBeforeCursor(), " ")
	return argumentsCompleter(args)
}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(options.InitialOptions, args[0], true)
	}
	first := args[0]
	switch first{
	case "plasma":
		if len(args) == 2{
			return prompt.FilterHasPrefix(options.PlasmaOptions, args[1], true)
		}
	case "eth":
		if len(args) == 2{
			return prompt.FilterHasPrefix(options.EthOptions, args[1], true)
		}
	case "main":
		if len(args) == 2{
			return prompt.FilterHasPrefix(options.MainOptions, args[1], true)
		}
	}
return []prompt.Suggest{}
}
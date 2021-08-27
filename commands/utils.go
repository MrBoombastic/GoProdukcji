package commands

import (
	"errors"
	"fmt"
)

var List = map[string]CommandData{
	"stats":  StatsCommand,
	"help":   HelpCommand,
	"search": SearchCommand,
	"last":   LastCommand,
}

func FindCommand(name string) (CommandData, error) {
	if List[name].Command != nil {
		return List[name], nil
	} else {
		for com := range List { //Commands loop
			for _, alias := range List[com].Aliases { //Aliases of command loop
				if alias == name {
					return List[com], nil
				}
			}
		}
	}
	return CommandData{}, errors.New("not found")
}

var GenerateHelpOutput = func(prefix string) {
	output := ""
	for com := range List {
		output += fmt.Sprintf("- `%v%v` - %v\n", prefix, com, List[com].Description)
		if len(List[com].Usage) > 0 {
			output += fmt.Sprintf("Użycie: `%v`\n", List[com].Usage)
		}
	}
	HelpOutput = output
}

var HelpOutput string

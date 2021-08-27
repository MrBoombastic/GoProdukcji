package commands

import "fmt"

var List = map[string]CommandData{
	"stats":  StatsCommand,
	"help":   HelpCommand,
	"search": SearchCommand,
	"last":   LastCommand,
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

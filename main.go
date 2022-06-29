package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"me.kryptk.overcommit/components"
	"os"
)

/*
feat: The new feature you're adding to a particular application
fix: A bug fix
style: Feature and updates related to styling
refactor: Refactoring a specific section of the codebase
test: Everything related to testing
docs: Everything related to documentation
chore: Regular code maintenance.

from : https://www.freecodecamp.org/news/writing-good-commit-messages-a-practical-guide/
*/

func main() {
	const defaultWidth = 20

	m := components.PageView{Page: components.SELECTION}
	
	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

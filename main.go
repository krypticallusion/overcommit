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
	hook := "$PWD/.git/hooks/prepare-commit-msg"

	_, err := os.ReadDir(os.ExpandEnv("$PWD/.git"))
	if err != nil {
		fmt.Println("not a git repository")
		return
	}

	if len(os.Args) <= 1 {
		fmt.Println("Hi, set up using -i or --init")
		return
	}

	if len(os.Args) > 1 {
		// check if initing
		isInit := os.Args[1] == "-i" || os.Args[1] == "--init"

		if isInit {
			//back up previous stuff
			_ = os.Rename(os.ExpandEnv(fmt.Sprintf("%s.sample", hook)), os.ExpandEnv(fmt.Sprintf("%s.bak", hook)))
			_ = os.WriteFile(os.ExpandEnv(hook), []byte("overcommit $1 $2"), 0755)

			fmt.Println(os.ExpandEnv("Successfully set up overcommit in $PWD!"))
		}

		return
	}

	m := components.PageView{Page: components.SELECTION}

	if err := tea.NewProgram(m, tea.WithAltScreen(), tea.WithANSICompressor()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

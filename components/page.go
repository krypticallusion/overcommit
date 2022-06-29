package components

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Page int

const (
	SELECTION = iota
	MSG
)

type PageView struct {
	Page     Page
	selected Provider
	message  string
}

func (p PageView) Init() tea.Cmd {
	return nil
}

func (p PageView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return p, tea.Quit
		}
	}

	if p.Page == SELECTION {
		return keywords.Update(msg, p)
	}

	return CommitViewInstance.Update(msg, p)
}

func (p PageView) View() string {
	switch p.Page {
	case SELECTION:
		return keywords.View()
	}

	return CommitViewInstance.View(p)
}

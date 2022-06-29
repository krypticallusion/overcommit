package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"io"
	"me.kryptk.overcommit/utils"
	"os"
)

var (
	term = termenv.TrueColor
)

type Provider struct {
	isSelected  bool
	Description string
	Keyword     string
}

func (p Provider) FilterValue() string {
	return p.Keyword
}

var keywords = genList()

var GLOBAL = []Provider{
	{
		Description: "introduce new features",
		Keyword:     "feat",
	},
	{
		Description: "fix a bug",
		Keyword:     "fix",
	}, {
		Description: "updates and features related to styling",
		Keyword:     "style",
	}, {
		Description: "refactor code",
		Keyword:     "refactor",
	}, {
		Description: "add or update tests",
		Keyword:     "test",
	}, {
		Description: "add or update documentation",
		Keyword:     "docs",
	}, {
		Description: "regular maintenance",
		Keyword:     "chore",
	},
}

func genList() TypeSelectorView {
	items := ProviderToItem(GLOBAL)

	li := list.New(items, listDelegate{}, 40, 14)
	li.SetShowTitle(true)
	li.Styles.StatusBar = lipgloss.NewStyle().UnsetPaddingLeft().UnsetMarginLeft().MarginBottom(1)
	li.Title = "What is the type of the commit?"
	li.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("#8AA8F9"))
	li.Styles.TitleBar = lipgloss.NewStyle().UnsetPaddingLeft().UnsetMarginLeft().Bold(true)
	li.Styles.HelpStyle = lipgloss.NewStyle().UnsetFaint()

	li.SetFilteringEnabled(true)
	return TypeSelectorView{
		view: li,
	}
}

type TypeSelectorView struct {
	view list.Model
}

type listDelegate struct{}

func (l listDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	selected := index == m.Index()

	i, ok := item.(Provider)
	if !ok {
		return
	}

	txt := fmt.Sprintf("(%s) - %s [%d]", i.Keyword, i.Description, index+1)

	if selected {
		txt = termenv.String(txt).Foreground(term.Color("#8AA8F9")).Underline().String()
	} else {
		txt = termenv.String(txt).Faint().String()
	}

	_, _ = fmt.Fprint(w, txt)
}

func (l listDelegate) Height() int { return 1 }

func (l listDelegate) Spacing() int {
	return 0
}

func (l listDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (l TypeSelectorView) View() string {
	return l.view.View()
}

func (l *TypeSelectorView) Update(msg tea.Msg, v PageView) (PageView, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter", tea.KeyRight.String():
			v.selected = l.view.SelectedItem().(Provider)

			if len(os.Args) >= 3 {
				fileName := os.Args[1]
				_ = utils.AddToCommitMsg(utils.BuildPrefixWithMsg(v.selected.Keyword, ""), fileName)

				return PageView{}, tea.Quit
			}

			v.Page = MSG
		}
	}

	l.view, cmd = l.view.Update(msg)

	return v, cmd
}

func ProviderToItem(p []Provider) []list.Item {
	items := make([]list.Item, len(p))

	for i, k := range p {
		items[i] = k
	}

	return items
}

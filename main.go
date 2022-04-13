package main

import (
	"os"

	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type ConfigItems struct {
	Item []item
}

type item struct {
	ShellCommand, Desc string
}

func (i item) Title() string       { return i.ShellCommand }
func (i item) Description() string { return i.Desc }
func (i item) FilterValue() string { return i.Desc }

type model struct {
	list   list.Model
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch key := msg.String(); key {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			m.choice = m.list.SelectedItem().(item).ShellCommand
			return m, tea.Quit

		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	blob, err := os.ReadFile(xdg.ConfigHome + "/shells/config.toml")
	if err != nil {
		log.Fatal(err)
	}
	var favorites ConfigItems
	if _, err := toml.Decode(string(blob), &favorites); err != nil {
		log.Fatal(err)
	}

	var items []list.Item
	for _, s := range favorites.Item {
		items = append(items, s)
	}

	log.Println("Home directory:", xdg.Home)

	m := model{
		list:   list.New(items, list.NewDefaultDelegate(), 0, 0),
		choice: "",
	}
	m.list.Title = "Shells"

	p := tea.NewProgram(m, tea.WithAltScreen())

	m2, err := p.StartReturningModel()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if m2, ok := m2.(model); ok && m2.choice != "" {
		writeToStdin(m2.choice)
	}
}

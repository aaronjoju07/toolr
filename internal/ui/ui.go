package ui

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/aaronjoju07/toolr/internal/registry"
)

type model struct {
    cursor int
    tools  []registry.Tool
}

func InitialModel() model {
    return model{
        tools: registry.ListTools(),
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q":
            return m, tea.Quit
        case "up":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down":
            if m.cursor < len(m.tools)-1 {
                m.cursor++
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    s := "Toolr CLI Tools:\n\n"
    for i, tool := range m.tools {
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor
        }
        s += cursor + " " + tool.Name + "\n"
    }
    s += "\nPress q to quit.\n"
    return s
}

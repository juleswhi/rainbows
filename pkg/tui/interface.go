package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	// "github.com/charmbracelet/log"
	"github.com/juleswhi/rainbows/pkg/renderer"
)

type model struct {
	width  int
	height int

	lines  []string
	border lipgloss.Style
	cols   []renderer.Colour
}

func New() *model {
	return &model{
		lines:  []string{"Hello"},
		border: lipgloss.NewStyle(),
		cols:   renderer.RayCast(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "e":
			return m, tea.ClearScreen
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m model) View() string {
	var sb strings.Builder

	lines := []string{}

	st := lipgloss.NewStyle()

    for j := range 192 {
        for i := range 255 {
			for range 2 {
				col := m.cols[(j*256)+i]
				idxstr := fmt.Sprintf(
                    "#%02X%02X%02X",
                    int(col.R),
                    int(col.G),
                    int(col.B),
                )

                s := st.Foreground(lipgloss.Color(idxstr)).Render("m")

				sb.WriteString(s)
			}

		}
		lines = append(lines, sb.String())
		sb.Reset()
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, lines...),
	)
}

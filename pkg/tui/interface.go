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
	rays   []renderer.Ray
}

func New() *model {
	return &model{
		lines:  []string{"Hello"},
		border: lipgloss.NewStyle(),
		rays:   renderer.RayCast(),
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
			for range 3 {
				ray := m.rays[(j*256)+i]
				idxstr := fmt.Sprintf("#%02X%02X%02X", int(ray.Direction.X), int(ray.Direction.Y), 50)

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
	// idxstr := fmt.Sprintf("#%02X%02X%02X", int(m.rays[0].Direction.X), int(m.rays[0].Direction.Y), int(200))
	// fmt.Println(idxstr)
	// return fmt.Sprintf("%d", sb.Cap())
}

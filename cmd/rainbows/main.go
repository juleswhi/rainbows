package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/juleswhi/rainbows/pkg/tui"
)

func main() {
    p := tea.NewProgram(tui.New(), tea.WithFPS(120), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Oopsiues")
        os.Exit(1)
    }
}

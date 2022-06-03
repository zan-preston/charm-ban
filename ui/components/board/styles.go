package board

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/zan-preston/charm-ban/ui/styles"
)

var (
	boardTextStyle = lipgloss.NewStyle().Foreground(styles.DefaultTheme.MainText)
	boardStyle     = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(styles.DefaultTheme.Border)
)

package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zan-preston/charm-ban/ui/components/help"
	"github.com/zan-preston/charm-ban/utils"
)

type Model struct {
  keys utils.KeyMap
  help help.Model

}


func NewModel() Model {
  return Model {
    keys: utils.Keys,
    help: help.NewModel(),
  }
}

// func initScreen() tea.Msg {
//   // config, err := config.ParseConfig()
//   // if err != nil {
//   //   return errMsg{err}
//   // }
//   // return initMsg{Config: config}
// }

func (m Model) Init() tea.Cmd {
  // return tea.Batch(initScreen, tea.EnterAltScreen)
  return tea.EnterAltScreen
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var (
    cmd tea.Cmd
    cmds []tea.Cmd
  )
  switch msg := msg.(type) {
  case tea.KeyMsg:
   switch  {
   case key.Matches(msg, m.keys.Quit):
     cmd = tea.Quit
   } 
 case tea.WindowSizeMsg:
   m.onWindowSizedChange(msg)
  }
  cmds = append(cmds, cmd)
  return m, tea.Batch(cmds...)
}

func (m *Model) onWindowSizedChange (msg tea.WindowSizeMsg) {
  m.help.SetWidth(msg.Width)
  // m.ctx.ScreenWidth = msg.Width
  // m.ctx.ScreenHeight = msg.Height
  // m.ctx.MainContentHeight = msg.Height
  // m.syncMainContentWidth()
}

func (m Model) View() string {
  s := strings.Builder{}
  // Main content will be the lanes of the kanban
  mainContent := "lanes here"
  s.WriteString(mainContent)
  s.WriteString("\n")
  s.WriteString(m.help.View())
  return s.String()

}

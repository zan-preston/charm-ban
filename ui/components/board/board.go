package board

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
  // array of lanes
}

func NewModel() Model {
  return Model {

  }
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
  return m, nil
}


func (m Model) View() string {
  var lanes = "Hopefully I'm in a box"
  return boardStyle.Copy().Render(lanes)
}

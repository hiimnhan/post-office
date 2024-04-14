package ui

func (m model) NewInputModel() string {
	return m.InputStyle().Render(" ")
}

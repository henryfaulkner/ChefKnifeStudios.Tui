package views

import (
	"encoding/json"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	lunchOrderTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("205"))

	labelStyle = lipgloss.NewStyle().
			Width(24).
			Align(lipgloss.Left)

	focusedLabelStyle = lipgloss.NewStyle().
				Width(24).
				Align(lipgloss.Left).
				Foreground(lipgloss.Color("205"))
)

var metadataKeys = []string{
	"score (1-10)",
	"food ordered",
	"order again?",
	"visit outside of work?",
	"notes",
}

type LunchOrderModel struct {
	// Phase tracking: 0 = restaurant, 1 = metadata
	phase int

	// Phase 0: Restaurant
	restaurantInput textinput.Model
	restaurant      string

	// Phase 1: Metadata
	metadataInputs []textinput.Model
	focusIndex     int

	// Submission
	submitted bool
}

func NewLunchOrderModel() LunchOrderModel {
	// Restaurant input
	ri := textinput.New()
	ri.Placeholder = "Torched Hop"
	ri.CharLimit = 156
	ri.Width = 30

	// Metadata inputs
	inputs := make([]textinput.Model, len(metadataKeys))
	for i := range metadataKeys {
		ti := textinput.New()
		ti.CharLimit = 256
		ti.Width = 30
		inputs[i] = ti
	}

	return LunchOrderModel{
		phase:           0,
		restaurantInput: ri,
		metadataInputs:  inputs,
		focusIndex:      0,
	}
}

func (m LunchOrderModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m LunchOrderModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.phase == 0 {
				// Phase 0: Save restaurant and move to phase 1
				m.restaurant = m.restaurantInput.Value()
				if m.restaurant == "" {
					return m, nil // Don't proceed without a restaurant name
				}
				m.phase = 1
				m.restaurantInput.Blur()
				m.metadataInputs[0].Focus()
				return m, textinput.Blink
			} else {
				// Phase 1: Move to next field or submit on last
				if m.focusIndex < len(m.metadataInputs)-1 {
					m.metadataInputs[m.focusIndex].Blur()
					m.focusIndex++
					m.metadataInputs[m.focusIndex].Focus()
					return m, textinput.Blink
				} else {
					// Last field - submit
					m.submitted = true
					return m, nil
				}
			}

		case "up":
			if m.phase == 1 && m.focusIndex > 0 {
				m.metadataInputs[m.focusIndex].Blur()
				m.focusIndex--
				m.metadataInputs[m.focusIndex].Focus()
				return m, textinput.Blink
			}

		case "down":
			if m.phase == 1 && m.focusIndex < len(m.metadataInputs)-1 {
				m.metadataInputs[m.focusIndex].Blur()
				m.focusIndex++
				m.metadataInputs[m.focusIndex].Focus()
				return m, textinput.Blink
			}

		case "tab":
			if m.phase == 1 && m.focusIndex < len(m.metadataInputs)-1 {
				m.metadataInputs[m.focusIndex].Blur()
				m.focusIndex++
				m.metadataInputs[m.focusIndex].Focus()
				return m, textinput.Blink
			}

		case "shift+tab":
			if m.phase == 1 && m.focusIndex > 0 {
				m.metadataInputs[m.focusIndex].Blur()
				m.focusIndex--
				m.metadataInputs[m.focusIndex].Focus()
				return m, textinput.Blink
			}
		}
	}

	// Update the active input
	var cmd tea.Cmd
	if m.phase == 0 {
		m.restaurantInput, cmd = m.restaurantInput.Update(msg)
	} else {
		m.metadataInputs[m.focusIndex], cmd = m.metadataInputs[m.focusIndex].Update(msg)
	}
	return m, cmd
}

func (m LunchOrderModel) View() string {
	var b strings.Builder

	if m.phase == 0 {
		b.WriteString(lunchOrderTitleStyle.Render("What restaurant made your lunch?"))
		b.WriteString("\n\n")
		b.WriteString(m.restaurantInput.View())
		b.WriteString("\n")
	} else {
		b.WriteString(lunchOrderTitleStyle.Render("Restaurant: " + m.restaurant))
		b.WriteString("\n\n")
		b.WriteString("Add details (↑/↓ to navigate, enter to submit):\n\n")

		for i, key := range metadataKeys {
			label := key + ":"
			if i == m.focusIndex {
				b.WriteString(focusedLabelStyle.Render(label))
			} else {
				b.WriteString(labelStyle.Render(label))
			}
			b.WriteString(" ")
			b.WriteString(m.metadataInputs[i].View())
			b.WriteString("\n")
		}
	}

	return b.String()
}

func (m *LunchOrderModel) Focus() {
	if m.phase == 0 {
		m.restaurantInput.Focus()
	} else {
		m.metadataInputs[m.focusIndex].Focus()
	}
}

func (m *LunchOrderModel) Blur() {
	if m.phase == 0 {
		m.restaurantInput.Blur()
	} else {
		m.metadataInputs[m.focusIndex].Blur()
	}
}

func (m LunchOrderModel) IsSubmitted() bool {
	return m.submitted
}

func (m LunchOrderModel) SubmittedRestaurant() string {
	return m.restaurant
}

func (m LunchOrderModel) SubmittedMetadata() string {
	data := make(map[string]string)
	for i, key := range metadataKeys {
		val := strings.TrimSpace(m.metadataInputs[i].Value())
		if val != "" {
			data[key] = val
		}
	}
	if len(data) == 0 {
		return ""
	}
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func (m *LunchOrderModel) Clear() {
	m.phase = 0
	m.restaurant = ""
	m.restaurantInput.SetValue("")
	m.focusIndex = 0
	m.submitted = false
	for i := range m.metadataInputs {
		m.metadataInputs[i].SetValue("")
		m.metadataInputs[i].Blur()
	}
}

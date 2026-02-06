package views

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var officeTallyTitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("205"))

type OfficeTallyModel struct {
	// Phase tracking: 0 = restaurant, 1 = tally
	phase int

	// Phase 0: Restaurant
	restaurantInput textinput.Model
	restaurant      string

	// Phase 1: Tally
	tallyInput textinput.Model
	tally      int64

	// Submission
	submitted bool
}

func NewOfficeTallyModel() OfficeTallyModel {
	// Restaurant input
	ri := textinput.New()
	ri.Placeholder = "e.g., Torched Hop"
	ri.CharLimit = 256
	ri.Width = 30

	// Tally input (for numbers)
	ti := textinput.New()
	ti.Placeholder = "e.g., 5"
	ti.CharLimit = 10
	ti.Width = 10

	return OfficeTallyModel{
		phase:           0,
		restaurantInput: ri,
		tallyInput:      ti,
	}
}

func (m OfficeTallyModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m OfficeTallyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.tallyInput.Focus()
				return m, textinput.Blink
			} else {
				// Phase 1: Parse tally and submit
				tallyStr := strings.TrimSpace(m.tallyInput.Value())
				if tallyStr == "" {
					return m, nil // Don't submit without a tally
				}
				tallyNum, err := strconv.ParseInt(tallyStr, 10, 64)
				if err != nil {
					return m, nil // Invalid number
				}
				m.tally = tallyNum
				m.submitted = true
				return m, nil
			}
		}
	}

	// Update the active input
	var cmd tea.Cmd
	if m.phase == 0 {
		m.restaurantInput, cmd = m.restaurantInput.Update(msg)
	} else {
		m.tallyInput, cmd = m.tallyInput.Update(msg)
	}
	return m, cmd
}

func (m OfficeTallyModel) View() string {
	var b strings.Builder

	b.WriteString(officeTallyTitleStyle.Render("Post the office tally"))
	b.WriteString("\n\n")

	if m.phase == 0 {
		b.WriteString("Restaurant:\n")
		b.WriteString(m.restaurantInput.View())
		b.WriteString("\n")
	} else {
		b.WriteString("Restaurant: " + m.restaurant + "\n\n")
		b.WriteString("Tally count:\n")
		b.WriteString(m.tallyInput.View())
		b.WriteString("\n")
	}

	return b.String()
}

func (m *OfficeTallyModel) Focus() {
	if m.phase == 0 {
		m.restaurantInput.Focus()
	} else {
		m.tallyInput.Focus()
	}
}

func (m *OfficeTallyModel) Blur() {
	if m.phase == 0 {
		m.restaurantInput.Blur()
	} else {
		m.tallyInput.Blur()
	}
}

func (m OfficeTallyModel) IsSubmitted() bool {
	return m.submitted
}

func (m OfficeTallyModel) SubmittedRestaurant() string {
	return m.restaurant
}

func (m OfficeTallyModel) SubmittedTally() int64 {
	return m.tally
}

func (m *OfficeTallyModel) Clear() {
	m.phase = 0
	m.restaurant = ""
	m.restaurantInput.SetValue("")
	m.tallyInput.SetValue("")
	m.tally = 0
	m.submitted = false
}

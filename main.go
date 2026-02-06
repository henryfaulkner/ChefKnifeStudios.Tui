package main

import (
	"database/sql"
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"chefknifestudios/tui/db"
	"chefknifestudios/tui/views"
)

type viewState int

const (
	viewHome viewState = iota
	viewLunchOrder
	viewOfficeTally
)

type model struct {
	currentView viewState
	home        views.HomeModel
	lunchOrder  views.LunchOrderModel
	officeTally views.OfficeTallyModel
	database    *sql.DB
	err         error
}

func initialModel(database *sql.DB) model {
	return model{
		currentView: viewHome,
		home:        views.NewHomeModel(),
		lunchOrder:  views.NewLunchOrderModel(),
		officeTally: views.NewOfficeTallyModel(),
		database:    database,
		err:         nil,
	}
}

func (m model) Init() tea.Cmd {
	return m.home.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			if m.currentView == viewLunchOrder {
				m.lunchOrder.Clear()
				m.currentView = viewHome
				m.home.Clear()
				return m, nil
			} else if m.currentView == viewOfficeTally {
				m.officeTally.Clear()
				m.currentView = viewHome
				m.home.Clear()
				return m, nil
			}
		}
	}

	switch m.currentView {
	case viewHome:
		var homeModel tea.Model
		homeModel, cmd = m.home.Update(msg)
		m.home = homeModel.(views.HomeModel)

		// Check if user made a selection
		choice := m.home.Choice()
		if choice != "" {
			m.home.Clear()
			if choice == "Post your own lunch order" {
				m.currentView = viewLunchOrder
				m.lunchOrder.Focus()
				return m, m.lunchOrder.Init()
			} else if choice == "Post the office tally" {
				m.currentView = viewOfficeTally
				m.officeTally.Focus()
				return m, m.officeTally.Init()
			}
		}

	case viewLunchOrder:
		var lunchOrderModel tea.Model
		lunchOrderModel, cmd = m.lunchOrder.Update(msg)
		m.lunchOrder = lunchOrderModel.(views.LunchOrderModel)

		// Check if user submitted a lunch order
		if m.lunchOrder.IsSubmitted() {
			order := &db.LunchOrder{
				SRestaurant: m.lunchOrder.SubmittedRestaurant(),
				JMetadata:   m.lunchOrder.SubmittedMetadata(),
			}
			_, err := db.InsertLunchOrder(m.database, order)
			if err != nil {
				m.err = err
			}
			m.lunchOrder.Clear()
			m.currentView = viewHome
			m.home.Clear()
		}

	case viewOfficeTally:
		var officeTallyModel tea.Model
		officeTallyModel, cmd = m.officeTally.Update(msg)
		m.officeTally = officeTallyModel.(views.OfficeTallyModel)

		// Check if user submitted an office tally entry
		if m.officeTally.IsSubmitted() {
			tally := &db.OfficeTally{
				SRestaurant: m.officeTally.SubmittedRestaurant(),
				ITally:      m.officeTally.SubmittedTally(),
			}
			_, err := db.InsertOfficeTally(m.database, tally)
			if err != nil {
				m.err = err
			}
			m.officeTally.Clear()
			m.officeTally.Focus()
		}
	}

	return m, cmd
}

func (m model) View() string {
	switch m.currentView {
	case viewHome:
		return m.home.View()

	case viewLunchOrder:
		s := m.lunchOrder.View()
		s += "\nPress escape to go back.\n"
		return s

	case viewOfficeTally:
		s := m.officeTally.View()
		s += "\nPress escape to go back.\n"
		return s
	}

	return ""
}

func main() {
	database, err := db.Open("tui.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	p := tea.NewProgram(initialModel(database))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

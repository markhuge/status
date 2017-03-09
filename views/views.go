// Package views provides management for all the display shiz
package views

import (
	"time"

	"github.com/gizak/termui"
	"github.com/markhuge/status/site"
)

// Table models a table of URL statuses
type Table struct {
	Header     []string
	Checks     []*site.Site
	View       *termui.Table
	Dimensions struct{ X, Y, W, H int }
}

// NewTable is a constructor for Table
func NewTable(sites []*site.Site) *Table {
	t := new(Table)
	t.Header = []string{"URL", "Status Code", "Response Time", "Last Checked"}
	// This is dumb
	t.Dimensions = struct{ X, Y, W, H int }{0, 0, 100, 100}
	t.View = termui.NewTable()
	t.Append(t.Header)

	t.Checks = sites
	for _, item := range sites {
		t.Append([]string{item.URL, "n/a", "n/a", "never"})
	}
	return t
}

// Append a row to Table
func (t *Table) Append(row []string) {
	t.View.Rows = append(t.View.Rows, row)
}

// Render Table
func (t *Table) Render() {
	t.View.FgColor = termui.ColorWhite
	t.View.BgColor = termui.ColorDefault
	t.View.X = t.Dimensions.X
	t.View.Y = t.Dimensions.Y
	t.View.Width = t.Dimensions.W
	t.View.Height = t.Dimensions.H
	termui.Render(t.View)
}

// Update iterates over checks and updates Table with output
func (t *Table) Update() {
	for i, check := range t.Checks {

		go func(i int, check *site.Site) {
			res := check.Check()
			t.View.Rows[i+1] = res.String()
			t.Render()
		}(i, check)

	}
}

// TODO implement status bar view
// type StatusBar struct {
// 	Stats string
// }

// Init View
func Init(views []*Table) {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	for _, view := range views {
		go view.Update()
	}
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/r", func(termui.Event) {
		for _, view := range views {
			go view.Update()
		}
	})

	termui.Merge("timer/30s", termui.NewTimerCh(time.Second*30))
	termui.Handle("/timer/30s", func(termui.Event) {
		for _, view := range views {
			go view.Update()
		}
	})
	termui.Loop()
}

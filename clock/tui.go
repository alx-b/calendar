package clock

import (
	"errors"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/alx-b/calendar/logger"
)

type ClockWidget struct {
	Layout *tview.Grid
	digits []*tview.TextView
}

// returns a newly created ClockWidget.
func NewClockWidget(app *tview.Application) ClockWidget {
	grid := tview.NewGrid().
		SetRows(0, 5, 0).
		SetColumns(0, 7, 7, 7, 7, 7, 0).
		AddItem(tview.NewBox(), 1, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewBox(), 1, 6, 1, 1, 0, 0, false).
		AddItem(tview.NewBox(), 0, 0, 1, 7, 0, 0, false).
		AddItem(tview.NewBox(), 2, 0, 1, 7, 0, 0, false)

	grid.SetBorder(true).SetBorderColor(tcell.ColorBlack)

	textViewForDigits := func() *tview.TextView {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetDynamicColors(true).
			SetChangedFunc(func() { app.Draw() })
	}

	digits := []*tview.TextView{}

	for idx := range [5]int{} {
		digits = append(digits, textViewForDigits())
		grid.AddItem(digits[idx], 1, idx+1, 1, 1, 0, 0, false)
	}

	return ClockWidget{
		Layout: grid,
		digits: digits,
	}
}

// Updates border color depending if widget has focus or not.
func (c *ClockWidget) UpdateBorderColor() {
	if c.Layout.HasFocus() {
		c.Layout.SetBorderColor(tcell.ColorRed)
	} else {
		c.Layout.SetBorderColor(tcell.ColorBlack)
	}
}

// Runs update() if no error happens, starts a goroutine.
// First goroutine create a channel, start loopCurrentTime as a goroutine.
// Listen to the channel for a new value.
func (c *ClockWidget) Run() {
	err := c.update(getFormattedTime(time.Now()))

	if err != nil {
		logger.Log.Error.Println(err)
		return
	}

	go func() {
		timeCh := make(chan string)
		go loopCurrentTime(timeCh)
		for s := range timeCh {
			err := c.update(s)
			if err != nil {
				logger.Log.Error.Println(err)
				return
			}
		}
	}()
}

// Updates value and redraw to display.
func (c *ClockWidget) update(newTime string) error {
	coloredNumbers := setColorToChars(newTime, "[red:red]")

	if len(coloredNumbers) != len(c.digits) {
		return errors.New("Needs as many numbers as tview.TextView to display them.")
	}

	for idx := range c.digits {
		c.digits[idx].Clear()
		fmt.Fprintf(
			c.digits[idx],
			"%s",
			coloredNumbers[idx],
		)
	}

	return nil
}

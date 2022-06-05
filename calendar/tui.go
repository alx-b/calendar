package calendar

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

type CalendarWidget struct {
	Layout       *tview.Grid
	monthTable   *tview.Table
	monthName    *tview.TextView
	fullDate     *tview.TextView
	currentTime  time.Time
	selectedTime time.Time
}

// Returns a newly created CalendarWidget.
func NewCalendarWidget(app *tview.Application) CalendarWidget {
	table := tview.NewTable().SetBorders(false)

	monthName := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetChangedFunc(func() { app.Draw() })

	fullDate := tview.NewTextView().SetTextAlign(tview.AlignCenter).SetChangedFunc(func() { app.Draw() })

	grid := tview.NewGrid().
		SetRows(2, 2, 0).
		SetColumns(0, 39, 0).
		AddItem(tview.NewBox(), 0, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewBox(), 0, 2, 1, 1, 0, 0, false).
		AddItem(tview.NewBox(), 1, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewBox(), 1, 2, 1, 1, 0, 0, false).
		AddItem(tview.NewBox(), 2, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewBox(), 2, 2, 1, 1, 0, 0, false).
		AddItem(fullDate, 0, 1, 1, 1, 0, 0, false).
		AddItem(monthName, 1, 1, 1, 1, 0, 0, false).
		AddItem(table, 2, 1, 1, 1, 0, 0, false)

	// SetBorder() returns a box so it can't be declared with tview.NewGrid()
	// without erasing what should be in it
	grid.SetBorder(true).SetBorderColor(tcell.ColorBlack)

	return CalendarWidget{
		Layout:       grid,
		monthTable:   table,
		monthName:    monthName,
		fullDate:     fullDate,
		currentTime:  time.Now(),
		selectedTime: time.Now(),
	}
}

// Updates border color depending if widget has focus or not.
func (c *CalendarWidget) UpdateBorderColor() {
	if c.Layout.HasFocus() {
		c.Layout.SetBorderColor(tcell.ColorRed)
	} else {
		c.Layout.SetBorderColor(tcell.ColorBlack)
	}
}

// Updates value and redraw to display.
func (c *CalendarWidget) update() {
	c.monthName.Clear()
	c.monthTable.Clear()
	c.fullDate.Clear()

	c.monthName.SetText(getMonthYearFromTime(c.selectedTime))
	c.fullDate.SetText(getStringOfDate(c.currentTime)).SetTextColor(tcell.ColorBlue)

	calendarLayout := getMonthDateInfo(c.currentTime, c.selectedTime)
	for _, i := range calendarLayout {
		c.monthTable.SetCell(i.row, i.col,
			tview.NewTableCell(i.text).
				SetTextColor(i.color).
				SetAlign(tview.AlignCenter))
	}
}

// Runs update() and starts a goroutine.
// First goroutine create a channel, start loopCurrentDay as a goroutine.
// Listen to the channel for a new value.
func (c *CalendarWidget) Run() {
	c.update()
	go func() {
		dayCh := make(chan int)
		go loopCurrentDay(dayCh)
		for range dayCh {
			c.currentTime = time.Now()
			c.update()
		}
	}()

	c.Layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'l' {
			// Go a day forward.
			c.selectedTime = c.selectedTime.AddDate(0, 0, 1)
			c.update()
			return nil
		} else if event.Rune() == 'h' {
			// Go a day backward.
			c.selectedTime = c.selectedTime.AddDate(0, 0, -1)
			c.update()
			return nil
		} else if event.Rune() == 'j' {
			// Go a month forward.
			c.selectedTime = c.selectedTime.AddDate(0, 1, 0)
			c.update()
			return nil
		} else if event.Rune() == 'k' {
			// Go a month backward.
			c.selectedTime = c.selectedTime.AddDate(0, -1, 0)
			c.update()
			return nil
		}
		return event
	})
}

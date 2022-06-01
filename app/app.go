package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/alx-b/calendar/clock"
	"github.com/alx-b/calendar/logger"
)

// Initialises the app UI, add widgets, run and display them.
// Also init keybinding for the main layout.
func Run() {
	// Close the log file when closing the app.
	defer logger.Log.CloseFile()
	app := tview.NewApplication()

	//= Widgets ==============================
	clock := clock.NewClockWidget(app)
	clock.Run()

	//= Main layout ==========================
	layoutFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(clock.Layout, 0, 1, false)

	//= Keybindings ==========================
	app.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		if e.Key() == tcell.KeyEsc {
			app.Stop()
		}
		return e
	})

	if err := app.SetRoot(layoutFlex, true).EnableMouse(false).SetFocus(layoutFlex.GetItem(0)).Run(); err != nil {
		panic(err)
	}
}

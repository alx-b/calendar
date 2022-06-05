package calendar

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Go time library start sunday as 1 so had to change that.
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Gets a day name and return its integer value (Monday == 1, Sunday == 7)
func getDayOfWeekNumber(firstDayNameOfMonth string) int {
	switch firstDayNameOfMonth {
	case "Monday":
		return Monday
	case "Tuesday":
		return Tuesday
	case "Wednesday":
		return Wednesday
	case "Thursday":
		return Thursday
	case "Friday":
		return Friday
	case "Saturday":
		return Saturday
	case "Sunday":
		return Sunday
	default:
		return 0
	}
}

// Compares old day with new day,
// if not a match, send the new day number to the channel.
// !! Run loopCurrentTime as a goroutine !!
func loopCurrentDay(dayCh chan int) {
	oldDay := time.Now().Day()

	for range time.Tick(time.Second) {
		newDay := time.Now().Day()

		if oldDay != newDay {
			oldDay = newDay
			dayCh <- newDay
		}
	}
	close(dayCh)
}

// Returns a string of Month and Year -> "<MonthName> <Year>"
func getMonthYearFromTime(date time.Time) string {
	return fmt.Sprintf("%s %d", date.Month().String(), date.Year())
}

// Returns a string of current date -> "<day-of-week> <day> <month> <year>"
func getStringOfDate(date time.Time) string {
	return fmt.Sprintf("%s, %02d %s %d", date.Weekday().String(), date.Day(), date.Month().String(), date.Year())
}

type DateInfo struct {
	row   int
	col   int
	text  string
	color tcell.Color
}

// Returns all the info needed to be display by the TUI.
func getMonthDateInfo(actualTime, date time.Time) []DateInfo {
	currentYear, currentMonth, currentDay := actualTime.Date()
	selectedYear, selectedMonth, selectedDay := date.Date()
	lastDateOfMonth := date.AddDate(0, 1, -(date.Day())).Day()
	firstDayNumber := getDayOfWeekNumber(
		date.AddDate(0, 0, -(selectedDay - 1)).Weekday().String(),
	)
	rows := func() int {
		if lastDateOfMonth == 31 && firstDayNumber == Saturday ||
			lastDateOfMonth > 29 && firstDayNumber == Sunday {
			return 7
		}
		return 6
	}()

	monthLayout := []DateInfo{}
	days := [8]string{"W#", "Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}
	// First (Top) row
	for idx, value := range days {
		monthLayout = append(monthLayout, DateInfo{
			row:   0,
			col:   idx,
			text:  fmt.Sprintf(" %s ", value),
			color: tcell.ColorHotPink,
		})
	}
	// First (Left) column minus the first row.
	_, firstWeekCounter := date.AddDate(0, 0, -(selectedDay - 1)).ISOWeek()
	for i := 1; i < rows; i++ {
		monthLayout = append(monthLayout, DateInfo{
			row:   i,
			col:   0,
			text:  fmt.Sprintf(" %02d ", firstWeekCounter),
			color: tcell.ColorHotPink,
		})
		firstWeekCounter++
		if firstWeekCounter > 52 {
			firstWeekCounter = 1
		}
	}
	// Days in the month.
	dayCounter := 1
	for row := 1; row < rows; row++ {
		for col := 1; col < len(days); col++ {
			if dayCounter > lastDateOfMonth {
				break
			}
			if row == 1 && col < firstDayNumber {
				continue
			}
			color := tcell.ColorWhite
			formatText := fmt.Sprintf(" %02d ", dayCounter)
			if dayCounter == selectedDay {
				formatText = fmt.Sprintf("[%02d]", dayCounter)
			}
			if dayCounter == currentDay && currentMonth == selectedMonth && currentYear == selectedYear {
				color = tcell.ColorBlue
			}
			monthLayout = append(monthLayout, DateInfo{
				row:   row,
				col:   col,
				text:  formatText,
				color: color,
			})
			dayCounter++
		}
	}
	return monthLayout
}

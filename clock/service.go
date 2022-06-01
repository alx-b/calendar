package clock

import (
	"fmt"
	"strings"
	"time"
)

// Returns time in 00:00 format as a slice.
func getFormattedTime(currentTime time.Time) []string {
	return strings.Split(
		fmt.Sprintf("%02d:%02d", currentTime.Hour(), currentTime.Minute()),
		"",
	)
}

// Add color for specific char in a slice.
// Color should be written like -> [foreground:background] -> [red:red]
func setColorToChars(formattedTime []string, color string) []string {
	textWithColor := []string{}
	for _, symbol := range formattedTime {
		newSliceOfString := []string{}
		for _, char := range getSymbol(symbol) {
			if char == '#' {
				newSliceOfString = append(newSliceOfString, fmt.Sprintf("%s%c%s", color, char, "[white:black]"))
			} else if char == '.' {
				newSliceOfString = append(newSliceOfString, fmt.Sprintf("%s%c%s", "[black:black]", char, "[white:black]"))
			} else {
				newSliceOfString = append(newSliceOfString, fmt.Sprintf("%c", char))
			}
		}
		textWithColor = append(textWithColor, strings.TrimSpace(strings.Join(newSliceOfString, "")))
	}

	return textWithColor
}

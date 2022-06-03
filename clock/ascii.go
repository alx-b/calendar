package clock

// Returns ascii 'art' representing numbers and symbols or an empty string.
func getSymbol(number rune) string {
	switch number {
	case ':':
		return `
......
..##..
......
..##..
......
`
	case '0':
		return `
######
##..##
##..##
##..##
######
`
	case '1':
		return `
..##..
..##..
..##..
..##..
..##..
`
	case '2':
		return `
######
....##
######
##....
######
`
	case '3':
		return `
######
....##
######
....##
######
`
	case '4':
		return `
##..##
##..##
######
....##
....##
`
	case '5':
		return `
######
##....
######
....##
######
`
	case '6':
		return `
######
##....
######
##..##
######
`
	case '7':
		return `
######
....##
....##
....##
....##
`
	case '8':
		return `
######
##..##
######
##..##
######
`
	case '9':
		return `
######
##..##
######
....##
######
`
	default:
		return ""
	}
}

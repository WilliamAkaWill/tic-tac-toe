package colors

import "fmt"

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Bold   = "\033[1m"
)

// Colorize wraps text in ANSI color codes
func Colorize(color, text string) string {
	return color + text + Reset
}

// Error prints text in red
func Error(text string) string {
	return Colorize(Red, text)
}

// Success prints text in green
func Success(text string) string {
	return Colorize(Green, text)
}

// Warning prints text in yellow
func Warning(text string) string {
	return Colorize(Yellow, text)
}

// Info prints text in cyan
func Info(text string) string {
	return Colorize(Cyan, text)
}

// PrintError prints error message in red
func PrintError(format string, args ...interface{}) {
	fmt.Print(Error(fmt.Sprintf(format, args...)))
}

// PrintSuccess prints success message in green
func PrintSuccess(format string, args ...interface{}) {
	fmt.Print(Success(fmt.Sprintf(format, args...)))
}

// PrintWarning prints warning message in yellow
func PrintWarning(format string, args ...interface{}) {
	fmt.Print(Warning(fmt.Sprintf(format, args...)))
}

// PrintInfo prints info message in cyan
func PrintInfo(format string, args ...interface{}) {
	fmt.Print(Info(fmt.Sprintf(format, args...)))
}

package logger

import (
	"fmt"
	"time"
)

// timeFormatLogger const
const timeFormatLogger = "2006/01/02 15:04:05"

func init() {
	InitZap()
}

// LogWithDefer return defer func for status
func LogWithDefer(str string) (deferFunc func()) {
	fmt.Printf("%s %s ", time.Now().Format(timeFormatLogger), str)
	return func() {
		if r := recover(); r != nil {
			fmt.Printf("\x1b[31;1mERROR: %v\x1b[0m\n", r)
			panic(r)
		}
		fmt.Println("\x1b[32;1mSUCCESS\x1b[0m")
	}
}

// LogYellow log with yellow color
func LogYellow(str string) {
	fmt.Printf("%s\n", StringYellow(str))
}

// LogRed log with red color
func LogRed(str string) {
	fmt.Printf("%s\n", StringRed(str))
}

// LogGreen log with green color
func LogGreen(str string) {
	fmt.Printf("%s\n", StringGreen(str))
}

// StringYellow func
func StringYellow(str string) string {
	return fmt.Sprintf("\x1b[33;2m%s\x1b[0m", str)
}

// StringGreen func
func StringGreen(str string) string {
	return fmt.Sprintf("\x1b[32;2m%s\x1b[0m", str)
}

// StringRed func
func StringRed(str string) string {
	return fmt.Sprintf("\x1b[31;2m%s\x1b[0m", str)
}

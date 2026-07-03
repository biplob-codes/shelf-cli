// internal/ui/print.go
package ui

import "fmt"

func PrintSuccess(format string, args ...interface{}) {
	fmt.Println(Success.Render("✓ " + fmt.Sprintf(format, args...)))
}

func PrintError(format string, args ...interface{}) {
	fmt.Println(Error.Render("✗ " + fmt.Sprintf(format, args...)))
}

func PrintErrorMsg(msg string) {
	fmt.Println(Error.Render("✗ " + msg))
}

func PrintInfo(format string, args ...interface{}) {
	fmt.Println(Muted.Render(fmt.Sprintf(format, args...)))
}
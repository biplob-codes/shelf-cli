// internal/ui/time.go
package ui

import "time"

func FormatTime(t time.Time) string {
	return t.Format("Jan 2, 2006 3:04 PM")
}
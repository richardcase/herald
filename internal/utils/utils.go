package utils

import (
	"math"
	"strconv"
	"time"
)

const (
	ApproxDaysInYear  = 365
	ApproxDaysInMonth = 28
	DaysInWeek        = 7
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TimeElapsed(then time.Time) string {
	var parts []string
	var text string

	now := time.Now()
	diff := now.Sub(then)
	day := math.Floor(diff.Hours() / 24)
	year := math.Floor(day / ApproxDaysInYear)
	month := math.Floor(day / ApproxDaysInMonth)
	week := math.Floor(day / DaysInWeek)
	hour := math.Floor(math.Abs(diff.Hours()))
	minute := math.Floor(math.Abs(diff.Minutes()))
	second := math.Floor(math.Abs(diff.Seconds()))

	if year > 0 {
		parts = append(parts, strconv.Itoa(int(year))+"y")
	}

	if month > 0 {
		parts = append(parts, strconv.Itoa(int(month))+"mo")
	}

	if week > 0 {
		parts = append(parts, strconv.Itoa(int(week))+"w")
	}

	if day > 0 {
		parts = append(parts, strconv.Itoa(int(day))+"d")
	}

	if hour > 0 {
		parts = append(parts, strconv.Itoa(int(hour))+"h")
	}

	if minute > 0 {
		parts = append(parts, strconv.Itoa(int(minute))+"m")
	}

	if second > 0 {
		parts = append(parts, strconv.Itoa(int(second))+"s")
	}

	if now.After(then) {
		text = " ago"
	} else {
		text = " after"
	}

	if len(parts) == 0 {
		return "just now"
	}

	return parts[0] + text
}

func BoolPtr(b bool) *bool       { return &b }
func StringPtr(s string) *string { return &s }
func UintPtr(u uint) *uint       { return &u }
func IntPtr(u int) *int          { return &u }

func BoolDeref(b *bool, def bool) bool {
	if b != nil {
		return *b
	}

	return def
}

func StringDeref(s *string, def string) string {
	if s != nil {
		return *s
	}

	return def
}

func IntDeref(i *int, def int) int {
	if i != nil {
		return *i
	}

	return def
}

package timeCalc

import (
	"fmt"
	"strings"
	"time"
)

func plural(n int, unit string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s", n, unit)
	}
	return fmt.Sprintf("%d %ss", n, unit)
}

func TimeUntil(target time.Time) string {
	now := time.Now().UTC()
	target = target.UTC()

	if !target.After(now) {
		return "0 seconds"
	}
	total := target.Sub(now)

	years := target.Year() - now.Year()
	if now.AddDate(years, 0, 0).After(target) {
		years--
	}
	base := now.AddDate(years, 0, 0)

	months := 0
	for !base.AddDate(0, months+1, 0).After(target) {
		months++
	}
	base = base.AddDate(0, months, 0)

	// The rest is a plain duration.
	rest := target.Sub(base)
	days := int(rest / (24 * time.Hour))
	rest -= time.Duration(days) * 24 * time.Hour
	hours := int(rest / time.Hour)
	rest -= time.Duration(hours) * time.Hour
	minutes := int(rest / time.Minute)
	rest -= time.Duration(minutes) * time.Minute
	seconds := int(rest / time.Second)

	var parts []string
	add := func(n int, unit string) {
		if n > 0 {
			parts = append(parts, plural(n, unit))
		}
	}
	add(years, "год")
	add(months, "месяц")
	add(days, "день")
	add(hours, "час")
	add(minutes, "минут")
	if total < 10*time.Minute {
		add(seconds, "секунд")
	}

	if len(parts) == 0 {
		return "less than a minute"
	}
	return strings.Join(parts, ", ")
}

// using additional tools for this function in commit, next i research and overwrite this func

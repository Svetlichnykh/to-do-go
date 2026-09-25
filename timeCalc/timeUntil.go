package timeCalc

import (
	"strconv"
	"strings"
	"time"
)

func TimeUntil(target time.Time) string {
	now := time.Now().UTC()
	target = target.UTC()

	if !target.After(now) {
		return "0 секунд"
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
		if n != 0 {
			parts = append(parts, unit+strconv.Itoa(n))
		}

	}
	add(years, "лет: ")
	add(months, "месяцев: ")
	add(days, "дней: ")
	add(hours, "часов: ")
	add(minutes, "минут: ")
	if total < 10*time.Minute {
		add(seconds, "секунд:")
	}

	if len(parts) == 0 {
		return "Меньше минуты"
	}
	return strings.Join(parts, ", ")
}

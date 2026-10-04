package logs

import "strconv"

var Logs []string
var Counter int = 1

func CounterInc() {
	Counter++
}

func NewLog(counter int, text string) {
	if counter != 0 {
		Logs = append(Logs, strconv.Itoa(counter)+". "+text)
	} else {
		Logs = append(Logs, text)
	}
}

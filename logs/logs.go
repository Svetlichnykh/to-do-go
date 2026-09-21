package logs

import "strconv"

var Logs []string

func NewLog(counter int, text string) {
	Logs = append(Logs, strconv.Itoa(counter)+". "+text)
}

package main

import "fmt"

type LogLevel int

const (
	warning LogLevel = iota
	error
	fatal
	log
)

var levels = []string{"WARNING", "ERROR", "FATAL", "LOG"}

func (level LogLevel) logLevel() string {
	if level < warning || level > 3 {
		return "UNKNOWN"
	}

	return levels[level]
}

func logger(level LogLevel) string {
	return level.logLevel()
}

func main() {
	fmt.Println(logger(-1))
	fmt.Println(logger(0))
	fmt.Println(logger(1))
	fmt.Println(logger(2))
	fmt.Println(logger(3))
	fmt.Println(logger(4))
}

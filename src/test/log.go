package main
import (
    "log"
    l4g "code.google.com/p/log4go"
    "time"
)

func main() {
    log.Print("Hello1")
    log.SetPrefix("myprefix ")
    log.Print("Hello2")
    log.SetFlags(log.LstdFlags|log.Lshortfile)
    log.Print("Hello3")

	logger := l4g.NewLogger()
	logger.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())
	logger.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
}

package main

import (
	"time"
)

//import l4g "code.google.com/p/log4go"
import l4g "github.com/wangjia3015/log4go"

func main() {
	log := make(l4g.Logger)
	log.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())
	log.Debug("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	time.Sleep(10000)
}

package main

import l4g "github.com/wangjia3015/log4go"
import "time"
import "fmt"

func main() {
	for i := 0; i < 10000; i++ {
		fmt.Printf(" %d %v\n", i, time.Now())
		time.Sleep(time.Second)
		if i % 7 == 0 {
			l4g.FlushAll()
		}
		l4g.Info("test")
	}
	l4g.FlushAll()
	l4g.Close()
}

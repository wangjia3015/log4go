package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

//import l4g "code.google.com/p/log4go"

import l4g "github.com/wangjia3015/log4go"

const (
	filename = "flw.log"
)

func main() {

	l4g.Error("123456")

	//	log := make(l4g.Logger)
	//	log.AddFilter("file", l4g.FINE, l4g.NewFileLogWriter(filename, false))
	//	log.Close()

	// Get a new logger instance
	//	log := make(l4g.Logger)

	// Create a default logger that is logging messages of FINE or higher
	//	log.AddFilter("file", l4g.FINE, l4g.NewFileLogWriter(filename, false))
	//	log.Close()
	//
	//	/* Can also specify manually via the following: (these are the defaults) */
	//	flw := l4g.NewFileLogWriter(filename, false)
	//	flw.SetFormat("[%D %T] [%L] (%S) %M")
	//	flw.SetRotate(false)
	//	flw.SetRotateSize(0)
	//	flw.SetRotateLines(0)
	//	flw.SetRotateDaily(false)
	//	log.AddFilter("file", l4g.FINE, flw)

	// Log some experimental messages
	//	log.Finest("Everything is created now (notice that I will not be printing to the file)")
	//	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	//	log.Critical("Time to close out!")
	//
	//	// Close the log
	//	log.Close()

	// Print what was logged to the file (yes, I know I'm skipping error checking)
	//	fd, _ := os.Open(filename)
	//	in := bufio.NewReader(fd)
	//	fmt.Print("Messages logged to file were: (line numbers not included)\n")
	//	for lineno := 1; ; lineno++ {
	//		line, err := in.ReadString('\n')
	//		if err == io.EOF {
	//			break
	//		}
	//		fmt.Printf("%3d:\t%s", lineno, line)
	//	}
	//	fd.Close()
	//
	//	// Remove the file so it's not lying around
	//	os.Remove(filename)
	//	time.Sleep(1000)
}

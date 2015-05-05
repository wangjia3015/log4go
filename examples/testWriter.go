package main

import (
	"flag"
	"fmt"
//	"github.com/golang/glog"
	l4g "github.com/wangjia3015/log4go"
	"runtime"
	"sync"
	"time"
)

type costTime struct {
	startTime int64
}

func (c *costTime) start() {
	c.startTime = time.Now().UnixNano()
}

func (c *costTime) getInterval() int64 {
	return time.Now().UnixNano() - c.startTime
}

type testT struct {
	threadNum         int
	writeNumPerThread int
	wg                sync.WaitGroup
}

var t testT

func (t *testT) add() {
	t.wg.Add(1)
}

func (t *testT) doThreadWork(num int, f func()) {
	for i := 0; i < t.writeNumPerThread; i++ {
		f()
	}
	defer t.wg.Done()
}

func (t *testT) waitAll() {
	t.wg.Wait()
}

func (t *testT) DoTest(name string, f func()) {
	c := &costTime{}
	c.start()
	for i := 0; i < t.threadNum; i++ {
		t.add()
		go t.doThreadWork(i, f)
	}
	t.waitAll()
	fmt.Printf("missing %s cost %d \n", name, c.getInterval()/int64(time.Millisecond))
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.IntVar(&t.threadNum, "tn", 0, "thread to write")
	flag.IntVar(&t.writeNumPerThread, "ln", 0, "number to write per thread")
	flag.Parse()

	//	t.threadNum = *flag.Int("tn", 10, "thread to write")
	//	t.writeNumPerThread = *flag.Int("ln", 10, "number to write per thread")

	//t.threadNum = 1
	//t.writeNumPerThread = 100000

/*
	fmt.Printf("thread Num %d write num %d per thread\n", t.threadNum, t.writeNumPerThread)
	t.DoTest("say hello", func() {
		//	fmt.Println("hello word")
		glog.Error("Just for test")
	})
*/

/*
	for i := 0; i < t.writeNumPerThread ; i++ {
		l4g.Error("just for test")
	}
*/
	t.DoTest("l4g", func() {
		//	fmt.Println("hello word")
		l4g.Error("Just for test")
	})

	l4g.Close()

	fmt.Println("work done")
}

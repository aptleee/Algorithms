package sync

import (
	"fmt"
	"sync"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var buf = make([]int, 40)
func main() {
	for i := 0; i < 40; i++ {
		go func(x int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			for buf[x] == 0 {
				cond.Wait()
			}
			fmt.Println(x, buf[x])
			time.Sleep(time.Second * 1)

		}(i)
	}
	for i := 0; i < 40; i++ {
		time.Sleep(time.Second * 1)
		cond.L.Lock()
		buf[i] = 2000+i
		//fmt.Println("buf[", i, "]", buf[i])
		cond.Signal()
		cond.L.Unlock()
	}

}

//type Cond struct {
//	noCopy noCopy
//
//	// L is held while observing or changing the condition
//	L Locker
//
//	notify  notifyList
//	checker copyChecker
//}



/*func (c *Cond) Wait() {
	c.checker.check()
	t := runtime_notifyListAdd(&c.notify)
	c.L.Unlock()
	runtime_notifyListWait(&c.notify, t)
	c.L.Lock()
}*/

//func (c *Cond) Signal() {
//	c.checker.check()
//	runtime_notifyListNotifyOne(&c.notify)
//}

// Broadcast wakes all goroutines waiting on c.
//
// It is allowed but not required for the caller to hold c.L
// during the call.


//func (c *Cond) Broadcast() {
//	c.checker.check()
//	runtime_notifyListNotifyAll(&c.notify)
//}


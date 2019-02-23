package sync

import (
"fmt"
"sync"
"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func main() {
	for i := 0; i < 40; i++ {
		go func(x int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait() // release the Lock and acquire lock when be signaled
			fmt.Println(x)
			time.Sleep(time.Second * 1)

		}(i)
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal()
	time.Sleep(time.Second * 1)
	cond.Signal()
	time.Sleep(time.Second * 3)
	cond.Broadcast()
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 60)
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


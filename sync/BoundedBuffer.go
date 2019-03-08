package sync

type semaphore chan int

func New(n int) semaphore {
	sem := make(chan int, n)
	return sem
}

func (s semaphore) V(n int) {
	e := 1
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) P(n int) {
	for i := 0; i < n; i++ {
		<- s
	}
}

func (s semaphore) Wait() {
	s.V(1)
}

func (s semaphore)  Signal(){
	s.P(1)
}

type BoundedBuf struct {
	empty semaphore
	full semaphore
	lock semaphore
	buf []int
}


func Init(n int) BoundedBuf {
	empty := New(n)
	full := New(0)
	lock := New(1)
	return BoundedBuf{
		empty:empty,
		full:full,
		lock:lock,
		buf:make([]int, n),
	}
}

func (this *BoundedBuf) produce(x int){
	this.empty.Wait()
	this.lock.Wait()
	this.buf = append(this.buf, x)

	this.lock.Signal()
	this.full.Signal()
}



func (this *BoundedBuf) consumer() {
	this.full.Wait()
	this.lock.Wait()
	this.buf = this.buf[1:]
	this.lock.Signal()
	this.empty.Signal()
}



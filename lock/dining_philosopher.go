package lock


var chopsticks = [5]semaphore{}

func I() {
	for _, e := range chopsticks {
		e = new(1)
	}
}

// if 5 philosophers pick up chopstick[i] at the same time, there is deadlock
func loop(i int) {

	for {
		chopsticks[i].Wait()
		chopsticks[(i+1)%5].Wait()

		// eating


		chopsticks[i].Signal()
		chopsticks[(i+1)%5].Signal()

		// thinking
	}

}
// remedies
// 1. only allow four philosophers to pick up chopsticks at the same time
// 2. philosopher will only pick chopsticks when there are two available
// 3. odd numbered philosopher first pick up i and then pick up i+1;
//    even numbered philosopher first pick up i+1 and then i

func main()  {
	loop(0)
}
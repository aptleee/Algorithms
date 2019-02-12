package lock



var chopsticks = [5]semaphore{}

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
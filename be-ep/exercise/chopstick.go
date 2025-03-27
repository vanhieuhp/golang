package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5
const maxConcurrentEaters = 2
const eatingRounds = 3

type Chopstick struct{ sync.Mutex }
type Philosopher struct {
	id                            int
	leftChopstick, rightChopstick *Chopstick
	eatCount                      int
}

func (p *Philosopher) eat(wg *sync.WaitGroup, host chan struct{}) {
	defer wg.Done()

	for p.eatCount < eatingRounds {
		//Request permission from the host to eat
		host <- struct{}{} // block if maxConcurrentEaters is reached

		// Pick up chopsticks
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		// Start eating
		fmt.Println("Starting to eat", p.id)
		time.Sleep(time.Millisecond * 500) /// simulate eating time
		fmt.Println("Finish eating", p.id)

		// Release chopsticks
		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()

		// Release the host slot
		<-host

		// Increment eat count
		p.eatCount++
	}
}

func main() {
	var wg sync.WaitGroup

	// Initialize chopsticks
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = &Chopstick{}
	}

	// Initialize philosophers
	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
			eatCount:       0,
		}
	}

	// Host channel controlling concurrent eaters
	host := make(chan struct{}, maxConcurrentEaters)

	// Start philosopher goroutines
	wg.Add(numPhilosophers)
	for _, p := range philosophers {
		go p.eat(&wg, host)
	}

	// Wait for all philosophers to finish
	wg.Wait()
}

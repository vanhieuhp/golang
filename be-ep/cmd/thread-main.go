package main

import (
	"fmt"
	"sync"
)

func addOne(count *int, intChannel chan int) {
	(*count)++
	intChannel <- *count
}

// synchronization
// semaphore

func main() {
	value := 0

	intChannel := make(chan int)

	wg := new(sync.WaitGroup)
	mutex := &sync.Mutex{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mutex.Lock()
			defer mutex.Unlock()
			addOne(&value, intChannel)
		}()
	}

	//wg.Wait()
	//fmt.Print(value)
	for {
		outputValue, oke := <-intChannel
		if !oke {
			break
		}
		fmt.Println(outputValue)
	}
	//select {}
}

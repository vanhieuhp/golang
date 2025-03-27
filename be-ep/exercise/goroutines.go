/**
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.
-------------------------------------------------------------------------------

The program creates 2 go routines which should increment a global counter
variable in a loop

The end result should be the number of loops x 2 (per go routine) however
this is not often the case when the code is run.
The end counter value is less than this value.

This occurs because the counter is not kept in sync between the go routines.

During execution of a Go routine, the counter variable is read and after
incremented there can be a context switch so the increment operation of the
first thread is lost. If this happens many times, many increment operations
can be lost resulting in a value much lower than expected.

It does not matter if a global variable is affected by context switching.
What matters is when the variable is read, and when it is written, by each thread,
because the value is "taken off the shelf" and modified, then "put back".
Once the first thread has read the global variable and is working with the
value "somewhere in space", the second thread must not read the global
variable until the first thread has written the updated value.

*/

package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 2; i++ {

		go func() {
			count()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter value is: ", counter)
}

func count() {
	for i := 1; i <= 100000; i++ {
		counter = counter + 1
	}
}

// Implement a thread-safe counter using Mutex.
package main

import (
	"fmt"
	"sync"
)

// WRONG CODE
// type Counter struct {
// 	value int
// }

// func (c *Counter) Increment() {
// 	c.value++ //Race Condition
// }

// func (c *Counter) Value() int {
// 	return c.value /// not thread safe
// }

// func main() {
// 	var wg sync.WaitGroup
// 	counter := Counter{} //single counter instance
// 	for i := 0; i <= 10000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done() //signal Completion
// 			counter.Increment()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("Final counter Value-----", counter.Value())
// }

// CORRECT SOLUTION

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() //signal Completion
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Final counter Value-----", counter.Value())
}

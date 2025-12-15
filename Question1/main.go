// Write a program that prints numbers from 1 to 10 using two goroutines in sequence
//  (odd-even problem).

package main

import "fmt"

func main() {
	oddCh := make(chan bool)
	evenCh := make(chan bool)
	doneCh := make(chan bool)

	// Odd Goroutine
	go func() {
		for i := 1; i <= 9; i += 2 {
			<-oddCh //Wait for signal
			fmt.Println(i)
			evenCh <- true
		}
	}()

	// Even Goroutine
	go func() {
		for i := 2; i <= 10; i += 2 {
			<-evenCh //Wait for signal
			fmt.Println(i)
			if i == 10 {
				doneCh <- true
				return
			}
			oddCh <- true
		}
	}()
	oddCh <- true

	<-doneCh //Wait for completion of the program
}

// 1,

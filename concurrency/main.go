package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Allocate 1 logical processor
	runtime.GOMAXPROCS(1);

	var wg sync.WaitGroup

	wg.Add(2);

	fmt.Println("Start a Goroutines")

	// Declare a anonymous function

	go func() {
		defer wg.Done();

		// display alphabet 3 times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a' + 26; char++ {
				fmt.Printf("%c", char)
			}

		}
	}()

	go func() {
		defer wg.Done();

		// display alphabet 3 times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A' + 26; char++ {
				fmt.Printf("%c", char)
			}

		}
	}()

	// Wait for go routines to finish
	fmt.Println("Waiting to finish")
	wg.Wait();

	fmt.Println("\n Terminating program")
}

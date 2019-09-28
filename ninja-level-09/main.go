package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)

	fmt.Println("Start")
	go routine01()
	go routine02()
	wg.Wait()
	fmt.Println("End")

	fmt.Println(runtime.NumGoroutine())

	fmt.Println()

	// Race condition and Mutex

	grs := 100
	counter := 0

	wg.Add(grs)

	var mu sync.Mutex

	for i := 0; i < grs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func routine01() {
	fmt.Println("In routine 01!")
	wg.Done()
}

func routine02() {
	fmt.Println("In routine 02!")
	wg.Done()
}

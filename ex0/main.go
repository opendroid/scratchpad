package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ex1: A go-routine example that increments a counter (safely) in nFuncs go-routines.
//
//	The main go-routine waits for all the go-routines to finish.
var nFuncs int

// Counter is a simple counter that can be incremented safely.
type Counter struct {
	mu    sync.Mutex
	count int
}

// main is the entry point for the program.
func main() {
	wg := sync.WaitGroup{}
	counter := Counter{count: 0}
	nFuncs = 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	wg.Add(nFuncs)
	for i := 0; i < nFuncs; i++ {
		go func(ctx context.Context, counter *Counter, i int) {
			loops := 0
			defer func() { // Print counter on exit
				counter.mu.Lock()
				cnt := counter.count
				counter.mu.Unlock()
				fmt.Printf("Routine:[%d]: Loops: %d, Last count by this routine:%d\n", i, loops, cnt)
				wg.Done()
			}()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					counter.mu.Lock()
					counter.count++
					counter.mu.Unlock()
					loops++
				}
			}

		}(ctx, &counter, i)
	}
	wg.Wait()
	fmt.Printf("Count when all done: %d\n", counter.count)

}

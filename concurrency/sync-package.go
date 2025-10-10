package main

import(
	"fmt"
	"sync"
	"time"
)

func SyncPackageDemo() {

	fmt.Println("\nSync Package Demo:")

	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {

		wg.Add(1) //adds work

		go func (id int)  {
			defer wg.Done() //calls until work done
			fmt.Println("Worker", id, "starting...")
			time.Sleep(300 * time.Millisecond)
			fmt.Println("Worker", id, "done.")
		}(i)
	}

	wg.Wait() //waits until all goroutines end
	fmt.Println("All workers done.")
}
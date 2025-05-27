package main

import (
	"fmt"
	"sync"
)

var trackWorker []string

func worker(i int, mut *sync.Mutex, wg *sync.WaitGroup){
	defer wg.Done()

	mut.Lock()
	// temp := fmt.Sprint("Worker")
	trackWorker = append(trackWorker, fmt.Sprintf("Worker %d", i))
	mut.Unlock()

	fmt.Printf("Worker %d started \n", i)
	fmt.Println("Some task is happening")
	fmt.Printf("Worker %d end \n", i)

}

func main(){
	fmt.Println("Explore goroutine")

	var wg sync.WaitGroup
	var mut sync.Mutex


	// Start 3 worked goroutines
	for i := 1; i <= 3; i++ {
		// Incrementing wait group
		wg.Add(1)
		go worker(i, &mut, &wg)
	}

	// Wait for Worker group
	wg.Wait()
	fmt.Println("Worker task ended ", trackWorker)

}
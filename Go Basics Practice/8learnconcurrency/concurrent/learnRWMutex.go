package concurrent

import (
	"fmt"
	"sync"
)

var trackWorker []string
var trackId []int

func worker(i int, mut *sync.Mutex, remut *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()

	mut.Lock()
	// temp := fmt.Sprint("Worker")
	trackWorker = append(trackWorker, fmt.Sprintf("Worker %d", i))
	mut.Unlock()

	remut.Lock()
	trackId = append(trackId, i)
	remut.Unlock()

	fmt.Printf("Worker %d started \n", i)
	fmt.Println("Some task is happening")
	fmt.Printf("Worker %d end \n", i)

}

func LearnMutex() {
	fmt.Println("Explore goroutine")

	var wg sync.WaitGroup
	var mut sync.Mutex
	var rwmut = sync.RWMutex{}

	// Start 3 worked goroutines
	for i := 1; i <= 3; i++ {
		// Incrementing wait group
		wg.Add(1)
		go worker(i, &mut, &rwmut, &wg)
	}

	// Wait for Worker group
	wg.Wait()
	fmt.Println("Worker task ended ", trackWorker)
}

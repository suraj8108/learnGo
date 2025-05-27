package concurrent

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doSomeWork() int {
	time.Sleep(1000 * time.Millisecond)
	return rand.Intn(100)
}
func LearnChannel() {
	fmt.Println("Channel in GoRoutine")

	myCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				result := doSomeWork()
				myCh <- result
				wg.Done()
			}()

		}
		wg.Wait()
		close(myCh) // Do not forget to close the channel
	}()

	for element := range myCh {
		fmt.Printf("Element is %d \n", element)
	}
}

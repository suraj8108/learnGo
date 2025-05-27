package concurrent

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Select in GoRoutine")

	firstCh := make(chan int)
	secondCh := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		time.Sleep(1000 * time.Millisecond)
		firstCh <- 5
		fmt.Println("First chan updated")
		wg.Done()
	}()

	go func() {
		time.Sleep(2000 * time.Millisecond)
		secondCh <- 12
		fmt.Println("Second chan updated")
		wg.Done()
	}()

	select {
	case firstVal := <-firstCh:
		fmt.Println("My first val", firstVal)
	case secondVal := <-secondCh:
		fmt.Println("My Second Val", secondVal)
	}
	wg.Wait()
}

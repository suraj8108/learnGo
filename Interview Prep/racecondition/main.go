package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Race Condition - Learnong");

	var score = []int {0}

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3)	

	go func(mut *sync.Mutex, wg *sync.WaitGroup) {
		
		defer wg.Done()
		fmt.Println("One R")
		mut.Lock()
		score := append(score, 1)
		mut.Unlock()

		fmt.Println(score)
	}(mut, wg)

	go func(mut *sync.Mutex, wg *sync.WaitGroup) {

		defer wg.Done()
		fmt.Println("Two R")
		mut.Lock()
		score := append(score, 2)
		mut.Unlock()
		fmt.Println(score)
	}(mut, wg)

	go func(mut *sync.Mutex, wg *sync.WaitGroup) {
		
		defer wg.Done()
		fmt.Println("Three R")
		mut.Lock()
		score := append(score, 3)
		mut.Unlock()
		fmt.Println(score)

	}(mut, wg)

	wg.Wait()
	fmt.Println(score)
}
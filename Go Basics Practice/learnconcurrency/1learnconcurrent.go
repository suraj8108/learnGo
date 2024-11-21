package learnconcurrency

import (
	"fmt"
	"time"
)

func greet(phrase string, chanDone chan bool) {
	fmt.Println("Hello!", phrase)
	chanDone <- true
}

func slowGreet(phrase string, chanDone chan bool){
	
	defer close(chanDone)

	time.Sleep(3*time.Second)
	fmt.Println("Hello", phrase)
	chanDone <- true
	
}
func Learnconcurrency() {
	done := make([]chan bool, 4)

	done[0] = make(chan bool)
	go greet("Nice to meet you",done[0])
	done[1] = make(chan bool)
	go greet("How are you",done[1])
	done[2] = make(chan bool)
	go slowGreet("How ... are ... you?", done[2])
	done[3] = make(chan bool)
	go greet("I hope you are doing good", done[3])

	for _, todo := range done {
		<- todo
	}
}
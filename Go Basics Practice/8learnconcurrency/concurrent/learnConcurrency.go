package concurrent

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func task(taskid int, sc *sync.WaitGroup) {
	result := 0
	defer sc.Done()
	for i := 0; i < 10; i++ {
		result += i
	}
	textToPush := fmt.Sprintf("Task %v has been completed \n", taskid)

	file, err := os.OpenFile("./concurrent/output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Issue in task %v", taskid)
	}
	defer file.Close()

	if _, err := file.WriteString(textToPush); err != nil {
		fmt.Println(err)
	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Executed task", taskid)

}

func InvokeCalls() {
	sc := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		sc.Add(1)
		go task(i, &sc)
	}
	sc.Wait()
}

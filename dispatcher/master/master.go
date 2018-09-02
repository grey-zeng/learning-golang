package main

import (
	"fmt"
	"github.com/learning-golang/dispatcher/task"
)

func fixedTrigger(intervalMillSecond int) {
	ch := make(chan task.Task)
	defer close(ch)
	go task.FixedTrigger(ch, 1000)
	for taskItem := range ch {
		fmt.Println(taskItem.Content)
	}
}

func main() {
	fmt.Println("I am a Master")
	go fixedTrigger(300)
	go fixedTrigger(700)
	go fixedTrigger(1000)
	fixedTrigger(2000)
}

package task

import (
	"fmt"
	"time"
)

type Task struct {
	id                                      int
	createTime, recTime, beginTime, endTime time.Time
	Content                                 string
}

var (
	globalTriggerNum, globalTaskNum int = 0, 0
)

func FixedTrigger(ch chan Task, intervalMillSecond int) {
	triggerNum := globalTriggerNum
	globalTriggerNum++
	tick := time.Tick(time.Duration(intervalMillSecond) * time.Millisecond)
	boom := time.After(60 * time.Second)
	for {
		select {
		case <-tick:
			taskId := globalTaskNum
			globalTaskNum++
			newTask := Task{
				id:         taskId,
				createTime: time.Now(),
				Content:    fmt.Sprintf("trigger:%d-task:%d", triggerNum, taskId),
			}
			ch <- newTask
		case <-boom:
			fmt.Printf("triger:%d dump\n", triggerNum)
			close(ch)
			return
		default:
		}
	}
}

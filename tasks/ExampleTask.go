package tasks

import (
	"fmt"
	"time"
)

func NeedVeryLongTimeToFinish(params ...interface{}) {
	time.Sleep(10 * time.Second)
	fmt.Println("task succeed")
}

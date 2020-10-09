package src

import (
	"sync"
)

type Task func(params ...interface{})
type CallBack func(params ...interface{}) interface{}
type callback func() interface{}

//callback wrapper
func Back(back CallBack, params ...interface{}) func() interface{} {
	return func() interface{} {
		return back(params...)
	}
}

func (f Task) callback(params ...interface{}) interface{} {
	return nil
}

var tasks chan *TaskExecutor //task list
var once sync.Once

func init() {
	taskChan := GetTaskChan()
	go func() {
		for t := range taskChan { //range chan won't be closed if we dont close it manually
			go func() {
				if t.back != nil {
					defer t.back()
				}
				t.Exec()
			}()
		}
	}()
}

//get tasks chan
func GetTaskChan() chan *TaskExecutor {
	once.Do(func() {
		tasks = make(chan *TaskExecutor)
	})
	return tasks
}

type TaskExecutor struct {
	f    Task
	p    []interface{}
	back callback
}

func (self *TaskExecutor) Exec() {
	self.f(self.p...)
}
func NewTaskExecutor(f Task, back callback, p []interface{}) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, back: back}
}
func AddTask(f Task, back callback, params ...interface{}) {
	GetTaskChan() <- NewTaskExecutor(f, back, params)
}

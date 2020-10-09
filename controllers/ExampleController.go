package controllers

import (
	"GinEasy/dependencies"
	"GinEasy/src"
	"GinEasy/tasks"
	"fmt"
	"github.com/gin-gonic/gin"
)

//controller struct
type ExampleController struct {
	*dependencies.XOrmAdapter
}

//new controller struct
func NewExampleController() *ExampleController {
	return &ExampleController{}
}

//you can set anytype response in controller handler func
func (self *ExampleController) ExampleStringTest(ctx *gin.Context) string {
	return "example test"
}
func (self *ExampleController) ExampleJsonName(ctx *gin.Context) interface{} {
	return struct {
		ExampleName string `json:"example_name"`
	}{"example name"}
}

func (self *ExampleController) Build(ge *src.GE) { //bind gin.handler and route
	ge.Handle("GET", "/exam_name", self.ExampleJsonName)
	ge.Handle("GET", "/exam_test", self.ExampleStringTest)
}

type Example2Controller struct {
	*dependencies.XOrmAdapter
}

//new controller struct
func NewExample2Controller() *Example2Controller {
	return &Example2Controller{}
}

//you can set anytype response in controller handler func
func (self *Example2Controller) Example2StringTest(ctx *gin.Context) string {
	return "example2 test"
}
func (self *Example2Controller) Example2JsonName(ctx *gin.Context) interface{} {
	return struct{ name string }{"example2 name"}
}

func (self *Example2Controller) Build(ge *src.GE) { //bind gin.handler and route
	ge.Handle("GET", "/exam2_name", self.Example2JsonName)
	ge.Handle("GET", "/exam2_test", self.Example2StringTest)
}

type ExampleUseTaskController struct{}

func NewExampleTaskController() *ExampleUseTaskController {
	return &ExampleUseTaskController{}
}

func (self *ExampleUseTaskController) ExampleUseTaskView(ctx *gin.Context) string {
	src.AddTask(tasks.NeedVeryLongTimeToFinish, src.Back(func(params ...interface{}) interface{} {
		fmt.Println(len(params))
		return nil
	}, 1, 2, 3, 4, 5), 'a', 'b', 'c') //mock time consuming
	return "ok"
}

func (self *ExampleUseTaskController) Build(ge *src.GE) {
	ge.Handle("GET", "/ExampleUseTask", self.ExampleUseTaskView)
}

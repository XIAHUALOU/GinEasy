package middleWares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ExampleMid struct{}

type Example2Mid struct{}

func NewExampleMid() *ExampleMid {
	return &ExampleMid{}
}
func NewExample2Mid() *Example2Mid {
	return &Example2Mid{}
}

func (this *ExampleMid) BeforeRequest(ctx *gin.Context) error {
	fmt.Println("this is Example middleware")
	return nil
}

func (this *Example2Mid) BeforeRequest(ctx *gin.Context) error {
	fmt.Println("this is Example2 middleware")
	return nil
}

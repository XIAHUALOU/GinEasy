package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

var ResponderList []Responder

func init() {
	ResponderList = append(ResponderList, new(StringResponder), new(JsonResponder))
}

//Responder interface
type Responder interface {
	RespondTo() gin.HandlerFunc
}

//convert our anytype returned func to gin.Handler
func Convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	for _, r := range ResponderList {
		r_ref := reflect.ValueOf(r).Elem()
		if h_ref.Type().ConvertibleTo(r_ref.Type()) {
			r_ref.Set(h_ref)
			return r_ref.Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type StringResponder func(*gin.Context) string

func (self StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, self(context))
	}
}

type JsonResponder func(ctx *gin.Context) interface{}

func (self JsonResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println(self(context))
		context.JSON(200, self(context))
	}
}

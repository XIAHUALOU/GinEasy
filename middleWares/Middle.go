package middleWares

import "github.com/gin-gonic/gin"

//middleware interface
type Mid interface {
	BeforeRequest(*gin.Context) error
}

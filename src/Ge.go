package src

import (
	"GinEasy/middleWares"
	"github.com/gin-gonic/gin"
)

//The core object of the whole scaffold
type GE struct { // The nae is the abbreviation of Gin Easy, which means it is more convenient to use gin
	*gin.Engine
	group *gin.RouterGroup
	*Dependency
}

//init GE and set global error handler middleware
func StartGE() *GE {
	g := &GE{Engine: gin.New(), Dependency: NewDependency()}
	g.Use(ErrorHandler())
	return g
}

//start server method
func (self *GE) Launch() {
	self.Run(SERVER_ADDRESS)
}

//This method is the core of the scaffold. It is mainly for the convenience of returning any type of business results,
// and the binding of groups is done here
func (self *GE) Handle(httpMethod, relativePath string, handler interface{}) *GE {
	if h := Convert(handler); h != nil {
		self.group.Handle(httpMethod, relativePath, h)
	}
	return self
}

//GE's middleware
func (self *GE) AddMid(middleWares ...middleWares.Mid) *GE {
	for _, mid := range middleWares {
		mid := mid
		self.Use(func(context *gin.Context) {
			err := mid.BeforeRequest(context)
			if err != nil {
				context.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			} else {
				context.Next()
			}
		})
	}
	return self
}

//This method is mainly used to add routes and inject dependencies into controllers
func (self *GE) AddController(group string, controllers ...Controller) *GE {
	self.group = self.Group(group)
	for _, controller := range controllers {
		controller.Build(self)
		self.BindDep(controller)
	}
	return self
}

func (self *GE) PrepareDeps(deps ...interface{}) *GE {
	self.SetDep(deps...)
	return self
}

package main

import (
	"github.com/XIAHUALOU/GinEasy/controllers"
	"github.com/XIAHUALOU/GinEasy/dependencies"
	"github.com/XIAHUALOU/GinEasy/middleWares"
	. "github.com/XIAHUALOU/GinEasy/src"
)

func main() {
	StartGE().
		PrepareDeps(dependencies.NewXOrmAdapter()).
		AddMid(middleWares.NewExampleMid(), middleWares.NewExample2Mid()).
		AddController("v1",
			controllers.NewExampleController(),
			controllers.NewExample2Controller()).
		AddController("v1", controllers.NewExampleTaskController()).
		Launch()
}

package main

import (
	"GinEasy/controllers"
	"GinEasy/dependencies"
	"GinEasy/middleWares"
	. "GinEasy/src"
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

package main

import (
	"github.com/XIAHUALOU/variou-gin/controllers"
	"github.com/XIAHUALOU/variou-gin/dependencies"
	"github.com/XIAHUALOU/variou-gin/middleWares"
	. "github.com/XIAHUALOU/variou-gin/src"
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

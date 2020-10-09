package src

type Controller interface {
	Build(ge *GE) //参数和方法名必须一致
}

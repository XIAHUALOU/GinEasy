package src

import (
	"reflect"
)

type Dependency struct {
	members []interface{}
}

//dependencies factory
func NewDependency() *Dependency {
	dep := &Dependency{members: make([]interface{}, 0)}
	return dep
}

//add dependencies
func (self *Dependency) SetDep(deps ...interface{}) {
	self.members = append(self.members, deps)
}

//get dependencies by reflect type
func (self *Dependency) GetDep(t reflect.Type) interface{} {
	for _, p := range self.members {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}

//inject dependencies to each controller
func (self *Dependency) BindDep(controller Controller) {
	ctrl := reflect.ValueOf(controller).Elem() //controllers are all pointer
	for i := 0; i < ctrl.NumField(); i++ {
		f := ctrl.Field(i)
		if f.Kind() != reflect.Ptr || !f.IsNil() {
			continue
		}
		if p := self.GetDep(f.Type()); p != nil && f.CanInterface() {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

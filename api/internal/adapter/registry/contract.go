package registry

type IRegistry interface {
	Init()
	Parse(interface{})
}

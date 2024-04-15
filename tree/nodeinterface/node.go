package nodeinterface

type Node interface {
	GetChildren() []Node
	GetValue() any
}

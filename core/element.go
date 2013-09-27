package core

type Element interface {
	Parent() Element
	SetParent(Element)
	Script() *Script
}

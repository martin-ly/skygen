package core

type elementImpl struct {
	parent Element
}

func (e *elementImpl) Parent() Element {
	return e.parent
}

func (e *elementImpl) SetParent(parent Element) {
	e.parent = parent
}

func (e *elementImpl) Script() *Script {
	if e.parent == nil {
		return nil
	}
	return e.parent.Script()
}

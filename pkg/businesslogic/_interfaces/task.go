package _interfaces

type Unloadable interface {
	Unload(unload float32)
}

type Weightable interface {
	Weight() float32
}

type IHasMoved interface {
	HasMoved() bool
}

type ITask interface {
	Unloadable
	Weightable
	IHasMoved
	Name() string
}

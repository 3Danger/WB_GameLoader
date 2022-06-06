package task

import (
	"math/rand"
	"sync"
)

type Task struct {
	mut    sync.RWMutex
	name   string
	weight float32
}

func (t *Task) Name() string {
	t.mut.RLock()
	defer t.mut.RUnlock()
	return t.name
}
func (t *Task) Weight() float32 {
	t.mut.RLock()
	defer t.mut.RUnlock()
	return t.weight
}
func (t *Task) HasMoved() bool {
	t.mut.RLock()
	defer t.mut.RUnlock()
	return t.weight <= 0.
}

func (t *Task) Unload(unload float32) {
	t.mut.Lock()
	t.weight -= unload
	t.mut.Unlock()
}

func NewTaskRand(name string) *Task {
	return &Task{name: name, weight: rand.Float32()*70 + 10}
}

func NewTask(name string, weight float32) *Task {
	return &Task{name: name, weight: weight}
}

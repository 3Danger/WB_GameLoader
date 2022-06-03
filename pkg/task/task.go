package task

import (
	"math/rand"
)

type Task struct {
	name   string
	weight float32
}

func (t Task) Name() string    { return t.name }
func (t Task) Weight() float32 { return t.weight }
func (t Task) HasMoved() bool  { return t.weight <= 0. }

func (t *Task) Unload(unload float32) {
	t.weight -= unload
}

func NewTaskRand(name string) *Task {
	return &Task{name: name, weight: rand.Float32()*90 + 10}
}

func NewTask(name string, weight float32) *Task {
	return &Task{name: name, weight: weight}
}

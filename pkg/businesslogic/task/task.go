package task

import (
	"math/rand"
)

type Task struct {
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
}

func (t *Task) GetName() string    { return t.Name }
func (t *Task) GetWeight() float32 { return t.Weight }

func (t *Task) HasMoved() bool {
	return t.Weight <= 0.
}

func (t *Task) Unload(unload float32) {
	t.Weight -= unload
}

func NewTaskRand(name string) *Task {
	return &Task{Name: name, Weight: rand.Float32()*70 + 10}
}

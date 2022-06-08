package customer

import (
	"GameLoaders/pkg/businesslogic/task"
	"sort"
	"sync"
)

type chainOfTaskBuilder struct {
	sync.RWMutex
	tasks []*task.Task
}

func (c *chainOfTaskBuilder) Add(task ...*task.Task) *chainOfTaskBuilder {
	if len(task) < 1 {
		return nil
	}
	c.tasks = append(c.tasks, task...)
	return c
}

func (c *chainOfTaskBuilder) Build() *chaiOfTask {
	var head, last *chaiOfTask
	if c == nil || len(c.tasks) == 0 {
		return nil
	}
	sort.Slice(c.tasks, func(i, j int) bool {
		return c.tasks[i].Weight > c.tasks[j].Weight
	})

	head = &chaiOfTask{c.tasks[0], nil}
	last = head
	for _, v := range c.tasks[1:] {
		last.next = &chaiOfTask{v, nil}
		last = last.next
	}
	return head
}

type chaiOfTask struct {
	*task.Task
	next *chaiOfTask
}

func (c *chaiOfTask) Unload(unload float32) {
	c.Task.Unload(unload)
	if c.HasMoved() {
		if c.next != nil {
			*c = *c.next
		}
	}
}

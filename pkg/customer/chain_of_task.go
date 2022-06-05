package customer

import (
	. "GameLoaders/pkg/_interfaces"
	"sort"
)

type ChainOfTaskBuilder struct {
	tasks []ITask
}

func (c *ChainOfTaskBuilder) Add(task ...ITask) *ChainOfTaskBuilder {
	c.tasks = append(c.tasks, task...)
	return c
}

func (c *ChainOfTaskBuilder) Build() *ChaiOfTask {
	var head, last *ChaiOfTask
	if len(c.tasks) == 0 {
		return nil
	}
	sort.Slice(c.tasks, func(i, j int) bool {
		return c.tasks[i].Weight() > c.tasks[j].Weight()
	})

	head = &ChaiOfTask{c.tasks[0], nil}
	last = head
	for _, v := range c.tasks[1:] {
		last.next = &ChaiOfTask{v, nil}
		last = last.next
	}
	return head
}

type ChaiOfTask struct {
	ITask
	next *ChaiOfTask
}

func (c *ChaiOfTask) Unload(unload float32) {
	c.ITask.Unload(unload)
	if c.HasMoved() {
		if c.next != nil {
			*c = *c.next
		}
	}
}

package customer

import (
	. "GameLoaders/pkg/_interfaces"
	"sort"
)

type chainOfTaskBuilder struct {
	tasks []ITask
}

func (c *chainOfTaskBuilder) Add(task ...ITask) *chainOfTaskBuilder {
	c.tasks = append(c.tasks, task...)
	return c
}

func (c *chainOfTaskBuilder) Build() *chaiOfTask {
	var head, last *chaiOfTask
	if len(c.tasks) == 0 {
		return nil
	}
	sort.Slice(c.tasks, func(i, j int) bool {
		return c.tasks[i].Weight() > c.tasks[j].Weight()
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
	ITask
	next *chaiOfTask
}

func (c *chaiOfTask) Unload(unload float32) {
	c.ITask.Unload(unload)
	if c.HasMoved() {
		if c.next != nil {
			*c = *c.next
		}
	}
}

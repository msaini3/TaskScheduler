package main

import (
	"fmt"
	"sync"
)

// Task represents a task to be executed
type Task interface {
	Execute()
}

// SimpleTask represents a simple task
type SimpleTask struct {
	Name string
}

// Execute prints the name of the task
func (t SimpleTask) Execute() {
	fmt.Printf("Executing task: %s\n", t.Name)
}

// Scheduler represents a task scheduler
type Scheduler struct {
	Tasks []Task
}

// AddTask adds a task to the scheduler
func (s *Scheduler) AddTask(task Task) {
	s.Tasks = append(s.Tasks, task)
}

// Run executes all tasks concurrently
func (s *Scheduler) Run() {
	var wg sync.WaitGroup
	wg.Add(len(s.Tasks))

	for _, task := range s.Tasks {
		go func(t Task) {
			defer wg.Done()
			t.Execute()
		}(task)
	}

	wg.Wait()
}

func main() {
	// Create a scheduler
	scheduler := Scheduler{}

	// Add tasks to the scheduler
	scheduler.AddTask(SimpleTask{Name: "Task 1"})
	scheduler.AddTask(SimpleTask{Name: "Task 2"})
	scheduler.AddTask(SimpleTask{Name: "Task 3"})
	scheduler.AddTask(SimpleTask{Name: "Task 4"})
	scheduler.AddTask(SimpleTask{Name: "Task 5"})

	// Run the scheduler
	scheduler.Run()
}

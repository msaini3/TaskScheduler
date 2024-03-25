# TaskScheduler
In this application, we have two interfaces: 'Task' and 'Scheduler'. The 'Task' interface defines the 'Execute()' method, which represents a task to be executed. We have a 'SimpleTask' struct implementing the 'Task' interface.

The 'Scheduler' struct contains a slice of tasks and methods to add tasks ('AddTask') and execute them concurrently ('Run') using goroutines and a wait group.

In the 'main' function, we create a scheduler, add some tasks to it, and then execute them concurrently by calling the 'Run' method.

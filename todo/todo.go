package todo

type TodoList interface {
	Run()
	AddTask(name string)
	ListTasks()
	DeleteTask(index int)
	MarkAsDone(index int)
}

type Task struct {
	Name   string
	Status bool
}

func Run(t TodoList) {
	t.Run()
}

package main

type Todos struct {
	completed bool
	item      string
}

var list = []Todos{}

func AddTodo(todo string) {
	list = append(list, Todos{completed: false, item: todo})
}

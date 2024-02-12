package main

type Todos struct {
	completed bool
	item      string
}

var list = []Todos{}

func AddTodo(todo string) {
	list = append(list, Todos{completed: false, item: todo})
}

func FindTodo(list []Todos, todo string) []Todos {
	var item []Todos
	for _, t := range list {
		if t.item == todo {
			item = []Todos{{completed: t.completed, item: t.item}}
		}
	}
	return item
}

func UpdateTodo(list []Todos, todo string, newTodo string) {
	for i, t := range list {
		if t.item == todo {
			list[i].item = newTodo
		}
	}
}

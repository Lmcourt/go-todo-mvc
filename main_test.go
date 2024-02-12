package main

import (
	"reflect"
	"testing"
)

func TestTodoList(t *testing.T) {
	t.Run("Adds a todo", func(t *testing.T) {
		list = []Todos{{completed: true, item: "test"}}
		AddTodo("do this")
		got := list
		want := []Todos{{completed: true, item: "test"}, {completed: false, item: "do this"}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Finds a todo", func(t *testing.T) {
		list = []Todos{{completed: true, item: "test"}}
		FindTodo(list, "test")
		got := FindTodo(list, "test")
		want := []Todos{{completed: true, item: "test"}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("updates a todo", func(t *testing.T) {
		list = []Todos{{completed: true, item: "test"}}
		updateTodo := "Updated test todo"
		UpdateTodo(list, "test", updateTodo)

		want := []Todos{{completed: true, item: "Updated test todo"}}

		if !reflect.DeepEqual(list, want) {
			t.Errorf("got %v want %v", list, want)
		}
	})

}

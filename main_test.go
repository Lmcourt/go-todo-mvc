package main

import (
	"reflect"
	"testing"
)

func TestTodoList(t *testing.T) {
	list = []Todos{{completed: true, item: "test"}}
	t.Run("Adds a todo", func(t *testing.T) {
		AddTodo("do this")
		got := list
		want := []Todos{{completed: true, item: "test"}, {completed: false, item: "do this"}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Finds an item", func(t *testing.T) {
		FindTodo(list, "test")
		got := FindTodo(list, "test")
		want := []Todos{{completed: true, item: "test"}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

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
		want := []Todos{
			{completed: true, item: "test"},
			{completed: false, item: "do this"},
		}

		AssertEqual(t, got, want)

	})

	t.Run("Finds a todo", func(t *testing.T) {
		list = []Todos{{completed: true, item: "test"}}
		FindTodo(list, "test")

		got := FindTodo(list, "test")
		want := []Todos{{completed: true, item: "test"}}

		AssertEqual(t, got, want)

	})

	t.Run("updates a todo", func(t *testing.T) {
		list = []Todos{{completed: true, item: "test"}}
		updateTodo := "Updated test todo"

		UpdateTodo(list, "test", updateTodo)

		want := []Todos{{completed: true, item: "Updated test todo"}}

		AssertEqual(t, list, want)
	})
	t.Run("toggles completed todo", func(t *testing.T) {
		list = []Todos{
			{completed: true, item: "test"},
			{completed: false, item: "do this"},
		}
		want := []Todos{
			{completed: true, item: "test"},
			{completed: true, item: "do this"},
		}

		ToggleCompleted(list, "do this")

		AssertEqual(t, list, want)
	})

	t.Run("marks all todo as completed", func(t *testing.T) {
		list := []Todos{
			{completed: true, item: "test"},
			{completed: false, item: "test 2"},
			{completed: false, item: "test 3"},
		}
		want := []Todos{
			{completed: true, item: "test"},
			{completed: true, item: "test 2"},
			{completed: true, item: "test 3"},
		}

		ToggleAllCompleted(list)

		AssertEqual(t, list, want)
	})

}

func AssertEqual(t testing.TB, got, want []Todos) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

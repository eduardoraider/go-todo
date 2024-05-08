package entity

import "sync"

// TodoItem represents a single todo item
type TodoItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// TodoList represents a collection of todo items
type TodoList struct {
	sync.RWMutex
	Items []TodoItem
}

// NewTodoList initializes a new TodoList
func NewTodoList() *TodoList {
	return &TodoList{}
}

// GetTodos returns all todo items
func (tl *TodoList) GetTodos() []TodoItem {
	tl.RLock()
	defer tl.RUnlock()
	return tl.Items
}

// AddTodoItem adds a new todo item to the list
func (tl *TodoList) AddTodoItem(item TodoItem) {
	tl.Lock()
	defer tl.Unlock()
	tl.Items = append(tl.Items, item)
}

// UpdateTodoItem updates an existing todo item
func (tl *TodoList) UpdateTodoItem(updatedItem TodoItem) {
	tl.Lock()
	defer tl.Unlock()

	for i, item := range tl.Items {
		if item.ID == updatedItem.ID {
			tl.Items[i] = updatedItem
			break
		}
	}
}

// DeleteTodoItem deletes a todo item by ID
func (tl *TodoList) DeleteTodoItem(id string) {
	tl.Lock()
	defer tl.Unlock()

	for i, item := range tl.Items {
		if item.ID == id {
			tl.Items = append(tl.Items[:i], tl.Items[i+1:]...)
			break
		}
	}
}

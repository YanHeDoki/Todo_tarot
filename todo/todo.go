package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Todo struct {
	Description string
	Done        bool
}

type Todos struct {
	TodoList []*Todo
}

func LoadTodo() *Todos {

	file, err := os.OpenFile("todo.json", os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("read file err")
		return nil
	}
	defer file.Close()
	todobytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("todo json read err:", err)
		return nil
	}
	td := []*Todo{}
	if len(todobytes) <= 0 {
		return &Todos{TodoList: []*Todo{}}
	}
	err = json.Unmarshal(todobytes, &td)
	if err != nil {
		panic(err)
	}
	ts := Todos{TodoList: td}
	return &ts
}

func AddTodo(list []interface{}) {

	file, err := os.OpenFile("todo.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	if err != nil {
		fmt.Println("read file err")
		return
	}
	defer file.Close()
	bytes, err := json.Marshal(list)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	file.Write(bytes)
}

func RemoveTodo(list []interface{}) {

	file, err := os.OpenFile("todo.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	if err != nil {
		fmt.Println("read file err")
		return
	}
	defer file.Close()

	bytes, err := json.Marshal(list)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	file.Write(bytes)
}

func NewTodo(description string) Todo {
	return Todo{description, false}
}

func (t Todo) String() string {
	return fmt.Sprintf("%s - %t", t.Description, t.Done)
}

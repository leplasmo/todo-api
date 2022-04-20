package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/leplasmo/todo-api"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	// (try to) read existing ToDos from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// depending of the number of arguments
	switch {
	// no arguments - print the list
	case len(os.Args) == 1: // note: arg 1 is the command itself
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	// any other number of arguments - add a new ToDo from args
	default:
		// concat all arguments using space
		item := strings.Join(os.Args[1:], " ")

		// add the ToDo
		l.Add(item)

		// save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

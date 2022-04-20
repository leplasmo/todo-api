package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leplasmo/todo-api"
)

var todoFileName = ".todo.json"

func main() {
	// parse command line flags
	task := flag.String("task", "", "Task to be included in the TODO list")
	list := flag.Bool("list", false, "List all TODOs")
	complete := flag.Int("complete", 0, "TODO item to mark as completed")

	flag.Parse()

	// get filename from env
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	l := &todo.List{}

	// (try to) read existing ToDos from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// depending on the provided flags
	switch {
	case *list:
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

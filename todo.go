package main

import ( 
	"errors"
    "fmt" 
	"time" )

type Todo struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func(todos *Todos) add(title string) {
	todo := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos ) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todo *Todos) toggle(index int) error {
	t :=(*todos)
	if err := t.validateIndex(index); err != nil {
	return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return  nil
} //we are making toggle method to change task status if completed and if not 

// here Todo is the type, Think of it as the blueprint for your list.
//*Todo means apointer to a Todos list, a pointer is just memory address which tells the function where to find the original list.
// also todos: this is just a variable name it is the name used inside the method to refer to that pointer

// also t := (*todos) here todos is a pointer ( a memory address) and *todos is a value which we get at that address just like saying that we are dereferncing the pointer. this line gets the actual list of todos


// isCompleted := t[index].Completed , t[index] gets the specific todo from the list.

// .Completed accesses its Completed field, which is either true or false.

// This true or false value is stored in a temporary variable called isCompleted.
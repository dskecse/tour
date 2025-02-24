package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
    next *List[T]
    val  T
}

func (list *List[T]) Push(value T) {
    if list.next != nil {
        list.next.Push(value)
        return
    }
    list.val = value
    list.next = &List[T]{}
}

func (list *List[T]) String() string {
    var elems []T
    for curr := list; curr.next != nil; curr = curr.next {
        elems = append(elems, curr.val)
    }
    return fmt.Sprintf("%v", elems)
}

func main() {
    list := List[int]{}
    list.Push(1)
    list.Push(2)
    list.Push(3)
    fmt.Println("List:", &list) // "List: [1, 2, 3]"
}

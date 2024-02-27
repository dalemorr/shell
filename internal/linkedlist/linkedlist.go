package main

type LinkedList[T any] struct {
	head *node[T]
	size int
}

type node[T any] struct {
	value T
	next  *node[T]
}

func (list *LinkedList[T]) Add(value T) {
	if list.head == nil {
		list.head = &node[T]{value, nil}
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}

		current.next = &node[T]{value, nil}
	}

	list.size++
}

func (list *LinkedList[T]) Insert(index int, value T) {
	if index < 0 || index > list.size {
		return
	}

	if list.head == nil {
		list.head = &node[T]{value, nil}
	} else {
		current := list.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}

		if current.next == nil {
			current.next = &node[T]{value, nil}
		} else {
			current.next = &node[T]{value, current.next}
		}
	}

	list.size++
}

func (list *LinkedList[T]) Remove(index int) {
	if index < 0 || index >= list.size {
		return
	}

	if index == 0 {
		list.head = list.head.next
	} else {
		current := list.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}

	list.size--
}

func (list *LinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index >= list.size {
		return *new(T), nil
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}

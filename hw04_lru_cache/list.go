package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v any) *ListItem
	PushBack(v any) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value any
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v any) *ListItem {
	elem := &ListItem{Value: v}

	switch l.len {
	case 0:
		l.front = elem
		l.back = elem
	case 1:
		l.front = elem
		l.front.Next = l.back
		l.back.Prev = l.front
	default:
		elem.Next = l.front
		l.front.Prev = elem
		l.front = elem
	}

	l.len++
	return elem
}

func (l *list) PushBack(v any) *ListItem {
	elem := &ListItem{Value: v}

	switch l.len {
	case 0:
		l.front = elem
		l.back = elem
	case 1:
		l.back = elem
		l.back.Prev = l.front
		l.front.Next = l.back
	default:
		elem.Prev = l.back
		l.back.Next = elem
		l.back = elem
	}

	l.len++
	return elem
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.len == 1:
		l.front = nil
		l.back = nil
	case i == l.front:
		l.front = l.front.Next
		l.front.Prev = nil
	case i == l.back:
		l.back = l.back.Prev
		l.back.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case l.front == i:
		return
	case l.back == i:
		l.back.Prev.Next = nil
		l.back = l.back.Prev
		l.front.Prev = i
		i.Prev = nil
		i.Next = l.front
		l.front = i
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		l.front.Prev = i
		i.Prev = nil
		i.Next = l.front
		l.front = i
	}
}

func NewList() List {
	return new(list)
}

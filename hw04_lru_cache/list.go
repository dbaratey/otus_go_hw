package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый Item
	Back() *listItem                   // последний Item
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	Value interface{} // значение
	Next  *listItem   // следующий элемент
	Prev  *listItem   // предыдущий элемент
}

type list struct {
	back  *listItem
	front *listItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *listItem {
	return l.front
}

func (l *list) Back() *listItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *listItem {
	res := listItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.len != 0 {
		res.Next = l.front
		l.front.Prev = &res
		l.front = &res
	} else {
		l.front = &res
		l.back = &res
	}
	l.len++
	return &res
}

func (l *list) PushBack(v interface{}) *listItem {
	res := listItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.len != 0 {
		res.Prev = l.back
		l.back.Next = &res
		l.back = &res
	} else {
		l.front = &res
		l.back = &res
	}
	l.len++
	return &res
}

func (l *list) Remove(i *listItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
		if l.front == i {
			l.front = i.Next
		}
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next

		if l.back == i {
			l.back = i.Prev
		}
	}
	if l.len == 1 {
		l.back = nil
		l.front = nil
	}
	i.Prev = nil
	i.Next = nil
	i.Value = nil
	l.len--
}

func (l *list) MoveToFront(i *listItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
		if l.back == i {
			l.back = i.Prev
		}
	} else {
		return
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i.Next = l.front
	i.Prev = nil
	l.front.Prev = i
	l.front = i
}

func NewList() List {
	return &list{}
}

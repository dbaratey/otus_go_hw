package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()
		require.Equal(t, l.Len(), 0)
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})
	t.Run("add first(single)", func(t *testing.T) {
		l := NewList()
		l.PushFront("3")
		require.Equal(t, l.Front(), l.Back())
	})
	t.Run("remove first(single)", func(t *testing.T) {
		l := NewList()
		l.PushFront("3")
		l.Remove(l.Front())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})
	t.Run("add last", func(t *testing.T) {
		l := NewList()
		l.PushFront("1")
		l.PushFront("2")
		l.PushFront("3")
		l.PushBack("5")
		require.Equal(t, "5", l.Back().Value)
		require.Equal(t, "1", l.Back().Prev.Value)
		require.Nil(t, l.Back().Next)
	})
	t.Run("remove middle", func(t *testing.T) {
		l := NewList()
		l.PushFront("1")
		l.PushFront("2")
		l.PushFront("3")
		l.Remove(l.Front().Next)
		require.Equal(t, "3", l.Front().Value)
		require.Equal(t, "1", l.Front().Next.Value)
		require.Equal(t, "1", l.Back().Value)
		require.Equal(t, "3", l.Back().Prev.Value)
	})
	t.Run("remove last", func(t *testing.T) {
		l := NewList()
		l.PushFront("3")
		l.PushBack("5")
		l.Remove(l.Back())
		require.Equal(t, "3", l.Front().Value)
		require.Equal(t, "3", l.Back().Value)
	})
	t.Run("complex", func(t *testing.T) {
		l := NewList()
		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, l.Len(), 3)

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, l.Len(), 2)

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, l.Len(), 7)
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

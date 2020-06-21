package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

type Key string

type Item struct {
	Value interface{} // значение
	Next  *Item       // следующий элемент
	Prev  *Item       // предыдущий элемент
}

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                              // Очистить кэш
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    list
	items    map[Key]*cacheItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.Lock()
	if _, ok := l.items[key]; ok {
		l.items[key].value = value
		l.queue.MoveToFront(l.items[key].el)
		l.Unlock()
		return true
	}
	if l.capacity == l.queue.len && l.queue.len != 0 {
		delete(l.items, l.queue.back.Value.(Key))
		l.queue.Remove(l.queue.back)
	}
	l.items[key] = &cacheItem{
		value: value,
		el:    l.queue.PushFront(key),
	}
	l.Unlock()
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.Lock()
	if _, ok := l.items[key]; ok {
		l.queue.MoveToFront(l.items[key].el)
		l.Unlock()
		return l.items[key].value, true
	}
	l.Unlock()
	return nil, false
}

func (l *lruCache) Clear() {
	l.Lock()
	l.queue = list{}
	l.items = make(map[Key]*cacheItem)
	l.Unlock()
}

type cacheItem struct {
	value interface{}
	el    *listItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    list{},
		items:    make(map[Key]*cacheItem),
	}
}

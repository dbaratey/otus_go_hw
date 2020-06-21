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

type RWMap struct {
	mu sync.RWMutex
	m  map[Key]*cacheItem
}

func (rwm *RWMap) Store(key Key, val interface{}, queue *list, capacity int) (*cacheItem, bool) {
	rwm.mu.Lock()
	var p *cacheItem
	var ok bool
	if p,ok = rwm.m[key];ok{
		rwm.m[key].SetVal(val)
		queue.MoveToFront(rwm.m[key].el)
	} else{
		if capacity == queue.len{
			delete(rwm.m,queue.back.Value.(Key))
			queue.Remove(queue.back)
		}
		rwm.m[key] = &cacheItem{
			value: val,
			el:  queue.PushFront(key),
		}
		p = rwm.m[key]
	}
	rwm.mu.Unlock()
	return p,ok
}

func (rwm *RWMap) Get(key Key) (*cacheItem,bool) {
	rwm.mu.RLock()
	v,ok := rwm.m[key]
	rwm.mu.RUnlock()
	return v,ok
}

func (rwm *RWMap) Remove(key Key) {
	rwm.mu.Lock()
	delete(rwm.m, key)
	rwm.mu.Unlock()
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    list
	items    RWMap
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.Lock()
	_,ok := l.items.Store(key,value,&l.queue,l.capacity)
	l.Unlock()
	return ok
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.Lock()
	if v, ok := l.items.Get(key); ok {
		l.queue.MoveToFront(v.el)
		l.Unlock()
		return v.Val(), true
	}
	l.Unlock()
	return nil, false
}

func (l *lruCache) Clear() {
	l.Lock()
	l.queue = list{}
	l.items = RWMap{m:make(map[Key]*cacheItem)}
	l.Unlock()
}

type cacheItem struct {
	mu sync.RWMutex
	value interface{}
	el    *listItem
}

func (ci *cacheItem) Val() interface{}{
	ci.mu.RLock()
	res := ci.value
	ci.mu.RUnlock()
	return res
}

func (ci *cacheItem) El() *listItem{
	ci.mu.RUnlock()
	res := ci.el
	ci.mu.RUnlock()
	return res
}

func (ci *cacheItem) SetVal(val interface{}){
	ci.mu.Lock()
	ci.value = val
	ci.mu.Unlock()
}

func (ci *cacheItem) SetEl(el *listItem){
	ci.mu.Lock()
	ci.el = el
	ci.mu.Unlock()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    list{},
		items:    RWMap{m:make(map[Key]*cacheItem)},
	}
}

package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value any) bool
	Get(key Key) (any, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value any
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value any) bool {
	if li, ok := c.items[key]; ok {
		c.queue.MoveToFront(li)
		li.Value = cacheItem{key: key, value: value}
		return true
	}

	item := cacheItem{key: key, value: value}
	insertedItemPtr := c.queue.PushFront(item)
	c.items[key] = insertedItemPtr

	if c.queue.Len() > c.capacity {
		lastItem := c.queue.Back()
		c.queue.Remove(lastItem)

		if valWithKey, ok := lastItem.Value.(cacheItem); ok {
			delete(c.items, valWithKey.key)
		}
	}

	return false
}

func (c *lruCache) Get(key Key) (any, bool) {
	if li, ok := c.items[key]; ok {
		c.queue.MoveToFront(li)

		if valWithKey, ok := li.Value.(cacheItem); ok {
			return valWithKey.value, true
		}
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

package Cache

import (
	"container/list"
)

/*
LRU(least Recently Used)最近最少使用策略缓存
2020.03.09
zhangliu
*/
//缓存
type Cache struct {
	cache    map[string]*list.Element
	ll       *list.List
	maxBytes int64
	nBytes   int64
	//回调方法
	OnEvicted func(key string, value Value)
}

//给elemnet的具体类型
type entry struct {
	key   string
	value Value
}

//值类型要实现返回内存大小
type Value interface {
	Len() int
}

func NewCache(maxBytes int64, onEvicted func(key string, value Value)) *Cache {
	c := &Cache{
		cache:     make(map[string]*list.Element),
		ll:        list.New(),
		maxBytes:  maxBytes,
		nBytes:    0,
		OnEvicted: onEvicted,
	}
	return c
}

//获取
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

//添加
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
		return
	}
	ele := c.ll.PushFront(&entry{
		key:   key,
		value: value,
	})
	c.cache[key] = ele
	c.nBytes += int64(len(key)) + int64(value.Len())
	for c.maxBytes > 0 && c.nBytes > c.maxBytes {
		c.RemoveOldest()
	}
}
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

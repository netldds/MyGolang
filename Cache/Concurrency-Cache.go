package Cache

import (
	"log"
	"sync"
)

//返回值类型

type ByteView struct {
	b []byte
}

//实现len()
func (b ByteView) Len() int {
	return len(b.b)
}

//实现返回副本
func (b ByteView) ByteSlice() []byte {
	c := make([]byte, len(b.b))
	copy(c, b.b)
	return c
}

//实现String()
func (b ByteView) String() string {
	return string(b.b)
}

//封装cache,支持互斥锁
type cache struct {
	mu       sync.Mutex
	lru      *Cache
	maxBytes int64
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = NewCache(c.maxBytes, nil)
	}
	c.lru.Add(key, value)
}

//回调函数,获取不存在的值
type Getter interface {
	Get(key string) (b []byte, err error)
}
type GetterFunc func(key string) ([]byte, error)

func (g GetterFunc) Get(key string) (b []byte, err error) {
	return g(key)
}

//核心结构,互斥锁,名字空间
type Group struct {
	name   string
	getter Getter
	gcache cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:   name,
		getter: getter,
		gcache: cache{
			mu:       sync.Mutex{},
			lru:      nil,
			maxBytes: cacheBytes,
		},
	}
	return g
}
func GetGroup(name string) *Group {
	if name == "" {
		return nil
	}
	mu.RLock()
	defer mu.RUnlock()
	return groups[name]
}
func (g *Group) Get(key string) (value ByteView, err error) {
	if key == "" {
		return
	}
	if v, ok := g.gcache.get(key); ok {
		log.Printf("key: %v hit", key)
		return v, nil
	}
	return g.load(key)
}

//从源获取数据
func (g *Group) load(key string) (value ByteView, err error) {
	return g.getlocally(key)
}

//本地
func (g *Group) getlocally(key string) (value ByteView, err error) {
	b, err := g.getter.Get(key)
	value = ByteView{b: b}
	g.populateCache(key, value)
	return
}

//加入
func (g *Group) populateCache(key string, value ByteView) {
	g.gcache.add(key, value)
}

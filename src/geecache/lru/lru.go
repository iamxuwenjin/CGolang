package lru

import "container/list"

type Cache struct {
	maxBytes  int64                    // 最大容量
	nBytes    int64                    // 当前容量
	ll        *list.List               // 有序列表
	cache     map[string]*list.Element //
	onEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// 值为所有实现了Len方法的类型
type Value interface {
	Len() int
}

// 实例化一个Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		onEvicted: onEvicted,
	}
}

// query
func (c *Cache) Get(key string) (value Value, ok bool) {
	// 如果存在,将其移动到队首，返回结果值
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, ok
	}
	return
}

// delete
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.onEvicted != nil {
			c.onEvicted(kv.key, kv.value)
		}
	}
}

//add
func (c *Cache) Add(key string, value Value) {
	// 如果存在的话，对其进行更新，将其移动到最前面
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		// 加上新增的，减去被替换的
		c.nBytes += int64(len(key)) + int64(value.Len())
		kv.value = value
	} else {
		// 不存在直接添加在最前面
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

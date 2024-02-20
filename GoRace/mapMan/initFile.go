package mapMan

import "sync"

/**
 * @Author shenfz
 * @Date 2022/2/15 16:15
 * @Email 1328919715@qq.com
 * @Description:
 **/

type ConcurrentMap struct {
	locker sync.RWMutex
	mMan   map[string]interface{}
}

var (
	m    ConcurrentMap
	once sync.Once
)

func GetConcurrentMap() *ConcurrentMap {
	once.Do(m.init)
	return &m
}

func (c *ConcurrentMap) init() {
	c.mMan = make(map[string]interface{})
}

func (c *ConcurrentMap) Get(key string) interface{} {
	// c.locker.RLock()
	// defer c.locker.RUnlock()
	if tmp, ok := c.mMan[key]; ok {
		return tmp
	}
	return nil
}

func (c *ConcurrentMap) Set(key string, val interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.mMan[key] = val
}

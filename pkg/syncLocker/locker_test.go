package SyncLocker

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
 * @Author shenfz
 * @Date 2021/8/11 15:08
 * @Email 1328919715@qq.com
 * @Description:
 **/

type Config struct {
	locker sync.Locker
	add    int
	del    int
	data   map[string]string
}

func (c *Config) Add(key, val string) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.data[key] = val
	c.add++
}

func (c *Config) Del(key string) {
	c.locker.Lock()
	defer c.locker.Unlock()
	delete(c.data, key)
	c.del++
}

func (c *Config) Len() (int, int, int) {
	return len(c.data), c.add, c.del
}

func Test_NewLocker(t *testing.T) {
	conf := Config{
		locker: NewSpinLock(),
		data:   make(map[string]string),
	}
	go func() {
		tricker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-tricker.C:
				t.Log(conf.Len())
			default:
			}
		}

	}()
	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(50 * time.Millisecond)
			conf.Add(fmt.Sprint(i), fmt.Sprint(i))
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(60 * time.Millisecond)
			conf.Del(fmt.Sprint(i))
		}
	}()

	select {}
}

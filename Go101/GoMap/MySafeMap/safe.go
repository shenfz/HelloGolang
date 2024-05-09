package MySafeMap

/**
 * @Author shenfz
 * @Date 2021/12/20 16:39
 * @Email 1328919715@qq.com
 * @Description:
 **/

import "sync"

const (
	copyThreshold = 1000  // 复制阀值       避免过多的迁移数据
	maxDeletion   = 10000 // 最大删除标记数  触发迁移条件
)

type SafeMap struct {
	lock        sync.RWMutex
	deletionOld int
	deletionNew int
	dirtyOld    map[interface{}]interface{}
	dirtyNew    map[interface{}]interface{}
}

// NewSafeMap returns a SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		dirtyOld: make(map[interface{}]interface{}),
		dirtyNew: make(map[interface{}]interface{}),
	}
}

func (s *SafeMap) Delete(key interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	// check key in old or new ,del it
	if _, ok := s.dirtyOld[key]; ok {
		delete(s.dirtyOld, key)
		s.deletionOld++
	} else if _, ok := s.dirtyNew[key]; ok {
		delete(s.dirtyNew, key)
		s.deletionNew++
	}

	// old 满足迁移条件 ，迁移后拷贝完整new给old , new 再清空
	if s.deletionOld >= maxDeletion && len(s.dirtyOld) <= copyThreshold {
		for key, val := range s.dirtyOld {
			s.dirtyNew[key] = val
		}
		// copy
		s.dirtyOld = s.dirtyNew
		s.deletionOld = s.deletionNew
		// init new
		s.dirtyNew = map[interface{}]interface{}{}
		s.deletionNew = 0
	}

	//new: 满足迁移条件 , 直接完成迁移 ， new 再清空
	if s.deletionNew >= maxDeletion && len(s.dirtyNew) <= copyThreshold {
		for key, val := range s.dirtyNew {
			s.dirtyOld[key] = val
		}
		s.deletionNew = 0
		s.dirtyNew = map[interface{}]interface{}{}
	}
}

func (s *SafeMap) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.dirtyOld) + len(s.dirtyNew)
}

func (s *SafeMap) Set(key interface{}, val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	// 主力old 标记小于阀值，做掉new中假定存在的旧值,反之亦然
	if s.deletionOld < maxDeletion {
		if _, ok := s.dirtyNew[key]; ok {
			delete(s.dirtyNew, key)
			s.deletionNew++
		}
		s.dirtyOld[key] = val
	} else {
		if _, ok := s.dirtyOld[key]; ok {
			delete(s.dirtyOld, key)
			s.deletionOld++
		}
		s.dirtyNew[key] = val
	}
}

func (s *SafeMap) Get(key interface{}) (val interface{}, ok bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if val, ok = s.dirtyOld[key]; ok {
		return
	}
	val, ok = s.dirtyNew[key]
	return
}

package m2rcache

import "time"

type nameSpace struct {
	cacherMap map[string]Cacher
	mgoDB     *mgoDB
	redisDB   *redisDB
	cacheTime time.Duration
}

func (f *nameSpace) RegNameSpace(namespace string, build Builder) Cacher {
	return f.RegNameSpaceForCacheTime(namespace, build, -1)
}
func (f *nameSpace) RegNameSpaceForCacheTime(namespace string, build Builder, cacheTime time.Duration) Cacher {
	if cacheTime == -1 {

		cacheTime = f.cacheTime
	}
	var bc = &baseCache{
		mgoDB:     f.mgoDB,
		redisDB:   f.redisDB,
		build:     build,
		cacheTime: cacheTime,
	}
	if f.cacherMap == nil {
		f.cacherMap = make(map[string]Cacher)
	}
	f.cacherMap[namespace] = bc
	return bc
}

func (f *nameSpace) RemoveNameSpace(namespace string) Cacher {
	if f.cacherMap == nil {
		return nil
	}
	c, ok := f.cacherMap[namespace]
	if ok {
		delete(f.cacherMap, namespace)
	}
	return c
}
func (f *nameSpace) GetNameSpace(namespace string) Cacher {
	if f.cacherMap == nil {
		return nil
	}
	return f.cacherMap[namespace]
}

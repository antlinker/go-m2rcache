package m2rcache

import "time"

var defaultFactory Factory

// RegNameSpace 注册命名空间
func RegNameSpace(namespace string, build Builder) Cacher {
	return defaultFactory.RegNameSpace(namespace, build)
}

// RegNameSpaceForCacheTime 注册命名空间
func RegNameSpaceForCacheTime(namespace string, build Builder, cachetime time.Duration) Cacher {
	return defaultFactory.RegNameSpaceForCacheTime(namespace, build, cachetime)
}

// RemoveNameSpace 移除命名空间
func RemoveNameSpace(namespace string) Cacher {
	return defaultFactory.RemoveNameSpace(namespace)
}

// GetNameSpace 获取命名空间
func GetNameSpace(namespace string) Cacher {
	return defaultFactory.GetNameSpace(namespace)
}

// Get 获取值
func Get(namespace string, key interface{}, result interface{}, fields ...string) error {
	return defaultFactory.Get(namespace, key, result, fields...)
}

// InitFactory 初始化工厂
func InitFactory(rc RedisCfg, mgocfg MgoCfg, cacheTime time.Duration) {
	defaultFactory = CreateFacotry(rc, mgocfg, cacheTime)
}

// CreateFacotry 创建缓存工厂
func CreateFacotry(rc RedisCfg, mgocfg MgoCfg, cacheTime time.Duration) Factory {
	f := &factory{}
	f.mgoDB = createMgoDB(mgocfg)
	f.redisDB = createRedisDB(rc)
	f.cacheTime = cacheTime
	return f
}

package m2rcache

var defaultFactory Factory

// RegNameSpace 注册命名空间
func RegNameSpace(namespace string, build Builder) Cacher {
	return defaultFactory.RegNameSpace(namespace, build)
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
func InitFactory(rc RedisCfg, mgocfg MgoCfg) {
	defaultFactory = CreateFacotry(rc, mgocfg)
}

// CreateFacotry 创建缓存工厂
func CreateFacotry(rc RedisCfg, mgocfg MgoCfg) Factory {
	f := &factory{}
	f.mgoDB = CreateMgoDB(mgocfg)
	f.redisDB = createRedisDB(rc)
	return f
}

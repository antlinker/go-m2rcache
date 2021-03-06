package m2rcache

type nameSpace struct {
	cacherMap map[string]Cacher
	mgoDB     MgoDB
	redisDB   RedisDB
}

func (f *nameSpace) RegNameSpace(namespace string, build Builder) Cacher {

	var bc = &baseCache{
		mgoDB:   f.mgoDB,
		redisDB: f.redisDB,
		build:   build,
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

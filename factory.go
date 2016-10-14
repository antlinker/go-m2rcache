package m2rcache

type factory struct {
	nameSpace
}

func (f *factory) Get(namespace string, key interface{}, result interface{}, fields ...string) error {
	c := f.GetNameSpace(namespace)
	if c == nil {
		return nil
	}
	return c.Get(key, result, fields...)
}

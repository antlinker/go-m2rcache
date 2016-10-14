package m2rcache

import "fmt"

type baseCache struct {
	mgoDB   MgoDB
	redisDB RedisDB
	build   Builder
}

func (e *baseCache) Get(id interface{}, result interface{}, fields ...string) error {
	key := e.build.Key(id)

	data, err := e.redisDB.Get(key, fields...)
	if err == nil {
		for _, v := range data {
			if v != nil {
				fullValue(result, fields, data)
				return nil
			}
		}

	}
	r, err := e.build.LoadData(e.mgoDB, e.redisDB, id, fields...)
	if err != nil {
		return err
	}

	fullValue(result, fields, r)

	return nil
}
func formatStr(data interface{}) string {
	return fmt.Sprintf("%v", data)
}

package m2rcache

import (
	"fmt"
	"time"
)

type baseCache struct {
	mgoDB     *mgoDB
	redisDB   *redisDB
	build     Builder
	cacheTime time.Duration
}

func (e *baseCache) Get(id interface{}, result interface{}, fields ...string) error {
	key := e.build.Key(id)
	handler := e.build.Handler()
	data, err := e.redisDB.get(key, fields...)
	if err == nil {
		for _, v := range data {
			if v != nil {
				fullValue(result, fields, data)
				return nil
			}
		}

	}
	if handler != nil {
		r, err := handler(e.mgoDB, id, fields...)
		if err != nil {
			return err
		}

		fmt.Println("mgo:", r)
		pairs := make([]string, 0, len(fields)-1)
		for i, v := range fields {
			if i == 0 {
				continue
			}
			pairs = append(pairs, v)
			pairs = append(pairs, formatStr(r[i]))
		}
		e.redisDB.save(key, e.cacheTime, fields[0], formatStr(r[0]), pairs...)
		fullValue(result, fields, r)
	}
	return nil
}
func formatStr(data interface{}) string {
	return fmt.Sprintf("%v", data)
}

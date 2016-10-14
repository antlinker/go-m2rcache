package m2rcache

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CreateDefaultBuild make
func CreateDefaultBuild(nameSpace, collname, idname string, cacheTime time.Duration, fields ...string) Builder {

	return &defaultBuild{
		collname:  collname,
		idname:    idname,
		fields:    fields,
		nameSpace: nameSpace,
		cacheTime: cacheTime,
	}
}

type defaultBuild struct {
	idname    string
	collname  string
	fields    []string
	nameSpace string
	cacheTime time.Duration
}

func (b *defaultBuild) Key(id interface{}) string {
	return fmt.Sprintf("%s:mgo:%s:%v", b.nameSpace, b.collname, id)
}

// func (b *defaultBuild) Handler() LoadDataHandler {
// 	return func(db MgoDB, id interface{}, fields ...string) ([]interface{}, error) {
// 		var result map[string]interface{}
// 		var query = bson.M{b.idname: id}
// 		var selecter = bson.M{}
// 		for _, f := range fields {
// 			selecter[f] = 1
// 		}
// 		err := db.Get(&result, b.collname, query, selecter)
// 		if err != nil {
// 			return nil, err
// 		}
// 		l := len(fields)
// 		out := make([]interface{}, l, l)
// 		for i, f := range fields {
// 			out[i] = result[f]
// 		}
// 		return out, nil
// 	}
// }
func (b *defaultBuild) loadMgo(db MgoDB, id interface{}, fields []string) (map[string]interface{}, error) {
	var result map[string]interface{}
	var query = bson.M{b.idname: id}
	var selecter = bson.M{}
	for _, f := range fields {
		selecter[f] = 1
	}
	err := db.Get(&result, b.collname, query, selecter)
	if err != nil {
		return nil, err
	}
	return result, nil

}
func (b *defaultBuild) LoadData(mgodb MgoDB, redisdb RedisDB, id interface{}, fields ...string) ([]interface{}, error) {

	result, err := b.loadMgo(mgodb, id, b.fields)
	if err != nil {
		return nil, err
	}
	l := len(b.fields)
	all := make([]interface{}, l, l)
	for i, f := range b.fields {
		all[i] = result[f]
	}
	//	fmt.Println("mgo:", r)
	pairs := make([]string, 0, len(fields)-1)
	for i, v := range fields {
		if i == 0 {
			continue
		}
		pairs = append(pairs, v)
		pairs = append(pairs, formatStr(all[i]))
	}
	redisdb.Save(b.Key(id), b.cacheTime, fields[0], formatStr(all[0]), pairs...)
	out := make([]interface{}, l, l)
	for i, f := range fields {
		out[i] = result[f]
	}
	return out, nil
}

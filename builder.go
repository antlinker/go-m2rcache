package m2rcache

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func CreateDefaultBuild(collname, idname string) Builder {

	return &defaultBuild{
		collname: collname,
		idname:   idname,
	}
}

type defaultBuild struct {
	idname   string
	collname string
}

func (b *defaultBuild) Key(id interface{}) string {
	return fmt.Sprintf("mgo:%s:%v", b.collname, id)
}

func (b *defaultBuild) Handler() LoadDataHandler {
	return func(db MgoDB, id interface{}, fields ...string) ([]interface{}, error) {
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
		l := len(fields)
		out := make([]interface{}, l, l)
		for i, f := range fields {
			out[i] = result[f]
		}
		return out, nil
	}
}

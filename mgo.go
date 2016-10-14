package m2rcache

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2"
)

func createMgoDB(cfg MgoCfg) *mgoDB {
	var m = &mgoDB{}
	m.initDB(cfg)
	return m
}

// MgoCfg mongodb　配置
type MgoCfg struct {
	URL, DBName string
}

// MgoDB mongodb操作
type MgoDB interface {
	Get(result interface{}, collname string, query interface{}, selector interface{}) error
}

type mgoDB struct {
	session   *mgo.Session
	defaultDB string
}

func (db *mgoDB) initDB(cfg MgoCfg) {
	db.defaultDB = cfg.DBName
	tmp, err := mgo.Dial(cfg.URL)
	if err != nil {
		panic(errors.New("创建mongodb连接失败1:" + fmt.Sprintf("%v\n", err)))
	}
	db.session = tmp
	// Optional. Switch the session to a monotonic behavior.
	db.session.SetMode(mgo.Eventual, true)

	return
}
func (db *mgoDB) getSession() *mgo.Session {
	return db.session.Clone()
}
func (db *mgoDB) Get(result interface{}, collname string, query interface{}, selector interface{}) error {
	return db.getSession().DB(db.defaultDB).C(collname).Find(query).Select(selector).One(result)
}

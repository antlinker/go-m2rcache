package m2rcache_test

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/antlinker/go-m2rcache"
)

var (
	redisCfg = m2rcache.RedisCfg{
		Address: []string{"127.0.0.1:6379"},
		DB:      0,

		MaxRetries: 3,

		DialTimeout:  10,
		ReadTimeout:  3,
		WriteTimeout: 3,

		PoolSize: 10,

		PoolTimeout: 5,

		IdleTimeout: 5,
	}
	mgoCfg = m2rcache.MgoCfg{
		URL:    "127.0.0.1:27017",
		DBName: "testCache",
	}
)

func TestMain(m *testing.M) {

	m2rcache.InitFactory(redisCfg, mgoCfg)

	m.Run()
}

const (
	collname = "test3"
)

func TestOk(t *testing.T) {
	var result map[string]interface{}
	m2rcache.RegNameSpace("test2", m2rcache.CreateDefaultBuild("test2", "test2", "_id", 5*time.Second, "A", "B", "C"))
	c := m2rcache.GetNameSpace("test2")
	c.Get(1, &result, "A", "B")
	t.Log(result)
	c.Get(1, &result, "A", "C")
	t.Log(result)
}
func BenchmarkMgo(b *testing.B) {
	b.StopTimer()
	mgoDB := m2rcache.CreateMgoDB(mgoCfg)

	var selecter = bson.M{"A": 1, "B": 1}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var query = bson.M{"_id": i}
		var result map[string]interface{}
		mgoDB.Get(&result, collname, query, selecter)
	}
}
func BenchmarkCache(b *testing.B) {
	b.StopTimer()
	m2rcache.RegNameSpace(collname, m2rcache.CreateDefaultBuild(collname, collname, "_id", 5*time.Minute, "A", "B", "C"))
	c := m2rcache.GetNameSpace(collname)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var result map[string]interface{}
		c.Get(i, &result, "A", "B")
	}
}

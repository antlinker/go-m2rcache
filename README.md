# mongodb　 redis缓存

* 使用reids缓存mongodb中的值

## 例子

```golang

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

  	m2rcache.InitFactory(redisCfg, mgoCfg, 5*time.Second)

  	m.Run()
  }

  func TestOk(t *testing.T) {
  	var result map[string]interface{}
  	m2rcache.RegNameSpace("test2", m2rcache.CreateDefaultBuild("test2", "_id"))
  	c := m2rcache.GetNameSpace("test2")
  	c.Get(1, &result, "A", "B", "C")
  	t.Log(result)
  	c.Get(1, &result, "A", "B", "C")
  	t.Log(result)
  }


```

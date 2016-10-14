package m2rcache

import "time"

// Cacher 缓存接口
type Cacher interface {
	// 获取一个缓存值
	// id 缓存唯一标识
	// result 返回值，是一个地址，可以是map struct slice的地址
	// fields 缓存的字段
	Get(id interface{}, result interface{}, fields ...string) error
}

// NameSpacer 缓存命名空间
type NameSpacer interface {
	// 注册一个命名空间
	RegNameSpace(namespace string, build Builder) Cacher
	// 注册带有缓存时间的命名空间
	// cacheTime -1继承整体缓存设置　0 不缓存
	RegNameSpaceForCacheTime(namespace string, build Builder, cacheTime time.Duration) Cacher
	// 移除一个命名空间
	RemoveNameSpace(namespace string) Cacher
	// 获取一个命名空间
	GetNameSpace(namespace string) Cacher
}

// Factory 缓存命名空间
// 缓存管理工厂方法
type Factory interface {
	NameSpacer
	// 获取一个缓存值
	// namespace 命名空间
	// id 缓存唯一标识
	// result 返回值，是一个地址，可以是map struct slice的地址
	// fields 缓存的字段
	Get(namespace string, id interface{}, result interface{}, fields ...string) error
}

// LoadDataHandler 加载数据方法
type LoadDataHandler func(db MgoDB, id interface{}, fields ...string) ([]interface{}, error)

// Builder 缓存创建者
// 用来构建缓存数据来源及缓存键名
type Builder interface {
	// 缓存键生成器
	// id 缓存唯一标识
	Key(id interface{}) string
	// 数据加载函数句柄
	Handler() LoadDataHandler
}

package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// Rdb 声明一个全局的rdb变量
var Rdb *redis.Client

// Init 初始化连接
func Init() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetString("redis.port"),
		),
		DB:       int(viper.GetInt64("redis.db")), // use default DB
		PoolSize: int(viper.GetInt64("redis.pool_size")),
	})

	_, err = Rdb.Ping().Result()
	return err
}

func Close() {
	_ = Rdb.Close()
}

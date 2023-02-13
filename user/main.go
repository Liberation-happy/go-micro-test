package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go.uber.org/zap"
	"user/config/config"
	"user/config/logger"
	"user/config/mysql"
	"user/config/redis"
)

var (
	service = "user"
	version = "latest"
)

// InitDB 初始化数据库
func InitDB() {
	db, err := mysql.MysqlInit(
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("dbname"),
	)
	if err != nil {
		logger.Error(err)
		return
	}
	defer db.Close()
}

// InitLogger 初始化日志
func InitLogger() {
	if err := logger.Init(); err != nil {
		fmt.Printf("Init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()
}

// InitRdb 初始化redis
func InitRdb() {
	if err := redis.Init(); err != nil {
		logger.Error(err)
		return
	}
	defer redis.Close()
}

// InitAll 初始化合集
func InitAll() {
	config.InitConfig()
	InitLogger()
	InitDB()
	InitRdb()
}

func main() {
	InitAll()

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

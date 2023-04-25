package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cn-lxy/music-api/tools/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// 建立连接
	conn, err := mongo.Connect(ctx,
		options.Client().
			// 连接地址
			ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.Cfg.Mongo.Host, config.Cfg.Mongo.Port)).
			// 设置验证参数
			SetAuth(
				options.Credential{
					// 用户名
					Username: config.Cfg.Mongo.Username,
					// 密码
					Password: config.Cfg.Mongo.Password,
				}).
			// 设置连接数
			SetMaxPoolSize(20))
	if err != nil {
		log.Println(err)
		return
	}
	// 测试连接
	err = conn.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("mongodb connect success!")
	Client = conn
}

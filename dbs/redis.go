package dbs

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RDBConfig struct {
	Driver          string        `json:"driver" yaml:"driver"`                   // 数据库驱动类型
	Host            string        `json:"host" yaml:"host"`                       // Redis 服务器地址
	Port            int           `json:"port" yaml:"port"`                       // Redis 服务器端口
	Pass            string        `json:"pass" yaml:"pass"`                       // Redis 服务器密码
	DB              int           `json:"name" yaml:"name"`                       // Redis 数据库索引
	PoolSize        int           `json:"poolSize" yaml:"poolSize"`               // 连接池大小
	MaxIdleConns    int           `json:"maxIdleConns" yaml:"maxIdleConns"`       // 最大空闲连接数
	MinIdleConns    int           `json:"minIdleConns" yaml:"minIdleConns"`       // 最小空闲连接数
	MaxActiveConns  int           `json:"maxActiveConns" yaml:"maxActiveConns"`   // 最大活动连接数
	ConnMaxIdleTime time.Duration `json:"connMaxIdleTime" yaml:"connMaxIdleTime"` // 连接最大空闲时间
	ConnMaxLifetime time.Duration `json:"connMaxLifetime" yaml:"connMaxLifetime"` // 连接最大生命周期
}

// NewRedis 新建 redis 连接
func NewRedis(rdb RDBConfig) (*redis.Client, error) {
	// 创建 Redis 客户端
	rdbClient := redis.NewClient(&redis.Options{
		Addr:            rdb.Host + ":" + strconv.Itoa(rdb.Port),
		Password:        rdb.Pass,
		DB:              rdb.DB,
		PoolSize:        rdb.PoolSize,
		MaxIdleConns:    rdb.MaxIdleConns,
		MinIdleConns:    rdb.MinIdleConns,
		MaxActiveConns:  rdb.MaxActiveConns,
		ConnMaxIdleTime: rdb.ConnMaxIdleTime,
		ConnMaxLifetime: rdb.ConnMaxLifetime,
	})

	// Ping 测试连接
	pong, err := rdbClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to Redis:", pong)

	// 关闭连接
	_ = rdbClient.Close()

	return rdbClient, nil
}

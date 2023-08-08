package global

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tietiexx/bot/code/backend/constant"
	"sync"
	"time"
)

var (
	redisOnce   sync.Once
	RedisClient *redis.Client
)

func redisConn(config *constant.RedisConfig) (*redis.Client, error) {
	var redisConn *redis.Client
	var err error

	redisOnce.Do(
		func() {
			redisOptions := &redis.Options{
				Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
				Password:     config.Password,
				DB:           config.Database,
				PoolSize:     config.PoolSize,
				MinIdleConns: config.MaxIdle,
				IdleTimeout:  time.Minute,
			}
			// 初始化Redis连接
			redisConn = redis.NewClient(redisOptions)
			// 连接测试
			ctx := context.Background()
			err = redisConn.Ping(ctx).Err()
		})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}
	return redisConn, nil
}

func InitRedis(conf *constant.RedisConfig) {
	conn, err := redisConn(conf)
	if err != nil {
		panic(err)
		return
	}
	RedisClient = conn
}

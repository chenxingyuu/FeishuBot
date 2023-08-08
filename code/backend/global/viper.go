package global

import (
	"fmt"
	"github.com/spf13/viper"
)

type ViperManager struct {
}

// NewViperManger 构造
func NewViperManger() *ViperManager {
	vm := &ViperManager{}
	return vm
}

func (vm *ViperManager) loadConfig() {
	loggerViper := viper.New()
	loggerViper.SetConfigName("config") // 配置文件名称（无需后缀）
	loggerViper.SetConfigType("yaml")   // 配置文件类型
	loggerViper.AddConfigPath("./")     // 配置文件路径

	// 读取配置文件
	err := loggerViper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 将配置文件中的数据加载到结构体中
	// 加载 mmd
	err = loggerViper.UnmarshalKey("bot_admin", BotAdminConf)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal mmd config: %s", err))
	}
	// 加载 zap_logger
	err = loggerViper.UnmarshalKey("zap_logger", ZapLoggerConf)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal zap_log config: %s", err))
	}
	// 加载 mysql
	err = loggerViper.UnmarshalKey("mysql", MySQLConf)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal mysql config: %s", err))
	}
	// 加载 redis
	err = loggerViper.UnmarshalKey("redis", RedisConf)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal redis config: %s", err))
	}
}

func InitViper() {
	viperManager := NewViperManger()
	viperManager.loadConfig()
}

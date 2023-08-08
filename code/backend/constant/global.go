package constant

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

type BotAdminConfig struct {
	Domain        string `mapStructure:"domain"`
	Host          string `mapStructure:"host"`
	Port          int    `mapStructure:"port"`
	DebugPassword string `mapStructure:"debugPassword"`
}

func (c *BotAdminConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type ZapLoggerConfig struct {
	Level            string                 `mapStructure:"level"`
	Encoding         string                 `mapStructure:"encoding"`
	OutputPaths      []string               `mapStructure:"outputPaths"`
	ErrorOutputPaths []string               `mapStructure:"errorOutputPaths"`
	InitialFields    map[string]interface{} `mapStructure:"initialFields"`
	EncoderConfig    zapcore.EncoderConfig  `mapStructure:"encoderConfig"`
}

type MySQLConfig struct {
	Host     string `mapStructure:"host"`
	Port     int    `mapStructure:"port"`
	Username string `mapStructure:"username"`
	Password string `mapStructure:"password"`
	Database string `mapStructure:"database"`
}

func (c *MySQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Database)
}

type RedisConfig struct {
	Host     string `mapStructure:"host"`
	Port     int    `mapStructure:"port"`
	Password string `mapStructure:"password"`
	Database int    `mapStructure:"database"`
	PoolSize int    `mapStructure:"poolSize"`
	MaxIdle  int    `mapStructure:"maxIdle"`
}

func (c *RedisConfig) DSN() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

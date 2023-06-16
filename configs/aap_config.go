package configs

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
)

type AuthenticateAndPostConfig struct {
	Port  int           `yaml:"port"`
	MySQL mysql.Config  `yaml:"my_sql"`
	Redis redis.Options `yaml:"redis"`
}

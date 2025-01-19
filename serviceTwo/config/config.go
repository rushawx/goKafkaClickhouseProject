package config

import (
	"log"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

// DbConfig struct
type DbConfig struct {
	Db *gorm.DB
}

// Config struct
type Config struct {
	DbConf DbConfig
}

// NewConfig func
func NewConfig(dsn string) *Config {
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Open database failed: %v", err)
	}
	return &Config{DbConf: DbConfig{Db: db}}
}

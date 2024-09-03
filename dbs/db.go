package dbs

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Mysql    string = "Mysql"
	Postgres        = "Postgres"
	Redis           = "Redis"
)

// DBConfig 数据库连接配置
type DBConfig struct {
	Driver          string        `json:"driver" yaml:"driver"`                   // 数据库驱动
	Host            string        `json:"host" yaml:"host"`                       // 数据库地址
	Port            int           `json:"port" yaml:"port"`                       // 数据库端口
	User            string        `json:"user" yaml:"user"`                       // 数据库用户名
	Pass            string        `json:"pass" yaml:"pass"`                       // 数据库密码
	Name            string        `json:"name" yaml:"name"`                       // 数据库名称
	ConnMaxLifetime time.Duration `json:"connMaxLifetime" yaml:"connMaxLifetime"` // 连接最大存活时间
	ConnMaxIdleTime time.Duration `json:"connMaxIdleTime" yaml:"connMaxIdleTime"` // 连接最大空闲时间
	MaxOpenConns    int           `json:"maxOpenConns" yaml:"maxOpenConns"`       // 最大打开连接数
	MaxIdleConns    int           `json:"maxIdleConns" yaml:"maxIdleConns"`       // 最大空闲连接数
}

// dbConfig 数据库连接配置
func dbConfig(dbc DBConfig) (dbConfig gorm.Dialector) {
	// 数据库连接配置
	switch dbc.Driver {
	case Mysql:
		dbConfig = mysql.New(mysql.Config{
			DriverName: dbc.Driver,
			DSN:        fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbc.User, dbc.Pass, dbc.Host, dbc.Port, dbc.Name),
		})
	case Postgres:
		dbConfig = postgres.New(postgres.Config{
			DriverName: dbc.Driver,
			DSN:        fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", dbc.User, dbc.Pass, dbc.Host, dbc.Port, dbc.Name),
		})
	}

	return dbConfig
}

// NewDatabase 新建数据库连接
func NewDatabase(dbc DBConfig) (*gorm.DB, error) {
	// 创建数据库连接
	db, err := gorm.Open(dbConfig(dbc))
	if err != nil {
		return nil, err
	}

	// 设置连接池
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(dbc.MaxIdleConns)
	sqlDb.SetMaxOpenConns(dbc.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(dbc.ConnMaxLifetime)
	sqlDb.SetConnMaxIdleTime(dbc.ConnMaxIdleTime)

	return db, nil
}

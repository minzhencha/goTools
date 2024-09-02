package dbs

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DriverType int8

const (
	Mysql DriverType = iota + 0
	Postgres
)

// DriverTypeString 驱动类型字符串
var DriverTypeString = []string{
	0: "Mysql",
	1: "Postgres",
}

// String 驱动类型字符串
func (dt DriverType) String() string {
	return DriverTypeString[dt]
}

// DBConnection 数据库连接配置
type DBConnection struct {
	Driver          DriverType    `json:"driver" yaml:"driver"`                   // 数据库驱动
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
func dbConfig(conn DBConnection) (dbConfig gorm.Dialector) {
	// 数据库连接配置
	switch conn.Driver {
	case Mysql:
		dbConfig = mysql.New(mysql.Config{
			DriverName: conn.Driver.String(),
			DSN:        fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conn.User, conn.Pass, conn.Host, conn.Port, conn.Name),
		})
	case Postgres:
		dbConfig = postgres.New(postgres.Config{
			DriverName: conn.Driver.String(),
			DSN:        fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", conn.User, conn.Pass, conn.Host, conn.Port, conn.Name),
		})
	}

	return dbConfig
}

// ConnectDatabase 连接数据库
func ConnectDatabase(conn DBConnection) (*gorm.DB, error) {
	// 创建数据库连接
	db, err := gorm.Open(dbConfig(conn))
	if err != nil {
		return nil, err
	}

	// 设置连接池
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(conn.MaxIdleConns)
	sqlDb.SetMaxOpenConns(conn.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(conn.ConnMaxLifetime)
	sqlDb.SetConnMaxIdleTime(conn.ConnMaxIdleTime)

	return db, nil
}

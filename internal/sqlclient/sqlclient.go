package sqlclient

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	MYSQL = "mysql"
)

type ISqlClientConn interface {
	GetDB() *gorm.DB
	GetDriver() string
}

type SqlConfig struct {
	Driver       string
	Host         string
	Port         int
	Database     string
	Username     string
	Password     string
	Timeout      int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	PoolSize     int
	MaxIdleConns int
	MaxOpenConns int
}

type SqlClientConn struct {
	SqlConfig
	DB *gorm.DB
}

func NewSqlClient(config SqlConfig) ISqlClientConn {
	client := &SqlClientConn{}
	client.SqlConfig = config
	if err := client.Connect(); err != nil {
		log.Fatal(err)
		return nil
	}
	//ping to check connection in gorm
	sqlDB, err := client.DB.DB()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
		return nil
	}

	return client
}

func (c *SqlClientConn) Connect() error {
	switch c.Driver {
	case MYSQL:
		//username:password@protocol(address)/dbname?param=value
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&timeout=%ds", c.Username, c.Password, c.Host, c.Port, c.Database, c.Timeout)
		db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
			return err
		}
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal(err)
			return err
		}
		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
		c.DB = db
		return nil
	default:
		log.Fatal("driver is missing")
		return errors.New("driver is missing")
	}
}

func (c *SqlClientConn) GetDB() *gorm.DB {
	return c.DB
}

func (c *SqlClientConn) GetDriver() string {
	return c.Driver
}

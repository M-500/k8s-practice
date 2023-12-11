package db_helper

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 12:12
//

var dbEngine *xorm.Engine

type DbCfg struct {
	Engine          string `mapstructure:"engine"       json:"engine"`             // mysql
	Username        string `mapstructure:"username"      json:"username"`          // root
	Password        string `mapstructure:"password"      json:"password"`          // 123456
	Host            string `mapstructure:"host"          json:"host"`              // 127.0.0.1
	Port            int    `mapstructure:"port"          json:"port"`              // 3306
	Database        string `mapstructure:"database"      json:"database"`          // user_growth
	Charset         string `mapstructure:"charset"       json:"charset"`           // utf8mb4
	ShowSQL         bool   `mapstructure:"showSQL"       json:"showSQL"`           // false
	MaxIdleConns    int    `mapstructure:"maxIdleConns"  json:"maxIdleConns"`      // 2
	MaxOpenConns    int    `mapstructure:"maxOpenConns"  json:"maxOpenConns"`      // 10
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime" json:"connMaxLifetime"` // 30m
}

func SetUpDB(cfg *DbCfg) {
	if dbEngine != nil {
		return
	}
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
	)
	if engine, err := xorm.NewEngine(cfg.Engine, sourceName); err != nil {
		log.Fatalf("db_helper.InitDb(%s),%v\n", sourceName, err)
	} else {
		dbEngine = engine
	}

	if cfg.MaxIdleConns > 0 {
		dbEngine.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	if cfg.MaxOpenConns > 0 {
		dbEngine.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.ConnMaxLifetime > 0 {
		dbEngine.SetConnMaxLifetime(time.Minute * time.Duration(cfg.ConnMaxLifetime))
	}
}

func GetDb(cfg *DbCfg) *xorm.Engine {
	if dbEngine == nil {
		SetUpDB(cfg)
	}
	return dbEngine
}

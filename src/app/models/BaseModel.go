package models

import (
	"fmt"
	"github.com/nasanasaQuQ/goProject/src/pkg/casbin"
	"github.com/nasanasaQuQ/goProject/src/pkg/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/soft_delete"
	"os"
	"time"

	"log"
)

var db *gorm.DB

type BaseModel struct {
	Id         int64                 `gorm:"primary_key" json:"id"`
	UpdateTime time.Time             `json:"updateTime" gorm:"autoUpdateTime"`
	CreateTime time.Time             `json:"createTime" gorm:"autoCreateTime"`
	IsDel      soft_delete.DeletedAt `json:"isDel" gorm:"softDelete:flag"`
}

func Setup() {
	var err error

	var connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		global.CONFIG.Database.User,
		global.CONFIG.Database.Password,
		global.CONFIG.Database.Host,
		global.CONFIG.Database.Name)

	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 启用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Printf("[info gorm %s]", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("[info gorm %s]", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.DB = db

	casbin.InitCasbin(db)
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

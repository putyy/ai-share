package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

var db *gorm.DB

func init() {
	var (
		dbName, user, password, host string
	)

	dbName = config.Mysql.Database
	user = config.Mysql.User
	password = config.Mysql.Password
	host = config.Mysql.Host

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)

	dbCon, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	db = dbCon

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
}

func Db() *gorm.DB {
	return db
}

func PaginateScope(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = config.App.PageSize
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func CreatedWhereScope(form form.CommonForm) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if form.CreatedAtStart != "" {
			db = db.Where("created_at >= ?", form.CreatedAtStart)
		}
		if form.CreatedAtEnd != "" {
			db = db.Where("created_at <= ?", form.CreatedAtEnd)
		}
		return db
	}
}

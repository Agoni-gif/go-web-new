package initialize

import (
	"fmt"
	"go-web-new/global"
	"go-web-new/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func inItDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	var err error
	global.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		NamingStrategy: schema.NamingStrategy{
			//使用单数表名，启用该选项，此时`User` 的表名应为 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据库链接失败。", err)
	}

	if utils.TransferTable {
		// 迁移数据表
		_ = global.Db.AutoMigrate()
	}

	sqlDB, _ := global.Db.DB()
	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(10 * time.Second)

}

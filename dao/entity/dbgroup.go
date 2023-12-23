package entity

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBGroup struct {
	// 主库
	Master *gorm.DB
	// 从库
	Slave *gorm.DB
}

// DB 配置
func (dbGruop *DBGroup) NewDBGroup() {
	// 主库
	dbGruop.Master = dbGruop.InitMaster()

	// 从库
	dbGruop.Slave = dbGruop.InitSlave()
}

func (dbGroup *DBGroup) SelectDB(isMaster bool) *gorm.DB {
	if isMaster {
		return dbGroup.Master
	}
	return dbGroup.Slave
}

// 主从区分，db log
func (dbGruop *DBGroup) InitMaster() *gorm.DB {
	username := "test_public_2" //账号
	password := "txy-mysql-123" //密码
	host := "123.207.254.12"    //数据库地址，可以是Ip或者域名
	port := 3306                //数据库端口
	Dbname := "app"             //数据库名
	timeout := "10s"            //连接超时，10秒
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, Dbname, timeout,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}

// 主从区分，db log
func (dbGruop *DBGroup) InitSlave() *gorm.DB {
	username := "test_public_2" //账号
	password := "txy-mysql-123" //密码
	host := "123.207.254.12"    //数据库地址，可以是Ip或者域名
	port := 3306                //数据库端口
	Dbname := "app"             //数据库名
	timeout := "10s"            //连接超时，10秒
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, Dbname, timeout,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}

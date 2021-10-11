package model

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "time"
)

var DB *gorm.DB

func InitMySql(username, password, host, port, database string, debug bool) {
    dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        username, password, host, port, database,
    )
    db, err := gorm.Open("mysql", dsn)
    if nil != err {
        panic("init gorm failed: " + err.Error())
    }
    // Then you could invoke `*sql.DB`'s functions with it
    err = db.DB().Ping()
    if err != nil {
        panic(fmt.Sprintf("init gorm failed: %s, %s", dsn, err.Error()))
    }
    db.DB().SetMaxIdleConns(64)
    db.DB().SetConnMaxLifetime(time.Second * 300)
    db.DB().SetMaxOpenConns(128)
    // Disable table name's pluralization
    db.SingularTable(true)

    //迁移数据库
    db = db.AutoMigrate(

    )
    if debug {
        db = db.Debug()
    }
    DB = db
}

func Conn() *gorm.DB {
    return DB
}

type BaseModel struct {
    ID        uint `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt time.Time `gorm:"index" json:"-"`
}

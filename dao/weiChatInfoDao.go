package dao

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/mango-svip/wechat-spider/entity"
    "time"
)

func getConnection() *gorm.DB {
    db, err := gorm.Open("mysql", "root:####@tcp(####:####)/spider?charset=utf8mb4&parseTime=True&loc=utc")

    if err != nil {
        panic(err)
    }
    return db
}


func FindAll() {
    db := getConnection()
    defer db.Close()

    chatInfo := entity.WeiChatInfo{Name: "test", CreateTime: time.Now(), UpdateTime: time.Now() }

    db.Create(&chatInfo)

    info := entity.WeiChatInfo{}
    db.First(&info)

    fmt.Println(info)
}


func Create(info *entity.WeiChatInfo) {
    db := getConnection()
    defer db.Close()
    info.CreateTime = time.Now()
    info.UpdateTime = time.Now()
    db.Create(info)
}

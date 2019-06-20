package entity

import (
    "time"
)

type WeiChatInfo struct {
    Id int              `json:id`
    Name string         `json:"name"`
    Account string      `json:"account"`
    Description string  `json:"description"`
    MonthArticle int    `json:"monthArticle"`
    CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (WeiChatInfo) TableName() string {
    return "wei_chat_info"
}


package main

import (
    "github.com/mango-svip/wechat-spider/dao"
    "github.com/mango-svip/wechat-spider/spider"
)

func main()  {

    spider.Spider()
    dao.FindAll()

}



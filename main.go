package main

import (
    "encoding/json"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/mango-svip/wechat-spider/dto"
    "io"
    "net/http"
    "strings"
    "time"
)

const (
    baseUrl = "https://weixin.sogou.com/weixin"
    param = "type=1&s_from=input&query=java&ie=utf8&_sug_=n&_sug_type_="
    page = "&page=1"
)

func client() *http.Client {
    return &http.Client{}
}

func request(method string, url string, body io.Reader,header http.Header) *http.Request {
    newRequest, e := http.NewRequest(method, url, body)
    if e != nil {
        panic(e)
    }
    for k, v := range header {
        newRequest.Header.Add(k, v[0])
    }
    return newRequest
}

func resp(resp *http.Response, err error) {
    if err != nil {
        panic(nil)
    }

    defer resp.Body.Close()

    document, err := goquery.NewDocumentFromReader(resp.Body)
    fmt.Println(document.Html())
    if err != nil {
        panic(err)
    }

    children := document.Find("ul.news-list2").Children()

    infos := make([]dto.WeiChatInfo, 0)

    children.Each( func(i int, s *goquery.Selection) {
        txtBox := s.Find("div.txt-box")
        account := txtBox.Find("p.info").Find("label").Text()
        name := strings.Trim(txtBox.Find("p.tit").Text(), "\n")

        description := s.Find("dl dd").Eq(0)

        info := dto.WeiChatInfo{
            Name:        name,
            Account:     account,
            Description: description.Text(),
        }

        infos = append(infos, info)

    })

    bytes, e := json.Marshal(infos)
    if e != nil {
        panic(e)
    }

    fmt.Println( string(bytes))
}


func header(baseUrl string, param string) http.Header {
    return http.Header{
        "User-Agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.90 Safari/537.36"},
        "Referer":    []string{baseUrl },
    }
}

func main()  {

    var  baseUrl, param = baseUrl, param
    page := ""
    client := client()
    for i := 1; i < 11; i ++ {
        page = fmt.Sprintf("&page=%d", i)
        fmt.Println(page)

        req := request("GET", baseUrl+"?"+param + page, nil, header(baseUrl, param + page))

        resp(client.Do(req));

        time.Sleep(1000 * time.Millisecond)
    }


}



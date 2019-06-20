package spider

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/mango-svip/wechat-spider/dao"
    "github.com/mango-svip/wechat-spider/entity"
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

func resp(resp *http.Response, err error) []entity.WeiChatInfo {
    if err != nil {
        panic(nil)
    }

    defer resp.Body.Close()

    document, err := goquery.NewDocumentFromReader(resp.Body)
    //fmt.Println(document.Html())
    if err != nil {
        panic(err)
    }

    children := document.Find("ul.news-list2").Children()

    infos := make([]entity.WeiChatInfo, 0)

    children.Each( func(i int, s *goquery.Selection) {
        txtBox := s.Find("div.txt-box")
        account := txtBox.Find("p.info").Find("label").Text()
        name := strings.Trim(txtBox.Find("p.tit").Text(), "\n")

        description := s.Find("dl dd").Eq(0)

        info := entity.WeiChatInfo{
            Name:        name,
            Account:     account,
            Description: description.Text(),
        }

        infos = append(infos, info)

    })

    //bytes, e := json.Marshal(infos)
    //if e != nil {
    //    panic(e)
    //}
    //
    //fmt.Println( string(bytes))
    return infos
}

func header(baseUrl string, param string) http.Header {
    return http.Header{
        "User-Agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.90 Safari/537.36"},
        "Referer":    []string{baseUrl },
    }
}


/**
 *   接口访问频率有限制， 有待解决验证码
 */
func Spider() {

    var  baseUrl, param = baseUrl, param
    page := ""
    client := client()
    for i := 1; i < 11; i ++ {
        page = fmt.Sprintf("&page=%d", i)
        //fmt.Println(page)
        req := request("GET", baseUrl+"?"+param + page, nil, header(baseUrl, param + page))
        for k,v := range resp(client.Do(req)) {
            fmt.Println(k,v)
            dao.Create(&v)
        }



        time.Sleep(1000 * time.Millisecond)
    }
}



package main

import(
    "flag"
    "net/http"
    "io/ioutil"
    "log"
    "regexp"
    "net/url"
)

// http://seagma.testmart.cn/
var starturl string
var found map[string]bool

func init() {
    flag.StringVar(&starturl, "starturl", "", "crawl from the start url")
    found = make(map[string]bool, 100)
}

// 抓取某个网站的所有link
func main() {
    flag.Parse()

    parseLinks(starturl)
}

func parseLinks(url string) {
    links := getDomainLinks(url)

    for _,link := range links {
        if _,ok := found[link]; !ok {
            found[link] = true
            log.Println(link)
            parseLinks(link)
        }
    }
}

func removeDuplicate(slis *[]string) {
    found := make(map[string]bool)
    j := 0
    for i,val := range *slis {
        if _,ok := found[val]; !ok {
            found[val] = true
            (*slis)[j] = (*slis)[i]
            j++
        }
    }
    *slis = (*slis)[:j]
}

func getDomainLinks(url string) []string {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    // 正则匹配
    regex := `<a[\s\S]*?href=["']([\s\S]*?)["'][\s\S]*?>[\s\S]*?</a>`
    matches := regexp.MustCompile(regex).FindAllSubmatch(body, -1);

    ret := make([]string, 0)
    rootUrl := getRootUrl(url)
    for _,match := range matches {
        if string(match[1][0:4]) == "http" {
            if getRootUrl(string(match[1])) == rootUrl {
                ret = append(ret, string(match[1]))
            }
        }
        if string(match[1][0:1]) == "/" {
            ret = append(ret, rootUrl + string(match[1]))
        }
    }
    return ret
}

// 返回根url
func getRootUrl(urlstr string) string {
    u, _ := url.Parse(urlstr)
    return u.Scheme + "://" + u.Host
}


/*
yejianfengtekiMacBook-Air:webscaner yejianfeng$ go run main.go -starturl=http://seagma.testmart.cn
2013/03/27 00:44:20 http://seagma.testmart.cn/?username=seagma&channel=1
2013/03/27 00:44:20 http://seagma.testmart.cn/?username=seagma&channel=2
2013/03/27 00:44:20 http://seagma.testmart.cn/?username=seagma&channel=5
2013/03/27 00:44:20 http://seagma.testmart.cn/?username=seagma&channel=9
2013/03/27 00:44:20 http://seagma.testmart.cn/?username=seagma&channel=10
2013/03/27 00:44:21 http://seagma.testmart.cn/?username=seagma&channel=4
2013/03/27 00:44:21 http://seagma.testmart.cn/?username=seagma&channel=14
2013/03/27 00:44:21 http://seagma.testmart.cn
2013/03/27 00:44:21 http://seagma.testmart.cn/detail_190531292.htm
2013/03/27 00:44:21 http://seagma.testmart.cn/
2013/03/27 00:44:22 http://seagma.testmart.cn/detail_158991292.htm
2013/03/27 00:44:22 http://seagma.testmart.cn/detail_252091292.htm
2013/03/27 00:44:22 http://seagma.testmart.cn/detail_479471292.htm
*/
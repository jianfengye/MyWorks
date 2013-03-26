package main

import(
    "flag"
    "net/http"
    "ioutil"
    "log"
    "regexp"
)

const(
    starturl string
)

func init() {
    flag.StringVar(&starturl, "starturl", "", "crawl from the start url")
}

// 抓取某个网站的所有link
func main() {
    flag.Parse()

    getDomainLinks(starturl)
}

func getDomainLinks(url string) []string {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    // 正则匹配
    regex := `<a[\s\S]*href="[\s\S]*"[\s\S]*></a>`
    matches := regexp.MustCompile(regex).findAll(body, -1);
    for match := range matches {
        log.Println(string(match))
    }
    return nil
}
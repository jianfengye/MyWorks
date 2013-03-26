package main

import(
    "net/url"
    "fmt"
)

func main() {
    u, _ := url.Parse("http://bing.com/search?q=dotnet")

    fmt.Println(u.Query())

    fmt.Println(u.RequestURI())

    fmt.Println(u.Scheme + "//" + u.Host)

    /*
        map[q:[dotnet]]
        /search?q=dotnet
        http//bing.com
    */
}
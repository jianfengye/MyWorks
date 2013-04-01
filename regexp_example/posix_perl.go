package main

import(
    "regexp"
    "fmt"
)

func main() {
    var origin1 = `tt`
    var reg = `t\z`

    matches1 := regexp.MustCompile(reg).FindString(origin1)
    fmt.Println(matches1)

/*
    matches2 := regexp.MustCompilePOSIX(reg).FindString(origin1)
    fmt.Println(matches2)
*/
}
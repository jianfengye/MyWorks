package main

import(
    "regexp"
    "fmt"
)

func main() {
    var origin1 = `<a href="test"></a><a href="test"></a>`
    var reg = `<a[\s\S]*?href="([\s\S]*?)"[\s\S]*?></a>`

    matches1 := regexp.MustCompile(reg).FindString(origin1)
    fmt.Println(matches1)

    matches2 := regexp.MustCompile(reg).FindAllString(origin1, 0)
    fmt.Println(matches2)

    matches3 := regexp.MustCompile(reg).FindAllString(origin1, 1)
    fmt.Println(matches3)

    matches4 := regexp.MustCompile(reg).FindAllStringIndex(origin1, 1)
    fmt.Println(matches4)

    matches5 := regexp.MustCompile(reg).FindAllStringSubmatch(origin1, -1)
    fmt.Println(matches5)

    matches6 := regexp.MustCompile(reg).FindAllStringSubmatchIndex(origin1, 2)
    fmt.Println(matches6)

    matches7 := regexp.MustCompile(reg).FindStringSubmatch(origin1)
    fmt.Println(matches7)

    matches8 := regexp.MustCompile(reg).FindStringSubmatchIndex(origin1)
    fmt.Println(matches8)

    fmt.Println(regexp.MustCompile(reg).SubexpNames())

    fmt.Println(regexp.MustCompile(reg).String())

    fmt.Println(regexp.MustCompile(reg).NumSubexp())
}
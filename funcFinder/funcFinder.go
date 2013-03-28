package main

import(
	"flag"
	"os"
	"io/ioutil"
	"regexp"
	"strings"
	"fmt"
)

// php\src\application\controllers
var folder string

func init() {
    flag.StringVar(&folder, "folder", "", "the folder which contain controller")
}

func main() {
	flag.Parse()

	// 遍历folder下的文件
	os.Chdir(folder)
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		filename := file.Name()
		index := strings.Index(filename, "Controller")
		if index == -1 {
			continue
		}

		controller := strings.ToLower(filename[0:1]) + filename[1:index]

		// 读取php文件获取action
		// example：
		//public function checkAction() {
		content, _ := ioutil.ReadFile(filename)
		regstr := `public function ([\w]*?)Action\(\)`
		matches := regexp.MustCompile(regstr).FindAllSubmatch(content, -1);
		for _,match := range matches {
			action := string(match[1])
			fmt.Println(controller + "/" + action)
		}
	}
}
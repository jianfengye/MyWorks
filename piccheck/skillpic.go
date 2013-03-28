package main

import(
	"io/ioutil"
	"strings"
	"os"
)

// 将目录中的所有文件名前面的0去掉
func main() {
	skillfolder := `C:\Users\yejianfeng\Desktop\mxm\skill\`
	os.Chdir(skillfolder)
	// 获取所有文件
	files, _ := ioutil.ReadDir(skillfolder)
	for _,file := range files {
		if file.IsDir() {
			continue
		} else {
			name := file.Name()
			os.Rename(name, strings.TrimLeft(name, "0"))
		}
	}
}

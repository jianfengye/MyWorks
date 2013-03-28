package main

import(
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	os.Chdir(`C:\Users\yejianfeng\Desktop\pchelper\`)

	file := "XXXX"

	apiFile, _ := os.Open(file + "_api")
	defer apiFile.Close()

	callFile, _ := os.Open(file + "_call")
	defer callFile.Close()

	calls := make(map[string]int, 1)
	reader := bufio.NewReader(callFile)
	for {
		if api, _, err := reader.ReadLine(); err == nil {
			apistr := strings.Trim(string(api), " ")
			splits := strings.Split(apistr, " ")
			if len(splits) != 2 {
				continue;
			} else {
				num, _ := strconv.Atoi(splits[0])
				if _, ok := calls[strings.ToLower(splits[1])]; ok {
					calls[strings.ToLower(splits[1])] = num + calls[strings.ToLower(splits[1])]
				} else {
					calls[strings.ToLower(splits[1])] = num
				}
				
			}
		} else {
			break
		}
	}

	reader = bufio.NewReader(apiFile)
	for {
		if api, _, err := reader.ReadLine(); err == nil {
			apistr := "/" + string(api)

			if _, ok := calls[strings.ToLower(apistr)]; ok {
				fmt.Println(apistr, calls[strings.ToLower(apistr)])
			} else {
				fmt.Println(apistr, 0)
			}
		} else {
			break
		}
		
	}
}

package main

import(
	"encoding/csv"
	"os"
	"fmt"
	"io"
	"strings"
)

func main() {
	filecsv := `C:\Users\yejianfeng\Desktop\mxm\role.csv`
	file, err := os.Open(filecsv)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	roleFolder := `C:\Users\yejianfeng\Desktop\mxm\role\`
	reader.Read()
	for {
		fields, err := reader.Read()
		if err == io.EOF {
			break;
		}

		picurl := fields[19]
		// PET/petic/41.png
		picnames := strings.Split(picurl, "/")
		picurl = roleFolder + picnames[2]

		_, err = os.Open(picurl)
		if err == nil || os.IsExist(err) {
			continue
		}

		fmt.Println(fields[0])
		fmt.Println(fields[19])
	}


}

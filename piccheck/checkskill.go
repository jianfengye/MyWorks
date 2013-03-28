package main

import(
	"encoding/csv"
	"os"
	"fmt"
	"io"
)

func main() {
	filecsv := `C:\Users\yejianfeng\Desktop\mxm\skill.csv`
	file, err := os.Open(filecsv)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	roleFolder := `C:\Users\yejianfeng\Desktop\mxm\skill\`
	reader.Read()
	for {
		fields, err := reader.Read()
		if err == io.EOF {
			break;
		}

		picurl := fields[15]

		jpg := roleFolder + picurl + ".jpg"
		_, err = os.Open(jpg)
		if err == nil || os.IsExist(err) {
			continue
		}

		png := roleFolder + picurl + ".png"
		_, err = os.Open(png)
		if err == nil || os.IsExist(err) {
			continue
		}

		fmt.Println(fields[0])
	}
}

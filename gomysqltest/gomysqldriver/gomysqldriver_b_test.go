package gomysqldriver

import (
	"testing"
	"fmt"
)

func Benchmark_getAdmin(b *testing.B){
	for i := 0; i < b.N; i++ { //use b.N for looping 
       	getAdmin(1)
    }
}

func Benchmark_insertAdmin(b *testing.B) {
	var username string
	var password string
	var adminid int
	for i := 0; i < b.N; i++ {
		username = fmt.Sprintf("admin%v", adminid)
		password = fmt.Sprintf("password%v", adminid)
		insertAdmin(username, password)
	}
}

func Benchmark_updateAdmin(b *testing.B) {
	var password string
	for i := 0; i < b.N; i++ {
		password = fmt.Sprintf("password%v", i+2)
		updateAdmin(1, password)
	}
}

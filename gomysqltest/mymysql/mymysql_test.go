package mymysql

import(
	"testing"
)

func Test_getAdmin(t *testing.T) {
    username, _ := getAdmin(1)

    if (username != "admin") {
    	t.Error("getAdmin get data error")
    }
}


func Test_insertAdmin(t *testing.T) {
	username2 := "admin2"
	password2 := "password2"

	insertAdmin(username2, password2);
}

func Test_updateAdmin(t *testing.T) {
	adminid := 1
	password2 := "password2"

	updateAdmin(adminid, password2)

	_, password := getAdmin(adminid)
	if (password != password2) {
		t.Error("updateAdmin error")
	}
}


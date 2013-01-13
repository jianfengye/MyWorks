package gomysqldriver

import (
	"log"
	"database/sql"
	_"code.google.com/p/go-mysql-driver/mysql"
)

var (
	dbusername = "root"
	dbpassowrd = "t00R"
	dbname = "test"
)

func getAdmin(adminid int)(string, string) {
	db, err := sql.Open("mysql", dbusername + ":" + dbpassowrd + "@tcp(127.0.0.1:3306)/" + dbname)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var id int
	var username string
	var password string
	err = db.QueryRow("select * from admin where adminid=?", adminid).Scan(&id, &username, &password)
	if err != nil {
		log.Fatalln(err)
	}
	
	return username,password
}

func insertAdmin(username, password string) {
	db, err := sql.Open("mysql", dbusername + ":" + dbpassowrd + "@tcp(127.0.0.1:3306)/" + dbname)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into admin set username=?,password=?", username, password);
	if err != nil {
		log.Fatalln(err)
	}
}

func updateAdmin(adminid int, password string) {
	db, err := sql.Open("mysql", dbusername + ":" + dbpassowrd + "@tcp(127.0.0.1:3306)/" + dbname)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	_, err = db.Exec("update admin set password=? where adminid=?", password, adminid);
	if err != nil {
		log.Fatalln(err)
	}
}

func deleteAdmin(adminid int) {
	db, err := sql.Open("mysql", dbusername + ":" + dbpassowrd + "@tcp(127.0.0.1:3306)/" + dbname)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	_, err = db.Exec("delete from admin where adminid=?", adminid);
	if err != nil {
		log.Fatalln(err)
	}
}

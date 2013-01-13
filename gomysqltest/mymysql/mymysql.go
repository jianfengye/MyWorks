package mymysql

import(
	"log"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

var (
	dbusername = "root"
	dbpassowrd = "t00R"
	dbname = "test"
)

func getAdmin(adminid int) (string, string){
	db := mysql.New("tcp", "", "127.0.0.1:3306", dbusername, dbpassowrd, dbname)

	err := db.Connect()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close();

	rows, res, err := db.Query("select * from admin where adminid=%d", adminid)
	if err != nil {
		log.Panic(err)
	}

	if len(rows) < 1 {
		log.Panic("rows error")
	}

	row := rows[0]

	first := res.Map("username")
	second := res.Map("password")

	username, password := row.Str(first), row.Str(second)
	return username, password
}

func insertAdmin(username,password string) {
	db := mysql.New("tcp", "", "127.0.0.1:3306", dbusername, dbpassowrd, dbname)

	err := db.Connect()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close();

	stmt, err := db.Prepare("insert into admin set username=?,password=?")
	if err != nil {
		log.Panic(err)
	}

	_, _, err = stmt.Exec(username, password)
	if err != nil {
		log.Panic(err)
	}
}

func updateAdmin(adminid int, password string) {
	db := mysql.New("tcp", "", "127.0.0.1:3306", dbusername, dbpassowrd, dbname)

	err := db.Connect()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close();

	stmt, err := db.Prepare("update admin set password=? where adminid=?")
	if err != nil {
		log.Panic(err)
	}

	_, err = stmt.Run(password, adminid)
}

func deleteAdmin(adminid int) {
	db := mysql.New("tcp", "", "127.0.0.1:3306", dbusername, dbpassowrd, dbname)

	err := db.Connect()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close();

	stmt, err := db.Prepare("delete admin where adminid=?")
	if err != nil {
		log.Panic(err)
	}

	_, err = stmt.Run(adminid)
}
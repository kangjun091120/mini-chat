package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() {
	database, err := sqlx.Open("mysql", "root:Admin@tcp(127.0.0.1:3306)/chatter")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	db = database
}

func Checkuser(username string) bool {
	type User struct {
		Userid   int    `db:"userid"`
		Username string `db:"username"`
		Password string `db:"password"`
	}
	var user User
	_ = db.Get(&user, "SELECT * FROM user WHERE  username = ?", username)
	if user.Userid == 0 {
		return false
	} else {
		return true
	}
}
func Register(username string, password string) {
	_, _ = db.Exec("insert into user(username, password)values (?,?)", username, password)
}

func Login(username string, password string) int {
	type User struct {
		Userid   int    `db:"userid"`
		Username string `db:"username"`
		Password string `db:"password"`
	}
	var user User
	_ = db.Get(&user, "SELECT * FROM user WHERE  username = ? AND password = ?", username, password)
	return user.Userid
}

package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Uid   int64  `gorm:"primary_key;column:U_id"`
	Rid   int64  `gorm:"column:R_id"`
	Upwd  string `gorm:"column:U_pwd"`
	Uname string `gorm:"column:U_name"`
	Uage  int    `gorm:"column:U_age"`
	Usex  string `gorm:"column:U_sex"`
	Utel  string `gorm:"column:U_tel"`
}

func main() {
	//	gorm 连接数据库
	db, err := gorm.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/supermanager")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)
	//	指定表名
	//var user = User{
	//	Rid:   1,
	//	Upwd:  "password",
	//	Uname: "钱七",
	//	Uage:  20,
	//	Usex:  "男",
	//	Utel:  "1048562948",
	//}
	//create := db.Table("t_user").Create(&user)
	//fmt.Println("受影响行数", create.RowsAffected)
	var users []User
	find := db.Table("t_user").Find(&users)
	fmt.Println(find.RowsAffected)
	fmt.Println(users)

}

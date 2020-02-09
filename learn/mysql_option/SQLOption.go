package main

import (
	"database/sql"
	"fmt"

	//引入mysql的驱动
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id   int64
	name string
	age  int16
	sex  string
}

var db *sql.DB

//初始化连接
func initConn() {
	open, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/learn")
	if err != nil {
		println(err)
		return
	}
	//	尝试建立连接
	err = open.Ping()
	if err != nil {
		println(err)
		return
	}
	//设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。
	//如果n<=0，不会限制最大开启连接数，默认为0（无限制）。
	open.SetMaxIdleConns(80)
	//设置连接池中的最大闲置连接数。 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。 如果n<=0，不会保留闲置连接。
	open.SetMaxOpenConns(100)
	//	这里不能直接将open替换成db，底层的地址是不同的。
	db = open
}

//单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。
//QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。
func selectUser() {
	var user User
	querySql := "SELECT user_id,user_name,user_age,user_sex FROM learn_user where user_id = ?"
	//	单行查询
	err := db.QueryRow(querySql, 1).Scan(&user.id, &user.name, &user.age, &user.sex)
	if err != nil {
		println(err)
		return
	}
	fmt.Println(user)
}

//多行查询db.Query()执行一次查询，返回多行结果（即Rows），一般用于执行select命令。
//参数args表示query中的占位参数。
func selectUserList() {
	queryListSql := "SELECT user_id,user_name,user_age,user_sex FROM learn_user"
	query, err := db.Query(queryListSql)
	if err != nil {
		println(err)
		return
	}
	defer query.Close()
	for query.Next() {
		var user User
		err = query.Scan(&user.id, &user.name, &user.age, &user.sex)
		if err != nil {
			println(err)
			return
		}
		fmt.Println(user)
	}
}

// 插入数据
func insertUser() {
	addSql := "insert into learn_user(user_name,user_age,user_sex) values (?,?,?)"
	exec, err := db.Exec(addSql, "王五", 30, "男")
	if err != nil {
		println(err)
		return
	}
	//	新插入的id
	fmt.Println(exec.LastInsertId())
}

//更新数据
func updateUser() {
	modifySql := "update learn_user set user_name=?,user_age=? where user_id = ?"
	exec, err := db.Exec(modifySql, "赵六", 18, 5)
	if err != nil {
		println(err)
		return
	}
	//	操作影响的行数
	n, err := exec.RowsAffected()
	if err != nil {
		println(err)
	}
	fmt.Println(n)
}

//删除数据
func deleteUser() {
	deleteSql := "delete from learn_user where user_id = ?"
	exec, err := db.Exec(deleteSql, 4)
	if err != nil {
		println(err)
		return
	}
	//	操作影响的行数
	n, err := exec.RowsAffected()
	if err != nil {
		println(err)
	}
	fmt.Println(n)
}
func init() {
	initConn()
}

//	预处理的好处
//	优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
//	避免SQL注入问题。
func prepareQuery() {
	querySQL := "SELECT user_id,user_name,user_age,user_sex FROM learn_user"
	prepare, err := db.Prepare(querySQL)
	if err != nil {
		println(err)
		return
	}
	defer prepare.Close()
	query, err := prepare.Query()
	if err != nil {
		println(err)
		return
	}
	defer query.Close()
	for query.Next() {
		var user User
		err = query.Scan(&user.id, &user.name, &user.age, &user.sex)
		if err != nil {
			println(err)
			return
		}
		fmt.Println(user)
	}
}

//	其他操作操作也是这样
func prepareInsert() {
	addSql := "insert into learn_user(user_name,user_age,user_sex) values (?,?,?)"
	prepare, err := db.Prepare(addSql)
	if err != nil {
		println(err)
		return
	}
	defer prepare.Close()
	exec, err := prepare.Exec("钱七", 90, "男")
	if err != nil {
		println(err)
		return
	}
	n, err := exec.RowsAffected()
	if err != nil {
		println(err)
		return
	}
	fmt.Println(n)
}

func main() {
	//selectUser()
	//selectUserList()
	//insertUser()
	//updateUser()
	//deleteUser()
	//prepareQuery()
	//prepareInsert()
}

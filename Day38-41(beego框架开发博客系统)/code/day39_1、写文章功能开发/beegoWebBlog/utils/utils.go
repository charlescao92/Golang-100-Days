package utils

import (
	"github.com/beego/beego/v2/core/config"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"crypto/md5"
)

var db *sql.DB

func InitMysql() {

	fmt.Println("InitMysql....")
	driverName, _ := config.String("driverName")

	//注册数据库驱动
	//orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user, _ := config.String("mysqluser")
	pwd, _ := config.String("mysqlpwd")
	host, _ := config.String("host")
	port, _ := config.String("port")
	dbname, _ := config.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/myblog?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
		CreateTableWithArticle()
	}
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

//传入的数据不一样，那么MD5后的32位长度的数据肯定会不一样
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

//创建文章表
func CreateTableWithArticle(){
	sql:=`create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}
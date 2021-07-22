package model

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func DBInit() error {
	var err error
	db, err = sql.Open("mysql", "root:123@tcp(localhost:3306)/geek")
	if err != nil {
		log.Panicf(" 数据库连接错误，errors: %+v  ",err)
		return err
	}

	if err = db.Ping(); err != nil {
		log.Printf("数据库连接失败 , errors: %+v  ",err)
		return err
	}
	return nil
}

type User struct {
	id  int
	name string
}

type errorString struct {
	err error
	errStr string
}

// db errors
func (e *errorString)dbErr() error {
	if e.err == sql.ErrNoRows {
		return errors.New(e.err.Error())
	}else{
		return e.err
	}
}

// 获取一条记录 By id
func GetUser(id int) (*User,error)  {
	if err := DBInit();err !=nil{
		return nil,err
	}
	defer db.Close()

	var user User
	err := db.QueryRow("SELECT id,name FROM users WHERE id = ?",id).Scan(&user.id,&user.name)
	if err != nil {
		var e errorString
		e.err = err
		return nil,e.dbErr()
	}
	return &user,nil
}





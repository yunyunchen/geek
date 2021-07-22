package main

import (
	"geek/model"
	"log"
)

var ErrNoRows = "sql: no rows in result set"

//
// 	我觉得不需要单独处理sql.ErrNoRows也可以
//	代码的实现方案是，重新定义的一个错误变量，来返回提醒，不使用"database/sql"库自带的，主要考虑业务层不导入sql包
//	errorString ，这个结构体，我觉得可以把mysql的错误，分几个等级来封包返回，暂时没有成熟的实现想法
//	如果在项目开发中，我更偏向于，返回的User为nil，判断 err !=nil,sql错误。判断数据集user==nil，表示数据结果为空


func main()  {
	// 获取一条记录
	user ,err :=model.GetUser(2)

	if err != nil{
		if err.Error() == ErrNoRows {
			log.Println("数据为空")
		}else{
			log.Println(err.Error())
		}
	}else{
		log.Println(user)
	}
}


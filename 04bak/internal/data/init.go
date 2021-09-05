package data

//import "fmt"
//import "github.com/go-gorm/gorm"
//
//var (
//	DB *gorm.DB
//	err error
//)
//
//func Init() {
//	DB, err = gorm.Open(
//		"mysql",
//		fmt.Sprintf(
//			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
//			config.Mysql.Username,
//			config.Mysql.Password,
//			config.Mysql.Host,
//			config.Mysql.Port,
//			config.Mysql.Database,
//		),
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	//设置连接池
//	DB.DB().SetMaxIdleConns(10)
//	DB.DB().SetMaxOpenConns(100)
//
//	//数据库迁移
//	migrate()
//}
//
//func migrate()  {
//	DB.AutoMigrate(&User{})
//}
//
//func Close()  {
//	defer DB.Close()
//}

package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySqlMaster *gorm.DB
var MysqlDsn string

func GormSql() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	MysqlDsn = "LogServer:LogServerpwd@tcp(127.0.0.1:3306)/logserver?charset=utf8mb4&parseTime=True&loc=Local"

	sql, err := gorm.Open(mysql.Open(MysqlDsn), &gorm.Config{})
	MySqlMaster = sql
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//自动判断是否存在表 并创建表
	e := MySqlMaster.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&LogType{}, &otherInfo{}, &requst{}, &HeadInfo{})
	if e != nil {
		fmt.Println(e)
	}
	//MySqlMaster.Create(&LogType{})

	//lo := []LogType{}
	//
	//MySqlMaster.Find(&lo, []int{1, 2})
	//for i, logType := range lo {
	//	fmt.Println(i, logType)
	//}
}

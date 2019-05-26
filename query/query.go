package query

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"model_generate/config"
	"os"
)

type Table struct {
	TableName string `gorm:"column:table_name"` //table name
}

type Column struct {
	ColumnNumber int    `gorm:"column_number"` // column index
	ColumnName   string `gorm:"column_name"`   // column_name
	ColumnType   string `gorm:"column_type"`   // column_type
}

var DB *gorm.DB

func InitDB() {
	dataSouce := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s",
		config.Conf.UserName,
		config.Conf.PassWord,
		config.Conf.Host,
		config.Conf.Port,
		config.Conf.Database,
		"charset="+config.Conf.Charset+"&parseTime=True&loc=Local",
	)
	var err error
	DB, err = gorm.Open("mysql", dataSouce)
	DB.SingularTable(true)
	if err != nil {
		fmt.Println("数据库连接错误")
		os.Exit(0)
	}
}

func FindTables() []Table {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(fmt.Sprintf("recover from a fatal error : %v", e))
		}
	}()
	var tables = make([]Table, 0, 10)
	sql := `SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE()`
	if config.Conf.Table != "ALL" {
		sql += " and table_name = '" + config.Conf.Table + "'"
	}
	DB.Raw(sql).Find(&tables)
	return tables
}

func FindColumns(tableName string) []Column {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(fmt.Sprintf("recover from a fatal error : %v", e))
		}
	}()

	var columns = make([]Column, 0, 10)
	sql := `SELECT 
	column_name, 
	data_type as column_type, 
	ordinal_position as column_number 
FROM 
	information_schema.COLUMNS 
WHERE  
	table_name = ? and 
	table_schema = DATABASE()
	`
	DB.Raw(sql, tableName).Find(&columns)
	return columns
}

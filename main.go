package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"model_generate/config"
	"model_generate/gernerate"
	"model_generate/query"
)

func main() {
	config.Parse()

	query.InitDB()

	gernerate.Run()
}

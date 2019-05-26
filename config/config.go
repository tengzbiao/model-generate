package config

import (
	"flag"
	"fmt"
	"os"
)

var Conf = &config{}

type config struct {
	UserName string
	PassWord string
	Host     string
	Port     string
	Database string
	Charset  string
	Table    string
	Path     string
}

func Parse() {
	flag.StringVar(&Conf.UserName, "u", "root", "用户名")
	flag.StringVar(&Conf.PassWord, "p", "", "密码")
	flag.StringVar(&Conf.Host, "h", "localhost", "host")
	flag.StringVar(&Conf.Port, "P", "3306", "port")
	flag.StringVar(&Conf.Database, "d", "", "数据库")
	flag.StringVar(&Conf.Charset, "c", "utf8mb4", "数据库编码")
	flag.StringVar(&Conf.Table, "t", "ALL", "表名")
	flag.StringVar(&Conf.Path, "path", "./models", "表名")
	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `
Usage: COMMAND [-h host] [-P port] [-u user] [-p password] [-d database] [-c charset] [-t table] [-path path]

Options:
`)
	flag.PrintDefaults()
}

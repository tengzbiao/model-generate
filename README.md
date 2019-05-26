### gorm model生成工具

- go get github.com/tengzbiao/model-generate
- cd $GOPATH/src/github.com/tengzbiao/model-generate;go build -o $GOPATH/bin/model-generate main.go

```
Usage: $GOPATH/bin/model-generate [-h host] [-P port] [-u user] [-p password] [-d database] [-c charset] [-t table] [-path path]

Options:
  -h string
    	host (default "localhost")
  -u string
    	用户名 (default "root")
  -P string
    	port (default "3306")
  -p string
    	密码
  -d string
    	数据库
  -t string
    	表名 (default "ALL")
  -path string
    	表名 (default "./models")
  -c string
    	数据库编码 (default "utf8mb4")
```

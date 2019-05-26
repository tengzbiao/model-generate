package gernerate

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/inflection"
	"io"
	"model_generate/config"
	. "model_generate/helper"
	"model_generate/query"
	"os"
	"os/exec"
	"strings"
)

func Run() {
	for _, t := range query.FindTables() {
		// 生成带json tag的结构体
		goModelWithTag := tableToStructWithTag(t.TableName)

		fileName := GetFile(config.Conf.Path, t.TableName)
		f, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		_, err = io.WriteString(f, goModelWithTag)
		if err != nil {
			panic(err)
		}
	}
	_ = exec.Command("gofmt", "-l", "-w", config.Conf.Path).Start()
}

func tableToStruct(tableName string) (string, bool) {
	columnString := ""
	tmp := ""
	columns := query.FindColumns(tableName)

	var timeImport bool

	for _, column := range columns {
		columnName := column.ColumnName
		typeConvert := typeConvert(column.ColumnType)
		tmp = fmt.Sprintf("    %s  %v\n", inflection.Singular(UnderLineToHump(columnName)), typeConvert)
		columnString = columnString + tmp
		if typeConvert == "time.Time" {
			timeImport = true
		}
	}

	rs := fmt.Sprintf("type %s struct{\n%s}", inflection.Singular(UnderLineToHump(HumpToUnderLine(tableName))), columnString)
	return rs, timeImport
}

func tableToStructWithTag(tableName string) string {
	str, timeImport := tableToStruct(tableName)

	var result string
	scanner := bufio.NewScanner(strings.NewReader(str))
	var oldLineTmp = ""
	var lineTmp = ""
	var propertyTmp = ""
	var seperateArr []string
	for scanner.Scan() {
		oldLineTmp = scanner.Text()
		lineTmp = strings.Trim(scanner.Text(), " ")
		if strings.Contains(lineTmp, "{") || strings.Contains(lineTmp, "}") {
			result = result + oldLineTmp + "\n"
			continue
		}
		seperateArr = Split(lineTmp, " ")
		// 接口或者父类声明不参与tag, 自带tag不参与tag
		if len(seperateArr) == 1 || len(seperateArr) == 3 {
			continue
		}
		propertyTmp = HumpToUnderLine(seperateArr[0])
		oldLineTmp = oldLineTmp + fmt.Sprintf("`gorm:\"%s\" json:\"%s\"`", propertyTmp, propertyTmp)
		result = result + oldLineTmp + "\n"
	}
	if timeImport {
		result = fmt.Sprintf("import \"time\" \n\n") + result
	}
	return fmt.Sprintf("package %s \n\n", GetPackageName(config.Conf.Path)) + result
}

// 类型转换mysql->go
func typeConvert(s string) string {
	if In(s, []string{"int", "bigint", "smallint", "tinyint", "mediumint"}) {
		return "int"
	}
	if In(s, []string{"decimal", "real", "double", "float"}) {
		return "float64"
	}
	if In(s, []string{"tinyblob", "blob", "mediumblob", "longblob"}) {
		return "[]byte"
	}
	if In(s, []string{"date", "timestamp", "datetime", "time", "year"}) {
		return "time.Time"
	}
	if In(s, []string{"char", "varchat", "tinytext", "text", "mediumtext", "longtext"}) {
		return "string"
	}
	return "string"
}

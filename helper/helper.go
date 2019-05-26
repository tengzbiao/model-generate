package helper

import (
	"os"
	"path/filepath"
	"strings"
)

// 获取导出文件的详细路径
func GetFile(path, name string) string {
	fileInfo, _ := os.Stat(path)
	if fileInfo == nil {
		os.MkdirAll(path, os.ModePerm)
	}
	return path + string(filepath.Separator) + name + ".go"
}

// Split 增强型Split，对  a,,,,,,,b,,c     以","进行切割成[a,b,c]
func Split(s string, sub string) []string {
	var rs = make([]string, 0, 20)
	tmp := ""
	Split2(s, sub, &tmp, &rs)
	return rs
}

// Split2 附属于Split，可独立使用
func Split2(s string, sub string, tmp *string, rs *[]string) {
	s = strings.Trim(s, sub)
	if !strings.Contains(s, sub) {
		*tmp = s
		*rs = append(*rs, *tmp)
		return
	}
	for i := range s {
		if string(s[i]) == sub {
			*tmp = s[:i]
			*rs = append(*rs, *tmp)
			s = s[i+1:]
			Split2(s, sub, tmp, rs)
			return
		}
	}
}

// FindUpperElement 找到字符串中大写字母的列表,附属于HumpToUnderLine
func FindUpperElement(s string) []string {
	var rs = make([]string, 0, 10)
	for i := 0; i < len(s); i++ {
		c := i
		if s[i] >= 65 && s[i] <= 90 {
			if c+1 < len(s) && s[c+1] >= 65 && s[c+1] <= 90 {
				i++
			}
			rs = append(rs, string(s[c:c+2]))
		}
	}
	return rs
}

// HumpToUnderLine 驼峰转下划线
func HumpToUnderLine(s string) string {
	if s == "ID" {
		return "id"
	}
	var rs string
	elements := FindUpperElement(s)
	for _, e := range elements {
		s = strings.Replace(s, e, "_"+strings.ToLower(e), -1)
	}
	rs = strings.Trim(s, " ")
	rs = strings.Trim(rs, "\t")
	return strings.Trim(rs, "_")
}

// UnderLineToHump 下划线转驼峰
func UnderLineToHump(s string) string {
	arr := strings.Split(s, "_")
	for i, v := range arr {
		if v == "id" {
			v = "ID"
		}
		arr[i] = strings.ToUpper(string(v[0])) + string(v[1:])
	}
	return strings.Join(arr, "")
}

// 包含
func In(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

// 获取package名
func GetPackageName(path string) string {
	name := filepath.Base(path)
	if strings.Contains(name, ".") {
		return "models"
	}
	return name
}

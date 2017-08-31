package test1

import (
	"fmt"
	"strconv"
	"strings"
)

//版本比较
func VersionCompare() (string, bool) {
	version := "1.0.0.14"
	newsion := "1.0.0.20"
	// v, err := strconv.Atoi(version)
	vvv := strings.Replace(version, ".", "", -1)
	newvvv := strings.Replace(newsion, ".", "", -1)
	v, err := strconv.Atoi(vvv)
	nv, err := strconv.Atoi(newvvv)
	if err != nil {
		fmt.Println("数据出错", err)
		return "", false
	}
	if v < nv {
		return "有新版本", true
	}
	return "已是最新版本", false
}

package test1

import (
	"fmt"
	"strings"
)

//去掉url前的其它字符
func StringsTest1() {
	url := "dddddddddddhttp://www.baidu.com"
	index := strings.Index(url, "http")
	fmt.Println("获取的索引是：", index)
	url = url[index:]
	fmt.Println("URL:", url)
}

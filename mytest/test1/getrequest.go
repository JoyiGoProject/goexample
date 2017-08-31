package test1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

//执行方法
func RequestTest() {
	for i := 1; i <= 10; i++ {
		httpGet2()
		fmt.Println("正在执行第", i, "次访问")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("任务执行完毕")
}

//执行get请求，并且获取请求的页面数据
func httpGet2() {
	url := "http://192.168.96.35:10000/inter"
	resp, err := http.Get(url)
	if err != nil {
		beego.Error("the request is error...", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("get the data error", err)
		return
	}

	fmt.Println(string(body))
}

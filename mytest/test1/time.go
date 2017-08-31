package test1

import (
	"fmt"
	"strings"
	"time"
)

//常量时间
const (
	date        = "2006-01-02"
	shortdate   = "06-01-02"
	times       = "15:04:02"
	shorttime   = "15:04"
	datetime    = "2006-01-02 15:04:02"
	newdatetime = "2006/01/02 15~04~02"
	newtime     = "15~04~02"
	time2       = "2017/08/14"
)

//日期格式化
func TimeFormat() {
	thisdate := "2014-03-17 14:55:06"
	timeformatdate, _ := time.Parse(datetime, thisdate)
	fmt.Println(timeformatdate)
	/* 	convdate := timeformatdate.Format(date)
	   	convshortdate := timeformatdate.Format(shortdate)
	   	convtime := timeformatdate.Format(times)
	   	convshorttime := timeformatdate.Format(shorttime)
	   	convnewdatetime := timeformatdate.Format(newdatetime)
	   	convnewtime := timeformatdate.Format(newtime) */
	time22 := strings.Replace(time2, "/", "-", -1)
	timstr, _ := time.Parse(date, time22)
	/* 	fmt.Println(convdate)
	   	fmt.Println(convshortdate)
	   	fmt.Println(convtime)
	   	fmt.Println(convshorttime)
	   	fmt.Println(convnewdatetime)
		fmt.Println(convnewtime)
	*/
	fmt.Println(timstr)
}

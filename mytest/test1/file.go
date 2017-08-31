package test1

import (
	"tlog"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

//创建文件并保存
func SaveFile() {
	excelFileName := "./test.xlsx"
	File, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		tlog.Debug("文件读取错误", err)
		return
	}
	saveerr := File.Save("D:/testbak.xlsx")
	if saveerr != nil {
		tlog.Debug("文件保存失败", saveerr)
		return
	}
	beego.Info("文件保存成功")
}

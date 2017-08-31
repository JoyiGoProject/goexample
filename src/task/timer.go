package task

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/toolbox"

	"github.com/astaxie/beego"
)

func Init() {
	AutoDeleteFile()
}

// 自动删除
/*这种自动删除的任务放在运行的程序上有问题，
例如：  每隔多久删除一次日志文件，但是运行的程序上可能这个日志文件属于正在被使用的状态，就上Windows上打开的文件无法删除一样，（待解决）

*/
func AutoDeleteFile() {
	// 每隔一个月删除一次
	deleteFile := toolbox.NewTask("deleteFile", "0 0 0 */30 * *", func() error {
		beego.Debug("run deletefile start...")
		DeleteFile()
		return nil
	})

	toolbox.AddTask("deletefile", deleteFile)
	toolbox.StartTask()
}

func DeleteFile() {
	file := "./test.txt"   //源文件路径
	err := os.Remove(file) //删除文件test.txt
	if err != nil {
		beego.Error("file delete Error!", err)
	} else {
		fmt.Print("file delete OK!")
	}

}

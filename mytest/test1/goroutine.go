package test1

import (
	"fmt"
)

/*
在主方法中执行两个loop()
	01234567890123456789
将其中一个放到goroutine中：
	0123456789
	因为goroutine还未来得及执行，主函数已经执行完并退出了
*/

func GoroutineTest() {
	//将其中一个loop放到goroutine中，使用go关键字来定义并启动一个goroutine:
	go loop() //启动一个goroutine
	fmt.Println(<-complete)
	// loop()
	//阻止过早退出，加上sleep等待
	// time.Sleep(time.Second) //等待一秒钟
	//使用以上方法等待执行体验并不好，所以需要采用阻塞主线的方法——信道
	// 信道：用户goroutine之间的通信，可以在进程间传递信息
	//使用make建立信道 ,创建方式：
	/* 	var channel chan int = make(chan int)
	   	channel2 := make(chan int) */

}

var complete chan string = make(chan string)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}

	complete <- "Over" //执行完毕，发个消息
}

//建立信道，并在里面存入消息和取出消息
func ChannelMain() {
	var messages chan string = make(chan string)
	func(message string) {
		messages <- message //存消息
	}("Hello")

	fmt.Println(<-messages)
}

package construction

import (
	"fmt"
	"time"
)

//装饰器模式
// go的装饰器不像python，可以@什么什么这样
// 装饰器其实就是在某个函数执行前后做一些固定的事情

// 我们第一个例子是对一个函数执行前后收集latency
func someworks() string {
	time.Sleep(5 * time.Second)
	return fmt.Sprintf("podcontroller %s", "pod")
}

func decorator(f func() string) {
	fmt.Println("collect latency")
	now := time.Now()
	f()
	duration := time.Now().Sub(now)
	fmt.Println(duration)
}

func Call() {
	decorator(someworks)
}

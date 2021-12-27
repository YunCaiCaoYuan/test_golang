package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	// 第一个初始化函数
	f1 := func() {
		fmt.Println("in f1")
	}
	once.Do(f1) // 打印出 in f1

	// 第二个初始化函数
	f2 := func() {
		fmt.Println("in f2")
	}
	once.Do(f2) // 无输出
}

// ==== 很多时候我们是要延迟进行初始化的，所以有时候单例资源的初始化，我们会使用下面的方法

/*
// 使用互斥锁保证线程(goroutine)安全
var connMu sync.Mutex
var conn net.Conn

func getConn() net.Conn {
	connMu.Lock()
	defer connMu.Unlock()

	// 返回已创建好的连接
	if conn != nil {
		return conn
	}

	// 创建连接
	conn, _ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	return conn
}

// 使用连接
func main() {
	conn := getConn()
	if conn == nil {
		panic("conn is nil")
	}
}
 */

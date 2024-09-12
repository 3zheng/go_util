package util

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

//公共函数

/*
*****************************************************************
几个关键的ascii值 48='0', 57='9', 65='A', 90='Z', 97='a', 122='z'
函数作用：创建随机的ascii字符串(需要自己预设随机种子)
参数说明：scope	创建字符串选择的字符范围

	length	创建的字符串的长度

例子：见PublicFun_test.go
*****************************************************************
*/
func CreateRandAsciiString(scope []byte, length int) string {
	if length < 0 {
		return ""
	}

	scopeLen := len(scope)
	out := make([]byte, length)
	for i := 0; i < length; i++ {
		num := rand.Intn(scopeLen)
		out[i] = scope[num]
	}

	return string(out)
}

/*
*****************************************************************
函数作用：返回东八区（即北京时间）明天的0点0分0秒时的UNIX时间戳
参数说明：days：表示多少天，例如-3表示3天前的午夜 9表示9天后的午夜

例子：见PublicFun_test.go
*****************************************************************
*/
func GetMidnightTimer(days int) (int64, error) {
	now := time.Now()
	tomorrow := now.Add(time.Duration(days*24) * time.Hour) //加一天
	//明天的午夜时间
	strMidnight := fmt.Sprintf("%04d-%02d-%02dT00:00:00+08:00", tomorrow.Year(), tomorrow.Month(), tomorrow.Day()) //+08:00表示东八区
	t1, err := time.Parse(time.RFC3339, strMidnight)
	if err != nil {
		fmt.Println("午夜日期转换出错:", err)
		return 0, err
	}
	return t1.Unix(), nil
}

/*
*****************************************************************
函数作用：打印内存块的值
参数说明：ptr	内存起始地址

	lenghth	内存块大小

例子：见PublicFun_test.go(待补充)
*****************************************************************
*/
func DumpHex(ptr unsafe.Pointer, length uintptr) {
	p := uintptr(ptr)
	fmt.Printf("0x%08x:", p)
	var i, offset uintptr
	for i = 0; i < length; i += 16 {
		for offset = i; offset < length && offset < i+16; offset++ {
			b := *((*byte)(unsafe.Pointer(p + offset)))
			fmt.Printf(" %02x", b)
		}
	}
	fmt.Println()
}

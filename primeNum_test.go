package util

import (
	"fmt"
	"testing"
	"time"
)

func TestShowPrimeNumberV01(t *testing.T) {
	before := time.Now().UnixMilli() //UnixMilli毫秒 UnixMicro微秒 UnixNano纳秒
	ShowPrimeNumberV01(100000)
	after := time.Now().UnixMilli()
	t.Logf("1/2数算法耗时：%E ns\n", float64(after-before))
	fmt.Printf("\n1/2数算法耗时：%E ms\n", float64(after-before))
}

func TestShowPrimeNumber(t *testing.T) {
	before := time.Now().UnixMilli()
	ShowPrimeNumber(100000)
	after := time.Now().UnixMilli()
	t.Logf("平方根算法耗时：%E ns\n", float64(after-before))
	fmt.Printf("\n平方根算法耗时：%E ms\n", float64(after-before))
}

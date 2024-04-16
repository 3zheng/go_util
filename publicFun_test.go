package util

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandString(t *testing.T) {
	rand.Seed(time.Now().Unix()) //rand.New(rand.NewSource(time.Now().Unix()))

	scope := make([]byte, 26)
	//生成四位大写字母组成的字符串
	for i := 0; i < 26; i++ {
		scope[i] = byte(65 + i)
	}
	t.Log("字符串选择范围为:", string(scope))

	var str string
	for i := 0; i < 500000; i++ {
		str = CreateRandAsciiString(scope, 4)
	}
	t.Log("生成的字符串为：", str)
}

func TestMidnight(t *testing.T) {
	After11, err := GetMidnightTimer(11)
	if err != nil {
		t.Error(err)
	}
	Before4, err := GetMidnightTimer(-4)
	if err != nil {
		t.Error(err)
	}
	Today, err := GetMidnightTimer(0)
	if err != nil {
		t.Error(err)
	}
	t.Log(" After11 midnight = ", time.Unix(After11, 0),
		"\n Before4 midnight = ", time.Unix(Before4, 0),
		"\n Today midnight = ", time.Unix(Today, 0),
		"\n now = ", time.Now())
}

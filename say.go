package util

import (
	"fmt"
	"math/rand"
)

func Sayhello() {
	a := "　"
	b := []byte(a)
	a = "你"
	b = []byte(a)
	fmt.Println("hello world!!", a, b)
}

func GetRandNum() {

	fmt.Println("My favorite number is", rand.Intn(10))
}

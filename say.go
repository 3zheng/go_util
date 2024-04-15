package util

import (
	"fmt"
	"math/rand"
)

func Sayhello() {
	fmt.Println("hello world!!")
}

func GetRandNum() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

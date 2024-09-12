package util

import (
	"fmt"
	"testing"
)

func TestNewtonIterate(t *testing.T) {
	ret := SqrtByNewtonMenthod(73342499.0)
	fmt.Println("return num is", ret)
	ret = SqrtByNewtonMenthod(76.825225)
	fmt.Println("return num is", ret)
	ret = SqrtByNewtonMenthod(0.5)
	fmt.Println("return num is", ret)
	ret = SqrtByNewtonMenthod(2)
	fmt.Println("return num is", ret)
	ret = SqrtByNewtonMenthod(-76.825225)
	fmt.Println("return num is", ret)
	ret = SqrtByNewtonMenthod(-100.0)
	fmt.Println("return num is", ret)
}

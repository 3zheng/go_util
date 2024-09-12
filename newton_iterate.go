package util

import "fmt"

func SqrtByNewtonMenthod(input float64) float64 {
	if input < 0 {
		fmt.Println("输入值为负，不合法")
		return 0
	}
	ret := input //设定初始值为1/2的输入值
	fmt.Println("input number is ", input)
	//迭代100次
	for i := 0; i < 50; i++ {
		ret = (ret + input/ret) / 2
	}
	return ret
}

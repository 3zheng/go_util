package util

import (
	"fmt"
	"math"
)

/*
*****************************************************************
函数作用：// 显示输入数字为止的所有质数，使用平方根法
参数说明：

	finalNum	截止数字

*****************************************************************
*/
func ShowPrimeNumber(finalNum int) {
	if finalNum <= 1 {
		fmt.Println("输入数字不是有效数字", finalNum)
	}

	var primeNums []int
	primeNums = append(primeNums, 2) //2是质数
	for i := 3; i <= finalNum; i++ { //从3开始判断
		endNum := int(math.Sqrt(float64(i))) //一个合数的最小质因数一定小于等于该合数的平方根
		for _, elem := range primeNums {
			if i%elem == 0 { //能被整除，不是质数
				break
			}
			//if elem >= i/2 { //当循环到i的一半时仍然是质数，那么这个数就是质数
			if elem >= endNum {
				primeNums = append(primeNums, i)
				break
			}
		}
	}

	fmt.Println("primeNums capacity is ", cap(primeNums))
	fmt.Printf("从1到%d的质数共有%d个，列表如下：\n", finalNum, len(primeNums)) //打印质数
	for _, elem := range primeNums {
		fmt.Printf("%d, ", elem)
	}

}

// 不用的原有函数，使用1/2法，相比平方根更耗时
func ShowPrimeNumberV01(finalNum int) {
	if finalNum <= 1 {
		fmt.Println("输入数字不是有效数字", finalNum)
	}

	var primeNums []int
	primeNums = append(primeNums, 2) //2是质数
	for i := 3; i <= finalNum; i++ { //从3开始判断
		//endNum := int(math.Sqrt(float64(i))) //一个合数的最小质因数一定小于等于该合数的平方根
		endNum := i / 2 //当循环到i的一半时仍然是质数，那么这个数就是质数
		for _, elem := range primeNums {
			if i%elem == 0 { //能被整除，不是质数
				break
			}
			if elem >= endNum {
				primeNums = append(primeNums, i)
				break
			}
		}
	}

	fmt.Println("primeNums capacity is ", cap(primeNums))
	fmt.Printf("从1到%d的质数共有%d个，列表如下：\n", finalNum, len(primeNums)) //打印质数
	for _, elem := range primeNums {
		fmt.Printf("%d, ", elem)
	}

}

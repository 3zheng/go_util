package util

import "fmt"

// 显示输入数字为止的所有质数
func ShowPrimeNumber(finalNum int) {
	var primeNums []int
	primeNums = append(primeNums, 2) //2是质数
	for i := 3; i <= finalNum; i++ { //从3开始判断
		for _, elem := range primeNums {
			if i%elem == 0 { //能被整除，不是质数
				break
			}
			if elem >= i/2 { //当循环到i的一半时仍然是质数，那么这个数就是质数
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

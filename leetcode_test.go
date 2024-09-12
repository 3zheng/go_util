package util

import (
	"fmt"
	"testing"
	"time"
)

func TestTwoSum(t *testing.T) {
	before := time.Now().UnixMilli() //UnixMilli毫秒 UnixMicro微秒 UnixNano纳秒
	/*nums := make([]int, 100000)
	sign := -1 //符号位
	for i := 0; i < 100000; i++ {
		sign = sign * -1
		nums[i] = sign * rand.Intn(400000)
	}*/
	nums := []int{3, 2, 4}
	fmt.Println("nums are ", nums)
	ret := TwoSum(nums, 6)
	fmt.Println("ret is", ret, nums[ret[0]], nums[ret[1]])
	after := time.Now().UnixMilli()
	t.Logf("1/2数算法耗时：%E ns\n", float64(after-before))
	fmt.Printf("\n1/2数算法耗时：%E ms\n", float64(after-before))
}

func TestLengthOfLongestSubstring(t *testing.T) {

	ret := LengthOfLongestSubstring("abcabcbb")
	t.Logf("返回长度：%d ns\n", ret)
}

func TestFindMedianSortedArrays(t *testing.T) {

	ret := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	t.Logf("返回长度：%f ns\n", ret)
}

func TestLongestPalindrome(t *testing.T) {
	ret := longestPalindrome("b")
	t.Logf("返回子串：%s ns\n", ret)
}

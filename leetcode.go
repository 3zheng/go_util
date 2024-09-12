package util

import (
	"fmt"
	"slices"
)

/*
************************************************************
*从输入的数组nums里找出两个数,使得这两个数的和等于输入目标数target，并把这两个数的数组下标返回
*
*例如：Input: nums = [2,7,11,15], target = 9

	Output: [0,1]
	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

*
*
*************************************************************
*/
func TwoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	length := len(nums)
	if length < 2 {
		fmt.Println("错误：输入数组元素小于2个")
		return ret
	}
	sortedNums := make([]int, length)
	copy(sortedNums, nums)
	fmt.Println("nums is ", nums)
	slices.Sort(sortedNums) //排序

	//var num1, num2 int //找出两个数,num1<=num2,num1存较小的数,num2存较大的数
	var isfound bool

	isfound = false
	for i, v := range nums {
		_, isfound = slices.BinarySearch(sortedNums, target-v)
		if isfound {
			ret[0] = i
			for j, v2 := range nums {
				if v2 == target-v {
					ret[1] = j
					if ret[0] == ret[1] {
						isfound = false //如果元素值是输入数的1/2，那么会出现同一个数被重复记录
						continue
					}
					isfound = true
					return ret
				}
			}
		}
	}
	if !isfound {
		//没找到
		fmt.Println("数组没找到相应元素")
		return ret
	}

	return ret
}

// 折半查找
func HalfIntervalSearch(nums []int, length int, target int) (bool, int) {
	low := 0
	high := length - 1
	var middle, midVal int
	for low <= high {
		middle = (low + high) / 2
		midVal = nums[middle]
		if midVal == target {
			return true, middle
		} else if target > midVal {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return false, -1
}

/*
*两链表数翻转相加

*例如：Input: l1 = [2,4,3], l2 = [5,6,4]

	Output: [7,0,8]
	Explanation: 342 + 465 = 807.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	var retList, curRetList *ListNode
	var val1, val2 int //记录l1，l2节点的值
	retList = new(ListNode)
	retList.Val = 0
	retList.Next = nil
	curRetList = retList
	nextNode1 := l1
	nextNode2 := l2
	val1 = nextNode1.Val
	val2 = nextNode2.Val
	flag := 0 //进位标志
	for {
		//curRetList.Val = flag
		sum := val1 + val2 + flag //先把进位值加上
		if sum > 9 {
			flag = 1
			sum = sum - 10
		} else {
			flag = 0
		}
		curRetList.Val = curRetList.Val + sum

		if nextNode1.Next == nil && nextNode2.Next == nil {
			//结束循环
			if flag == 1 { //如果有进位，需要new节点后再跳出循环
				curRetList.Next = new(ListNode)
				curRetList = curRetList.Next
				curRetList.Next = nil
				curRetList.Val = 1
			}
			break
		}
		if nextNode1.Next != nil {
			nextNode1 = nextNode1.Next
			val1 = nextNode1.Val
		} else {
			val1 = 0
		}

		if nextNode2.Next != nil {
			nextNode2 = nextNode2.Next
			val2 = nextNode2.Val
		} else {
			val2 = 0
		}
		curRetList.Next = new(ListNode)
		curRetList = curRetList.Next
		curRetList.Next = nil
		curRetList.Val = 0
	}
	return retList
}

/*
无重复字符的最长子串
*/
func LengthOfLongestSubstring(s string) int {
	array := []byte(s)
	var subString, maxString string
	mapChar := make(map[byte]bool)

	for _, v := range array {
		_, ok := mapChar[v]
		if ok {
			if len(subString) > len(maxString) {
				//更新最大子串
				maxString = subString
			}
			//把该重复元素之前子串字符清空
			subArray := []byte(subString)
			for i1, v1 := range subArray {
				delete(mapChar, v1)
				if v1 == v {
					subString = string(subArray[i1+1 : len(subArray)])
					break
				}
			}
		}
		mapChar[v] = true
		subString += string(v)

	}

	//遍历完字符串后最后再判断一次maxString
	if len(subString) > len(maxString) {
		//更新最大子串
		maxString = subString
	}

	return len(maxString)
}

// 中值
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j, k int
	length1 := len(nums1)
	length2 := len(nums2)
	totalLength := length1 + length2 //总长度
	twoNums := make([]int, totalLength)
	for k < totalLength {
		if i >= length1 {
			twoNums[k] = nums2[j]
			k++
			j++
			continue
		}
		if j >= length2 {
			twoNums[k] = nums1[i]
			k++
			i++
			continue
		}
		if nums1[i] <= nums2[j] {
			twoNums[k] = nums1[i]
			k++
			i++
		} else {
			twoNums[k] = nums2[j]
			k++
			j++
		}
	}

	if totalLength%2 == 0 {
		//偶数取中间两个数的平均值
		return float64(twoNums[totalLength/2-1]+twoNums[totalLength/2]) / float64(2)
	} else {
		//奇数直接取中间数的值
		return float64(twoNums[totalLength/2])
	}
}

func longestPalindrome(s string) string {
	var subStr string
	var maxSubstrLen int
	length := len(s)
	array := []byte(s)
	for i := 0; i < length; i++ {
		for j := 1; j+i <= length; j++ {
			if IsPalindromic(array[i : i+j]) {
				if j > maxSubstrLen {
					maxSubstrLen = j
					subStr = string(array[i : i+j])
				}
			}
		}
	}
	return subStr
}

// A string is palindromic if it reads the same forward and backward.
// 判断一个字符串是否是回文，回文就是从前往后读和从后往前读都是一样的
func IsPalindromic(s []byte) bool {
	length := len(s)
	if length < 1 {
		return false //需要1个以上的字符
	}
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

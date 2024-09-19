package util

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Operation int

const (
	BeginToEnd Operation = iota
	WordOnlyWholeText
	WordOnlySpecifiedSection
)

func TextProcessing(text string, op Operation, para ...string) string {
	var outText string
	switch op {
	case BeginToEnd: //根据起始字符串和结束字符串位置，来替换包括其实字符串和结束字符串之间的所有字符
		//比如输入的字符串是"key:hello1, value:world1.jpg, key:hello2, value:world2.jpg, key:hello3, value:world3.jpg"
		//起始字符串是"key:",结束字符串是".jpg",替换内容是"hello"
		if len(para) != 3 {
			log.Panicln("BeginToEnd输入的参数数量不对")
			return outText
		}
		substr1 := para[0]
		length1 := len(substr1)
		substr2 := para[1]
		length2 := len(substr2)
		strRepalce := para[2]
		currentIndex := 0 //记录text的当前搜索位置
		//var indices []int
		for {
			index1 := strings.Index(text[currentIndex:], substr1)
			if index1 != -1 {
				//找到substr1,继续找substr2
				index2 := strings.Index(text[currentIndex+index1+length1:], substr2)
				if index2 != -1 {
					//同时找到substr1,substr2,进行内容替换
					outText += text[currentIndex:currentIndex+index1] + strRepalce
					currentIndex += index1 + length1 + index2 + length2
				} else {
					break
				}
			} else {
				//找不到直接跳出
				break
			}
		}
		//跳出循环后，把剩下的尾部内容补上
		outText += text[currentIndex:]
	case WordOnlyWholeText:
		//在指定范围内只进行单词替换
		if len(para)%2 != 0 {
			log.Println("WordOnlyWholeText输入的参数不是偶数")
		}
		r := strings.NewReplacer(para[2:]...)
		outText = r.Replace(text)
	case WordOnlySpecifiedSection:
		//在指定范围内只进行单词替换
		if len(para)%2 != 0 {
			log.Println("WordOnlySpecifiedSection输入的参数不是偶数")
		}
		begin := para[0]
		end := para[1]
		indexBegin := strings.Index(text, begin)
		indexEnd := strings.Index(text, end)
		if indexBegin == -1 || indexEnd == -1 || indexBegin > indexEnd {
			log.Printf("查找出错：indexBegin = %d, indexEnd = %d", indexBegin, indexEnd)
			break
		}
		r := strings.NewReplacer(para[2:]...)
		outText = text[:indexBegin] + r.Replace(text[indexBegin:indexEnd+len(end)]) + text[indexEnd+len(end):]
	}
	return outText
}

func OpenLocalFile(filepath string, isOverwrite bool, op Operation, para ...string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	outText := TextProcessing(string(content), op, para...)
	var outFileName string
	if isOverwrite {
		//覆盖原有文件
		outFileName = filepath
	} else {
		//保留原文件，把内容写入新建文件
		index := strings.LastIndex(filepath, ".")
		if index == -1 {
			outFileName = filepath + "_new"
		} else {
			outFileName = filepath[0:index] + "_new" + filepath[index:]
		}
	}

	fmt.Println("outFileName = ", outFileName)
	os.WriteFile(outFileName, []byte(outText), 0644)
}

package util

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func TextProcessing(text string, operation string, para ...string) {
	switch operation {
	case "substr": //根据起始字符串和结束字符串位置，来替换包括其实字符串和结束字符串之间的所有字符
		//比如输入的字符串是"key:hello1, value:world1.jpg, key:hello2, value:world2.jpg, key:hello3, value:world3.jpg"
		//起始字符串是"key:",结束字符串是".jpg",替换内容是"hello"
	}
	var indices []int
	index := strings.Index(str, substr)
	for index != -1 {
		indices = append(indices, index)
		index = strings.Index(str[index+len(substr):], substr)
		if index != -1 {
			index += indices[len(indices)-1] + len(substr)
		}
	}
	fmt.Println("子字符串出现的索引:", indices)
}

func OpenLocalFile(filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	TextProcessing(string(content))
}

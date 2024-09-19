package util

import (
	"testing"
)

func TestOpenLocalFile(t *testing.T) {
	//把img:' 到 ' 之间的所有字符替换成 'TBD1',含 img'和 '
	//OpenLocalFile("新建文本文档.txt", BeginToEnd, "image: '", "'", "'TBD1'")
	//OpenLocalFile("新建文本文档.txt", false, WordOnlyWholeText, "{", "[", "}", "]")
	OpenLocalFile("新建文本文档_new - 副本.txt", true, WordOnlySpecifiedSection,
		"WP-01", "WP-24", "TBD1", "Wall Panel de PVC Ripado 16cm*24mm*2.9m")
}

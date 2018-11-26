package main

import (
	"fmt"
	// "os"
	// "path/filex/path"
)

func main() {
	root := `/Users/yangchengzhi/yangchengzhi/test/zip/`
	path1 := `/Users/yangchengzhi/yangchengzhi/test/zip/1.txt`
	path2 := `/Users/yangchengzhi/yangchengzhi/test/zip/test/1.txt`
	content1 := path1[len(root):]
	content2 := path2[len(root):]
	fmt.Printf("%v\n%v\n", content1, content2)
}

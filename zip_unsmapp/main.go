// package main

// import (
// 	// "flag"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("index.php")
// 	if err != nil {
// 		fmt.Printf("%+v",err)
// 		return
// 	}
// 	fmt.Printf("%+v",file)	
// 	data := make([]byte, 10)
// 	// 文件句柄.Read(待放入的 byte 数组中)
// 	count, err := file.Read(data)
// 	if err != nil {
// 		fmt.Printf("%+v",err)
// 	}
// 	fmt.Printf("read %d bytes: %q\n", count, data[:count])

// 	// data = make([]byte, 10)
// 	// 文件句柄.Read(待放入的 byte 数组中)
// 	count, err = file.Read(data)
// 	if err != nil {
// 		fmt.Printf("%+v",err)
// 	}
// 	fmt.Printf("read %d bytes: %q\n", count, data[:count])
// }

package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)

}
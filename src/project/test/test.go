package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	input := json.RawMessage(`z\u0007\"\u0005\u00aa\u0010\u0002hi`)
	t := ""
	for _, r := range input {
		rint := int(r)
		if rint < 128 {
			t += string(r)
		} else {
			fmt.Printf("\n--||%d--||\n", rint)
			t += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	fmt.Printf("%x", t)
}

/*
Proto Marshal:============
z"ªhi
7a 07 22 05 aa 10 02 68 69
[122 7 34 5 170 16 2 104 105]
z bel " enq aa(特殊) dle stx h i

=========
22 7a 5c 75 30 30 30 37 5c 225c75303030355c75 61 61 5c75303031305c7530303032686922
"z\u0007\"\u0005\uaa\u0010\u0002hi"
*/

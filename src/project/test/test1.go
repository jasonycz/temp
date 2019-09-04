package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Model struct {
	N json.RawMessage `json:"name" bson:"name"`
}

func (m *Model) Name() string {
	return string(m.N)
}

func main() {
	s := "{\"name\": \"m\303\203ead\"}"
	r := strings.NewReader(s)
	d := json.NewDecoder(r)
	m := Model{}

	fmt.Println(d.Decode(&m))
	fmt.Println(m.Name())
}

/*
Proto Marshal:============
z"ªhi
7a 07 22 05 aa 10 02 68 69
[122 7 34 5 170 16 2 104 105]
z bel " enq aa(特殊) dle stx h i

=========
227a5c75303030375c225c75303030355c7561615c75303031305c7530303032686922
"z\u0007\"\u0005\uaa\u0010\u0002hi"
*/

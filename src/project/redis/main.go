package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

const (
	// ConnmandSet set
	ConnmandSet = "SET"
	// ConnmandGet get
	ConnmandGet = "GET"
	// ConnmandDel del
	ConnmandDel = "DEL"
)

// User ...
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// connect
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	cli := client{
		c,
	}
	key := "obj"
	// val := &User{
	// 	Name: "ycz",
	// 	Age:  28,
	// }
	val := "val"

	// set
	cli.Set(key, val)

	// get
	cli.Get(key)
}

// client ...
type client struct {
	redis.Conn
}

func (c client) Set(key string, val interface{}) {
	// string
	res, err := c.Do(ConnmandSet, key, val)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nres is %+v, err is %+v\n", res, err)
}

func (c client) Get(key string) {
	// string
	res, err := c.Do(ConnmandGet, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nres is %+v\nstring is %s\nerr is %+v\n", res, res, err)
}

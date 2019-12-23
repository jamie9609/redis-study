package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {


	//读写json到redis
	var mymapGet map[string]string

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("Connect to redis ok!")
	}
	defer c.Close()

	mykey := "profile"
	mymap := map[string]string{"username":"jamiezhangming", "id":"88888"}
	value, _ := json.Marshal(mymap)

	n, err := c.Do("setnx",mykey,value)
	if err != nil{
		fmt.Println("redis set failed:", err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}

	valueGet, err := redis.Bytes(c.Do("get", mykey))
	if err != nil{
		fmt.Println("redis get failed:", err)
	}
	errShal := json.Unmarshal(valueGet, &mymapGet)
	if errShal != nil{
		fmt.Println("redis json unmarshal failed:", err)
	}


}


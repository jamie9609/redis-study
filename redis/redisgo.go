package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main1() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("Connect to redis ok!")
	}
	defer c.Close()

//设置过期呢，可以使用SET的附加参数c.Do("set","mykey", "jamiezhangming","EX","5")
	_, err = c.Do("set","mykey", "jamiezhangming","EX","5")
	if err != nil{
		fmt.Println("redis set failed:", err)
	}
	username, err := redis.String(c.Do("get", "mykey"))
	if err != nil{
		fmt.Println("redis set failed1:", err)
	} else {
		fmt.Printf("get mykey: %v \n", username)
	}

	time.Sleep(7*time.Second)

	//批量写入读取
	//
	//MGET key [key …]
	//MSET key value [key value …]
	//
	//批量写入读取对象(Hashtable)
	//HMSET key field value [field value …]
	//HMGET key field [field …]
	//
	//检测值是否存在
	//EXISTS key
	//删除
	//DEL key [key …]

	username, err = redis.String(c.Do("get", "mykey"))
	if err != nil{
		fmt.Println("redis set failed1:", err)
	} else {
		fmt.Printf("get mykey: %v \n", username)
	}

}

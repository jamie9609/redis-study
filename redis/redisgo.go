package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("Connect to redis ok!")
	}
	defer c.Close()

	_, err = c.Do("set","mykey", "jamiezhangming")
	if err != nil{
		fmt.Println("redis set failed:", err)
	}
	username, err := redis.String(c.Do("get", "mykey"))
	if err != nil{
		fmt.Println("redis set failed1:", err)
	} else {
		fmt.Printf("get mykey: %v \n", username)
	}
}

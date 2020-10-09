package myredis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		// Addr:     "localhost:6379",
		Addr:     "192.168.31.229:32768",
		Password: "", //默认空密码
		DB:       0,  //使用默认数据库
	})

	//defer client.Close() //最后关闭
}

//ConnectRedis 测试redis连接
func ConnectRedis() {
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected result: ", pong)

}

//SetKey ...
func SetKey(k, v string) {
	client.Set(k, v, 0)
}

//GetKey ...
func GetKey(k string) string {
	strcmd := client.Get(k)
	return strcmd.Val()
}

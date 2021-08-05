package myredis

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client
var count int

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
	client.Set(k, v, 0).Err()
}

//GetKey ...
func GetKey(k string) string {
	ret, err := client.Get(k).Result()
	if err == redis.Nil {
		fmt.Printf("name do not exist")
		return ""
	} else if err != nil {
		fmt.Printf("get name faild err:%v\n", err)
		return ""
	} else {
		return ret
	}
}

// Add ...
func Add() {
	var wg sync.WaitGroup
	var l sync.Mutex
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			l.Lock()
			count++
			l.Unlock()
		}()
	}
	wg.Wait()
	println(count)
}

//AddWithRedisLock ...
func AddWithRedisLock() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 60)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}

func incr() {
	lockKey := "count_key"
	countKey := "counter"
	//加锁
	//defer client.Close()
	cmdLock := client.SetNX(lockKey, 1, time.Second*3)
	issucceed, err := cmdLock.Result()
	if err != nil || !issucceed {
		//client.Del(lockKey)
		fmt.Println("创建锁失败", err)
		return
	}
	//获取数值
	cmdGet := client.Get(countKey)
	v, err := cmdGet.Int64()
	if err != nil {
		//client.Del(lockKey)
		fmt.Println("get失败", err)
	}
	v++

	//设置数值
	cmdSet := client.Set(countKey, v, 0)
	if _, err := cmdSet.Result(); err != nil {
		//client.Del(lockKey)
		println("set value error!")
	}

	println("current counter is ", v)

	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed")
	}
}

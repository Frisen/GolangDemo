package myredis_test

import (
	"GolangDemo/myredis"
	"fmt"
	"testing"
)

func Test_ConnectRedis(t *testing.T) {
	myredis.ConnectRedis()
	myredis.SetKey("huang", "黄圆圆0001")
	v := myredis.GetKey("huang")
	fmt.Println("getkey value ---------->", v)
}

func Test_Add(t *testing.T) {
	myredis.Add()
}

func Test_AddWithRedisLock(t *testing.T) {
	myredis.AddWithRedisLock()
}

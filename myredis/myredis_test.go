package myredis_test

import (
	"GolangDemo/myredis"
	"fmt"
	"testing"
)

func Test_ConnectRedis(t *testing.T) {
	myredis.ConnectRedis()
	myredis.SetKey("huang.xinhui", "我是黄馨慧")
	v := myredis.GetKey("huang.xinhui")
	fmt.Println("getkey value ---------->", v)
}

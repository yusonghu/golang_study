package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的redisdb变量
var redisdb *redis.Client

//	redis连接<初始化客户端>
func initClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		println(err)
		return err
	}
	return err
}

//	设置值
func setStringVal() {
	err := redisdb.Set("count", 0, -1).Err()
	if err != nil {
		println(err)
		return
	}
}

//获取值
func getStringVal() {
	result, err := redisdb.Get("count").Result()
	if err == redis.Nil {
		fmt.Println("没有该key")
	} else if err != nil {
		println(err)
	} else {
		fmt.Println(result)
	}
}

func getListVal() {
	result, err := redisdb.Get("list").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

//设置list
func setListVal() {
	err := redisdb.LPush("list", 10).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func init() {
	initClient()
}

func main() {
	setListVal()
	getListVal()
}

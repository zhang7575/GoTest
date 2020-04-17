package main

import (
	"fmt"
	redis "github.com/gomodule/redigo/redis"
	"time"
)

func getChannel() <-chan time.Time{
	c:=make(chan time.Time)
	go func(){
		for i:=0; i<100;i++{
			time.Sleep(200*time.Millisecond)
			c<-time.Now()
		}
	}()
	return c
}

func main() {
	//for t:=range getChannel(){
	//	fmt.Println("start")
	//	fmt.Println(t)
	//}
	redisRead()
	go first(56)
	fmt.Println("sleep 0")

	go func(){
		fmt.Println("tick")
		//time.Tick(200 * time.Millisecond)
		//c<-time.Now()
	}()
	fmt.Println("sleep1")
	//go first()
	time.Sleep(time.Second*1)
	//go first()
	//t:=time.Tick(200 * time.Millisecond)
	//fmt.Println(reflect.TypeOf(t))
	//for t1 := range time.Tick(200 * time.Millisecond) {
	//	fmt.Println(t1)
	//	fmt.Println(reflect.TypeOf(t1))
	//}

	//go func(ch chan time.Time) {
	//	j:=<-time.After(time.Second*1)
	//	ch<-j
	//}(c)

	////close(c)
	//i,err:=<-c
	//fmt.Print(err)
	//fmt.Print(i)
	//time.Tick()
}

func redisRead(){
	c,_ := redis.Dial("tcp","localhost:6379")
	defer c.Close()
	ret,_:=redis.String(c.Do("get","jzhang13"))
	fmt.Println(ret)
}

func first(i int){
	fmt.Println("sleep")
	//time.Sleep(time.Second*5)
}





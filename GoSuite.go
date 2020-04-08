package main

import (
	"fmt"
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
	for t:=range getChannel(){
		fmt.Println("start")
		fmt.Println(t)
	}


	//go func(){
	//	time.Tick(200 * time.Millisecond)
	//	//c<-time.Now()
	//}()
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
	//go first()
	////close(c)
	//i,err:=<-c
	//fmt.Print(err)
	//fmt.Print(i)
	//time.Tick()
}

func first(){
	time.Sleep(time.Second*30)
}





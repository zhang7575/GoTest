package main

import (
	"GoServer/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func queryEmployee(c *gin.Context){
	ids := c.Param("id")
	id,_:=strconv.Atoi(ids)
	emp,err:=model.QueryEmployee(int(id))
	if(err!=nil){
		c.JSON(http.StatusInternalServerError,gin.H{"code":1, "message":err})
	}
	ret,_ := json.Marshal(emp)
	c.String(http.StatusOK,string(ret))
}

func addEmployee(c *gin.Context){
	var emp model.Employee
	data,_:=c.GetRawData()
	json.Unmarshal(data, &emp)
	err:=model.InsertEmployee(emp.Name,emp.Manager,emp.Dept)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"code":1, "message":err})
	}
	c.JSON(http.StatusOK,gin.H{"code":0, "message":"insert successfully"})
}

func main() {
	model.Init()
	defer model.CloseDB()
	router := gin.Default()
	router.GET("/employee/:id", queryEmployee)
	router.POST("/employees/",addEmployee)
	router.Run()
}

//type Company struct{
//	name string
//	code int
//}
//
//func (comp Company) getName() string{
//	return comp.name
//}
//
//func main() {
//	department:=new(model.Department)
//	dep:=model.Department{1,"abc"}
//
//	fmt.Println(department)
//	//fmt.Print("go test")
//	flag.Parse()
//	defer glog.Flush()
//
//	glog.Info("This is info message")
//	glog.Error("fatal message")
//	var list = list.New()
//	list.PushBack(1)
//	list.PushBack("test")
//	for e:=list.Front();e!=nil;e=e.Next(){
//		fmt.Println(e.Value)
//	}
//	var company Company
//	company.name = "origin"
//	fmt.Println("ok "+company.getName())
//	mymap:=make(map[string]Company)
//	mymap["mycompany"] = company
//	for cc:=range mymap{
//		fmt.Println("map")
//		fmt.Println(mymap[cc])
//	}
//	//changeComp(&company)
//
//	fmt.Println(company.name)
//	fmt.Printf("test %v  ",company)
//	//var p *string
//	//var abc string = "OK"
//	//var abc1 string = "sdsdffd"
//	//p = &abc
//	//fmt.Println(p)
//	//fmt.Printf("%d %d", unsafe.Sizeof(abc),unsafe.Sizeof(abc1))
//	////var array  = [4]int{1,2,3,4}
//	//for i:=0;i<10;i++{
//	//	fmt.Println(i);
//	//}
//	mytest(20,5)
//
//}
//
//func mytest(x int, y int) string{
//	z:=x+y
//	var ret string
//	ret = strconv.Itoa(z)
//	fmt.Println(ret)
//	return ret
//}
//
//func changeComp(comp *Company){
//	comp.name = "changed"
//}

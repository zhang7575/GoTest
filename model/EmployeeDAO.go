package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var db *sql.DB

func Init(){
	var err error
	var mux sync.Mutex
	mux.Lock()
	db, err = sql.Open("mysql", "admin:password@tcp(127.0.0.1:3306)/tiktok")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	mux.Unlock()
}

func CloseDB(){
	if db!=nil {
		db.Close()
	}
}

func InsertEmployee(name string,manager int, department int) error{
	stmtIns, err := db.Prepare("INSERT INTO `EMPLOYEE`(`id`, `name`, `manager`, `department`) VALUES (NUll,?,?,?)")
	defer stmtIns.Close()
	if err!=nil{
		panic(err.Error())
	}
	_, err = stmtIns.Exec(name,manager,department)
	return err
}

func QueryEmployee(id int) (*Employee,error) {
	stmtOut, err := db.Prepare("SELECT * FROM `EMPLOYEE` WHERE id = ?")
	//l:=list.New()
	//l.Init()
	var emp Employee
	//ret:=new(Employee)
	defer stmtOut.Close()
	if err!=nil{
		panic(err.Error())
	}
	//result,err:= stmtOut.Query(id)
	result:= stmtOut.QueryRow(id)
	err = result.Scan(&emp.Id,&emp.Name,&emp.Manager,&emp.Dept)
	if(err!=nil){
		fmt.Println(err)
	}
	return &emp,err
}


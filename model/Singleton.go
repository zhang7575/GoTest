package model
import(
	"sync"
)

type Singleton struct{
	employee Employee
}

var instance Singleton

func (singleton *Singleton) getInstance() *Singleton{
	var once sync.Once
	once.Do(func(){
		instance = Singleton{employee:Employee{}}
	})
	return &instance
}

package model

type Department struct {
	Id   int
	Name string
}

type Employee struct {
	Id      int `json:"id"`
	Name    string `json:"name"`
	Manager int	`json:"manager"`
	Dept    int `json:"department"`
}

//func NewEmployee(id int,name string,manager int,department *Department) *Employee{
//	employee:=new()
//}

func (employee *Employee) GetManager() *Employee {
	return new(Employee)
}

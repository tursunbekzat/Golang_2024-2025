package main

import (
	"fmt"
	"math"
	 "encoding/json"
)

type Person struct{
	name string
	age int
}

func Greet(p Person) string {
	return "Hello, " + p.name
}

type Employee struct {
	Name string
	ID string
}

func (e Employee) Work() {
	fmt.Printf("Employee name : %v, ID : %v\n", e.Name, e.ID)
}

type Manager struct {
	Employee
	Department string
}

type Shape interface {
	Area() float64
}

type Circle struct{
	radius float64
}

func (c Circle) Area() float64{
	return (c.radius * c.radius * math.Pi)
}

type Rectangle struct{
	a float64 
	b float64
}

func (r Rectangle) Area() float64{
	return (r.a * r.b)
}

func PrintArea(sh Shape) float64{
	return sh.Area()
}

type Product struct{
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`

}

func encode(p Product) string{
	jsonData, err := json.Marshal(p)
    if err != nil {
        return("Error encoding JSON:")
    }

    return string(jsonData)
}

func decode(jsonString string){
    var p Product
    err := json.Unmarshal([]byte(jsonString), &p)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    fmt.Printf("Decoded struct: %+v\n", p)
}

func main(){
	var person Person
	person.name = "Bekzat"
	person.age = 20
	fmt.Println(Greet(person))

	var manager Manager
	manager.Employee = Employee{Name: "Bekzat", ID: "21b030726"}
	manager.Department = "KBTU"
	manager.Work()

	circle := Circle{radius: 3}
	rectangle := Rectangle{a: 4, b: 5}
	fmt.Println(PrintArea(circle))
	fmt.Println(PrintArea(rectangle))

	p := Product{Name: "laptop",Price: 400000, Quantity: 1}
	jsonData := encode(p)
	fmt.Println("JSON encoded: ", jsonData)
	decode(jsonData)
}
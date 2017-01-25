package main

import "fmt"
	
import t "strings"

var p = fmt.Println
	
type geometry interface {
    area() float64

}

type person struct{

name string
age int

}

type Response1 struct {
    Page   int
    Fruits []string
}
type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}


type rect struct {

width, height float64

}

func (r rect) area() float64 {
    return r.width * r.height
}

func measure(g geometry) {
 
    fmt.Println(g.area())

}

func main(){
fmt.Println("Hello World")
a := "This is a variable"
fmt.Println(a)

s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

p("Contains:  ", t.Contains("test", "es"))

r := rect{width: 3, height: 4}
measure(r)

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))
    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

}
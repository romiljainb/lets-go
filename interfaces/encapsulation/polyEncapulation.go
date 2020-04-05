package main

import (
	"fmt"
)

// thest ut Interfaces provide generic algos, hidden implementation, interception points
//implicit satisfaction break dependencies
// type assertions to extend behaviors, classify errors, and maintain compatibility

type Person struct {
	name string
	age int

}


type Lion struct {
	id int
	age int

}

func (p *Person) Jump() {
	fmt.Println("Person jumped!")
}

func (p *Person) Speak() {
	fmt.Println("Person spoke!")
}



func (l *Lion) Jump(){
	fmt.Println("Lion jumped!")
}

func (l *Lion) Roar(){
	fmt.Println("Lion roared!")
}


func GetLion(id int, age int) Jumper {
	l := Lion{id:id, age:age}
	return &l
}


func GetPerson(name string, age int) Jumper {
	p := Person{name:name, age:age}
	return &p
}


type Jumper interface {
	Jump()
}

type Roarer interface {
	Roar()
}

type Speaker interface {
	Speak()
}

func main() {
    var jumpers []Jumper
	j1 := GetPerson("asdf",5)
	j2 := GetLion(1099,3)
	j3 := GetPerson("anshu",3)

    jumpers = append(jumpers, j1,j2,j3)

    for _,jumper := range jumpers{

        if s, ok := jumper.(Roarer); ok {
            s.Roar()
        }
        if s, ok := jumper.(Jumper); ok {
            s.Jump()
        }
        if s, ok := jumper.(Speaker); ok {
            s.Speak()
        }

        /*
        Cannot do fallthrough because switch can only  select one type
        switch s := jumper.(type) {
        case Roarer:
            s.Roar()
            fallthrough
        case Jumper:
            s.Jump()
            fallthrough
        case Speaker:
            s.Speak()
        default:
            fmt.Println("not sure what type")
        }
        //fmt.Println(jumper)
        */

    }


    /*
	j1.Jump()
    j1.(Speaker).Speak()

	j2.Jump()
    j2.(Roarer).Roar()
    */

}

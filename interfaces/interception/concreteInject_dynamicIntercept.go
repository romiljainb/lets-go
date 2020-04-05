package main

import (
    "fmt"
)


type Peasant struct {
    name string
    sc Scriber

}

type Philosopher struct {
    name string
}

type Journalist struct {
    name string
}

func (j *Journalist) Scribe(s string) {
    fmt.Println("Journalist: "  + s)
}

func (p *Philosopher) Scribe(s string) {
    fmt.Println("Philosopher: "  + s)
}

type Scriber interface {
    Scribe(s string)
}

func main(){

    var phil Scriber = &Philosopher{name:"Angelo"}
    var journo Scriber = &Journalist{name:"loser"}

    p1 := Peasant{
                    name:"no more peasant",
                    sc: phil,
                }
    p1.sc.Scribe(p1.name)

    p2 := Peasant{
                    name:"no more peasant",
                    sc: journo,
                }
    p2.sc.Scribe(p1.name)

    p3 := Peasant{
                    name:"no more peasant",
                    sc: &Journalist{name:"Mical"},
                }
    p3.sc.Scribe(p3.name)

}

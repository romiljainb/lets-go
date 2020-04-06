package main

import (
    "fmt"
    "strconv"
    "math/rand"
    "time"
)

    // collect results from game action/ops from main gamers.
    // based on incoming results/action types do something this will trigger other action and send those in another channels (game mangaer)
    // concurreently log the action/results from players
    // second part of this is to use batch tasks/job that worker pools will take care of

var Action  = make(chan string)
var Events  = make(chan string)

type Character struct {
    name string
    charType string
    hp int
    attack int
}

func (c *Character) loseHealth(i int) string {
    c.hp -= i*10
    return fmt.Sprintf("Character %s has lowered hp to %d ", c.name, c.hp)
}

func (c1 *Character)  Attack (c2 *Character) string {
    c2.hp -= c1.attack
    return fmt.Sprintf(c1.name + " makes " + c2.name + " lose " + strconv.Itoa(c1.attack)  + " hp" )
}

func NewCharacter(name string, charType string, hp int, attack int) *Character {
    return &Character{name: name, charType: charType, hp :hp  , attack : attack} 
}

func main() {

    var chars []*Character
    rand.Seed(time.Now().UnixNano())

    for i := 0 ; i < 10; i++ {
        n :=  "k" + strconv.Itoa(i)
        chars = append(chars, NewCharacter(n, "knight", 100, i*10))
    }

    a := 0
    b := 9
    for i, c := range chars {
        go func(c *Character, i int) {
            otherChar := chars[a + rand.Intn(b-a+1)]
            Action <- c.Attack(otherChar)
            if otherChar.hp <= 20 {
                Events <- fmt.Sprintf("Character %s is event with current hp of only %d left.", otherChar.name, otherChar.hp) 
            }
        }(c,i)
    }


    //gamestate := true
    for i:= 0 ; i < len(chars); i++{
        select {
        case action := <-Action :
            fmt.Println(action)
        case event := <-Events :
            fmt.Println(event)
        }
        
    }
}

package main

import "fmt"

type Applicant interface {
	Apply(a *Agent)
}

type Agent struct {
	Name         string
	Applications []Applicant
	Weapon       string
	OffHand      string
}

type Thing struct {
	Name   string
	Value  string
	Action func(a *Agent)
}

func (t Thing) Apply(a *Agent) {
	t.Action(a)
}

var Sword = Thing{
	Name: "Sword",
	Action: func(a *Agent) {
		a.Weapon = "Sword"
	},
}

var Shield = Thing{
	Name: "Shield",
	Action: func(a *Agent) {
		a.OffHand = "Shield"
	},
}

var Hero = Agent{
	Name:         "JoeBobHenryBob",
	Applications: []Applicant{Sword, Shield},
}

func main() {
	fmt.Printf("%+v\n", Hero)
	for i, a := range Hero.Applications {
		if v, ok := a.(Applicant); ok {
			fmt.Printf("%+T\n", v)
			fmt.Println(i, a)
			a.Apply(&Hero)
		}
		//fmt.Println(i, a)
		//a.Apply(&Hero)
	}
	fmt.Printf("%+v\n", Hero)
}

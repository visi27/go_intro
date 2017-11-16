package main

import "fmt"

func main() {
	goku := &Sayan{
		Person: &Person{"Goku"},
		Power: 14000,
		Friends: map[string]*Sayan{
			"Krillin": {&Person{"Krillin"}, 3000, nil},
			"Vegheta": {&Person{"Vegheta"}, 7000, nil},
		},
	}


	goku.Super()
	fmt.Printf("Power of %s is %d\n", goku.Person.Name, goku.Power)
	goku.Person.Yell()
	goku.Yell()
	goku.PrintFriends()
}

type Person struct {
	Name string
}

type Sayan struct {
	*Person
	Power int
	Friends map[string]*Sayan
}

func (s *Sayan) Super() {
	s.Power += 10000
}

func (p *Person) Yell() {
	fmt.Printf("I am a MAN, my name is %s\n", p.Name)
}

func (s *Sayan) Yell() {
	fmt.Printf("I am SUPER SAYAAAAAAN %s\n", s.Name)
}

func (s *Sayan) PrintFriends() {
	for _, sayan := range s.Friends {
		fmt.Println("Friend:", sayan.Name, "Power:", sayan.Power)
	}
}

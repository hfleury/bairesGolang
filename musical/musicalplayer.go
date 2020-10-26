package musical

import "fmt"

// Trumpeter represent a Trumpeter character
type Trumpeter struct {
	name string
}

// Greetings is a function to print a greeting
func (t *Trumpeter) Greetings() string {
	return fmt.Sprintf("Hi, I am %s a Trumpeter.", t.name)
}

// Violinist represent a Violinist character
type Violinist struct {
	name string
}

// Greetings is a function to print a greeting
func (v *Violinist) Greetings() string {
	return fmt.Sprintf("Hi, I am %s a Violinist.", v.name)
}

type MusicalPlayer interface {
	Greetings() string
}

func Greet(m MusicalPlayer) string {
	return fmt.Sprintf(m.Greetings())

}

// Musiclaplayermain is the main function of this file
func Musiclaplayermain() {
	bone := &Violinist{"Bone"}
	maison := &Trumpeter{"Maison"}

	players := [...]MusicalPlayer{bone, maison}

	for _, player := range players {
		fmt.Println(Greet(player))
	}

}

package typeswitch

import "fmt"

type Exercise interface {
	TypeSwitch(i interface{})
}

type BairesDev string

func (b BairesDev) TypeSwitch(i interface{}) {
	switch receiveType := i.(type) {
	case int:
		fmt.Printf("%v is a integer type\n", receiveType)
	case string:
		fmt.Printf("%q is a string type\n", receiveType)
	case bool:
		fmt.Printf("%t is a boolean type\n", receiveType)
	case float32, float64:
		fmt.Printf("%f is a float type\n", receiveType)
	default:
		fmt.Printf("Other type %T!\n", receiveType)
	}
}

// Typemain is the main function of this file
func Typemain() {
	var exercise Exercise = BairesDev("Teste")
	exercise.TypeSwitch(20)
	exercise.TypeSwitch("HEnrique")
	exercise.TypeSwitch(2.222)
}

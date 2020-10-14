package strings

import "fmt"

const (
	strLit0 = "this is a string"
	strLit1 = `this is a "string"
	
	sdfasdf
	`
	// @R("")
)

//Run this should be commented
func Run() {
	runUTF8()
}

func runAt() {
	s := "01234"
	fmt.Println(s[0:4])
}

func runUTF8() {
	s := "😀Hi there 😀!!"
	fmt.Println(s)
	fmt.Println(s[0:4])

	s = strLit1
	fmt.Println(s)

}

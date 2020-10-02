package main

import (
	"fmt"
)

func main() {
	exercise_one()
	exercise_two()
	exercise_three()
	exercise_four()
	exercise_five()
}

func exercise_one() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
}

func exercise_two()  {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)
}

func exercise_three()  {
	const (
		a = 42
		b int = 43
	)

	fmt.Println(a, b)
}

func exercise_four() {
	a := `here is something
	as 
	a 
	raw string
	literal
	"you see"
	another thing`
	fmt.Println(a)
}

func exercise_five()  {
	const (
		a = 2017 + iota
		b = 2017 + iota
		c = 2017 + iota
		d = 2017 + iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
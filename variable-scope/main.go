package main

//global
var x = 4

func f() {
	println("The value of x:", x)
}

func g() {
	println("The value of x:", x)
}

//local or set in a function
func n() {
	var a = 10
	println("The value of a:", a)
}

// func m() {
// 	println("The value of a:", a)
// }

func main() {
	f()
	g()
	n()
}
package main

func foo(c chan int, vals int) {
	c <- vals * 5
}
func main() {
	val := make(chan int)

	go foo(val, 4)
	go foo(val, 6)

	v1 := <- val
	v2 := <- val
}

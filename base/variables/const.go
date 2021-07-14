package main

/*
	const常量

 */

const(
	ERROR_500 = 500
	ERROR_404 = 404
)
func main(){
	const a, b, c = 1, false, "c"
	println(a, b, c)
	println(ERROR_404, ERROR_500)
}

package main
import "fmt"
var x =0
func main() {
	incr := func () int{   //anonymous function
		x++
		return x
	}
	fmt.Println(incr())
	fmt.Println(incr())
}
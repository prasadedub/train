package main
import "fmt"

func main(){
	x := 0
	incr := func() int{
		x ++
		return x
	}
	fmt.Println(incr())
	fmt.Println(incr())
}
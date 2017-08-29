package main

func wrapper() func() int{
	x := 0
	return func() int{
		x += 2
		return x
	}
}
func main(){
incr := wrapper()
fmt.Println(incr())
fmt.Println(incr())
}

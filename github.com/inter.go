package main
import "fmt"
type bird struct {
	name string 
	colour string
}
type pet struct {
	bird
	animal string
}
type life interface{
     speak()	
}
func (b bird) speak(){
	fmt.Println(b.name ,"looks" , b.colour)
}
func (p pet) speak(){
	fmt.Println(p.animal ,"looks" , p.colour)
}
func let(l life){
	l.speak()  // calling that method
}
func main(){
	b1 := bird {
		"pegion",
		"white",
	}
	let(b1)

}
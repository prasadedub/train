package main
	import (
            "fmt"
           "math"
           )   
type geo interface{
	area() float64
	peri() float64
}
type rec struct{
	len float64
	bre float64
}
type cir struct {
	radi float64
}
func (r rec) area() float64{
	return r.len*r.bre
}
func (r rec) peri() float64{
	return 2*r.len*r.bre
}
func (c cir) area() float64{
	return math.Pi*c.radi*c.radi
}
func (c cir) peri() float64{
	return 2*math.Pi*c.radi
}
func measure (g geo){
	fmt.Println(g.area())
	fmt.Println(g.peri())
}
func main(){
	r1:= rec{1 ,2}
	c1:= cir{2}
    measure(r1)
	measure(c1)
	

}

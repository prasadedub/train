package main
import "fmt"
func max(x int ) int {
	return 42 + x
}
func main (){
	max :=max(7)
	fmt.Printf("value of max %v", max)   //max func is shadowed by variable max

}
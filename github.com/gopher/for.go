package main
 import "fmt"
func main() {
	for i :=0;i<=5 ; i++{
		fmt.Printf("value of i is %d\n", i)
	}
    j :=1
for j<=5{
    j ++
    fmt.Println(j)
}
for {
	fmt.Println("loop")
	break
}
for n :=0; n<=5 ; n ++{
	if n%2 == 0{
		continue
	}
	fmt.Println(n)
}

}
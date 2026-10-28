package main
import"fmt"
func main(){
var n,k int
fmt.Scan(&n,&k)
	if n==1 || n == 2 {
	fmt.Println(1)
	return
	//break the main
	}
a := int64(1)
b := int64(1)
var c int64
	for i := 3; i<=n; i++{
	c = a + int64(k)*b
	b = a
	a = c
fmt.Println(i,"masih jalan",a," ",b," ",c)
	}
fmt.Println(a)
}

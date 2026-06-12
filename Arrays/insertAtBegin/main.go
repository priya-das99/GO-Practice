package main
import "fmt"
func main(){
	var arr=[5]int{1,2,3,5}
	size:=4
	value:=11
	for i:=size;i>0;i--{
		arr[i]=arr[i-1]
	}
	arr[0]=value
	fmt.Println("Fullarray---->",arr)
}
package main

import "fmt"

func main (){
	arr :=  []string{"satu","dua","tiga","empat"}
	perulangan(arr)
}

func perulangan(arr []string){
	//Basic general For
	for i:= 0; i< len(arr);i++{
		fmt.Println(arr[i])
	}

	//For dengan prinsip Range
	for i,v:= range arr{
		fmt.Println(fmt.Sprintf("Index ke %d  : %s",i,v))
	}

	//Simplified For (While)
	i:=0
	for i<len(arr) {
		fmt.Println(fmt.Sprintf("Index ke %d  : %s",i,arr[i]))
		i++
	}
}

func ifstatement (x,y,max int) bool{

	//Uncomment kode ini untuk melihat tipe 1
	//if x < y {
	//	return true
	//}else {
	//	return false
	//}

	//Uncomment kode ini untuk melihat tipe 2
	//if v:=x*y;v < max{
	//	return true
	//}else if v>max {
	//	return false
	//}else {
	//	return true
	//}

	//Uncomment kode ini untuk melihat tipe 3
	//if ok := x<y;ok{
	//	return ok
	//}else {
	//	return false
	//}
}



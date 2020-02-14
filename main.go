package main

import (
	"fmt"
	//"github.com/vorellan/gocurso/maps"
	

)



func main(){
	
	
	pointerTest()


}

func pointerTest(){
	a := 100
	var b *int
	b = &a
	fmt.Println(a,*b)
	fmt.Println(&a,b)
	pointerModify(b)
	fmt.Println(a,*b)
	fmt.Println(&a,b)

}

func pointerModify(c *int){
	*c = 10


}

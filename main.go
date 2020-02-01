package main

import (
	"fmt"
	//"github.com/vorellan/gocurso/maps"
	//"github.com/vorellan/gocurso/structs"

)




func main(){
	name := getName()
	fmt.Printf(name)

	a,b,c := getVariables()
	fmt.Println(a,b,c)

	numer:= sum(50,50)
	fmt.Println(numer)



}



func getName() string{
	var name string
	name = "Sin nombre"
	fmt.Printf("Ingresa tu nombre:")
	fmt.Scanf("%s", &name)
	return name


}

func getVariables() (int,int,int){
	return 1 , 2, 3

}

func sum(a int, b int) int{
	return a+b
}

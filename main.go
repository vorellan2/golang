package main

import (
	"fmt"
	//"github.com/vorellan/gocurso/maps"
	//"github.com/vorellan/gocurso/structs"

)




func main(){
	fmt.Print("hola mundo")
	/*var name string
	fmt.Println("Ingresa tu nombre:")
	fmt.Scanf("%s", &name)
	fmt.Printf("holamundo %s", "bienvenido")

	getArray()*/
	//fmt.Println(maps.GetMap())
	//fmt.Println(maps.GetMap2("Juan"))

	//PlatziCourse := PlatziCourse{Name: "Go", Slug:"go", Skills: []string{"1","2"}}
	//PlatziCourse1 := new(PlatziCourse)
	//PlatziCourse1.Name="Go1"
	//PlatziCourse.Subscribe("Cristina")
	//fmt.Println(PlatziCourse)
//	structs.InterfaceTest()

}

func getArray(){

	var arr1 [2]string
	//arr2 := [3]int{1,2,3}
	arr1[0] = "array"
	arr1[1] = "array2"

	fmt.Println(arr1)


}

func getSlice(){
	var slice1 []string 
	slice1 = append(slice1, "mi","slice")

}

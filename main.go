package main

import (
	"fmt"
	//"github.com/vorellan/gocurso/maps"
	//"github.com/vorellan/gocurso/structs"

)

type PlatziCourse struct{
	Name string 
	Slug string
	Skills []string

}

type PlatziCareer struct{
	PlatziCourse
}

func (p PlatziCourse) Subscribe(name string){
	fmt.Printf("La persona %s se ha registrado al curso %s\n ", name, p.Name )
}

func main(){
	
	//platziCourse := PlatziCourse{Name: "Go", Slug:"go", Skills: []string{"1","2"}}
	platziCourse1 := new(PlatziCourse)
	platziCourse1.Name = "Go1"
	platziCourse1.Slug = "go1"
	platziCourse1.Skills = []string{"backend1"}
	platziCourse1.Subscribe("Yohan")
	fmt.Println(platziCourse1)


}


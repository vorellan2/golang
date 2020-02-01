package structs

import "fmt"


type PlatziCourse struct{
	Name string
	Slug string
	Skills []string


}

func (p PlatziCourse) Subscribe(name string){
	fmt.Printf("La persona %s se ha registrado al curso %s", name, p.Name)
}

type PlatziCareer struct {

	PlatziCourse
}
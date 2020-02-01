package structs

type Platzi interface{

	Subscribe(name string)
}

func InterfaceTest(){
	PlatziCourse := PlatziCourse{Name: "Go", Slug:"go", Skills: []string{"1","2"}}
	PlatziCareer := new(PlatziCareer)
	PlatziCareer.Name ="GoCareer"
	PlatziCareer.Slug = "go"
	callSubstribe(PlatziCareer)
	callSubstribe(PlatziCourse)


}

func callSubstribe(p Platzi){
	p.Subscribe("Juan")
}


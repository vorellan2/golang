package maps

//GetMap devuelve un map con llave tipo string y valores tipo entero
func GetMap(name string) int{
	//mapTest := make(map[string]int)
	mapTest := map[string]int{
		"Juan":18,
		"Yohan":24,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 2
	delete(mapTest, "llave1")

	return mapTest[name]

}
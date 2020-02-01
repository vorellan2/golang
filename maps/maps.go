package maps

//devuelve un map con llave string y valores enteros
func GetMap() map[string]int {

	//los mapas necesitan ser inicializados
	//mapTest := make(map[string]int)
	mapTest := map[string]int{

		"Juan":18,
		"Yohan":20,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 2
	delete(mapTest, "llave1")
	delete(mapTest, "llave2")
	return mapTest

}


func GetMap2(name string) int {

	//los mapas necesitan ser inicializados
	//mapTest := make(map[string]int)
	mapTest := map[string]int{

		"Juan":18,
		"Yohan":20,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 2
	delete(mapTest, "llave1")
	delete(mapTest, "llave2")
	return mapTest[name]

}


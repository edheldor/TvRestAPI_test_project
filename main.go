package main

//Пример запроса {"id": 2,"brand":"LG","manufacturer":"LG","model":"1","year":2015}

func main() {
	a := App{}
	a.Initialize("root", "", "tv")
	a.Run(":8085")
}

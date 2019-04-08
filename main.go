package main

//Пример запроса {"id": 2,"brand":"LG","manufacturer":"LG","model":"1","year":2015}

func main() {
	a := App{}
	a.Initialize("root", "", "tv") //логин, пароль, навание БД в MYSQL
	a.Run(":8080")                 //порт для запуска веб сервера
}

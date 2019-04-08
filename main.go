package main

func main() {
	a := App{}
	a.Initialize("root", "", "tv")
	a.Run(":8085")
}

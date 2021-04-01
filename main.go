package main

var BuildVersion string = "0.0.0"

func main() {
	a := App{}
	a.Initialize()
	a.Run(":8080")
}

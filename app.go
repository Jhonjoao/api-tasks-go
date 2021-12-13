package main

func main() {
	a := App{}

	a.initializeRoutes()

	a.Run(":8000")
}

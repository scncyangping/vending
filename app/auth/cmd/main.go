package main

func main() {
	app, err := NewApp()
	if err != nil {
		panic(err.(interface{}))
	}
	app.RunApp()
}

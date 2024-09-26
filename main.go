package main

import (
	"shorty/api/app"
)

func main() {
	app := app.MakeApp()
	app.Start()
}

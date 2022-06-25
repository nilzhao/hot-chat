package main

import (
	"red-server/initializer"
)

func main() {
	app := initializer.New()
	app.Start()
}

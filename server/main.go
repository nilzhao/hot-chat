package main

import (
	"server/initializer"
)

func main() {
	app := initializer.New()
	app.Start()
}

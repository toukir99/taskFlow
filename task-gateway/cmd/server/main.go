package main

import (
	"task-gateway/app"
)

func main() {
	app := app.NewApplication()
	app.Init()
	app.Run()
	// app.Wait()
	// app.Cleanup()
}

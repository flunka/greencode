package main

import "github.com/flunka/greencode/router"

func main() {
	r := router.SetupRouter()
	r.Run()
}

package main

import (
	"fmt"
	"home-api/Routes"
)

func main() {
	fmt.Print("test 1 2 3")
	r := Routes.SetupRouter()
	//running
	r.Run()
}

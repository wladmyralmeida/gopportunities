package main

import "github.com/wladmyralmeida/gopportunities/router"

func main() {

	// Initialize Router
	r := router.SetupRouter()
	r.Run() // listen and serve on

}

package main

import (
	_ "go-web-new/initialize"

	"go-web-new/routes"
	_ "go-web-new/utils"
)

func main() {

	routes.InitRouter()
}

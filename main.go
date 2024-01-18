package main

import (
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/router"
)

func main() {
	r := router.SetupRouter()

	r.Run()
}
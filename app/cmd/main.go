package main

import (
	"OCsemmerApp/pkg/handler"
	"OCsemmerApp/pkg/initfunc"
)

func main() {
	initfunc.InitUpdatePassword()
	router := handler.NewRouter()
	router.Run()
}

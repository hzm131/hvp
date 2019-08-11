package main

import (
	"com/routers"
)

func main() {
	r := routers.InitRouter()

	r.Run("127.0.0.1:3000")
}

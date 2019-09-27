package main

import (
	"com/routers"
)

func main() {
	r := routers.InitRouter()

	r.Run("192.168.2.166:3000")
}

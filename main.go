package main

import (
	"com/routers"
)

func main() {
	r := routers.InitRouter()

	r.Run("169.254.78.223:3000")
}

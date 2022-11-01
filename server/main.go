package main

import (
	"blog/global"
	"blog/initialize"
)

func main() {
	initialize.MySQL()
	defer global.DB.Close()
}

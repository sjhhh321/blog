package main

import (
	"blogx_server/core"
	"fmt"
)

func main() {
	core.InitIPDB()

	fmt.Println(core.GetIpAddr("175.0.201.207"))

}

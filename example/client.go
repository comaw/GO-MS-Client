package main

import (
	"../client"
	"../client/resources"
	"fmt"
)

func main() {
	client.SetConfigPath("E:/Go/bin/spheremall/client/example/config.json")
	//var gateway = client.Config()

	body := resources.GetBrandByid(50)

	fmt.Println(body)
	fmt.Println("OK")
}

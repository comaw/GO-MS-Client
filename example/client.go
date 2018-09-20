package main

import (
	"../client"
	"fmt"
)

func main() {
	client.SetConfigPath("E:/Go/bin/spheremall/client/example/config.json")
	var gateway = client.Config()

	fmt.Println("GW url " + gateway.Url)
	fmt.Println("OK")
}

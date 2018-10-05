package main

import (
	"fmt"
	"ms-client/client"
	"ms-client/client/resources"
)

func main() {
	client.SetConfigPath("/home/shaman/go/src/ms-client/example/config.json")
	//var gateway = client.Config()

	body := resources.GetBrandByid(50)

	fmt.Println(body.Id)
	fmt.Println("OK")
}

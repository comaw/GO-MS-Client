package client

import (
	"./interfaces"
	"./makers"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetEntityById(id int32, entity string) interfaces.Entity {

	var gateway = Config()
	var url = fmt.Sprintf("%s/v1/%s/%d", gateway.Url, entity, id)

	return makers.ParsContent(HttpClient(url))
}

func HttpClient(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

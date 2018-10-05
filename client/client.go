package client

import (
	"fmt"
	"io/ioutil"
	"ms-client/client/entities"
	"ms-client/client/makers"
	"net/http"
)

func GetEntityById(id int32, entity string) entities.Brand {

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

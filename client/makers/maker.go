package makers

import "encoding/json"
import "ms-client/client/entities"

type Content struct {
	Data    []Data
	Ver     string `json:"ver"`
	Status  string `json:"status"`
	Service string `json:"service"`
}

type Data struct {
	Type       string          `json:"type"`
	Id         string          `json:"id"`
	Attributes json.RawMessage `json:"attributes"`
}

func ParsContent(content string) entities.Brand {

	c := []byte(content)
	m := Content{}

	err := json.Unmarshal(c, &m)
	if err != nil {
		panic(err)
	}

	brand := entities.Brand{}
	err2 := json.Unmarshal(m.Data[0].Attributes, &brand)
	if err2 != nil {
		panic(err2)
	}

	return brand
}

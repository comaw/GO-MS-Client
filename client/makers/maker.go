package makers

import "encoding/json"
import "../interfaces"
import "../entities"

type Content struct {
	Data    Data
	Ver     string `json:"ver"`
	Status  string `json:"status"`
	Service string `json:"service"`
}

type Data struct {
	Type       string          `json:"type"`
	Id         string          `json:"id"`
	Attributes json.RawMessage `json:"attributes"`
}

func ParsContent(content string) interfaces.Entity {
	return entities.Brand{}
}

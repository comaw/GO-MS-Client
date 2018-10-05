package resources

import (
	"ms-client/client"
	"ms-client/client/entities"
)

func GetBrandByid(int322 int32) entities.Brand {

	client.GetEntityById(int322, "brands")

	brand := client.GetEntityById(int322, "brands")

	return brand
}

package resources

import (
	"../../client"
	"../interfaces"
)

func GetBrandByid(int322 int32) interfaces.Entity {

	client.GetEntityById(int322, "brands")

	brand := client.GetEntityById(int322, "brands")

	return brand
}

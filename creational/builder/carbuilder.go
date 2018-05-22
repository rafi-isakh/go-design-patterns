// Copyright (c) 2018 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package builder

type CarBuilder struct{
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess  {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess  {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess  {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}

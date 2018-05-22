// Copyright (c) 2018 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess)  {
	// Implementation
	f.builder = b
}

func (f *ManufacturingDirector) Construct()  {
	//Implementation
	f.builder.SetStructure().SetSeats().SetWheels()
}

type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}


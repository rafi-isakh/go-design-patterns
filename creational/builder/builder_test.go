// Copyright (c) 2018 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package builder

import "testing"

func TestBuilderPattern(t *testing.T)  {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("Wheels on car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorBike := bikeBuilder.Build()
	motorBike.Seats = 1

	if motorBike.Wheels != 2 {
		t.Errorf("Wheels on motorbike must be 2 and they were %d\n", motorBike.Wheels)
	}

	if motorBike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike and was %s\n", motorBike.Structure)
	}

}

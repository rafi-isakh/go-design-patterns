// Copyright (c) 2018 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package main

func main() {

	swimmer := CompositeSwimmer{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmer.Train()
	swimmer.Swim()

	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()

}

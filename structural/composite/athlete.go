// Copyright (c) 2018 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package main

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

func Swim()  {
	fmt.Println("Swimming")
}

type Swimmer interface {
	Swim ()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct {}
func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming")
}

type CompositeSwimmer struct {
	Trainer
	Swimmer
}

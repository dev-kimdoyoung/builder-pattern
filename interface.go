package main

import "strings"

type CarBuilder interface {
	SetRegion()
	SetPrice()
	SetProducer()
	SetName()
	GetCar() Car
}

func GetBuilder(producer string) CarBuilder {
	switch strings.ToUpper(producer) {
	case "HYUNDAI":
		return NewNormalBuilder()
	case "BMW":
		return NewBMWBuilder()
	}
	return nil
}

package main

type Director struct {
	Builder CarBuilder
}

func NewDirector(builder CarBuilder) *Director {
	return &Director{Builder: builder}
}

func (d *Director) setBuilder(b CarBuilder) {
	d.Builder = b
}

func (d *Director) buildCar() Car {
	d.Builder.SetPrice()
	d.Builder.SetProducer()
	d.Builder.SetProducer()
	d.Builder.SetName()
	return d.Builder.GetCar()
}

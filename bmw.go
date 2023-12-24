package main

type BMWBuilder struct {
	Region   string
	Price    int
	Producer string
	Name     string
}

func (b *BMWBuilder) SetRegion() {
	b.Region = "germany"
}

func (b *BMWBuilder) SetPrice() {
	b.Price = 300000000
}

func (b *BMWBuilder) SetProducer() {
	b.Producer = "BMW"
}

func (b *BMWBuilder) SetName() {
	b.Name = "i330d"
}

func (b *BMWBuilder) GetCar() Car {
	return Car{
		Region:   b.Region,
		Name:     b.Name,
		Price:    b.Price,
		Producer: b.Producer,
	}
}

func NewBMWBuilder() *BMWBuilder {
	return &BMWBuilder{}
}

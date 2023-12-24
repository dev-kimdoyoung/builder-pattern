package main

type HyundaiBuilder struct {
	Region   string
	Price    int
	Producer string
	Name     string
}

func NewNormalBuilder() *HyundaiBuilder {
	return &HyundaiBuilder{}
}

func (h *HyundaiBuilder) SetRegion() {
	h.Region = "korea"
}

func (h *HyundaiBuilder) SetPrice() {
	h.Price = 2000000
}

func (h *HyundaiBuilder) SetProducer() {
	h.Producer = "hyundai"
}

func (h *HyundaiBuilder) SetName() {
	h.Name = "cn7"
}

func (h *HyundaiBuilder) GetCar() Car {
	return Car{
		Region:   h.Region,
		Price:    h.Price,
		Producer: h.Producer,
		Name:     h.Name,
	}
}

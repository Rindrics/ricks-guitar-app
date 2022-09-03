package models

type Builder int

const (
	ESP Builder = iota
	FENDER
	GIBSON
	GRETSCH
	IBANEZ
	MARTIN
	PRS
	SUHR
)

type Type int

const (
	ELECTRIC Type = iota
	ACOUSTIC
)

type Wood int

const (
	ALDER Wood = iota
	ASH
	EBONY
	MAHOGANY
	MAPLE
	ROSEWOOD
	SPRUCE
)

type Guitar struct {
	SerialNumber string
	Price        float32
	Builder      Builder
	Model        string
	Type         Type
	BackWood     Wood
	TopWood      Wood
}

func (g *Guitar) getSerialNumber() string {
	return g.SerialNumber
}

func (g *Guitar) getPrice() float32 {
	return g.Price
}

func (g *Guitar) getBuilder() Builder {
	return g.Builder
}

func (g *Guitar) getModel() string {
	return g.Model
}

func (g *Guitar) getType() Type {
	return g.Type
}

func (g *Guitar) getBackWood() Wood {
	return g.BackWood
}

func (g *Guitar) getTopWood() Wood {
	return g.TopWood
}

package guitar

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

func (g *Guitar) GetSerialNumber() string {
	return g.SerialNumber
}

func (g *Guitar) GetPrice() float32 {
	return g.Price
}

func (g *Guitar) GetBuilder() Builder {
	return g.Builder
}

func (g *Guitar) GetModel() string {
	return g.Model
}

func (g *Guitar) GetType() Type {
	return g.Type
}

func (g *Guitar) GetBackWood() Wood {
	return g.BackWood
}

func (g *Guitar) GetTopWood() Wood {
	return g.TopWood
}

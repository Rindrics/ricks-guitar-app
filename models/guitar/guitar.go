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
	Spec         *GuitarSpec
}

func (g *Guitar) GetSerialNumber() string {
	return g.SerialNumber
}

func (g *Guitar) GetPrice() float32 {
	return g.Price
}

func (g *Guitar) GetSpec() *GuitarSpec {
	return g.Spec
}

type GuitarSpec struct {
	builder  Builder
	model    string
	gtype    Type
	backWood Wood
	topWood  Wood
}

func NewGuitarSpec(builder Builder, model string, gtype Type, backWood, topWood Wood) *GuitarSpec {
	return &GuitarSpec{
		builder:  builder,
		model:    model,
		gtype:    gtype,
		backWood: backWood,
		topWood:  topWood,
	}
}

func (gs *GuitarSpec) GetBuilder() Builder {
	return gs.builder
}

func (gs *GuitarSpec) GetModel() string {
	return gs.model
}

func (gs *GuitarSpec) GetType() Type {
	return gs.gtype
}

func (gs *GuitarSpec) GetBackWood() Wood {
	return gs.backWood
}

func (gs *GuitarSpec) GetTopWood() Wood {
	return gs.topWood
}

package models

type Guitar struct {
	SerialNumber string
	Price        float32
	Builder      string
	Model        string
	Gtype        string
	BackWood     string
	TopWood      string
}

func (g *Guitar) getSerialNumber() string {
	return g.SerialNumber
}

func (g *Guitar) getPrice() float32 {
	return g.Price
}

func (g *Guitar) getBuilder() string {
	return g.Builder
}

func (g *Guitar) getModel() string {
	return g.Model
}

func (g *Guitar) getGtype() string {
	return g.Gtype
}

func (g *Guitar) getBackWood() string {
	return g.BackWood
}

func (g *Guitar) getTopWood() string {
	return g.TopWood
}

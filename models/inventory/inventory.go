package inventory

import "github.com/Rindrics/ricks-guitar-app/models/guitar"

type Inventory struct {
	guitars []*guitar.Guitar
}

func (i *Inventory) AddGuitar(
	serialNumber string,
	price float32,
	builder guitar.Builder,
	model string,
	gtype guitar.Type,
	backWood guitar.Wood,
	topWood guitar.Wood,
) {
	g := &guitar.Guitar{
		SerialNumber: serialNumber,
		Price:        price,
		Spec: guitar.NewGuitarSpec(
			builder,
			model,
			gtype,
			backWood,
			topWood,
		),
	}

	i.guitars = append(i.guitars, g)
}

func (i *Inventory) getGuitar(serialNumber string) *guitar.Guitar {
	for _, g := range i.guitars {
		if g.SerialNumber == serialNumber {
			return g
		}
	}
	return nil
}

func (i *Inventory) Search(spec guitar.GuitarSpec) []*guitar.Guitar {
	var matched []*guitar.Guitar
	for _, g := range i.guitars {
		if !g.GetSpec().Matches(spec) {
			continue
		}
		matched = append(matched, g)
	}
	return matched
}

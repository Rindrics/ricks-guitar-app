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
		Spec: guitar.GuitarSpec{
			Builder:  builder,
			Model:    model,
			Type:     gtype,
			BackWood: backWood,
			TopWood:  topWood,
		},
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
		gs := g.GetSpec()
		if builder := spec.GetBuilder(); builder != gs.GetBuilder() {
			continue
		}

		if model := spec.GetModel(); model != gs.GetModel() {
			continue
		}

		if gtype := spec.GetType(); gtype != gs.GetType() {
			continue
		}

		if backWood := spec.GetBackWood(); backWood != gs.GetBackWood() {
			continue
		}

		if topWood := spec.GetBackWood(); topWood != gs.GetBackWood() {
			continue
		}
		matched = append(matched, g)
	}
	return matched
}

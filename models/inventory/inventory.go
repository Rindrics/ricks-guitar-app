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
		Builder:      builder,
		Model:        model,
		Type:         gtype,
		BackWood:     backWood,
		TopWood:      topWood,
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

func (i *Inventory) Search(searchGuitar guitar.Guitar) *guitar.Guitar {
	for _, g := range i.guitars {
		if builder := searchGuitar.GetBuilder(); builder != g.GetBuilder() {
			continue
		}

		if model := searchGuitar.GetModel(); model != g.GetModel() {
			continue
		}

		if gtype := searchGuitar.GetType(); gtype != g.GetType() {
			continue
		}

		if backWood := searchGuitar.GetBackWood(); backWood != g.GetBackWood() {
			continue
		}

		if topWood := searchGuitar.GetBackWood(); topWood != g.GetBackWood() {
			continue
		}
		return g
	}
	return nil
}

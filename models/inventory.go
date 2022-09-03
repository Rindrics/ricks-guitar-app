package models

type Inventory struct {
	guitars []*Guitar
}

func (i *Inventory) AddGuitar(
	serialNumber string,
	price float32,
	builder Builder,
	model string,
	gtype Type,
	backWood Wood,
	topWood Wood,
) {
	g := &Guitar{
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

func (i *Inventory) getGuitar(serialNumber string) *Guitar {
	for _, g := range i.guitars {
		if g.SerialNumber == serialNumber {
			return g
		}
	}
	return nil
}

func (i *Inventory) Search(searchGuitar Guitar) *Guitar {
	for _, g := range i.guitars {
		if builder := searchGuitar.getBuilder(); (builder != "") && (builder != g.getBuilder()) {
			continue
		}

		if model := searchGuitar.getModel(); (model != "") && (model != g.getModel()) {
			continue
		}

		if gtype := searchGuitar.getGtype(); (gtype != "") && (gtype != g.getGtype()) {
			continue
		}

		if backWood := searchGuitar.getBackWood(); (backWood != "") && (backWood != g.getBackWood()) {
			continue
		}

		if topWood := searchGuitar.getBackWood(); (topWood != "") && (topWood != g.getBackWood()) {
			continue
		}
		return g
	}
	return nil
}

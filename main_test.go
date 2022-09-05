package main

import (
	"testing"

	"github.com/Rindrics/ricks-guitar-app/models/guitar"
	"github.com/Rindrics/ricks-guitar-app/models/inventory"
)

func TestFindGuitar(t *testing.T) {
	whatErinLikes := guitar.GuitarSpec{
		Builder:  guitar.FENDER,
		Model:    "Stratocastor",
		Type:     guitar.ELECTRIC,
		BackWood: guitar.ALDER,
		TopWood:  guitar.ALDER,
	}
	serial1 := "V95693"
	serial2 := "V9512"

	t.Run("ok", func(t *testing.T) {
		inventory := &inventory.Inventory{}

		inventory.AddGuitar(
			serial1,
			1499.95,
			whatErinLikes.Builder,
			whatErinLikes.Model,
			whatErinLikes.Type,
			whatErinLikes.BackWood,
			whatErinLikes.TopWood,
		)

		foundGuitar := inventory.Search(whatErinLikes)
		if foundGuitar == nil {
			t.Error("sorry, Erin")
		}
	})

	t.Run("ng", func(t *testing.T) {
		inventory := &inventory.Inventory{}

		foundGuitar := inventory.Search(whatErinLikes)
		if foundGuitar != nil {
			t.Errorf("expected nil, but got %v", foundGuitar)
		}
	})

	t.Run("multi-hit", func(t *testing.T) {
		inventory := &inventory.Inventory{}

		inventory.AddGuitar(
			serial1,
			1499.95,
			whatErinLikes.Builder,
			whatErinLikes.Model,
			whatErinLikes.Type,
			whatErinLikes.BackWood,
			whatErinLikes.TopWood,
		)

		inventory.AddGuitar(
			serial2,
			1549.95,
			whatErinLikes.Builder,
			whatErinLikes.Model,
			whatErinLikes.Type,
			whatErinLikes.BackWood,
			whatErinLikes.TopWood,
		)

		foundGuitars := inventory.Search(whatErinLikes)
		if len(foundGuitars) != 2 {
			t.Errorf("expect two guitars, found %v", len(foundGuitars))
		}
	})

}

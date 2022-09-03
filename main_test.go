package main

import (
	"testing"

	"github.com/Rindrics/ricks-guitar-app/models/guitar"
	"github.com/Rindrics/ricks-guitar-app/models/inventory"
)

func TestFindGuitar(t *testing.T) {
	whatErinLikes := &guitar.Guitar{
		Builder:  guitar.FENDER,
		Model:    "Stratocastor",
		Type:     guitar.ELECTRIC,
		BackWood: guitar.ALDER,
		TopWood:  guitar.ALDER,
	}

	t.Run("ok", func(t *testing.T) {
		inventory := &inventory.Inventory{}

		inventory.AddGuitar(
			whatErinLikes.SerialNumber,
			1499.95,
			whatErinLikes.Builder,
			whatErinLikes.Model,
			whatErinLikes.Type,
			whatErinLikes.BackWood,
			whatErinLikes.TopWood,
		)

		foundGuitar := inventory.Search(*whatErinLikes)
		if foundGuitar == nil {
			t.Error("sorry, Erin")
		}
	})

	t.Run("ng", func(t *testing.T) {
		inventory := &inventory.Inventory{}

		foundGuitar := inventory.Search(*whatErinLikes)
		if foundGuitar != nil {
			t.Errorf("expected nil, but got %v", foundGuitar)
		}
	})
}

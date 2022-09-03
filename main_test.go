package main

import (
	"testing"

	"github.com/Rindrics/ricks-guitar-app/models"
)

func TestFindGuitar(t *testing.T) {
	whatErinLikes := &models.Guitar{
		Builder:  "fender",
		Model:    "Stratocastor",
		Gtype:    "electric",
		BackWood: "Alder",
		TopWood:  "Alder",
	}

	t.Run("ok", func(t *testing.T) {
		inventory := &models.Inventory{}

		inventory.AddGuitar(
			whatErinLikes.SerialNumber,
			1499.95,
			whatErinLikes.Builder,
			whatErinLikes.Model,
			whatErinLikes.Gtype,
			whatErinLikes.BackWood,
			whatErinLikes.TopWood,
		)

		foundGuitar := inventory.Search(*whatErinLikes)
		if foundGuitar == nil {
			t.Error("sorry, Erin")
		}
	})

	t.Run("ng", func(t *testing.T) {
		inventory := &models.Inventory{}

		foundGuitar := inventory.Search(*whatErinLikes)
		if foundGuitar != nil {
			t.Errorf("expected nil, but got %v", foundGuitar)
		}
	})
}

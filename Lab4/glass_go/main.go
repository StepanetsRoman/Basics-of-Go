package main

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// Ціни за 1 см² з PDF: 6 типів склопакетів
var glassPrices = []float64{2.5, 3, 0.5, 1, 1.5, 2}

func calc(area float64, price float64, sill bool) float64 {
	sum := area * price
	if sill {
		sum += 350
	}
	return sum
}

func initGUI() {
	w := ui.NewWindow("Склопакет (Go)", 320, 320, false)
	w.SetMargined(true)

	heightEntry := ui.NewEntry()
	widthEntry := ui.NewEntry()
	combo := ui.NewCombobox()
	combo.Append("Однокамерний, дерев'яний")
	combo.Append("Двокамерний, дерев'яний")
	combo.Append("Однокамерний, металевий")
	combo.Append("Двокамерний, металевий")
	combo.Append("Однокамерний, металопластиковий")
	combo.Append("Двокамерний, металопластиковий")
	combo.SetSelected(0)

	check := ui.NewCheckbox("Підвіконня (+350 грн)")
	result := ui.NewLabel("")
	btn := ui.NewButton("Розрахувати")

	btn.OnClicked(func(*ui.Button) {
		var h, w float64
		fmt.Sscanf(heightEntry.Text(), "%f", &h)
		fmt.Sscanf(widthEntry.Text(), "%f", &w)
		area := h * w
		idx := combo.Selected()
		if idx < 0 {
			idx = 0
		}
		if idx >= len(glassPrices) {
			idx = len(glassPrices) - 1
		}
		res := calc(area, glassPrices[idx], check.Checked())
		result.SetText(fmt.Sprintf("Вартість: %.2f грн", res))
	})

	box := ui.NewVerticalBox()
	box.SetPadded(true)
	box.Append(ui.NewLabel("Висота (см):"), false)
	box.Append(heightEntry, false)
	box.Append(ui.NewLabel("Ширина (см):"), false)
	box.Append(widthEntry, false)
	box.Append(combo, false)
	box.Append(check, false)
	box.Append(btn, false)
	box.Append(result, false)

	w.SetChild(box)
	w.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	w.Show()
}

func main() {
	err := ui.Main(initGUI)
	if err != nil {
		panic(err)
	}
}

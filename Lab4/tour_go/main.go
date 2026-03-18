package main

import (
	"fmt"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// Ціни за 1 день: [країна][сезон]. Країни: 0 Болгарія, 1 Німеччина, 2 Польща. Сезон: 0 літо, 1 зима.
var tourPrices = [3][2]float64{
	{100, 150}, // Болгарія: літо, зима
	{160, 200}, // Німеччина
	{120, 180}, // Польща
}

func initGUI() {
	w := ui.NewWindow("Тур", 340, 360, false)
	w.SetMargined(true)

	daysEntry := ui.NewEntry()
	tripsEntry := ui.NewEntry()

	countryCombo := ui.NewCombobox()
	countryCombo.Append("Болгарія")
	countryCombo.Append("Німеччина")
	countryCombo.Append("Польща")
	countryCombo.SetSelected(0)

	seasonCombo := ui.NewCombobox()
	seasonCombo.Append("літо")
	seasonCombo.Append("зима")
	seasonCombo.SetSelected(0)

	guide := ui.NewCheckbox("Індивідуальний гід (+50 $/день)")
	lux := ui.NewCheckbox("Номер люкс (+20%)")
	res := ui.NewLabel("")
	btn := ui.NewButton("Обчислити")

	btn.OnClicked(func(*ui.Button) {
		var d, n int
		fmt.Sscanf(daysEntry.Text(), "%d", &d)
		fmt.Sscanf(tripsEntry.Text(), "%d", &n)
		if d <= 0 {
			res.SetText("Введіть кількість днів > 0")
			return
		}
		if n <= 0 {
			n = 1
		}
		ci := countryCombo.Selected()
		si := seasonCombo.Selected()
		if ci < 0 {
			ci = 0
		}
		if ci > 2 {
			ci = 2
		}
		if si < 0 {
			si = 0
		}
		if si > 1 {
			si = 1
		}
		pricePerDay := tourPrices[ci][si]
		sum := pricePerDay * float64(d) * float64(n)
		if guide.Checked() {
			sum += 50 * float64(d)
		}
		if lux.Checked() {
			sum *= 1.2
		}
		res.SetText(fmt.Sprintf("Ціна: %.2f $", sum))
	})

	box := ui.NewVerticalBox()
	box.SetPadded(true)
	box.Append(ui.NewLabel("Кількість днів:"), false)
	box.Append(daysEntry, false)
	box.Append(ui.NewLabel("Країна:"), false)
	box.Append(countryCombo, false)
	box.Append(ui.NewLabel("Сезон:"), false)
	box.Append(seasonCombo, false)
	box.Append(ui.NewLabel("Кількість путівок:"), false)
	box.Append(tripsEntry, false)
	box.Append(guide, false)
	box.Append(lux, false)
	box.Append(btn, false)
	box.Append(res, false)

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

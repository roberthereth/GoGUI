package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"math/rand"
)

// Callable function that creates and operates GUI
func goMD() {
	var diagnosis [5]string = [5]string{"Common Cold", "Ebola 2", "Radiation Poisoning", "Munchausen's Syndrome", "Teditis"}
	a := app.New()                                         // Create new Fyne app
	win := a.NewWindow("GoMD: Instant Medical Diagnostic") // Name window
	illness := widget.NewLabel("Diagnosis:")               // Creates and adds label to window
	win.SetContent(widget.NewVBox(                         // Create vertical alignment box, depreciated but needed
		widget.NewLabel("GoMD"),
		widget.NewLabel("Instant Medical Diagnostic"),
		illness, // Label declared outside of box, allows changes within other functions
		widget.NewButton("I'm feelin woozy!", func() { // Button that changes diagnosis
			illness.SetText("Diagnosis: " + diagnosis[rand.Intn(len(diagnosis))]) // Changes diagnosis label, adds random diagnosis
		}),
	))
	win.ShowAndRun() // Creates and runs Fyne window
}

func main() {
	goMD()
}

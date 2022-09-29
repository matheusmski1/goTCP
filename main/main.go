package main

import (
	"a3-prototipo/model"
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	//	process.Server()
	//	print("before")
	//	time.Sleep(200)
	//	print("after")
	//	go process.Client()

	myApp := app.New()
	myWindow := myApp.NewWindow("Nova Janela")

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	var p = model.Pessoa{}

	nome := widget.NewEntry()
	email := widget.NewEntry()
	telefone := widget.NewEntry()
	idade := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Nome", Widget: nome},
			{Text: "Email", Widget: email},
			{Text: "Telefone", Widget: telefone},
			{Text: "Idade", Widget: idade},
		},
		OnSubmit: func() {
			//idadeI, _ := strconv.ParseInt(idade.Text, 64, 64)
			p = model.Pessoa{Nome: nome.Text, Email: email.Text, Telefone: telefone.Text, Idade: 24}
			dialog.ShowCustom("Titulo", "fechar", container.NewHBox(canvas.NewText("lalala", green)), myWindow)
			fmt.Println(p)
			myWindow.Close()
		},
	}

	// we can also append items
	//form.Append("Text", nome)

	myWindow.SetContent(form)
	myWindow.ShowAndRun()

}

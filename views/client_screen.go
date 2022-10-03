package views

import (
	"a3-prototipo/model"
	"a3-prototipo/process"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/sys/unix"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var myApp fyne.App
var initialUnixScreen fyne.Window
var startClientBtn, startServerBtn widget.Button
var inputPort widget.Entry
var formArea *fyne.Container

func InitialScreenTCP() {
	wg.Add(2)
	renderScreen()

	inputPort = createInput()
	startServerBtn = createStartServerBtn()
	startClientBtn = createStartClientBtn()

	formArea = container.NewVBox()
	initialUnixScreen.SetContent(container.NewVBox(container.NewHBox(&inputPort, &startServerBtn), &startClientBtn, formArea))

	setHandleExit()
	initialUnixScreen.ShowAndRun()

	wg.Wait()
}

func renderScreen() {
	myApp = app.New()
	initialUnixScreen = myApp.NewWindow("Comunicação via Unix Sockets")
	initialUnixScreen.Resize(fyne.NewSize(400, 400))
}

func createInput() widget.Entry {
	return widget.Entry{PlaceHolder: "Porta"}
}

func createStartServerBtn() widget.Button {
	return widget.Button{Text: "Iniciar Servidor", OnTapped: func() {
		port, e := strconv.Atoi(inputPort.Text)
		if e != nil || port < 1024 {
			showDialogErrorPort()
		} else {
			go process.ServerTCP(&wg, inputPort.Text, initialUnixScreen)
			showDialogServerRunning(inputPort.Text)
			inputPort.Disable()
			startServerBtn.Disable()
			startClientBtn.Enable()
		}
	}}
}

func createStartClientBtn() widget.Button {
	return widget.Button{Text: "Mostrar Formulário", OnTapped: func() {
		startClientBtn.Disable()
		createForm()
	}}
}

func showDialogServerRunning(port string) {
	dialog.ShowCustom(
		"ServerTCP Started",
		"entendido",
		&canvas.Text{Text: "ServerTCP rodando na porta " + port},
		initialUnixScreen)
}

func setHandleExit() {
	initialUnixScreen.SetOnClosed(func() {
		wg.Done()
		wg.Done()
		err := unix.Unlink("/tmp/ex_go.sock")
		if err != nil {
			print(err)
		}
		print("done")
	})
}

func showDialogErrorPort() {
	dialog.ShowCustom(
		"Erro",
		"entendido",
		&canvas.Text{Text: "Você deve inserir uma porta igual ou superior a 1024"},
		initialUnixScreen)
}

func createForm() {
	var student = model.Student{}

	inputName := widget.NewEntry()
	inputEmail := widget.NewEntry()
	inputNotes := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Nome", Widget: inputName},
			{Text: "Email", Widget: inputEmail},
			{Text: "Notas", Widget: inputNotes, HintText: "separadas por -"},
		},
		OnSubmit: func() {
			notasArray := strings.Split(inputNotes.Text, "-")
			var notas []float64
			for _, arg := range notasArray {
				if n, err := strconv.ParseFloat(arg, 64); err == nil {
					notas = append(notas, n)
				}
			}
			student = model.Student{Nome: inputName.Text, Email: inputEmail.Text, Notas: notas}
			go process.ClientTCP(&wg, inputPort.Text, student)
		},
	}
	formArea.Add(form)
}

package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type VideoController struct {
	currentTime int
	totalTime   int
	label       *widget.Label
	onChange    func(int)
}

func NewVideoController(totalTime int, onChange func(int)) *VideoController {
	vc := &VideoController{
		currentTime: 0,
		totalTime:   totalTime,
		label:       widget.NewLabel("0:00 / " + formatTime(totalTime)),
		onChange:    onChange,
	}
	return vc
}

func (vc *VideoController) Render() *fyne.Container {
	backButton := widget.NewButton("<<", func() {
		vc.currentTime -= 10
		if vc.currentTime < 0 {
			vc.currentTime = 0
		}
		vc.update()
	})

	forwardButton := widget.NewButton(">>", func() {
		vc.currentTime += 10
		if vc.currentTime > vc.totalTime {
			vc.currentTime = vc.totalTime
		}
		vc.update()
	})

	controls := container.NewHBox(backButton, vc.label, forwardButton)
	return controls
}

func (vc *VideoController) update() {
	vc.label.SetText(formatTime(vc.currentTime) + " / " + formatTime(vc.totalTime))
	if vc.onChange != nil {
		vc.onChange(vc.currentTime)
	}
}

func formatTime(seconds int) string {
	min := seconds / 60
	sec := seconds % 60
	return fmt.Sprintf("%d:%02d", min, sec)
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	mainWindow := myApp.NewWindow("Exemplo Completo - Fyne")

	// Textos
	label := widget.NewLabel("Bem-vindo ao Fyne!")
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Digite algo...")

	// Botões
	button := widget.NewButton("Clique Aqui", func() {
		dialog.ShowInformation("Ação", "Você clicou no botão!", mainWindow)
	})

	// CheckBox
	check := widget.NewCheck("Ativar Opção", func(checked bool) {
		if checked {
			dialog.ShowInformation("Checkbox", "Opção ativada!", mainWindow)
		} else {
			dialog.ShowInformation("Checkbox", "Opção desativada!", mainWindow)
		}
	})

	// Seletores (DropDown)
	selectMenu := widget.NewSelect([]string{"Opção 1", "Opção 2", "Opção 3"}, func(value string) {
		dialog.ShowInformation("Seleção", "Você selecionou: "+value, mainWindow)
	})

	// Sliders
	slider := widget.NewSlider(0, 100)
	slider.OnChanged = func(value float64) {
		label.SetText(fmt.Sprintf("Valor do slider: %f", value))
	}

	// Progresso
	progress := widget.NewProgressBar()
	progress.SetValue(0.5)

	// Contêiner com separadores e abas
	tabContainer := container.NewAppTabs(
		container.NewTabItem("Geral", widget.NewLabel("Conteúdo da aba Geral")),
		container.NewTabItem("Configurações", widget.NewLabel("Configurações do App")),
	)

	// Gráficos e canvas
	rectangle := canvas.NewRectangle(color.NRGBA{R: 200, G: 100, B: 50, A: 255})
	rectangle.SetMinSize(fyne.NewSize(100, 100))

	text := canvas.NewText("Texto em Canvas", color.Gray{Y: 0x66})
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 20

	// Novo widget
	videoDuration := 300
	videoController := NewVideoController(videoDuration, func(newTime int) {
		fmt.Printf("Novo tempo: %d segundos\n", newTime)
	})

	// Layout
	layout := container.NewVBox(
		label,
		entry,
		button,
		check,
		selectMenu,
		slider,
		progress,
		rectangle,
		text,
		tabContainer,
		videoController.Render(),
	)

	mainWindow.SetContent(layout)
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
}

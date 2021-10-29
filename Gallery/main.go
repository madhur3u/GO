package main

import (
	"image/color"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var imagePath []string
var imgName []string

func main() {

	a := app.New()
	w := a.NewWindow("Gallery")
	w.Resize(fyne.NewSize(720, 600))

	head := canvas.NewText("File Name", color.White)
	head.TextStyle = fyne.TextStyle{Bold: true}
	head.Alignment = fyne.TextAlignCenter

	pathInput := widget.NewEntry()
	pathInput.SetPlaceHolder("Input a valid directory and press enter")

	pathVar := canvas.NewText("/home/madhur3u/Pictures", color.White)

	path := pathVar.Text

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if !file.IsDir() {
			ext := strings.Split(file.Name(), ".")[1]

			if ext == "png" || ext == "jpg" || ext == "jpeg" {

				imagePath = append(imagePath, path+"/"+file.Name())
				imgName = append(imgName, file.Name())
			}
		}
	}

	i := 0
	image := canvas.NewImageFromFile(imagePath[0])
	image.FillMode = canvas.ImageFillContain

	prev := widget.NewButton("<", func() {
		i--

		if i < 0 {
			i = len(imagePath) - 1
		}

		image.File = imagePath[i]
		image.Refresh()

		head.Text = imgName[i]
		head.Refresh()

	})
	next := widget.NewButton(">", func() {
		i++

		if i >= len(imagePath) {
			i = 0
		}

		image.File = imagePath[i]
		image.Refresh()

		head.Text = imgName[i]
		head.Refresh()
	})

	enter := widget.NewButtonWithIcon("Enter", theme.ConfirmIcon(), func() {

		pathVar.Text = pathInput.Text
		pathVar.Refresh()

		path := pathVar.Text

		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		imagePath = nil
		imgName = nil

		for _, file := range files {

			if !file.IsDir() {
				ext := strings.Split(file.Name(), ".")[1]

				if ext == "png" || ext == "jpg" || ext == "jpeg" {

					imagePath = append(imagePath, path+"/"+file.Name())
					imgName = append(imgName, file.Name())
				}
			}
		}

		image.File = imagePath[i]
		image.Refresh()

		head.Text = imgName[i]
		head.Refresh()
	})

	blue_color := canvas.NewRectangle(color.NRGBA{R: 0, G: 96, B: 215, A: 255})

	header := container.NewVBox(head, pathInput, enter)
	buttons := container.NewHBox(layout.NewSpacer(), container.New(layout.NewMaxLayout(), blue_color, prev), container.New(layout.NewMaxLayout(), blue_color, next), layout.NewSpacer())

	head.Text = imgName[i]
	head.Refresh()

	w.SetContent(container.New(layout.NewBorderLayout(header, buttons, nil, nil), header, image, buttons))

	w.ShowAndRun()
}

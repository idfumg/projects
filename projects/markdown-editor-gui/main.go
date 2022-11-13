package main

import (
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
	FileFilter    storage.FileFilter
}

func main() {
	cfg := NewConfig()

	// create an app
	a := app.New()

	// create a window
	window := a.NewWindow("Markdown")

	// get the user interface
	mainMenu := cfg.createMenuItems(window)
	window.SetMainMenu(mainMenu)

	// set content of the window
	window.SetContent(container.NewHSplit(cfg.EditWidget, cfg.PreviewWidget))

	// show the window and run the app
	window.Resize(fyne.Size{Width: 800, Height: 500})
	window.CenterOnScreen()
	window.ShowAndRun()
}

func NewConfig() *Config {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	edit.OnChanged = preview.ParseMarkdown
	return &Config{
		EditWidget:    edit,
		PreviewWidget: preview,
		FileFilter:    storage.NewExtensionFileFilter([]string{".md", ".MD", ".mD", ".Md"}),
	}
}

func (cfg *Config) createMenuItems(window fyne.Window) *fyne.MainMenu {
	openMenuItem := fyne.NewMenuItem("Open...", cfg.openFunc(window))
	saveMenuItem := fyne.NewMenuItem("Save", cfg.saveFunc(window))
	saveAsMenuItem := fyne.NewMenuItem("Save As...", cfg.saveAsFunc(window))
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)
	menu := fyne.NewMainMenu(fileMenu)
	cfg.SaveMenuItem = saveMenuItem
	cfg.SaveMenuItem.Disabled = true
	return menu
}

func (cfg *Config) saveAsFunc(window fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(w fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if w == nil {
				return
			}
			defer w.Close()
			if !strings.HasSuffix(strings.ToLower(w.URI().Name()), ".md") {
				dialog.ShowInformation("Error", "Please name your file with a .md extension!", window)
			}
			w.Write([]byte(cfg.EditWidget.Text))
			cfg.CurrentFile = w.URI()
			window.SetTitle(window.Title() + " " + w.URI().Name())
			cfg.SaveMenuItem.Disabled = false
		}, window)

		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(cfg.FileFilter)
		saveDialog.Show()
	}
}

func (cfg *Config) openFunc(window fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if r == nil {
				return
			}
			defer r.Close()
			data, err := ioutil.ReadAll(r)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			cfg.EditWidget.SetText(string(data))
			cfg.CurrentFile = r.URI()
			window.SetTitle(window.Title() + " " + r.URI().Name())
			cfg.SaveMenuItem.Disabled = false
		}, window)

		openDialog.SetFilter(cfg.FileFilter)
		openDialog.Show()
	}
}

func (cfg *Config) saveFunc(window fyne.Window) func() {
	return func() {
		if cfg.CurrentFile != nil {
			w, err := storage.Writer(cfg.CurrentFile)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			w.Write([]byte(cfg.EditWidget.Text))
			defer w.Close()
		}
	}
}

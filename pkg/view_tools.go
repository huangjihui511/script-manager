package pkg

import (
	"fmt"
	"jihui/script-manager/model"

	"github.com/andlabs/ui"
	"github.com/atotto/clipboard"
)

func CreateTools(box *ui.Box, getData func() (model.View, error)) *ui.Box {
	toolsbox := ui.NewHorizontalBox()
	toolsbox.SetPadded(true)
	toolsbox.Append(ui.NewLabel("tools:"), false)
	quitButton := ui.NewButton("Quit")
	quitButton.OnClicked(func(b *ui.Button) {
		ui.Quit()
	})
	toolsbox.Append(quitButton, false)
	restartButton := ui.NewButton("Update")
	toolsbox.Append(restartButton, false)
	restartButton.OnClicked(func(b *ui.Button) {
		newViewData, err := getData()
		if err != nil {
			CreateErrorUI(err)
		}
		newContentBox := CreateWindowContent(newViewData)
		box.Delete(1)
		box.Append(newContentBox, false)
	})
	demoButton := ui.NewButton("Demo")
	toolsbox.Append(demoButton, false)
	demoButton.OnClicked(func(b *ui.Button) {
		clipboard.WriteAll(model.DemoYaml)
	})
	helpButton := ui.NewButton("Help")
	toolsbox.Append(helpButton, false)
	helpButton.OnClicked(func(b *ui.Button) {
		helpWindow := CreateHelpWindow()
		helpWindow.Show()
	})
	return toolsbox
}

func CreateHelpWindow() *ui.Window {
	helpWindow := ui.NewWindow("help", 100, 50, false)
	helpWindow.SetMargined(true)
	box := ui.NewVerticalBox()
	box.SetPadded(true)
	qaBox := ui.NewVerticalBox()
	qaBox.SetPadded(true)
	qaBox.Append(ui.NewLabel("Questions & Answers"), false)
	for _, qa := range model.QAs {
		qaBox.Append(ui.NewLabel(fmt.Sprintf("Q: %s", qa.Question)), false)
		qaBox.Append(ui.NewLabel(fmt.Sprintf("A: %s", qa.Answer)), false)
	}
	box.Append(qaBox, false)
	helpWindow.SetChild(box)
	helpWindow.OnClosing(func(*ui.Window) bool {
		helpWindow.Hide()
		return true
	})
	return helpWindow
}

package pkg

import (
	"fmt"
	"jihui/script-manager/model"

	"github.com/andlabs/ui"
	"github.com/atotto/clipboard"
)

func CreateUI(getData func() (model.View, error)) bool {
	viewData, err := getData()
	if err != nil {
		CreateErrorUI(err)
		return false
	}
	err = ui.Main(func() {
		window := ui.NewWindow(fmt.Sprintf("script-management: %s", viewData.ViewName), 200, 200, false)
		window.SetMargined(true)
		box := ui.NewVerticalBox()
		toolsBox := CreateTools(box, getData)
		contentBox := CreateWindowContent(viewData)
		box.Append(toolsBox, false)
		box.Append(contentBox, false)
		window.SetChild(box)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
	})
	if err != nil {
		panic(err)
	}
	return true
}

func CreateWindowContent(viewData model.View) *ui.Box {
	box := ui.NewVerticalBox()
	box.SetPadded(true)
	for _, group := range viewData.CommandGroups {
		box.Append(ui.NewLabel(fmt.Sprintf("group-name: %s", group.GroupName)), false)
		for _, command := range group.Commands {
			hbox := CreateCommand(command)
			box.Append(hbox, false)
		}
	}
	return box
}

func CreateCommand(command model.Command) *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	copyBotton := ui.NewButton("Copy")
	detailBotton := ui.NewButton("Detail")
	hbox.Append(copyBotton, false)
	hbox.Append(detailBotton, false)
	hbox.Append(ui.NewLabel(fmt.Sprintf("%s", shortenSentence(command.Script, 50))), false)
	cmd := command.Script
	copyBotton.OnClicked(func(b *ui.Button) {
		clipboard.WriteAll(cmd)
	})
	detailBotton.OnClicked(func(b *ui.Button) {
		detailsWindow := ui.NewWindow("detail", 100, 50, false)
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel(fmt.Sprintf("Script: %s", command.Script)), false)
		if len(command.Info) == 0 {
			box.Append(ui.NewLabel(fmt.Sprintf("Info: no info")), false)
		} else {
			box.Append(ui.NewLabel(fmt.Sprintf("Info: %s", command.Info)), false)
		}
		detailsWindow.SetMargined(true)
		detailsWindow.SetChild(box)
		detailsWindow.OnClosing(func(*ui.Window) bool {
			detailsWindow.Hide()
			return true
		})
		detailsWindow.Show()
	})
	return hbox
}

func CreateErrorUI(err error) {
	err = ui.Main(func() {
		window := ui.NewWindow("script-management:error msg", 100, 100, false)
		window.SetMargined(true)
		box := ui.NewVerticalBox()
		box.SetPadded(true)
		box.Append(ui.NewLabel("error msg:"), false)
		box.Append(ui.NewLabel(err.Error()), false)
		window.SetChild(box)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

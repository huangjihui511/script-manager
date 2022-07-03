package pkg

import (
	"fmt"
	"huangjihui511/script-manager/model"

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
		window := ui.NewWindow(fmt.Sprintf("script-management: %s", viewData.ViewName), 30, 100, false)
		window.SetMargined(true)
		box := ui.NewVerticalBox()
		box.SetPadded(true)
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
		groupBox := ui.NewHorizontalBox()
		groupBox.SetPadded(true)
		hideBottom := ui.NewButton("Expand")
		groupBox.Append(hideBottom, false)
		groupBox.Append(ui.NewLabel(fmt.Sprintf("group-name: %s", group.GroupName)), false)
		box.Append(groupBox, false)
		commandsBox := ui.NewVerticalBox()
		commandsBox.SetPadded(true)
		for _, command := range group.Commands {
			commandBox := CreateCommand(command)
			commandsBox.Append(commandBox, false)
		}
		box.Append(commandsBox, false)
		commandsBox.Hide()
		hideBottom.OnClicked(func(b *ui.Button) {
			if commandsBox.Visible() {
				commandsBox.Hide()
				hideBottom.SetText("Expand")
			} else {
				commandsBox.Show()
				hideBottom.SetText("Hide")
			}
		})
	}
	return box
}

func CreateCommand(command model.Command) *ui.Box {
	commandBox := ui.NewVerticalBox()
	scriptBox := ui.NewHorizontalBox()
	scriptBox.SetPadded(true)
	copyBotton := ui.NewButton("Copy")
	detailBotton := ui.NewButton("Detail")
	scriptBox.Append(ui.NewLabel(" "), false)
	scriptBox.Append(copyBotton, false)
	scriptBox.Append(detailBotton, false)
	scriptBox.Append(ui.NewLabel(fmt.Sprintf("%s", shortenSentence(command.Script, 50))), false)
	cmd := command.Script
	copyBotton.OnClicked(func(b *ui.Button) {
		clipboard.WriteAll(cmd)
	})
	commandBox.Append(scriptBox, false)
	detailForm := ui.NewForm()
	detailForm.SetPadded(true)
	detailForm.Append("   Script:", ui.NewLabel(fmt.Sprintf("%s", command.Script)), false)
	if len(command.Info) != 0 {
		detailForm.Append("   Info:", ui.NewLabel(fmt.Sprintf(" %s", command.Info)), false)
	}
	commandBox.Append(detailForm, false)
	detailForm.Hide()
	detailBotton.OnClicked(func(b *ui.Button) {
		if detailForm.Visible() {
			detailForm.Hide()
		} else {
			detailForm.Show()
		}
	})
	return commandBox
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

package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const tableHeaders = `Category|Name|Link|Size|Date|Seeders|Leechers|Downloads`

const fakeData = tableHeaders +
	`
Anime|Kimetsu no yaiba|https://nyaa.si/download/1232132121|20MiB|11.21.2010|95|42|189
Literature|Naruto manga|https://nyaa.si/download/9098086551|1GiB|22.01.2022|34|97|300
Software|Doki Doki lit club|https://nyaa.si/download/1596796751|23MiB|06.03.2015|66|2|19`

var app = tview.NewApplication()

var query = ""
var category = ""
var filter = ""

func main() {
	table := tview.NewTable().
		SetFixed(1, 1).
		SetSelectable(true, false)
	setTableData(table, tableHeaders)
	table.SetBorder(true).SetTitle("nyaa.si")

	form := tview.NewForm().
		AddInputField("Search", "", 25, nil, func(text string) {
			query = text
		}).
		AddDropDown("Category", []string{
			"All",
			"Anime",
			"- Anime Music Video",
			"- English-translated",
			"- Non-English-translated",
			"- Raw",
			"Audio",
			"- Lossless",
			"- Lossy",
			"Literature",
			"- English-translated",
			"- Non-English-translated",
			"- Raw",
			"Live Action",
			"- English-translated",
			"- Idol/Promotional Video",
			"- Non-English-translated",
			"- Raw",
			"Pictures",
			"- Graphics",
			"- Photos",
			"Software",
			"- Applications",
			"- Games",
		}, 0, func(option string, optionIndex int) {
			category = option
		}).
		AddDropDown("Filter", []string{
			"No filter",
			"No remakes",
			"Trusted only",
		}, 0, func(option string, optionIndex int) {
			filter = option
		}).
		AddButton("Search", func() {
			setTableData(table, fakeData)
			app.SetFocus(table)
		})

	form.SetCancelFunc(func() {
		app.SetFocus(table)
	})

	table.SetSelectedFunc(func(row int, column int) {
		//a := table.GetCell(row, column)
	})

	table.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(form)
	})

	flex := tview.NewFlex().
		AddItem(form, 0, 1, true).SetDirection(tview.FlexColumn).
		AddItem(table, 0, 3, true).SetDirection(tview.FlexColumn)

	app.SetRoot(flex, true)
	app.EnableMouse(true)
	app.Run()
}

func setTableData(table *tview.Table, data string) {
	for row, line := range strings.Split(data, "\n") {
		for column, cell := range strings.Split(line, "|") {
			color := tcell.ColorWhite
			if row == 0 {
				color = tcell.ColorYellow
			} else if column == 0 {
				color = tcell.ColorLightCyan
			}
			align := tview.AlignLeft
			if row == 0 {
				align = tview.AlignCenter
			} else if column == 0 || column >= 4 {
				align = tview.AlignLeft
			}
			tableCell := tview.NewTableCell(cell).
				SetTextColor(color).
				SetAlign(align).
				SetSelectable(row != 0 && column != 0 && column != 1 && column != 3 && column != 4 && column != 5 && column != 6 && column != 7)
			if column >= 1 && column <= 3 {
				tableCell.SetExpansion(1)
			}
			table.SetCell(row, column, tableCell)
		}
	}
}

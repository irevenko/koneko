package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

const tableHeaders = `Category|Name|Link|Size|Date|Seeders|Leechers|Downloads`

const fakeData = tableHeaders +
	`
Anime|Kimetsu no yaiba|https://nyaa.si/download/321|20MiB|11.21.2010|95|42|189
Literature|Naruto manga|https://nyaa.si/download/3321|1GiB|22.01.2022|34|97|300
Software|Doki Doki lit club|https://nyaa.si/22221|23MiB|06.03.2015|66|2|19`

var query = ""
var category = ""
var filter = ""
var sortBy = ""

var nyaaCategories = []string{
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
}

var filters = []string{
	"No filter",
	"No remakes",
	"Trusted only",
}

var sortOptions = []string{
	"Downloads",
	"Comments",
	"Date",
	"Seeders",
	"Leechers",
	"Size",
}

const torrentText = `Description: lol
Comments: xd`

func main() {
	table := tview.NewTable().
		SetFixed(1, 1).
		SetSelectable(true, false)
	setTableData(table, tableHeaders)
	table.SetBorder(true).SetTitle("nyaa.si")

	form := tview.NewForm().
		AddInputField("Query", "", 24, nil, func(text string) {
			query = text
		}).
		AddDropDown("Category", nyaaCategories, 0, func(option string, optionIndex int) {
			category = option
		}).
		AddDropDown("Filter", filters, 0, func(option string, optionIndex int) {
			filter = option
		}).
		AddDropDown("Sort By", sortOptions, 0, func(option string, optionIndex int) {
			sortBy = option
		}).
		AddButton("Search", func() {
			setTableData(table, fakeData)
			app.SetFocus(table)
		})

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	numSelections := 0

	fmt.Fprintf(textView, "%s ", torrentText)
	textView.SetDoneFunc(func(key tcell.Key) {
		currentSelection := textView.GetHighlights()
		if key == tcell.KeyEnter {
			if len(currentSelection) > 0 {
				textView.Highlight()
			} else {
				textView.Highlight("0").ScrollToHighlight()
			}
		} else if len(currentSelection) > 0 {
			index, _ := strconv.Atoi(currentSelection[0])
			if key == tcell.KeyTab {
				index = (index + 1) % numSelections
			} else if key == tcell.KeyBacktab {
				index = (index - 1 + numSelections) % numSelections
			} else {
				return
			}
			textView.Highlight(strconv.Itoa(index)).ScrollToHighlight()
		}
	})
	textView.SetBorder(true)
	textView.SetTitle("Torrent Info")

	form.SetCancelFunc(func() {
		app.SetFocus(table)
	})

	table.SetSelectedFunc(func(row int, column int) {
		//a := table.GetCell(row, column)
	})

	table.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(form)
	})

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlD {
			app.SetFocus(textView)
		}
		return event
	})

	flex := tview.NewFlex().
		AddItem(form, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(textView, 0, 1, true).
			AddItem(table, 0, 2, true), 0, 4, true)

	app.SetRoot(flex, true).EnableMouse(true)
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

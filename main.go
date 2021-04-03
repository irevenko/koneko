package main

import (
	"log"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/irevenko/go-nyaa/nyaa"
	h "github.com/irevenko/koneko/helpers"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

var query = ""
var category = 0
var filter = 0
var sortBy = 0

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
	"Date",
	"Downloads",
	"Size",
	"Seeders",
	"Leechers",
	"Comments",
}

func main() {
	table := tview.NewTable().
		SetSelectable(true, false)
	table.SetBorder(true).SetTitle("nyaa.si")

	form := tview.NewForm().
		AddInputField("Query", "", 24, nil, func(text string) {
			query = text
		}).
		AddDropDown("Category", nyaaCategories, 0, func(option string, optionIndex int) {
			category = optionIndex
		}).
		AddDropDown("Filter", filters, 0, func(option string, optionIndex int) {
			filter = optionIndex
		}).
		AddDropDown("Sort By", sortOptions, 0, func(option string, optionIndex int) {
			sortBy = optionIndex
		}).
		AddButton("Search", func() {
			c := h.ConvertCategory(category)
			s := h.ConvertSort(sortBy)
			f := h.ConvertFilter(filter)
			torrents := fetchTorrents("nyaa", query, c, s, f)
			setTableData(table, torrents[:len(torrents)-1]) // remove last \n
			app.SetFocus(table)
		})

	form.SetBorder(true)
	form.SetHorizontal(true)

	form.SetCancelFunc(func() {
		app.SetFocus(table)
	})

	table.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(form)
	})

	table.SetSelectedFunc(func(row int, column int) {
		//a := table.GetCell(row, column)
	})

	flex := tview.NewFlex().
		AddItem(form, 0, 1, true).SetDirection(tview.FlexRow).
		AddItem(table, 0, 5, true).SetDirection(tview.FlexRow)

	app.SetRoot(flex, true).EnableMouse(true)
	app.Run()
}

func setTableData(table *tview.Table, data string) {
	for row, line := range strings.Split(data, "\n") {
		for column, cell := range strings.Split(line, "{}") {
			textColor := tcell.ColorWhiteSmoke
			bgColor := tcell.ColorBlack
			align := tview.AlignLeft

			if column == 0 {
				textColor = tcell.ColorYellow
			} else if column == 1 {
				textColor = tcell.ColorLightGreen
			} else if column == 2 {
				textColor = tcell.ColorRed
			} else if column == 3 {
				textColor = tcell.ColorLightSalmon
			} else if column == 4 {
				textColor = tcell.ColorLightCyan
			}

			tableCell := tview.NewTableCell(cell).
				SetTextColor(textColor).
				SetBackgroundColor(bgColor).
				SetAlign(align).
				SetSelectable(true)
			table.SetCell(row, column, tableCell)
		}
	}
}

func fetchTorrents(p string, q string, c string, s string, f string) string {
	var torrents string

	opt := nyaa.SearchOptions{
		Provider: p,
		Query:    q,
		Category: c,
		SortBy:   s,
		Filter:   f,
	}

	t, err := nyaa.Search(opt)
	if err != nil {
		log.Fatal(err)
	}

	initialLayout := "Mon, 02 Jan 2006 15:04:05 -0700"
	dateLayout := "2006-01-02"

	for _, v := range t {
		t, _ := time.Parse(initialLayout, v.Date)
		date := t.Format(dateLayout)

		name := strings.Trim(v.Name, " ")

		//link := strings.Split(v.Link, "download/")
		torrents += v.Downloads + "{}" + v.Seeders + "{}" + v.Leechers + "{}" + v.Size + "{}" + date + "{}" + name + "\n"
	}

	return torrents
}
